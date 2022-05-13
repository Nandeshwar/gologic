package main

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	_ "github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/logic-building/functional-go/fp"

	"tba/pkg/auth"
	"tba/pkg/cosmosapi"
)

// Get it from mongodb
type Channel struct {
	SourceId   string  `bson:"source_id,omitempty"`
	CallLetter string  `bson:"call_letter,omitempty"`
	ImsAliases []Alias `bson:"ims_aliases,omitempty"`
}

// Get it from mongodb for failed sources
type FailedChannel struct {
	RequestBody FailedRequestBody     `bson:"request_body,omitempty"`
	Failures    FailedSourcesFailures `bson:"failures,omitempty"`
}
type FailedRequestBody struct {
	SourceId          string `bson:"source_id,omitempty"`
	ChannelProviderId string `bson:"channel_provider_id,omitempty"`
}

type FailedSourcesFailures struct {
	ResponseBody FailedSourcesFailuresResponseBody `bson:"response_body,omitempty"`
}

type FailedSourcesFailuresResponseBody struct {
	ErrorStrList []interface{} `bson:"errors,omitempty"`
	//FailedSourcesError []FailedSourcesError `bson:"errors,omitempty"`
}

type FailedSourcesError struct {
	Overlaps []FailedSourcesOverlap `bson:"overlaps,omitempty"`
}

type FailedSourcesOverlap struct {
	EventName string `bson:"event_name,omitempty"`
	BeginTime string `bson:"begin_time,omitempty"`
	Duration  int    `bson:"duration,omitempty"`
}

type FailedSourceOverlapGap struct {
	SourceId        int64
	BeginTime       time.Time
	BeginTimeStrUTC string
	DurationMin     int
	EndTime         time.Time
	EndTimeStrUTC   string
}

// Get it from mongodb
type Alias struct {
	ImsSvcId   string `bson:"ims_svc_id,omitempty"`
	Mode       string `bson:"mode,omitempty"`
	Processing string `bson:"processing,omitempty"`
	LoadedTime string `bson:"loaded,omitempty"`
}

// get data from csv
type TBA_Data struct {
	ChannelProviderId string
	BeginTimeESTStr   string
	BeginTimeUTCStr   string
}

type SourceId_To_BeginTimeUTC_Cosmos struct {
	SourceId        int64
	BeginTimeUTCStr string
}

type tbaSimsEvent struct {
	SourceId     int64
	ProgramId    string
	BeginTime    string
	Name         string
	Episode      string
	ShortName    string
	EchoUniqueId string
	Description  string
}

type simsEvent struct {
	SourceId          string  `json:"source_id"`
	ChannelProviderId string  `json:"channel_provider_id"`
	Events            []event `json:"events"`
}

type event struct {
	BeginTime    string `json:"begin_time"`
	Name         string `json:"name"`
	Episode      string `json:"episode"`
	ShortName    string `json:"short_name"`
	EchoUniqueId string `json:"echo_unique_id"`
	Description  string `json:"description"`
}

type SLReqBody struct {
	Channelsources []SLSourceInfo `json:"channelSources"`
}

type SLSourceInfo struct {
	SourceId                     int64  `json:"sourceId"`
	ProviderId                   string `json:"providerId"`
	Mode                         string `json:"mode"`
	Processing                   string `json:"processing"`
	ReplaceEntireServiceSchedule bool   `json:"replace_entire_service_schedule"`
}

type SupairProviderScheduleCosmos struct {
	SupairId        int64
	ProgramId       string
	BeginTimeStrUTC string
	EndTimeStrUTC   string
	Duration        int64
}

var AliasSourceIdLookup = make(map[string]int64)

var sourceId_programIds = make(map[int64]map[string]struct{})

var simsFileLineCh chan string
var sourceId_programId_lookup_sims_chan chan map[int64]map[string]struct{}
var simsTBAProgramIdsCh = make(chan string, 50000)
var simsTBAEventCh = make(chan tbaSimsEvent, 50000)
var simsTBAPogramIds []string
var configInsertSimsTBA bool
var configParserStatus bool
var fullResetImsData bool
var simsApiFullResetURL string
var configRecordFetchCount int
var (
	gnUsername   string
	gnPassword   string
	gnApiKey     string
	gnProgramApi string
)

var smtpClient *mail.SMTPClient
var emailHost string
var emailPort int
var emailRecipients []string

var soupDistressHours time.Duration

var cosmosApiObj *cosmosapi.CosmosApi

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("error reading config file", err)
			return
		} else {
			fmt.Println("error occurred while reading config file", err)
			return
		}
	}

	emailHost = viper.GetString("email.host")
	emailPort = viper.GetInt("email.port")
	emailRecipients = viper.GetStringSlice("email.recipients")

	server := mail.NewSMTPClient()
	server.Host = emailHost
	server.Port = emailPort
	server.KeepAlive = true
	server.SendTimeout = time.Second * 10
	//server.Encryption = mail.EncryptionTLS

	smtpClient, err = server.Connect()
	if err != nil {
		mulSection := `
Host igor-prod
    HostName 10.70.114.52
    Port 12002

    LocalForward 2525 172.29.1.39:25:25
		`
		fmt.Println("error connecting email server=", err.Error())
		fmt.Println("make sure to ssh tunnel to mul. The line given below should be in your mul section of ~/.ssh/config")
		fmt.Println("LocalForward 2525 172.29.1.39:25:25")
		fmt.Println("Ex: ")
		fmt.Println(mulSection)
		os.Exit(1)
	}

	fmt.Println("Email server connection setup successfully")
	configRecordFetchCount = viper.GetInt("tba.recordFetchCount")
	configParserStatus = viper.GetBool("tba.parserStatus")
	configInsertSimsTBA = viper.GetBool("tba.insertSimsTBA")
	fullResetImsData = viper.GetBool("tba.fullResetImsData")

	simsApiFullResetURL = viper.GetString("sims.api.fullResetUrl")

	maxGoRoutines := viper.GetInt("tba.goRoutineCntSims")
	simsFileLineCh = make(chan string, maxGoRoutines)
	sourceId_programId_lookup_sims_chan = make(chan map[int64]map[string]struct{}, maxGoRoutines)

	mongoHost := viper.GetString("mongo.host")
	mongoDb := viper.GetString("mongo.db")
	mongoChannel := viper.GetString("mongo.channel")

	soupDistressHours = viper.GetDuration("mongo.distressHours")

	// default distressHours is 36
	if soupDistressHours == 0 {
		soupDistressHours = 36
	}

	client, ctx, cancel, err := connect("mongodb://" + mongoHost)
	if err != nil {
		panic(err)
	}

	defer closeMongo(client, ctx, cancel)

	ping(client, ctx)

	// col := client.Database("channels").Collection("gn_02_14_2022")
	col := client.Database(mongoDb).Collection(mongoChannel)
	failedCollection := client.Database(mongoDb).Collection("failed")

	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	totalDocuments, aliasSourceIdLookUp, distressedAliasProviderSourceId, err := prepareAliasSourceLookup(col, ctx)
	if err != nil {
		fmt.Printf("error preparing ims alias and source id lookup. error=%s", err.Error())
		return
	}
	//fmt.Println("aliasSourceIdLookUp=", aliasSourceIdLookUp, "\ntotalAliasSourceCnt=", len(aliasSourceIdLookUp), "\ntotalDocuments=", totalDocuments)
	fmt.Println("\ntotalAliasSourceCnt in mongodb=", len(aliasSourceIdLookUp), "\ntotalDocuments=", totalDocuments)

	mysqlUser := viper.GetString("mysql.user")
	mysqlPassword := viper.GetString("mysql.password")
	mysqlHost := viper.GetString("mysql.host")
	mysqlDb := viper.GetString("mysql.db")
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDb)

	mySqlDb, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		fmt.Printf("\n error connecting to mysql cosmos. error=%s \n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Connected to cosmos db successfully", mySqlDb)

	err = connectToSims(simsApiFullResetURL)
	if err != nil {
		fmt.Println("Can not connect to SIMS. error=", err)
		fmt.Println("Add line below in file .ssh/config under prod section")
		fmt.Println("LocalForward 9210 che-imssked01:9200")
		os.Exit(1)
	}

	gnUsername = viper.GetString("gn.username")
	gnPassword = viper.GetString("gn.password")
	gnApiKey = viper.GetString("gn.apiKey")
	gnProgramApi = viper.GetString("gn.programUrl")
	gnValidProgramId := viper.GetString("gn.validProgramIdForConnectionTest")

	baseUrl, err := url.Parse(gnProgramApi)
	if err != nil {
		fmt.Println("Malformed GN URL=", err.Error())
		os.Exit(1)
	}

	params := url.Values{}
	params.Add("limit", "1000")
	params.Add("tmsId", gnValidProgramId)
	params.Add("api_key", gnApiKey)

	baseUrl.RawQuery = params.Encode()
	gnUrl := baseUrl.String()

	var gnConnectionError string
	err = connectToGNAPI(gnUsername, gnPassword, gnUrl)
	if err != nil {
		fmt.Println("error connecting GN API=%s. error=%s", gnUrl, err.Error())
		if strings.Contains(err.Error(), "Client.Timeout exceeded while awaiting headers") {
			gnConnectionError = err.Error()
		} else {
			os.Exit(1)
		}
	} else {
		fmt.Println("Connected to GN API successfully")
	}

	if len(gnConnectionError) > 1 {
		fmt.Println("WARN: Gn connection taking too long, GN functionaly may not work")
		gnConnectionError = "except GN connection. error=" + gnConnectionError
	}

	cosmosHost := viper.GetString("cosmos.host")
	cosmosTokenApi := viper.GetString("cosmos.tokenApi")
	cosmosScheduleApi := viper.GetString("cosmos.scheduleApi")
	cosmosScheduleIngestSource := viper.GetString("cosmos.ingestSource")
	cosmosUser := viper.GetString("cosmos.user")
	cosmosPassword := viper.GetString("cosmos.password")

	accessToken, err := auth.GetToken(cosmosHost, cosmosTokenApi, cosmosUser, cosmosPassword)
	if err != nil {
		fmt.Println("Can not connect to cosmos schedule API. error=", err)
		os.Exit(1)
	}

	cosmosApiObj = cosmosapi.New(cosmosHost, cosmosScheduleApi, accessToken, cosmosScheduleIngestSource)

	fmt.Println("Cosmos schdule api connetion successfull")

	fmt.Println("#############Congratulations! All connection done successfully #############" + gnConnectionError)

	tbaAction := viper.GetString("tba.action")

	//soupStatus(mySqlDb, getAllSoupSourceIds(aliasSourceIdLookUp))

	if tbaAction == "fix_soup_overlap" {
		err = fixSoupOverlap(mySqlDb, failedCollection, ctx)
		if err != nil {
			fmt.Println("error handling fix_soup_overlap in failed sources in soup. error=", err.Error())
		}
		return
	}

	if tbaAction == "soup_failed_full_reset" {
		err = soupFailedFullReset(failedCollection, ctx)
		if err != nil {
			fmt.Println("error handling full reset for failed sources in soup. error=", err.Error())
		}
		return
	}

	if tbaAction == "soup_distressed_full_reset" {
		soupDistressFullReset(distressedAliasProviderSourceId)
		return
	}

	if tbaAction == "tba_full_reset_file" {
		fullResetFile := viper.GetString("sims.fullResetFile")
		err := tbaFullResetFile(fullResetFile, simsApiFullResetURL, aliasSourceIdLookUp)
		if err != nil {
			fmt.Println("error full reset from file. error=", err)
			os.Exit(1)
		}
		return
	}
	if tbaAction == "tba_cosmos" {
		configInsertRecord := viper.GetBool("tba.insertRecord")
		configInsertRecordFromCache := viper.GetBool("tba.insertRecordFromCache")

		// _, allProgramIdsFor6000SourcesInSoup, err := getAllProgramIdsFor6000SourcesInSOUP(mySqlDb, aliasSourceIdLookUp, maxGoRoutines)
		// if err != nil {
		// 	fmt.Println("error preparing all programIds for 6000 sources in soup. error=%s", err.Error())
		// 	os.Exit(1)
		// }

		// fmt.Println("all program ids for all sources in soup=", allProgramIdsFor6000SourcesInSoup)
		// fmt.Println("all program ids for all sources in soup count=", len(allProgramIdsFor6000SourcesInSoup))

		allTbaProgramIds := getAllTBAProgramIds(mySqlDb)

		allSoupSourceIds := getAllSoupSourceIds(aliasSourceIdLookUp)
		eligibleTBAProgramIds, eligibleProgramIdsToSourceIds, err := findEligibleProgramIdsForSoupGnSources(mySqlDb, allTbaProgramIds, allSoupSourceIds)
		if err != nil {
			fmt.Printf("error finding eligibleTBAProgramIds. error=%s", err.Error())
			os.Exit(1)
		}

		allTbaProgramIds = eligibleTBAProgramIds

		fmt.Println("allTbaProgramIds=", allTbaProgramIds)
		fmt.Println("len(allTbaProgramIds)=", len(allTbaProgramIds))
		fmt.Println("\nBegin: all TBA programIds -> sourceIds in soup\n")
		for programId, sourceIds := range eligibleProgramIdsToSourceIds {
			fmt.Printf("\n%s -> %v", programId, sourceIds)
		}
		fmt.Println("\nEnd: all TBA programIds -> sourceIds in soup\n")

		if !configInsertRecordFromCache {
			writeToFile(allTbaProgramIds)
		}

		allTbaProgramIdsLookup := readFromFile()
		fmt.Println("allTbaProgramIdsLookup_fromFile=", allTbaProgramIdsLookup)
		fmt.Println("len(allTbaProgramIdsLookup_fromFile)=", len(allTbaProgramIdsLookup))

		if configInsertRecord {

			// allTbaProgramIdsLookup cotains remaining records
			allTbaProgramIdsLookup = insert_100_records_to_custom_table(mySqlDb, allTbaProgramIdsLookup)

			writeToFile(allTbaProgramIdsLookup)
		}

		programIdsTBAGN, programIdsNoDataGN, programIdsErrorGN := gnTbaReport(allTbaProgramIds)

		if configParserStatus {
			fmt.Printf("\n Parser status")
			programIdsNotInStagintCatalog, programIdsInStagingTable, err := findProgramIdsNotInStagingTable(mySqlDb, fp.DistinctStr(allTbaProgramIds))
			if err != nil {
				fmt.Printf("\ncosmos db error. error=%s", err.Error())
				os.Exit(1)
			}
			fmt.Printf("\nParsed Program ids count =%d. Parsed ProgramIds=%v", len(programIdsInStagingTable), programIdsInStagingTable)
			fmt.Printf("\nNon-Parsed Program ids count =%d. Non-Parsed ProgramIds=%v\n", len(programIdsNotInStagintCatalog), programIdsNotInStagintCatalog)
			if len(programIdsNotInStagintCatalog) > 0 {

				fmt.Printf("\nThese are the programIds not in staging tables=%v", programIdsNotInStagintCatalog)
				fmt.Println("\nWhat does this mean?")
				fmt.Println("Data parser did not parse the above program ids")
				fmt.Println("Either GN might not have data or you might have forgotten to run argo workflow: cwf-gracenote-programs-custom")
				fmt.Printf("\nRun the argo workflow: cwf-gracenote-programs-custom\n")
			} else {
				if len(fp.DistinctStr(allTbaProgramIds)) > 0 {
					fmt.Printf("\nAll these sims tba programs ids parsed correclty: %v\n", fp.DistinctStr(allTbaProgramIds))
				} else {
					fmt.Printf("Parser status: There is no TBA\n")
				}
			}
		}

		subject := "action: tba_cosmos"
		body := `
all TBA program_ids                 = <TBA_programIds>
all TBA program_ids_count           = <TBA_programIds_cnt>
all TBA program_ids_sourceIds       = <TBA_programIds_sourceIds>
all TBA program_ids_sourceIds_count = <TBA_programIds_sourceIds_cnt>

GN TBA program_ids                  = <GN_TBA_program_ids>
GN TBA program_ids_count            = <GN_TBA_program_ids_cnt>
GN has no data for program_ids      = <GN_NO_DATA_program_ids>
GN has no data for program_ids_count= <GN_NO_DATA_program_ids_cnt>
GN_error_fetch_program_ids          = <GN_ERROR_FETCH_PROGRAM_ID> 
GN_error_fetch_program_ids_COUNT    = <GN_ERROR_FETCH_PROGRAM_ID_cnt> 
		`

		body = strings.ReplaceAll(body, " <TBA_programIds>", strings.Join(allTbaProgramIds, ","))
		body = strings.ReplaceAll(body, " <TBA_programIds_cnt>", fmt.Sprintf("%d", len(allTbaProgramIds)))
		body = strings.ReplaceAll(body, " <TBA_programIds_sourceIds>", fmt.Sprintf("%v", eligibleProgramIdsToSourceIds))
		body = strings.ReplaceAll(body, " <TBA_programIds_sourceIds_cnt>", fmt.Sprintf("%d", len(eligibleProgramIdsToSourceIds)))

		body = strings.ReplaceAll(body, " <GN_TBA_program_ids>", strings.Join(programIdsTBAGN, ","))
		body = strings.ReplaceAll(body, " <GN_TBA_program_ids_cnt>", fmt.Sprintf("%d", len(programIdsTBAGN)))
		body = strings.ReplaceAll(body, " <GN_NO_DATA_program_ids>", strings.Join(programIdsNoDataGN, ","))
		body = strings.ReplaceAll(body, " <GN_NO_DATA_program_ids_cnt>", fmt.Sprintf("%d", len(programIdsNoDataGN)))
		body = strings.ReplaceAll(body, " <GN_ERROR_FETCH_PROGRAM_ID>", strings.Join(programIdsErrorGN, ","))
		body = strings.ReplaceAll(body, " <GN_ERROR_FETCH_PROGRAM_ID_cnt>", fmt.Sprintf("%d", len(programIdsErrorGN)))

		sendEmail(subject, body)

		return

	}
	if tbaAction == "tba_sims" {
		getSourceProgramIdLookUpFromSimsJson(mySqlDb)
		return
	}

	// ##################### TBA CSV logic ###################

	//csvFile := "tba2.csv"
	csvFile := viper.GetString("tba.csvFile")
	f, err := os.Open(csvFile)
	if err != nil {
		panic(fmt.Sprintf("error opening file=%s. error=%s", csvFile, err.Error()))
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("error reading csv file. error=%s", csvFile, err.Error())
	}

	tbaDataList, err := prepareChannelProviderIdAndBeginTime_UTC_lookup_TBA_CSV(data)
	if err != nil {
		fmt.Printf("error parsing TBA CSV. error=%s", err.Error())
		return
	}
	//fmt.Println("channelProviderIdBeginTimeUTC_TBA_lookup=", tbaDataList)
	fmt.Println("channelProviderIdBeginTimeUTC_TBA_lookup length=", len(tbaDataList))

	sourceIdToBeginTimeUTCCosmos := findSourceIdToBeginTimeForCosmos(aliasSourceIdLookUp, tbaDataList)
	//fmt.Println("sourceIdToBeginTimeUTCCosmos=", sourceIdToBeginTimeUTCCosmos)
	fmt.Println("sourceIdToBeginTimeUTCCosmos length=", len(sourceIdToBeginTimeUTCCosmos))

	//#######################  Cosmos MySql Connection ####################################

	tbaProgramIdList, err := getTBA_ProgramId(mySqlDb, sourceIdToBeginTimeUTCCosmos)
	if err != nil {
		fmt.Println("error fetching TBA program id from cosmos db. error=%s", err.Error())
		os.Exit(1)
	}

	distinctTBAProgramList := fp.DistinctStr(tbaProgramIdList)
	//fmt.Println("tbaProgramIdList=", tbaProgramIdList)
	programIdsTBAGN, programIdsNoDataGN, programIdsErrorGN := gnTbaReport(distinctTBAProgramList)

	fmt.Println("len(tbaProgramIdList)=", len(tbaProgramIdList))
	distinctTbaProgramIdList := fp.DistinctStr(tbaProgramIdList)
	fmt.Println("distinct tba program id list", distinctTbaProgramIdList)
	fmt.Println("distinct tba program id list len=", len(distinctTbaProgramIdList))

	programIds_noCatalogs, err := getProgramIdsNotInCatalog(mySqlDb, distinctTbaProgramIdList)
	if err != nil {
		fmt.Println("error runnings queries to find programIds with no catalog. exiting...", err.Error())
		os.Exit(1)
	}
	fmt.Println("\n programIds having no catalogs=", programIds_noCatalogs)
	fmt.Println("programIds having no len(catalogs)=", len(programIds_noCatalogs))

	programIds_noDescriptions, err := getProgramIdsNotInDescription(mySqlDb, distinctTbaProgramIdList)
	if err != nil {
		fmt.Println("error runnings queries to find programIds with no descriptions. exiting...", err.Error())
		os.Exit(1)
	}
	fmt.Println("\n programIds having no descriptions=", programIds_noDescriptions)
	fmt.Println("programIds having no len(programIds_noDescriptions)=", len(programIds_noDescriptions))

	var programIds_nocatelog_nodescriptions []string
	programIds_nocatelog_nodescriptions = append(programIds_nocatelog_nodescriptions, programIds_noCatalogs...)
	programIds_nocatelog_nodescriptions = append(programIds_nocatelog_nodescriptions, programIds_noDescriptions...)
	uniqeProgramIds_nocatelog_nodescriptions := fp.DistinctStr(programIds_nocatelog_nodescriptions)

	fmt.Println("uniqeProgramIds_nocatelog_nodescriptions=", uniqeProgramIds_nocatelog_nodescriptions)
	fmt.Println("len(uniqeProgramIds_nocatelog_nodescriptions)=", len(uniqeProgramIds_nocatelog_nodescriptions))

	configInsertCsvProgramIds := viper.GetBool("tba.insertCsvProgramIds")
	//insertIntoCustomDataCosmos(mySqlDb, distinctTbaProgramIdList)
	if configInsertCsvProgramIds {
		insertIntoCustomDataCosmos(mySqlDb, uniqeProgramIds_nocatelog_nodescriptions)
	}
	fmt.Println("----------SourceId -> ProgramIds")
	fmt.Println("sourceId_programIds", sourceId_programIds)
	fmt.Println("len(sourceId_programIds)", len(sourceId_programIds))

	subject := "action: tba_ims"
	body := `
all TBA program_ids                 = <TBA_programIds>
all TBA program_ids_count           = <TBA_programIds_cnt>
all TBA program_ids_sourceIds       = <TBA_programIds_sourceIds>
all TBA program_ids_sourceIds_count = <TBA_programIds_sourceIds_cnt>
Program_ids_no_catalog_no_descriptions = <TBA_NO_CATALOG_NO_DESCRIPTIONS>
Program_ids_no_catalog_no_descriptions_CNT = <TBA_NO_CATALOG_NO_DESCRIPTIONS_CNT>
<FULL_RESET_LIST>
<FULL_RESET_LIST_CNT>

GN TBA program_ids                  = <GN_TBA_program_ids>
GN TBA program_ids_count            = <GN_TBA_program_ids_cnt>
GN has no data for program_ids      = <GN_NO_DATA_program_ids>
GN has no data for program_ids_count= <GN_NO_DATA_program_ids_cnt>
GN_error_fetch_program_ids          = <GN_ERROR_FETCH_PROGRAM_ID> 
GN_error_fetch_program_ids_COUNT    = <GN_ERROR_FETCH_PROGRAM_ID_cnt> 
		`

	body = strings.ReplaceAll(body, " <TBA_programIds>", strings.Join(distinctTbaProgramIdList, ","))
	body = strings.ReplaceAll(body, " <TBA_programIds_cnt>", fmt.Sprintf("%d", len(distinctTbaProgramIdList)))
	body = strings.ReplaceAll(body, " <TBA_programIds_sourceIds>", fmt.Sprintf("%v", sourceId_programIds))
	body = strings.ReplaceAll(body, " <TBA_programIds_sourceIds_cnt>", fmt.Sprintf("%d", len(sourceId_programIds)))
	body = strings.ReplaceAll(body, " <TBA_NO_CATALOG_NO_DESCRIPTIONS>", strings.Join(uniqeProgramIds_nocatelog_nodescriptions, ","))
	body = strings.ReplaceAll(body, " <TBA_NO_CATALOG_NO_DESCRIPTIONS_CNT>", fmt.Sprintf("%d", len(uniqeProgramIds_nocatelog_nodescriptions)))

	body = strings.ReplaceAll(body, " <GN_TBA_program_ids>", strings.Join(programIdsTBAGN, ","))
	body = strings.ReplaceAll(body, " <GN_TBA_program_ids_cnt>", fmt.Sprintf("%d", len(programIdsTBAGN)))
	body = strings.ReplaceAll(body, " <GN_NO_DATA_program_ids>", strings.Join(programIdsNoDataGN, ","))
	body = strings.ReplaceAll(body, " <GN_NO_DATA_program_ids_cnt>", fmt.Sprintf("%d", len(programIdsNoDataGN)))
	body = strings.ReplaceAll(body, " <GN_ERROR_FETCH_PROGRAM_ID>", strings.Join(programIdsErrorGN, ","))
	body = strings.ReplaceAll(body, " <GN_ERROR_FETCH_PROGRAM_ID_cnt>", fmt.Sprintf("%d", len(programIdsErrorGN)))

	if fullResetImsData {
		aliasSourceIdMapTBA := make(map[string]int64)

		for _, tbaData := range tbaDataList {
			sourceId, ok := aliasSourceIdLookUp[tbaData.ChannelProviderId]
			if ok {
				aliasSourceIdMapTBA[tbaData.ChannelProviderId] = sourceId
			}
		}

		if len(aliasSourceIdMapTBA) < 1 {
			fmt.Println("No alias and sourceid to full reset. so no full reset")
		} else {
			fmt.Println("Full reset will be done these aliases(channel-providers=", aliasSourceIdMapTBA)
			fmt.Println("Full reset will be done these aliases(channel-providers) count=", len(aliasSourceIdMapTBA))
			err := fullResetSources(simsApiFullResetURL, aliasSourceIdMapTBA)
			if err != nil {
				fmt.Println("error in full reseting. error=%s", err.Error())
			} else {
				body = strings.ReplaceAll(body, "<FULL_RESET_LIST>", fmt.Sprintf("Full Reset list=%v", aliasSourceIdMapTBA))
				body = strings.ReplaceAll(body, "<FULL_RESET_LIST_CNT>", fmt.Sprintf("Full Reset list count=%d", len(aliasSourceIdMapTBA)))
				fmt.Println("----------->Full reset is scheduled to sims. Check sims log /var/log/sims/sims.log<---------------")
			}
		}
	}

	sendEmail(subject, body, csvFile)

	if configParserStatus {
		fmt.Printf("\n Parser status")
		programIdsNotInStagintCatalog, programIdsInStagingTable, err := findProgramIdsNotInStagingTable(mySqlDb, uniqeProgramIds_nocatelog_nodescriptions)

		if err != nil {
			fmt.Printf("\ncosmos db error. error=%s", err.Error())
			os.Exit(1)
		}
		fmt.Printf("\nParsed Program ids count =%d. Parsed ProgramIds=%v", len(programIdsInStagingTable), programIdsInStagingTable)
		fmt.Printf("\nNon-Parsed Program ids count =%d. Non-Parsed ProgramIds=%v\n", len(programIdsNotInStagintCatalog), programIdsNotInStagintCatalog)
		if len(programIdsNotInStagintCatalog) > 0 {

			fmt.Printf("\nThese are the programIds not in staging tables=%v", programIdsNotInStagintCatalog)
			fmt.Println("\nWhat does this mean?")
			fmt.Println("Data parser did not parse the above program ids")
			fmt.Println("Either GN might not have data or you might have forgotten to run argo workflow: cwf-gracenote-programs-custom")
			fmt.Printf("\nRun the argo workflow: cwf-gracenote-programs-custom\n")
		} else {
			if len(uniqeProgramIds_nocatelog_nodescriptions) > 0 {
				fmt.Printf("\nAll these sims tba programs ids parsed correclty: %v\n", fp.DistinctStr(uniqeProgramIds_nocatelog_nodescriptions))
			} else {
				fmt.Printf("Parser status: There is no TBA\n")
			}
		}
	}

}

func fullResetSources(url string, aliasSourceId map[string]int64) error {

	var slSourceInfoList []SLSourceInfo
	for alias, sourceId := range aliasSourceId {
		slSourceInfo := SLSourceInfo{
			SourceId:                     sourceId,
			ProviderId:                   alias,
			Mode:                         "full",
			Processing:                   "sl",
			ReplaceEntireServiceSchedule: true,
		}
		slSourceInfoList = append(slSourceInfoList, slSourceInfo)
	}
	sLReqBody := SLReqBody{
		Channelsources: slSourceInfoList,
	}

	sLReqBodyJson, err := json.Marshal(sLReqBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(sLReqBodyJson))

	if err != nil {
		return err
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("SL response=", resp.StatusCode)

	return nil
}

func readFromFile() []string {
	var allTbaProgramIds []string
	// open file
	f, err := os.Open("allTba.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		allTbaProgramIds = append(allTbaProgramIds, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return allTbaProgramIds
}

func getSourceProgramIdLookUpFromSimsJson(db *sql.DB) map[int64]map[string]struct{} {

	var simsTBAEvents []tbaSimsEvent

	f, err := os.Open(viper.GetString("tba.jsonFile"))
	if err != nil {
		log.Fatalf("\nerror reading sims json file: %. error=%s", viper.GetString("tba.jsonFile"), err.Error())
	}

	defer f.Close()

	r := bufio.NewReader(f)
	var fileErr error
	var line string

	var lineCounter int64

	go processSimsFile(db)
	go sendSimsTBAProgramIdsToDB(db)

	go func() {
		for simsTbaProgramId := range simsTBAProgramIdsCh {
			simsTBAPogramIds = append(simsTBAPogramIds, simsTbaProgramId)
		}
	}()

	go func() {
		for simsTBAevent := range simsTBAEventCh {
			simsTBAEvents = append(simsTBAEvents, simsTBAevent)
		}
	}()

	for fileErr == nil {
		lineCounter++
		line, fileErr = r.ReadString('\n')
		fmt.Printf("\n######################processing line=%d###############\n", lineCounter)
		if len(line) < 10 {
			close(simsFileLineCh)
			break
		}
		simsFileLineCh <- line
	}

	distinctSimsTBAProgramids := fp.DistinctStr(simsTBAPogramIds)

	programIdsTBAGN, programIdsNoDataGN, programIdsErrorGN := gnTbaReport(distinctSimsTBAProgramids)

	subject := "action: tba_sims"
	body := `
all TBA program_ids                 = <TBA_programIds>
all TBA program_ids_count           = <TBA_programIds_cnt>
all TBA program_ids_sourceIds       = <TBA_programIds_sourceIds>
all TBA program_ids_sourceIds_count = <TBA_programIds_sourceIds_cnt>

GN TBA program_ids                  = <GN_TBA_program_ids>
GN TBA program_ids_count            = <GN_TBA_program_ids_cnt>
GN has no data for program_ids      = <GN_NO_DATA_program_ids>
GN has no data for program_ids_count= <GN_NO_DATA_program_ids_cnt>
GN_error_fetch_program_ids          = <GN_ERROR_FETCH_PROGRAM_ID> 
GN_error_fetch_program_ids_COUNT    = <GN_ERROR_FETCH_PROGRAM_ID_cnt> 
		`

	body = strings.ReplaceAll(body, " <TBA_programIds>", strings.Join(distinctSimsTBAProgramids, ","))
	body = strings.ReplaceAll(body, " <TBA_programIds_cnt>", fmt.Sprintf("%d", len(distinctSimsTBAProgramids)))

	var tbaEventInfo string
	for _, simsTbaEvent := range simsTBAEvents {
		tbaEventInfo += fmt.Sprintf("\nsourceid=%d \nprogramIds=%s \neventName=%s \nshortName=%s \ndescription=%s \nepisode=%s \nechoUniqueId=%s \nbeginTime=%s", simsTbaEvent.SourceId, simsTbaEvent.ProgramId, simsTbaEvent.Name, simsTbaEvent.ShortName, simsTbaEvent.Description, simsTbaEvent.Episode, simsTbaEvent.EchoUniqueId, simsTbaEvent.BeginTime)
	}

	body = strings.ReplaceAll(body, " <TBA_programIds_sourceIds>", fmt.Sprintf("%v", tbaEventInfo))
	body = strings.ReplaceAll(body, " <TBA_programIds_sourceIds_cnt>", fmt.Sprintf("%d", len(simsTBAEvents)))

	body = strings.ReplaceAll(body, " <GN_TBA_program_ids>", strings.Join(programIdsTBAGN, ","))
	body = strings.ReplaceAll(body, " <GN_TBA_program_ids_cnt>", fmt.Sprintf("%d", len(programIdsTBAGN)))
	body = strings.ReplaceAll(body, " <GN_NO_DATA_program_ids>", strings.Join(programIdsNoDataGN, ","))
	body = strings.ReplaceAll(body, " <GN_NO_DATA_program_ids_cnt>", fmt.Sprintf("%d", len(programIdsNoDataGN)))
	body = strings.ReplaceAll(body, " <GN_ERROR_FETCH_PROGRAM_ID>", strings.Join(programIdsErrorGN, ","))
	body = strings.ReplaceAll(body, " <GN_ERROR_FETCH_PROGRAM_ID_cnt>", fmt.Sprintf("%d", len(programIdsErrorGN)))

	sendEmail(subject, body)

	fmt.Println("All sims TBA programIds=", distinctSimsTBAProgramids)

	if configParserStatus {
		fmt.Println("Parser status:")

		programIdsNotInStagintCatalog, programIdsInStagingTable, err := findProgramIdsNotInStagingTable(db, fp.DistinctStr(simsTBAPogramIds))
		if err != nil {
			fmt.Printf("\ncosmos db error. error=%s", err.Error())
			os.Exit(1)
		}
		fmt.Printf("\nParsed Program ids count =%d. Parsed ProgramIds=%v", len(programIdsInStagingTable), programIdsInStagingTable)
		fmt.Printf("\nNon-Parsed Program ids count =%d. Non-Parsed ProgramIds=%v\n", len(programIdsNotInStagintCatalog), programIdsNotInStagintCatalog)
		if len(programIdsNotInStagintCatalog) > 0 {

			fmt.Printf("\nThese are the programIds not in staging tables=%v", programIdsNotInStagintCatalog)
			fmt.Printf("\nRun the argo workflow: cwf-gracenote-programs-custom and  once it's done then \n change in config.yaml - insertSimsTBA:false \nthen run tba program to verify if cwf-gracenote-programs-custom parsed your xml for the above program ids correctly")
		} else {
			if len(fp.DistinctStr(simsTBAPogramIds)) > 0 {

				fmt.Printf("\nAll these sims tba programs ids parsed correclty: %v\n", fp.DistinctStr(simsTBAPogramIds))
			} else {
				fmt.Println("There is no TBA record\n")
			}
		}
	}

	return nil
}

func processSimsFile(db *sql.DB) {
	var wg sync.WaitGroup
	for i := 0; i < viper.GetInt("tba.goRoutineCntSims"); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range simsFileLineCh {
				sourceId_programIds := make(map[int64]map[string]struct{})

				sEvent := &simsEvent{}
				err := json.Unmarshal([]byte(line), sEvent)
				if err != nil {
					fmt.Printf("Unable to unmarshal sims json file=%s, error=%s", viper.GetString("tba.jsonFile"), err.Error())
					return
				}

				var programId string
				for _, event := range sEvent.Events {

					if strings.Contains(event.Name, "TBA") || strings.Contains(event.Description, "TBA") || strings.Contains(event.Episode, "TBA") || strings.Contains(event.ShortName, "TBA") {
						event.BeginTime = strings.Replace(event.BeginTime, "T", " ", -1)
						event.BeginTime = strings.Replace(event.BeginTime, "Z", "", -1)

						fmt.Println("Source_id=", sEvent.SourceId)
						fmt.Println("Begin: TBA Data in SIMS_____________________________>")
						fmt.Println("event.Name=", event.Name)
						fmt.Println("event.Description: ", event.Description)
						fmt.Println("event.Episode: ", event.Episode)
						fmt.Println("event.ShortName: ", event.ShortName)
						fmt.Println("event.BeginTime: ", event.BeginTime)
						fmt.Println("event.EchoUniqueId: ", event.EchoUniqueId)
						fmt.Println("End: TBA Data in SIMS_____________________________>")

						tbaEvent := tbaSimsEvent{}
						tbaEvent.Name = event.Name
						tbaEvent.Description = event.Description
						tbaEvent.Episode = event.Episode
						tbaEvent.ShortName = event.ShortName
						tbaEvent.BeginTime = event.BeginTime
						tbaEvent.EchoUniqueId = event.EchoUniqueId

						qry := "SELECT skd_program_id FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_source_mapping_id = ? and skd_begin_time=?"
						stmt, err := db.Prepare(qry)
						if err != nil {
							fmt.Printf("error preparing mysql query = %s. error=%s", qry, err.Error())
						}

						rows, err := stmt.Query(sEvent.SourceId, event.BeginTime)
						if err != nil {
							fmt.Printf("error executing query=%s. error=%s", qry, err.Error())
							os.Exit(1)
						}

						for rows.Next() {
							err := rows.Scan(&programId)
							if err != nil {
								fmt.Println("error scanning program id from cosmos mysql db. error=%s", err.Error())
								os.Exit(1)
							}

							sourceId, err := strconv.ParseInt(sEvent.SourceId, 10, 64)
							if err != nil {
								fmt.Errorf("error string to in64 conversion while preparing sims source id and program id lookup table. error=%s", err.Error())
								os.Exit(1)
							}

							programIds, ok := sourceId_programIds[sourceId]
							if ok {
								programIds[programId] = struct{}{}
								sourceId_programIds[sourceId] = programIds
							} else {
								programIds := make(map[string]struct{})
								programIds[programId] = struct{}{}
								sourceId_programIds[sourceId] = programIds
							}

							fmt.Println("sourceId_programIds: ", sourceId_programIds)

						} // db for loop
						sourceId, _ := strconv.ParseInt(sEvent.SourceId, 10, 64)
						tbaEvent.SourceId = sourceId
						for _, programIdsMap := range sourceId_programIds[sourceId] {
							tbaEvent.ProgramId = fmt.Sprintf("%v", programIdsMap)

						}
						simsTBAEventCh <- tbaEvent

					} // closing if event contains TBA
				} // closing for events

				//fmt.Println(sourceId_programIds)
				sourceId_programId_lookup_sims_chan <- sourceId_programIds

			} // chan range for loop

		}()
	} // 1000 goroutine for loop

	wg.Wait()
	close(sourceId_programId_lookup_sims_chan)

}

func sendSimsTBAProgramIdsToDB(db *sql.DB) {
	var wg sync.WaitGroup
	for i := 0; i < viper.GetInt("tba.goRoutineCntSims"); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for sourceId_programId_lookup_sims := range sourceId_programId_lookup_sims_chan {
				var tbaProgramIds []string
				for sourceId, program_ids := range sourceId_programId_lookup_sims {
					fmt.Printf("\nChecking program Id exitence in category: source_id=%s. program_id=%v", sourceId, program_ids)

					var programIds []string
					for programId, _ := range program_ids {
						programIds = append(programIds, programId)
					}
					programIdsNotInCatalog, err := getProgramIdsNotInCatalog(db, programIds)
					if err != nil {
						fmt.Printf("\nerror. sims program id not in catalog. err=%s", err.Error())
						os.Exit(1)
					}
					programIdsNotInDescription, err := getProgramIdsNotInDescription(db, programIds)
					if err != nil {
						fmt.Printf("\nerror. sims program id not in description. err=%s", err.Error())
						os.Exit(1)
					}
					if len(programIdsNotInCatalog) > 0 {
						tbaProgramIds = append(tbaProgramIds, programIdsNotInCatalog...)
					}
					if len(programIdsNotInDescription) > 0 {
						tbaProgramIds = append(tbaProgramIds, programIdsNotInDescription...)
					}

					distinctTbaProgramids := fp.DistinctStr(tbaProgramIds)
					if len(distinctTbaProgramids) > 0 {
						for _, tbaProgramId := range distinctTbaProgramids {
							simsTBAProgramIdsCh <- tbaProgramId
						}

						if configInsertSimsTBA {
							fmt.Printf("\n******************Begin: Sims Inserting program ids in custome table. source=%d, programId=%v", sourceId, tbaProgramIds)
							insertIntoCustomDataCosmos(db, distinctTbaProgramids)
							fmt.Printf("\n******************End: Sims Inserting program ids in custome table. source=%d, programId=%v", sourceId, tbaProgramIds)
						}

					}
				}
			} // for channel range
		}()
	} // for end 1000 goroutines
	wg.Wait()
	close(simsTBAProgramIdsCh)
}

func writeToFile(tbaProgramIds []string) {
	filePath := "allTba.txt"
	os.Remove(filePath)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range tbaProgramIds {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}

func insertIntoCustomDataCosmos(db *sql.DB, programIdList []string) {
	cnt := 0
	fmt.Printf("-------> Inserting TBA program ids to cosmos mysql db's custom table....")
	qry := "insert into supair_staging.load_stg_api_customdata (ingest_data_id, ingest_category, source_id) values(84, 'gnapi-custom-metadata', ?)"
	stmt, err := db.Prepare(qry)
	if err != nil {
		fmt.Printf("\n error preparing statement while inserting record to cosmos custom db. error=%s", err.Error())
	}

	for _, programId := range programIdList {
		cnt++
		fmt.Printf("\ninserting %d out of %d. program id=%s", cnt, len(programIdList), programId)
		_, err := stmt.Exec(programId)
		if err != nil {
			fmt.Printf("\nerror inserting program id=%s. error=%s", programId, err.Error())
		}
	}

}

func insert_100_records_to_custom_table(db *sql.DB, programIdList []string) []string {
	cnt := 0
	fmt.Printf("-------> Inserting 100 TBA program ids to cosmos mysql db's custom table....")
	qry := "insert into supair_staging.load_stg_api_customdata (ingest_data_id, ingest_category, source_id) values(84, 'gnapi-custom-metadata', ?)"
	stmt, err := db.Prepare(qry)
	if err != nil {
		fmt.Printf("\n error preparing statement while inserting record to cosmos custom db. error=%s", err.Error())
	}

	var insertedProgramIds []string
	for _, programId := range programIdList {
		cnt++
		if cnt > configRecordFetchCount {
			break
		}
		fmt.Printf("\ninserting %d out of %d. program id=%s", cnt, len(programIdList), programId)
		_, err := stmt.Exec(programId)
		if err != nil {
			fmt.Printf("\nerror inserting program id=%s. error=%s exiting", programId, err.Error())
			os.Exit(1)
		}
		insertedProgramIds = append(insertedProgramIds, programId)
	}
	fmt.Printf("\nInserted %d record successfully", len(insertedProgramIds))

	return fp.DifferenceStr(programIdList, insertedProgramIds)

}

func getAllTBAProgramIds(db *sql.DB) []string {
	fmt.Println("--------> Fetching all TBA program ids exist in supair_relationships and not exist in catalog from cosmos db ........")
	var tbaProgramIds []string
	qry := "SELECT source_id FROM supair.supair_relationships WHERE supair_id NOT IN (SELECT supair_id FROM supair.supair_catalog)  and ingest_source = 'gn'"
	stmt, err := db.Prepare(qry)
	if err != nil {
		fmt.Printf("error preparing mysql query = %s. error=%s", qry, err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("error executing query=%s. error=%s", qry, err.Error())
		os.Exit(1)
	}

	var programId string
	for rows.Next() {
		err := rows.Scan(&programId)
		if err != nil {
			fmt.Println("error scanning program id from cosmos mysql db. error=%s", err.Error())
			os.Exit(1)
		}
		tbaProgramIds = append(tbaProgramIds, programId)

	}
	return fp.DistinctStr(tbaProgramIds)
}

func getTBA_ProgramId(db *sql.DB, sourceIdToBeginTimeUTCCosmosList []SourceId_To_BeginTimeUTC_Cosmos) ([]string, error) {
	fmt.Println("--------> Fetching TBA program ids from cosmos mysql db.......")
	var tbaProgramIdList []string
	qry := "SELECT skd_program_id FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_source_mapping_id = ? and skd_begin_time=?"
	stmt, err := db.Prepare(qry)
	if err != nil {
		return nil, fmt.Errorf("error preparing mysql query = %s. error=%s", qry, err.Error())
	}

	cnt := 0
	for _, sourceId_BeginTime := range sourceIdToBeginTimeUTCCosmosList {
		cnt++
		fmt.Printf("\ncollecting TBA program ids from cosmos db. counter=%d out of %d", cnt, len(sourceIdToBeginTimeUTCCosmosList))
		rows, err := stmt.Query(sourceId_BeginTime.SourceId, sourceId_BeginTime.BeginTimeUTCStr)
		if err != nil {
			fmt.Printf("error executing query=%s. error=%s", qry, err.Error())
		}
		var programId string
		for rows.Next() {
			err := rows.Scan(&programId)
			if err != nil {
				fmt.Println("error scanning program id from cosmos mysql db. error=%s", err.Error())
			}
			tbaProgramIdList = append(tbaProgramIdList, programId)

			programIds, ok := sourceId_programIds[sourceId_BeginTime.SourceId]
			if ok {
				programIds[programId] = struct{}{}
				sourceId_programIds[sourceId_BeginTime.SourceId] = programIds
			} else {
				programIdsNew := make(map[string]struct{})
				programIdsNew[programId] = struct{}{}
				sourceId_programIds[sourceId_BeginTime.SourceId] = programIdsNew
			}

		}
	}
	return tbaProgramIdList, nil

}

func getProgramIdsNotInCatalog(db *sql.DB, programIds []string) ([]string, error) {
	fmt.Println("--------> collecting program ids which do not have catalog entry but entry in relationships....")
	var programIds_noCatelog []string
	qry := "SELECT source_id FROM supair.supair_relationships rs WHERE  supair_id=(SELECT supair_id FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_program_id = ? limit 1) and ingest_source='gn' and  NOT exists (SELECT 1 FROM supair.supair_catalog where supair_id=(SELECT supair_id FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_program_id = ? limit 1))  and rs.ingest_source = 'gn';"
	stmt, err := db.Prepare(qry)
	if err != nil {
		return nil, fmt.Errorf("error preparing mysql query = %s. error=%s", qry, err.Error())
	}

	cnt := 0
	for _, programId := range programIds {
		cnt++
		fmt.Printf("\nfinding program id exist in relationships and not in catalog. counter=%d out of %d", cnt, len(programIds))
		rows, err := stmt.Query(programId, programId)
		if err != nil {
			fmt.Printf("error executing query=%s. error=%s", qry, err.Error())
		}
		var programId string
		for rows.Next() {
			err := rows.Scan(&programId)
			if err != nil {
				fmt.Println("error scanning program id from cosmos mysql db. error=%s", err.Error())
			}
			programIds_noCatelog = append(programIds_noCatelog, programId)
		}
	}
	return programIds_noCatelog, nil

}

func getProgramIdsNotInDescription(db *sql.DB, programIds []string) ([]string, error) {
	fmt.Println("--------> collecting program ids which do not have description entry but entry in relationships....")
	var programIds_noDescriptions []string
	qry := "SELECT source_id FROM supair.supair_relationships rs WHERE  supair_id=(SELECT supair_id FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_program_id = ? limit 1) and ingest_source='gn' and  NOT exists (SELECT 1 FROM supair.supair_descriptions where supair_id=(SELECT supair_id FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_program_id = ? limit 1))  and rs.ingest_source = 'gn';"
	stmt, err := db.Prepare(qry)
	if err != nil {
		return nil, fmt.Errorf("error preparing mysql query = %s. error=%s", qry, err.Error())
	}

	cnt := 0
	for _, programId := range programIds {
		cnt++
		fmt.Printf("\nfinding program id exist in relationships and not in descriptions. counter=%d out of %d", cnt, len(programIds))
		rows, err := stmt.Query(programId, programId)
		if err != nil {
			fmt.Printf("error executing query=%s. error=%s", qry, err.Error())
		}
		var programId string
		for rows.Next() {
			err := rows.Scan(&programId)
			if err != nil {
				fmt.Println("error scanning program id from cosmos mysql db. error=%s", err.Error())
			}
			programIds_noDescriptions = append(programIds_noDescriptions, programId)
		}
	}
	return programIds_noDescriptions, nil

}

func findSourceIdToBeginTimeForCosmos(aliasSourceIdLookUp map[string]int64, tbaDataList []TBA_Data) []SourceId_To_BeginTimeUTC_Cosmos {
	var sourceIdToBeginTimeCosmosUtcList []SourceId_To_BeginTimeUTC_Cosmos

	for _, tbaData := range tbaDataList {
		sourceId, ok := aliasSourceIdLookUp[tbaData.ChannelProviderId]
		if ok {
			sourceIdToBeginTimeCosmosUtc := SourceId_To_BeginTimeUTC_Cosmos{SourceId: sourceId, BeginTimeUTCStr: tbaData.BeginTimeUTCStr}
			sourceIdToBeginTimeCosmosUtcList = append(sourceIdToBeginTimeCosmosUtcList, sourceIdToBeginTimeCosmosUtc)
		}
	}
	return sourceIdToBeginTimeCosmosUtcList
}

// channelProviderId of TBA CSV is same as alias in SOUP's mongodb
// SELECT * FROM supair.supair_provider_schedules where skd_ingest_source = 'GN'
// and skd_source_mapping_id = 54513 and skd_begin_time='2022-03-13 01:00:00';

// TODO
// read TBA csv
// get channel providerid and utc time lookup
// other function can find program id and insert to table
// insert into supair_staging.load_stg_api_customdata (ingest_data_id, ingest_category, source_id)
// values(84, 'gnapi-custom-metadata', 'SH035406860000');
func prepareChannelProviderIdAndBeginTime_UTC_lookup_TBA_CSV(data [][]string) ([]TBA_Data, error) {
	fmt.Println("-------------> reading TBA csv and preparing channelProviderId and BeginTime UTC lookup")
	var TBADataList []TBA_Data

	for i, line := range data {
		var channelProviderId string
		var beginTimeEST string

		if i > 0 { // omit header line
			for j, field := range line {
				switch j {
				case 0:
					channelProviderId = field
				case 2:
					beginTimeEST = field
				}
			}

			EST, err := time.LoadLocation("America/New_York")
			if err != nil {
				return nil, fmt.Errorf("error parsing EST to UTC. error=%s", err.Error())
			}

			const longForm = "1/2/06 15:04 EST"

			// ZANCH1:3/25/22 14:00 ZANF1H:3/26/22 12:00 ZANNH:3/24/22 23:35
			s, err1 := time.ParseInLocation(longForm, beginTimeEST+" EST", EST)
			if err1 != nil {
				longForm2 := "1/2/06 EST"
				s, err1 = time.ParseInLocation(longForm2, beginTimeEST+" EST", EST)
				if err1 != nil {
					return nil, fmt.Errorf("error parsing date=%s. error=%s", beginTimeEST, err1.Error())
				}

			}
			//fmt.Printf("EST to UTC: %v\n\n", s.UTC())
			//fmt.Printf("EST to UTC: %v\n\n", s.UTC().Format("2006-01-02 15:04:05"))
			tbaData := TBA_Data{
				ChannelProviderId: channelProviderId,
				BeginTimeESTStr:   beginTimeEST,
				BeginTimeUTCStr:   s.UTC().Format("2006-01-02 15:04:05"),
			}
			TBADataList = append(TBADataList, tbaData)
		}
	}

	return TBADataList, nil
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func closeMongo(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

func prepareAliasSourceLookup(col *mongo.Collection, ctx context.Context) (int, map[string]int64, map[string]int64, error) {
	fmt.Println("--------> Preparing ims alias and source id lookup table from Mongodb")
	cursor, err := col.Find(context.TODO(), bson.D{})

	totalDocuments := 0
	loadedTimeFormat := "2006-01-02T15:04:05.000000Z"
	distressedAliasProviderSourceId := make(map[string]int64)
	utcTimeNow := time.Now().UTC()

	if err != nil {
		defer cursor.Close(ctx)
		return -1, nil, nil, fmt.Errorf("error finding all documents for alias source lookup. error=%s", err.Error())

	} else {

		for cursor.Next(ctx) {
			totalDocuments++

			// declare a result BSON object
			//var result bson.M
			var result Channel
			err := cursor.Decode(&result)

			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				return -1, nil, nil, fmt.Errorf("cursor.Next() error while preparing ims alias source id look up table. error=%s", err.Error())

			} else {
				sourceId, err := strconv.ParseInt(result.SourceId, 10, 64)
				if err != nil {
					return -1, nil, nil, fmt.Errorf("error string to in64 conversion while preparing ims alias source id lookup table. error=%s", err.Error())
				}

				for _, alias := range result.ImsAliases {
					if alias.Processing == "sl" {
						AliasSourceIdLookup[alias.ImsSvcId] = sourceId
						loadedTimeStr := alias.LoadedTime
						loadedTime, err := time.Parse(loadedTimeFormat, loadedTimeStr)
						if err != nil {
							fmt.Printf("\nerror converting loaded time of soup=%s, error=%s", loadedTimeStr, err.Error())
							continue
						}

						cutoffDateTime := utcTimeNow.Add(-time.Hour * soupDistressHours)

						if loadedTime.Before(cutoffDateTime) {
							distressedAliasProviderSourceId[alias.ImsSvcId] = sourceId
						}
					}
				}
			}
		}
	}

	return totalDocuments, AliasSourceIdLookup, distressedAliasProviderSourceId, nil
}

func ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected to mongodb successfully")
	return nil
}

func findProgramIdsNotInStagingTable(db *sql.DB, programIds []string) ([]string, []string, error) {
	if len(programIds) == 0 {
		return []string{}, []string{}, nil
	}

	fmt.Println("--------> checking if programIds are in staging tables or not---->")
	var programIdsNotInStagingTable []string
	var programIdsInStagingTable []string
	qry := "SELECT source_id FROM supair_staging.load_stg_supair_catalog where source_id=?"
	stmt, err := db.Prepare(qry)
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing mysql query = %s. error=%s", qry, err.Error())
	}

	cnt := 0
	for _, programId := range programIds {
		cnt++
		fmt.Printf("\nfinding program id exist in supair_staging.load_stg_supair_catalog table . counter=%d out of %d", cnt, len(programIds))
		rows, err := stmt.Query(programId)
		if err != nil {
			fmt.Printf("findProgramIdsInStagingTable: error executing query=%s. error=%s", qry, err.Error())
			return nil, nil, err
		}
		var programIdFromDb string
		rowsCnt := 0
		for rows.Next() {
			rowsCnt++
			err := rows.Scan(&programIdFromDb)
			if err != nil {
				fmt.Println("findProgramIdsInStagingTable: error scanning program id from cosmos mysql db. error=%s", err.Error())
				return nil, nil, err
			}
			fmt.Printf("\n programId=%s exist in table - supair_staging.load_stg_supair_catalog", programIdFromDb)
			programIdsInStagingTable = append(programIdsInStagingTable, programIdFromDb)

		}
		if rowsCnt == 0 {
			programIdsNotInStagingTable = append(programIdsNotInStagingTable, programId)
		}

	}
	return fp.DistinctStr(programIdsNotInStagingTable), fp.DistinctStr(programIdsInStagingTable), nil

}

func connectToSims(url string) error {
	url = strings.Replace(url, "channelsFullReset", "status", -1)
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println("connected to sims")
	fmt.Println(string(responseData))

	return nil
}

func tbaFullResetFile(csvFile string, url string, aliasSourceId map[string]int64) error {
	if _, err := os.Stat(csvFile); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file for full reset does not exist. %s", err.Error())
	}

	f, err := os.Open(csvFile)
	if err != nil {
		return fmt.Errorf("error opening file=%s. error=%s", csvFile, err.Error())
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading csv file. error=%s", csvFile, err.Error())
	}

	channelProviderSourceIdFromCSV := make(map[string]int64)

	if len(data) < 1 {
		return fmt.Errorf("%s should not be empty", csvFile)
	}

	for i, line := range data {

		if len(line) != 2 {
			return fmt.Errorf("There must be two column at row=%d", i+1)
		}

		if i > 0 {
			var channel string
			var providerSourceId int64
			for j, field := range line {

				switch j {
				case 0:
					channel = field
				case 1:
					n, err := strconv.ParseInt(field, 10, 64)
					if err != nil {
						return fmt.Errorf("error at row=%d of file=%s. 2nd column must be integer. error=%s", j+1, csvFile, err.Error())
					}
					providerSourceId = n
				}
			}
			channelProviderSourceIdFromCSV[channel] = providerSourceId
		} else if i == 0 {

			for j, field := range line {
				switch j {
				case 0:
					if field != "channel" {
						return fmt.Errorf("1st column of header must be: channel ")
					}
				case 1:
					if field != "provider_source_id" {
						return fmt.Errorf("2nd column of header must be: provider_source_id")
					}
				}
			}
		}
	}

	fmt.Println("full reset data from file", channelProviderSourceIdFromCSV)

	aliasSourceIdMapTBA := make(map[string]int64)
	for channel, providerSourceId := range channelProviderSourceIdFromCSV {
		providerSourceIdInMongo, ok := aliasSourceId[channel]
		if ok {
			if providerSourceIdInMongo == providerSourceId {
				aliasSourceIdMapTBA[channel] = providerSourceId
			} else {
				fmt.Printf("\n sourceId=%d for channel=%s in CSV does not match with mongodb for the same channel. No full reset for this.", providerSourceId, channel)
			}
		} else {
			fmt.Printf("\nchannel=%s in csv does not exist in mongodb. There will not be full reset for channel=%s, provider_source_id=%d", channel, channel, providerSourceId)
		}
	}

	if len(aliasSourceIdMapTBA) < 1 {
		return fmt.Errorf("No record for csv-full-reset in the file=%s", csvFile)
	}
	err = fullResetSources(simsApiFullResetURL, aliasSourceIdMapTBA)
	if err != nil {
		fmt.Println("error in csv-full-reset. error=%s", err.Error())
	} else {
		fmt.Println("----------->Full reset is scheduled to sims. Check sims log /var/log/sims/sims.log<---------------")
	}

	subject := "action: tba_full_reset_file"
	body := `
List of sources that is sent to SIMS for full reset=<RESET_SOURCES>
List of sources that is sent to SIMS for full reset_count=<RESET_SOURCES_COUNT>		
	`

	body = strings.ReplaceAll(body, "<RESET_SOURCES>", fmt.Sprintf("%v", aliasSourceIdMapTBA))
	body = strings.ReplaceAll(body, "<RESET_SOURCES_COUNT>", fmt.Sprintf("%v", len(aliasSourceIdMapTBA)))

	sendEmail(subject, body, csvFile)

	return nil
}

func getAllSoupSourceIds(aliasSourceId map[string]int64) []int64 {
	var sources []int64
	for _, sourceId := range aliasSourceId {
		sources = append(sources, sourceId)
	}

	uniqeSourceIds := fp.DistinctInt64(sources)
	return uniqeSourceIds
}

// Check if program id is related to soup source id
func findEligibleProgramIdsForSoupGnSources(db *sql.DB, programIds []string, soupSourceIds []int64) (eligibleProgramIds []string, eligibleProgramidToSourceId map[string][]int64, err error) {
	eligibleProgramidToSourceId = make(map[string][]int64)

	var dbSourceId string
	fmt.Println("Finding eligible program Ids that is related to provider source id in soup")
	programIdsLen := len(programIds)
	for _, programId := range programIds {
		fmt.Printf("%d ", programIdsLen)
		programIdsLen--

		qry := "SELECT distinct(skd_source_mapping_id) FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_program_id = ?"
		stmt, err1 := db.Prepare(qry)
		if err1 != nil {
			err = fmt.Errorf("findEligibleProgramIdsForSoupGnSources: error preparing mysql query = %s. error=%s", qry, err1.Error())
			return
		}

		rows, err2 := stmt.Query(programId)
		if err2 != nil {
			err = fmt.Errorf("findEligibleProgramIdsForSoupGnSources: error executing query=%s. error=%s", qry, err2.Error())
			return
		}

		for rows.Next() {
			err3 := rows.Scan(&dbSourceId)
			if err3 != nil {
				err = fmt.Errorf("getAllProgramIdsFor6000SourcesInSOUP: error scanning program id from cosmos mysql db. error=%s", err3.Error())
				return
			}

			sourceId, err4 := strconv.ParseInt(dbSourceId, 10, 64)
			if err4 != nil {
				err = fmt.Errorf("error string to in64 conversion while preparing sims source id and program id lookup table. error=%s", err4.Error())
				return
			}

			if fp.ExistsInt64(sourceId, soupSourceIds) {
				eligibleProgramIds = append(eligibleProgramIds, programId)
				sourceIds, ok := eligibleProgramidToSourceId[programId]
				if ok {
					sourceIds = append(sourceIds, sourceId)
					eligibleProgramidToSourceId[programId] = fp.DistinctInt64(sourceIds)
				} else {
					eligibleProgramidToSourceId[programId] = []int64{sourceId}
				}
			}

		} // db for loop
	}
	return fp.DistinctStr(eligibleProgramIds), eligibleProgramidToSourceId, nil
}

func getAllProgramIdsFor6000SourcesInSOUP(db *sql.DB, aliasSourceId map[string]int64, maxGoroutines int) ([]int64, []string, error) {
	var sources []int64
	for _, sourceId := range aliasSourceId {
		sources = append(sources, sourceId)
	}

	uniqeSourceIds := fp.DistinctInt64(sources)
	//fmt.Println("unique source ids in Soup=", uniqeSourceIds)
	fmt.Println("All source ids count in Soup=", len(sources))
	fmt.Println("unique source ids count in Soup=", len(uniqeSourceIds))

	sourceIdsCh := make(chan int64, maxGoroutines)
	programIdsCh := make(chan string, maxGoroutines*maxGoroutines)

	var programIds []string
	var wg sync.WaitGroup

	// keep adding sourceIds in channel
	wg.Add(1)
	go func(sourceIds []int64) {
		defer wg.Done()

		for _, sourceId := range sourceIds {
			sourceIdsCh <- sourceId
		}
		close(sourceIdsCh)

	}(uniqeSourceIds)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for sourceId := range sourceIdsCh {
				var programId string

				qry := "SELECT distinct(skd_program_id) FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_source_mapping_id = ?"
				stmt, err := db.Prepare(qry)
				if err != nil {
					fmt.Printf("getAllProgramIdsFor6000SourcesInSOUP: error preparing mysql query = %s. error=%s", qry, err.Error())
					os.Exit(1)
				}

				rows, err := stmt.Query(sourceId)
				if err != nil {
					fmt.Printf("getAllProgramIdsFor6000SourcesInSOUP: error executing query=%s. error=%s", qry, err.Error())
					os.Exit(1)

				}

				for rows.Next() {
					err := rows.Scan(&programId)
					if err != nil {
						fmt.Printf("getAllProgramIdsFor6000SourcesInSOUP: error scanning program id from cosmos mysql db. error=%s", err.Error())
						os.Exit(1)
					}
					programIdsCh <- programId

				} // db for loop
			} // for loop end getting info from sourcesCh

		}()
	} // end of for loop creating 100 goroutines

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for programId := range programIdsCh {
				fmt.Println("Data transfering to global variables", programId)
				programIds = append(programIds, programId)
			}
		}()
	}

	wg.Wait()
	fmt.Println("I am after watch")
	close(programIdsCh)

	uniqueProgramIds := fp.DistinctStr(programIds)
	fmt.Println("all unique programIds uniqueProgramIds=", uniqueProgramIds)
	fmt.Println("all unique programIds len(uniqueProgramIds)=", len(uniqueProgramIds))

	return uniqeSourceIds, uniqueProgramIds, nil
}

func connectToGNAPI(gnUsername, gnPassword, gnUrl string) error {
	_, err := callGnAPI(gnUrl, "GET", gnUsername, gnPassword)
	return err
}

func callGnAPI(url, method, username, password string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("gn api=%s, Got error %s", err.Error(), url)
	}
	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("gn api=%s. Got error %s", url, err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("gn api=%s. Error while reading the response bytes:", url, err)
	}

	defer response.Body.Close()
	return string([]byte(body)), nil
}

func TbaInGN(programIds []string) (programIdsTBA []string, programIdsNoData []string, programIdsError []string, err error) {
	cnt := 0
	totalCnt := len(programIds)

	if err != nil {
		err = fmt.Errorf("TbaInGN: Malformed GN URL=", err.Error())
		return
	}

	for _, programId := range programIds {
		retryCounter := 0
		params := url.Values{}
		params.Add("limit", "1000")

		params.Add("api_key", gnApiKey)

		baseUrl, err := url.Parse(gnProgramApi)

		cnt++
		fmt.Printf("\nchecking progamid=%s has TBA in GN. %d out of %d", programId, cnt, totalCnt)
		params.Add("tmsId", programId)
		baseUrl.RawQuery = params.Encode()
		gnUrl := baseUrl.String()

		fmt.Println(gnUrl)

		bodyStr, err := callGnAPI(gnUrl, "GET", gnUsername, gnPassword)

		for !strings.Contains(bodyStr, fmt.Sprintf("<program TMSId=\"%s\"", programId)) {
			retryCounter++
			time.Sleep(time.Second * 1)
			bodyStr, err = callGnAPI(gnUrl, "GET", gnUsername, gnPassword)
			if retryCounter == 3 {
				break
			}

		}

		if err != nil {
			fmt.Printf("\nerror fetching record from GN. programId=%s, err=%s", programId, err.Error())
			programIdsError = append(programIdsError, programId)
		} else if strings.Contains(bodyStr, "TBA") {
			fmt.Printf("\nFound TBA in GN. programId=%s \n xml body=%s", programId, bodyStr)
			programIdsTBA = append(programIdsTBA, programId)
		} else if !strings.Contains(bodyStr, fmt.Sprintf("<program TMSId=\"%s\"", programId)) {
			fmt.Printf("\nGN- no data in GN for programId=%s", programId)

			programIdsNoData = append(programIdsNoData, programId)
		}

	}
	return
}

func gnTbaReport(programIds []string) (programIdsTBA, programIdsNoData, programIdsError []string) {
	if len(programIds) > 0 {
		var err error
		programIdsTBA, programIdsNoData, programIdsError, err = TbaInGN(programIds)
		if err != nil {
			fmt.Println("error getting TBA from GN. error=%s", err.Error())
		}

		fmt.Println("\n-------------Begin: GN Report-------------\n")
		fmt.Println("TBA in GN. programIds      =", programIdsTBA)
		fmt.Println("TBA in GN. programIds count=", len(programIdsTBA))

		fmt.Println("Program Ids do not have data in GN. cause of TBA. program-ids      =", programIdsNoData)
		fmt.Println("Program Ids do not have data in GN. cause of TBA. program-ids count=", len(programIdsNoData))

		fmt.Println("Program Ids have error while fetching data from GN. Cause of TBA. program-ids      =", programIdsError)
		fmt.Println("Program Ids have error while fetching data from GN. Cause of TBA. program-ids count=", len(programIdsError))
		fmt.Println("\n-------------End: GN Report-------------\n")
	}

	return
}

func sendEmail(subject string, body string, attachments ...string) {
	var emailSentList []string
	for _, recipient := range emailRecipients {

		email := mail.NewMSG()
		email.SetFrom("From Me <tba@noreply.com>")
		email.AddTo(recipient)

		email.SetSubject(subject)

		email.SetBody(mail.TextPlain, body)

		if len(attachments) > 0 {
			email.AddAttachment(attachments[0])
		}

		for cnt := 1; cnt <= 3; cnt++ {
			// Send email
			err := email.Send(smtpClient)
			if err != nil {
				fmt.Println("error sending mail sending failed to recipient=", recipient, "error=", err)
				fmt.Println("Retrying sending mail to recipient = %s", recipient, " again")
				if cnt == 3 {
					continue
				}
			}
			break
		}

		emailSentList = append(emailSentList, recipient)
	}

	fmt.Println("email sent successfully to ", emailSentList)

}

func soupDistressFullReset(distressedAliasProviderSourceId map[string]int64) {
	if len(distressedAliasProviderSourceId) == 0 {
		fmt.Println("There are no distress records in soup")
		return
	}
	fmt.Println("Full Reset aliases count=", len(distressedAliasProviderSourceId))
	fmt.Println("Full reset will be scheduled for these alias and sources", distressedAliasProviderSourceId)
	err := fullResetSources(simsApiFullResetURL, distressedAliasProviderSourceId)
	if err != nil {
		fmt.Println("error in csv-full-reset. error=%s", err.Error())
	} else {
		fmt.Println("----------->Full reset is scheduled to sims. Check sims log /var/log/sims/sims.log<---------------")
	}

	subject := "action: soup_distressed_full_reset"
	body := `
List of sources that is sent to SIMS for full reset=<RESET_SOURCES>	
List of sources that is sent to SIMS for full reset_count=<RESET_SOURCES_CNT>	
	`

	body = strings.ReplaceAll(body, "<RESET_SOURCES>", fmt.Sprintf("%v", distressedAliasProviderSourceId))
	body = strings.ReplaceAll(body, "<RESET_SOURCES_CNT>", fmt.Sprintf("%v", len(distressedAliasProviderSourceId)))

	sendEmail(subject, body)
}

func soupFailedFullReset(fcol *mongo.Collection, ctx2 context.Context) error {
	var aliasSourceFailed = make(map[string]int64)
	fmt.Println("--------> Preparing ims alias and source id lookup table from Mongodb for failed sources")

	var ctx context.Context
	var ctxCancelFunc context.CancelFunc
	var timeTillContextDeadline = time.Now().Add(600 * time.Second)

	ctx, ctxCancelFunc = context.WithDeadline(context.Background(), timeTillContextDeadline)
	cursor, err := fcol.Find(ctx, bson.D{})

	totalDocuments := 0

	defer ctxCancelFunc()

	if err != nil {
		defer cursor.Close(ctx)
		return fmt.Errorf("error finding all documents for failed sources in soup. error=%s", err.Error())

	} else {
		for cursor.Next(ctx) {
			totalDocuments++

			// declare a result BSON object
			//var result bson.M
			var result FailedChannel

			err := cursor.Decode(&result)

			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				return fmt.Errorf("cursor.Next() error while preparing ims alias source id look up table for soup failed sources. error=%s", err.Error())

			} else {
				sourceId, err := strconv.ParseInt(result.RequestBody.SourceId, 10, 64)
				if err != nil {
					return fmt.Errorf("error string to in64 conversion while preparing ims alias source id lookup table for failed sources. error=%s", err.Error())
				}

				aliasSourceFailed[result.RequestBody.ChannelProviderId] = sourceId

				// for _, alias := range result.ImsAliases {
				// 	if alias.Processing == "sl" {
				// 		aliasSourceFailed[alias.ImsSvcId] = sourceId

				// 	}
				// }
			}
		}
	}

	fmt.Println("aliasSourceFailed=", aliasSourceFailed)
	fmt.Println("aliasSourceFailed=", len(aliasSourceFailed))

	// var aliasSourceFailed1000 = make(map[string]int64)
	// cnt := 0
	// for alias, sourceId := range aliasSourceFailed {
	// 	cnt++
	// 	if cnt == 1000 {
	// 		break
	// 	}
	// 	aliasSourceFailed1000[alias] = sourceId
	// }

	// err = fullResetSources(simsApiFullResetURL, aliasSourceFailed1000)
	err = fullResetSources(simsApiFullResetURL, aliasSourceFailed)
	if err != nil {
		fmt.Println("error in csv-full-reset for failed sources. error=%s", err.Error())
	} else {
		fmt.Println("----------->Full reset for failed sources is scheduled to sims. Check sims log /var/log/sims/sims.log<---------------")
	}

	subject := "action: soup_failed_full_reset"
	body := `
List of sources that is sent to SIMS for full reset=<RESET_SOURCES>
List of sources that is sent to SIMS for full reset_count=<RESET_SOURCES_COUNT>		
	`

	body = strings.ReplaceAll(body, "<RESET_SOURCES>", fmt.Sprintf("%v", aliasSourceFailed))
	body = strings.ReplaceAll(body, "<RESET_SOURCES_COUNT>", fmt.Sprintf("%v", len(aliasSourceFailed)))

	sendEmail(subject, body)

	return nil
}

func soupStatus(db *sql.DB, soupEnabledSources []int64) {
	var sourceLessThan14DaysOfData []int64
	counter := 0
	totalSources := len(soupEnabledSources)
	for _, sourceId := range soupEnabledSources {
		counter++
		fmt.Println("checking for 14 days of record in cosmos db for source id = ", sourceId, ". ", counter, "out of ", totalSources)
		if checkForsourceLessThan14DaysOfData(db, sourceId) {
			sourceLessThan14DaysOfData = append(sourceLessThan14DaysOfData, sourceId)
		}
	}
	fmt.Println("sourcesLessThanOf14DaysOfData=", sourceLessThan14DaysOfData)
}

func checkForsourceLessThan14DaysOfData(db *sql.DB, sourceId int64) bool {
	now := time.Now().UTC()
	nowPlus14day := now.AddDate(0, 0, 14)
	dateFormat := "2006-01-02"
	nowPlus14dayStr := nowPlus14day.Format(dateFormat)

	var dbDataCnt int64
	qry := "SELECT count(skd_source_mapping_id)  providerSourceIdCount FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_source_mapping_id = ? and skd_date=?"
	stmt, err := db.Prepare(qry)
	if err != nil {
		fmt.Printf("checkForsourceLessThan14DaysOfData: error preparing mysql query = %s. error=%s", qry, err.Error())
		os.Exit(1)
	}

	rows, err := stmt.Query(sourceId, nowPlus14dayStr)
	if err != nil {
		fmt.Printf("checkForsourceLessThan14DaysOfData: error executing query=%s. error=%s", qry, err.Error())
		os.Exit(1)
	}

	if rows.Next() {
		err := rows.Scan(&dbDataCnt)
		if err != nil {
			fmt.Printf("checkForsourceLessThan14DaysOfData: scanning rows query=%s. error=%s", qry, err.Error())
			os.Exit(1)
		}
		if dbDataCnt > 0 {
			return true
		}
	}

	return false
}

func fixSoupOverlap(db *sql.DB, fcol *mongo.Collection, ctx2 context.Context) error {
	fmt.Println("--------> Preparing ims alias and source id lookup table from Mongodb for failed sources to fix soup overlap")

	var ctx context.Context
	var ctxCancelFunc context.CancelFunc
	var timeTillContextDeadline = time.Now().Add(600 * time.Second)

	ctx, ctxCancelFunc = context.WithDeadline(context.Background(), timeTillContextDeadline)
	cursor, err := fcol.Find(ctx, bson.D{})

	totalDocuments := 0

	defer ctxCancelFunc()

	var failedSourceOverlapGapList []FailedSourceOverlapGap

	if err != nil {
		defer cursor.Close(ctx)
		return fmt.Errorf("error finding all documents for failed sources in soup. error=%s", err.Error())

	} else {
		for cursor.Next(ctx) {
			totalDocuments++

			// declare a result BSON object
			//var result bson.M
			var result FailedChannel

			err := cursor.Decode(&result)

			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				return fmt.Errorf("cursor.Next() error while preparing ims alias source id look up table for soup failed sources for overlap issue. error=%s", err.Error())

			} else {

				sourceId, err := strconv.ParseInt(result.RequestBody.SourceId, 10, 64)
				if err != nil {
					return fmt.Errorf("error string to in64 conversion while preparing ims alias source id lookup table for failed sources for overlap issue. error=%s", err.Error())
				}

				fmt.Println(totalDocuments, "-------------------------------")
				fmt.Println("sourceId=", sourceId)
				//r, _ = regexp.Compile("([\{a-z_]+)=([a-z|0-9|_]+)")
				r, _ := regexp.Compile("(\\{begin_time.+?\\}\\])")

				if len(result.Failures.ResponseBody.ErrorStrList) == 2 && strings.Contains(fmt.Sprintf("%v", result.Failures.ResponseBody.ErrorStrList[0]), "validation failed, gaps/overlaps found") {
					gapOverLapStr := fmt.Sprintf("%v", result.Failures.ResponseBody.ErrorStrList[1])
					foundItems := r.FindAllString(gapOverLapStr, -1)
					for _, v := range foundItems {
						v = strings.ReplaceAll(v, "{", "")
						v = strings.ReplaceAll(v, "}", "")
						v = strings.ReplaceAll(v, "]", "")
						fmt.Println("FailedSourcesOverlapGap in soup=", v)

						// v: should be like this -> begin_time 2022-01-11T03:00:00Z duration 60
						infoList := strings.Split(v, " ")
						beginTimeStr := infoList[1]
						durationMinStr := infoList[3]
						durationInMin, err := strconv.Atoi(durationMinStr)
						if err != nil {
							return fmt.Errorf("fixSoupOverlap: error converting durationInStr=%s to int. error=%s", durationMinStr, err.Error())
						}

						//s.UTC().Format("2006-01-02 15:04:05"),
						beginTime, err := time.Parse("2006-01-02T15:04:05Z", beginTimeStr)
						if err != nil {
							return fmt.Errorf("fixSoupOverlap: error parsing beginTimeStr=%s. error=%s", beginTimeStr, err.Error())
						}
						beginTimeStrUTC := beginTime.Format("2006-01-02 15:04:05")
						fmt.Println("beginTimeStrUTC=", beginTimeStrUTC)

						endTime := beginTime.Add(time.Duration(durationInMin) * time.Minute)
						endTimeStrUTC := endTime.Format("2006-01-02 15:04:05")
						fmt.Println("End time str=", endTimeStrUTC)

						failedSourceOverlapGap := FailedSourceOverlapGap{SourceId: sourceId, BeginTime: beginTime, BeginTimeStrUTC: beginTimeStrUTC, EndTime: endTime, EndTimeStrUTC: endTimeStrUTC, DurationMin: durationInMin}
						failedSourceOverlapGapList = append(failedSourceOverlapGapList, failedSourceOverlapGap)
					}

					//fmt.Println(gapOverLapStr)
				}

			}
		}
	}

	var fsOverlapList []FailedSourceOverlapGap
	var fsGapList []FailedSourceOverlapGap
	for i := 0; i < len(failedSourceOverlapGapList); i += 2 {
		if failedSourceOverlapGapList[i+1].BeginTime.Before(failedSourceOverlapGapList[i].EndTime) {

			if failedSourceOverlapGapList[i+1].SourceId == 15569 {
				fmt.Println(failedSourceOverlapGapList[i+1].BeginTime)
				fmt.Println(failedSourceOverlapGapList[i].EndTime)
			}
			fsOverlapList = append(fsOverlapList, failedSourceOverlapGapList[i])
			fsOverlapList = append(fsOverlapList, failedSourceOverlapGapList[i+1])
		} else {
			fsGapList = append(fsGapList, failedSourceOverlapGapList[i])
			fsGapList = append(fsGapList, failedSourceOverlapGapList[i+1])
		}
	}

	fmt.Println("overlap report")
	fmt.Println("===================")

	cnt := 0
	for i, v := range fsOverlapList {
		cnt++

		if cnt%2 == 0 {
			analyzeSoupOverlapAndGap(db, "overlap", fsOverlapList[i-1].SourceId, fsOverlapList[i-1].BeginTimeStrUTC, fsOverlapList[i-1].EndTimeStrUTC, fsOverlapList[i-1].DurationMin, fsOverlapList[i].BeginTimeStrUTC, fsOverlapList[i].EndTimeStrUTC, fsOverlapList[i].DurationMin)
		}

		if i+1 < len(fsOverlapList) && v.SourceId != fsOverlapList[i+1].SourceId || i+1 == len(fsOverlapList) {
			fmt.Printf("\nTotal mongdb entry for overlap for source=%d is %d\n", v.SourceId, cnt/2)
			cnt = 0
		}
	}

	cnt = 0
	fmt.Println("gap report")
	fmt.Println("===================")
	for i, v := range fsGapList {
		cnt++
		//fmt.Printf("\n %d. sourceId=%d, beginTimeUtcStr=%s, duration=%d, endTimeUtcStr=%s", cnt, v.SourceId, v.BeginTimeStrUTC, v.DurationMin, v.EndTimeStrUTC)

		// always take 1st and 2nd records
		if cnt%2 == 0 {
			analyzeSoupOverlapAndGap(db, "gap", fsGapList[i-1].SourceId, fsGapList[i-1].BeginTimeStrUTC, fsGapList[i-1].EndTimeStrUTC, fsGapList[i-1].DurationMin, fsGapList[i].BeginTimeStrUTC, fsGapList[i].EndTimeStrUTC, fsGapList[i].DurationMin)
		}
		/*
			schedules := getSupair_provider_schedulesBySourceIdBeginTime(db, v.SourceId, v.BeginTimeStrUTC)
			if len(schedules) <= 0 {
				fmt.Println(" No schedules found in the cosmos db for the above record")
			} else {
				fmt.Println("\nSchedules in cosmos for the record given above")
				for _, s := range schedules {
					fmt.Printf("\n   sourceId=%d, beginTimeUtcStr=%s, duration=%d, endTimeUtcStr=%s, supairId=%d, programId=%s", v.SourceId, s.BeginTimeStrUTC, s.Duration, s.EndTimeStrUTC, s.SupairId, s.ProgramId)

				}
			}
		*/

		if i+1 < len(fsGapList) && v.SourceId != fsGapList[i+1].SourceId || i+1 == len(fsGapList) {
			fmt.Printf("\nTotal mongdb entry for gap for source=%d is %d\n", v.SourceId, cnt/2)
			cnt = 0
		}
	}

	// var aliasSourceFailed1000 = make(map[string]int64)
	// cnt := 0
	// for alias, sourceId := range aliasSourceFailed {
	// 	cnt++
	// 	if cnt == 1000 {
	// 		break
	// 	}
	// 	aliasSourceFailed1000[alias] = sourceId
	// }

	// err = fullResetSources(simsApiFullResetURL, aliasSourceFailed1000)
	// err = fullResetSources(simsApiFullResetURL, aliasSourceFailed)
	// if err != nil {
	// 	fmt.Println("error in csv-full-reset for failed sources. error=%s", err.Error())
	// } else {
	// 	fmt.Println("----------->Full reset for failed sources is scheduled to sims. Check sims log /var/log/sims/sims.log<---------------")
	// }

	// subject := "action: soup_failed_full_reset"
	// body := `
	// List of sources that is sent to SIMS for full reset=<RESET_SOURCES>
	// List of sources that is sent to SIMS for full reset_count=<RESET_SOURCES_COUNT>
	// `

	// body = strings.ReplaceAll(body, "<RESET_SOURCES>", fmt.Sprintf("%v", aliasSourceFailed))
	// body = strings.ReplaceAll(body, "<RESET_SOURCES_COUNT>", fmt.Sprintf("%v", len(aliasSourceFailed)))

	// sendEmail(subject, body)

	return nil
}

func getSupair_provider_schedulesBySourceIdBeginTime(db *sql.DB, sourceId int64, beginTimeUTC string) []SupairProviderScheduleCosmos {
	var schedules []SupairProviderScheduleCosmos

	qry := "SELECT supair_id, skd_program_id, skd_begin_time, skd_end_time, skd_duration FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_source_mapping_id = ? and skd_begin_time=?"
	stmt, err := db.Prepare(qry)
	if err != nil {
		fmt.Printf("error preparing mysql query = %s. error=%s", qry, err.Error())
	}

	rows, err := stmt.Query(sourceId, beginTimeUTC)
	if err != nil {
		fmt.Printf("error executing query=%s. error=%s", qry, err.Error())
		os.Exit(1)
	}

	for rows.Next() {
		schedule := SupairProviderScheduleCosmos{}
		err := rows.Scan(&schedule.SupairId, &schedule.ProgramId, &schedule.BeginTimeStrUTC, &schedule.EndTimeStrUTC, &schedule.Duration)
		if err != nil {
			fmt.Println("error scanning program id from cosmos mysql db. error=%s", err.Error())
			os.Exit(1)
		}
		schedule.Duration /= 60
		schedules = append(schedules, schedule)
	}

	return schedules
}

func analyzeSoupOverlapAndGap(db *sql.DB, issueType string, sourceId int64, beginTime1, endTime1 string, duration1 int, beginTime2, endTime2 string, duration2 int) {
	findEndTimeForCosmosApi := func(beginDateStr, beginTimeStr string, durationInMin int64) string {
		utcDateStr := beginDateStr + " " + beginTimeStr
		utcDateTime, err := time.Parse("2006-01-02 15:04:05", utcDateStr)
		if err != nil {
			fmt.Printf("\nerror converting schedule api datetime string to datetime. datetime=%s, error=%s\n", utcDateStr, err.Error())
			return ""
		}
		newUtcTime := utcDateTime.Add(time.Duration(durationInMin) * time.Minute)
		return newUtcTime.Format("2006-01-02 15:04:05")
	}

	overlapInCosmos := false
	overlapGoodRecordInCosmos := false

	fmt.Printf("\n---------- Begin: soup %s analysis: source=%d ----------\n", issueType, sourceId)
	fmt.Println("In Soup")
	fmt.Printf("\n\t source=%d, beginTime=%s endTime=%s durationInMin=%d", sourceId, beginTime1, endTime1, duration1)
	fmt.Printf("\n\t source=%d, beginTime=%s endTime=%s durationInMin=%d\n", sourceId, beginTime2, endTime2, duration2)

	// Record in Cosmos API
	var endTimeRecord1CosmosApi string
	fmt.Println("In the cosmos schdule api: ")
	cosmosSchedule, exist := cosmosApiObj.CheckStartDateExist(sourceId, beginTime1)
	if !exist {
		fmt.Println("\tNo record in Cosmos API")
	} else {
		endTimeRecord1CosmosApi = findEndTimeForCosmosApi(cosmosSchedule.StartDate, cosmosSchedule.StartTime, cosmosSchedule.DurationInSec/60)
		fmt.Printf("\n\t source=%d, beginTime=%s %s, endTime=%s, durationInMin=%d\n", cosmosSchedule.SourceId, cosmosSchedule.StartDate, cosmosSchedule.StartTime, endTimeRecord1CosmosApi, cosmosSchedule.DurationInSec/60)
	}

	cosmosSchedule, exist = cosmosApiObj.CheckStartDateExist(sourceId, beginTime2)
	if !exist {
		fmt.Println("\t No overlap record in Cosmos API")
	} else {
		fmt.Printf("\n\t source=%d, beginTime=%s %s, endTime=%s, durationInMin=%d\n", cosmosSchedule.SourceId, cosmosSchedule.StartDate, cosmosSchedule.StartTime, findEndTimeForCosmosApi(cosmosSchedule.StartDate, cosmosSchedule.StartTime, cosmosSchedule.DurationInSec/60), cosmosSchedule.DurationInSec/60)
	}

	if len(endTimeRecord1CosmosApi) > 1 {
		cosmosSchedule, exist = cosmosApiObj.CheckStartDateExist(sourceId, endTimeRecord1CosmosApi)
		if !exist {
			fmt.Println("\t No good record in Cosmos API as per the endtime of first record")
		} else {
			fmt.Printf("\n\t source=%d, beginTime=%s %s, endTime=%s, durationInMin=%d\n", cosmosSchedule.SourceId, cosmosSchedule.StartDate, cosmosSchedule.StartTime, findEndTimeForCosmosApi(cosmosSchedule.StartDate, cosmosSchedule.StartTime, cosmosSchedule.DurationInSec/60), cosmosSchedule.DurationInSec/60)
		}
	}

	// overlap Record in Cosmos DB
	fmt.Println("\nIn Cosmos DB: ", issueType)
	schedule1 := getSupair_provider_schedulesBySourceIdBeginTime(db, sourceId, beginTime1)

	if len(schedule1) == 0 {
		fmt.Println("------> No record in cosmos <----------")
	}

	for _, s := range schedule1 {
		fmt.Printf("\n   sourceId=%d, beginTimeUtcStr=%s, endTimeUtcStr=%s, duration=%d, supairId=%d, programId=%s", sourceId, s.BeginTimeStrUTC, s.EndTimeStrUTC, s.Duration, s.SupairId, s.ProgramId)
	}

	schedule2 := getSupair_provider_schedulesBySourceIdBeginTime(db, sourceId, beginTime2)
	if len(schedule2) == 0 {
		fmt.Println("\n------> No overlap record in cosmos db <----------\n")
	}

	for _, s := range schedule2 {
		overlapInCosmos = true
		fmt.Printf("\n   sourceId=%d, beginTimeUtcStr=%s, endTimeUtcStr=%s, duration=%d, supairId=%d, programId=%s", sourceId, s.BeginTimeStrUTC, s.EndTimeStrUTC, s.Duration, s.SupairId, s.ProgramId)
	}

	// since there is no overlap in cosmos db
	// find good records based on cosmos db endtime

	var cosmosDbEndtimeOfFirstRecord string
	if issueType == "overlap" && len(schedule1) >= 1 && len(schedule2) >= 1 && schedule1[0].EndTimeStrUTC == schedule2[0].BeginTimeStrUTC {
		fmt.Println("NO OVERLAP_IN_COSMOS_DB")
		overlapInCosmos = false
		cosmosDbEndtimeOfFirstRecord = schedule1[0].EndTimeStrUTC
	}

	// Finding Good records in cosmos db
	if len(cosmosDbEndtimeOfFirstRecord) < 1 { // This is prepared once we find there is no overlap in cosmos
		cosmosDbEndtimeOfFirstRecord = endTime1
	}
	fmt.Println("\n\n In Cosmos: Good Records - as per soup beginTime and endTime")
	if len(schedule1) == 0 {
		fmt.Println("\n------> No good record in cosmos <----------\n")
	}

	// print 1st record what we found above
	for _, s := range schedule1 {
		fmt.Printf("\n   sourceId=%d, beginTimeUtcStr=%s, endTimeUtcStr=%s, duration=%d, supairId=%d, programId=%s", sourceId, s.BeginTimeStrUTC, s.EndTimeStrUTC, s.Duration, s.SupairId, s.ProgramId)
	}

	// 2nd good record's start time will be end time of the 1st record
	schedule1End := getSupair_provider_schedulesBySourceIdBeginTime(db, sourceId, cosmosDbEndtimeOfFirstRecord)
	if len(schedule1End) == 0 {
		fmt.Println("\n------> No good record in cosmos <----------\n")
	}

	for _, s := range schedule1End {
		overlapGoodRecordInCosmos = true
		fmt.Printf("\n   sourceId=%d, beginTimeUtcStr=%s, endTimeUtcStr=%s, duration=%d, supairId=%d, programId=%s", sourceId, s.BeginTimeStrUTC, s.EndTimeStrUTC, s.Duration, s.SupairId, s.ProgramId)
	}

	fmt.Println("")
	if overlapGoodRecordInCosmos == true && overlapInCosmos == true {
		ans := "what"

		fmt.Println("\n since we found good records, do you want to delete the record given below from table: supair_provider_schedules?")
		for _, s := range schedule2 {
			overlapInCosmos = true
			fmt.Printf("\n   sourceId=%d, beginTimeUtcStr=%s, endTimeUtcStr=%s, duration=%d, supairId=%d, programId=%s", sourceId, s.BeginTimeStrUTC, s.EndTimeStrUTC, s.Duration, s.SupairId, s.ProgramId)
		}
		for {
			fmt.Println("\nY:N?")
			fmt.Scanf("%s", &ans)
			if strings.ToLower(ans) == "y" || strings.ToLower(ans) == "yes" {
				fmt.Println("-----------------------------------------------------------------------------")
				fmt.Println("|After deletion, run the query given below in Cosmos DB, and change delivery_load_type=FULL")
				fmt.Println(`|
|Select delivery_parameters from supair.config_delivery_assoc where track_id =6 and delivery_group = 'provider_schedule_onlyscheduledata';`)
				fmt.Println("|Then Run Argoflow: cwf-datadelivery-provider-curated-schedule")
				fmt.Println("|Argoflow will update ES index")
				fmt.Println("|once argoflow is done. Revert back to DELTA from FULL.")
				err := deleteSchedule(db, sourceId, schedule2[0].BeginTimeStrUTC, schedule2[0].EndTimeStrUTC, schedule2[0].Duration*60)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("\n-----------------------------------------------------------------------------")
				break

			} else if strings.ToLower(ans) == "n" || strings.ToLower(ans) == "no" {
				break
			}
		}
	}
	fmt.Printf("\n---------- End: soup %s analysis ----------\n", issueType)

}

func deleteSchedule(db *sql.DB, sourceId int64, beginDateTime string, endDateTime string, durationInSec int64) error {

	qry := "delete FROM supair.supair_provider_schedules where skd_ingest_source = 'GN' and skd_source_mapping_id = ? and skd_begin_time=? and skd_end_time = ? and skd_duration = ?"
	stmt, err := db.Prepare(qry)
	if err != nil {
		fmt.Printf("error preparing mysql query = %s. error=%s", qry, err.Error())
	}

	result, err := stmt.Exec(sourceId, beginDateTime, endDateTime, durationInSec)
	if err != nil {
		return fmt.Errorf("error executing query=%s. error=%s", qry, err.Error())
	}

	rowCnt, _ := result.RowsAffected()
	fmt.Println("Rows deleted count=", rowCnt)

	return nil
}
