package csv

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/logic-building/functional-go/fp"
	"github.com/spf13/viper"
	"github.com/tushar2708/altcsv"

	"ppv-transformer/pkg/eventdetail"
)

type PPV struct {
	Fields72         [72]string
	FieldsLenUnknown []string

	Fields81Output [81]string
	MissingTMSId   string
}

func TransformCSV(csvFile, cosmosHost, cosmosEventDetailApi, token, outputFailureFile string) error {

	ppvList, err := processCSV(csvFile)
	if err != nil {
		return err
	}

	fmt.Println(len(ppvList))
	tmsIdEchostarEventIdLookup := fillEchostar3Id(ppvList, cosmosHost, cosmosEventDetailApi, token)
	outputData, missingTmsIds, outputUnknownLenFields := replaceTmsIdWithEchostar3Id(ppvList, tmsIdEchostarEventIdLookup)
	missingTmsIds = fp.RemoveStr(func(id string) bool { return strings.TrimSpace(id) == "" }, missingTmsIds)

	//fmt.Println("outputData=", outputData)
	//fmt.Println("missing tmsids", missingTmsIds)
	//fmt.Println("outputUnknownLenFields", outputUnknownLenFields)
	fmt.Println("processing with converted records")

	err = generateCSV(outputData, missingTmsIds, outputUnknownLenFields, csvFile, outputFailureFile)
	if err != nil {
		return err
	}

	return nil
}

func processCSV(csvFile string) ([]PPV, error) {
	if _, err := os.Stat(csvFile); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("csv file=%s does not exist. error=%s", csvFile, err.Error())
	}

	f, err := os.Open(csvFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file=%s. error=%s", csvFile, err.Error())
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)

	var ppvList []PPV

	lineCnt := 0
	for {
		ppv := PPV{}
		lineCnt++
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "wrong number of fields") {
				fmt.Printf("\nline number=%d has columns=%d\n", lineCnt, len(line))
				fmt.Println("expected fields length in CSV is 52. other than this is considered as wrong number of fields.")
				fmt.Println("This line will be appended at the end of CSV")
				fmt.Println(line)
			} else {
				return nil, fmt.Errorf("error reading csv line by line. line_no=%d. error=%s", lineCnt, err.Error())
			}
		}

		colCnt := len(line)
		switch colCnt {
		case 72:
			i := 0
			for j, fieldVal := range line {
				ppv.Fields72[j] = fieldVal

				// Removing 52th field for the ouput
				if j == 51 {
					if len(strings.TrimSpace(fieldVal)) == 0 {
						return nil, fmt.Errorf("52th field is empty at line=%d", lineCnt)
					}
					continue
				}
				ppv.Fields81Output[i] = fieldVal
				i++
			}

			// Add 10 empty fields
			for i < 81 {
				ppv.Fields81Output[i] = ""
				i++
			}

		default:
			ppv.FieldsLenUnknown = make([]string, colCnt)
			for j, fieldVal := range line {
				ppv.FieldsLenUnknown[j] = fieldVal
			}
		}
		ppvList = append(ppvList, ppv)
	}
	return ppvList, nil

}

func fillEchostar3Id(ppvList []PPV, cosmosHost, cosmosEventDetailApi, token string) map[string]string {
	var wg sync.WaitGroup

	totalGoroutines := viper.GetInt("cosmos.apiThreads")
	threadCnt := 0

	tmsIdEchostarEventIdLookup := make(map[string]string)

	tmsIdEchostarMapCh := make(chan map[string]string, 1000)

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for m := range tmsIdEchostarMapCh {
			for k, v := range m {
				tmsIdEchostarEventIdLookup[k] = v
			}
		}
	}()

	tmsIdMap := make(map[string]struct{})
	fmt.Println("fethcing echostar3 id")
	for _, ppv := range ppvList {

		if len(ppv.Fields72[51]) > 1 {
			_, ok := tmsIdMap[ppv.Fields72[51]]
			if !ok { // make sure processing for unique tmsId
				threadCnt++

				wg.Add(1)
				go func(tmsId, cosmosHost, cosmosEventDetailApi, token string) {
					defer wg.Done()
					fmt.Print(".")

					tmsIdEchostarEventIdMap := make(map[string]string)
					echostar3Id, err := eventdetail.GetEchostar3Id(tmsId, cosmosHost, cosmosEventDetailApi, token)
					if err != nil {
						//fmt.Printf("\n---->error finding echostar3id for programId=%s. error=%s", ppv.Fields72[51], err)
						tmsIdEchostarEventIdMap[tmsId] = ""
					} else {
						//fmt.Println("echostar3Id-> ", echostar3Id)
						tmsIdEchostarEventIdMap[tmsId] = echostar3Id
					}

					tmsIdEchostarMapCh <- tmsIdEchostarEventIdMap

				}(ppv.Fields72[51], cosmosHost, cosmosEventDetailApi, token)

				if threadCnt == totalGoroutines {
					wg.Wait()
					fmt.Printf("\nprocessed %d records for echostar3id.\n", threadCnt)
					threadCnt = 0
				}
			}

		} // end if

	}

	wg.Wait()
	close(tmsIdEchostarMapCh)

	wg2.Wait()

	// fmt.Println("tmsIdEchostarEventIdLookup->:", tmsIdEchostarEventIdLookup)

	fmt.Println("Done with echostar3 id fetching")
	return tmsIdEchostarEventIdLookup
}

func replaceTmsIdWithEchostar3Id(ppvList []PPV, tmsIdEchostarEventIdLookup map[string]string) ([][]string, []string, [][]string) {
	var outputData [][]string
	var missingIds []string
	var outputUnknownLenFields [][]string

	for _, ppv := range ppvList {
		echostar3Id := tmsIdEchostarEventIdLookup[ppv.Fields72[51]]
		if len(echostar3Id) > 1 {
			ppv.Fields81Output[80] = echostar3Id
			var field81 []string

			for _, v := range ppv.Fields81Output {
				field81 = append(field81, fmt.Sprintf("%s", v))
			}
			outputData = append(outputData, field81)
		} else {
			missingIds = append(missingIds, ppv.Fields72[51])
		}

		if len(ppv.FieldsLenUnknown) > 1 {
			var fieldsLenUnknown []string
			for _, f := range ppv.FieldsLenUnknown {
				fieldsLenUnknown = append(fieldsLenUnknown, f)
			}
			outputUnknownLenFields = append(outputUnknownLenFields, fieldsLenUnknown)
		}

	}
	return outputData, missingIds, outputUnknownLenFields
}

func generateCSV(outputData [][]string, missingTmsIds []string, outputUnknownLenFields [][]string, csvFile, outputFailureFile string) error {
	if len(fp.DistinctStr(missingTmsIds)) > 0 {
		f3, err := os.OpenFile(outputFailureFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("can not create output failure file=%s. error=%s", outputFailureFile, err)
		}

		datawriter := bufio.NewWriter(f3)

		for _, data := range fp.DistinctStr(missingTmsIds) {
			_, err = datawriter.WriteString(data + "\n")
			if err != nil {
				return fmt.Errorf("error writing missingIds to file=%s", outputFailureFile)
			}
		}

		datawriter.Flush()
		f3.Close()

		return fmt.Errorf("output zip file will not be generated because of missing ids. see the file=%s", outputFailureFile)
	}

	fileName := csvFile
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("can not create output csv file=%s. error=%s", fileName, err)
	}

	csvWtr := altcsv.NewWriter(f)
	csvWtr.Quote = '"'      // use " as "quote"
	csvWtr.AllQuotes = true // surround each field with '"'

	csvWtr.WriteAll(outputData)
	f.Close()

	f2, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("error opening file for csv write. file=%s. error=%s", fileName, err)
	}

	csvWtr2 := altcsv.NewWriter(f2)
	csvWtr2.Quote = '"'      // use " as "quote"
	csvWtr2.AllQuotes = true // surround each field with '"'

	csvWtr2.WriteAll(outputUnknownLenFields)
	f2.Close()

	return nil
}
