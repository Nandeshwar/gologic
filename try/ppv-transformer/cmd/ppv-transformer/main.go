package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"path/filepath"

	"ppv-transformer/pkg/auth"
	"ppv-transformer/pkg/csv"
	"ppv-transformer/pkg/ziputil"
)

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

	cosmosHost := viper.GetString("cosmos.host")
	cosmosTokenApi := viper.GetString("cosmos.tokenApi")
	cosmosEventDetailApi := viper.GetString("cosmos.eventDetailApi")
	cosmosUser := viper.GetString("cosmos.user")
	cosmosPassword := viper.GetString("cosmos.password")

	//inputPpvFileZip := viper.GetString("cosmos.inputPPVZipFile")

	mainArgsLen := len(os.Args)
	if mainArgsLen != 2 {
		fmt.Println("There must be argument. Pass filename as argument")
		fmt.Println("exiting...")
		os.Exit(1)
	}
	inputPpvFileZip := os.Args[1]
	inputExt := filepath.Ext(inputPpvFileZip)

	if strings.ToLower(inputExt) != ".zip" {
		fmt.Println("pass zip file as argument")
		fmt.Println("exiting...")
		os.Exit(1)
	}

	fmt.Println(inputPpvFileZip)

	baseFileName := filepath.Base(inputPpvFileZip)
	baseFileOnly := strings.Split(baseFileName, ".")[0]

	outputSuccessPPVZipFile := strings.ReplaceAll(inputPpvFileZip, baseFileOnly, baseFileOnly+"_f")
	outputFailureFile := strings.ReplaceAll(inputPpvFileZip, baseFileOnly, baseFileOnly+"_Missing_TMS_IDs.txt")

	extOutputFailureFile := filepath.Ext(outputFailureFile)
	outputFailureFile = strings.ReplaceAll(outputFailureFile, extOutputFailureFile, "")

	os.Remove(outputFailureFile)
	os.Remove(outputSuccessPPVZipFile)

	fileInsideZip, err := ziputil.FirstFileNameInsideZip(inputPpvFileZip)
	if err != nil {
		fmt.Printf("\nerror reading contents of zip file. filename=%s, error=%s\n", inputPpvFileZip, err.Error())
		return
	}

	defer func() {
		os.Remove(fileInsideZip)
	}()

	os.Remove(fileInsideZip)
	fmt.Println("fileInsideZip=", fileInsideZip)

	err = ziputil.UnzipSource(inputPpvFileZip, "")
	if err != nil {
		log.Fatalf("error unziping file=%s, error=%s\n", inputPpvFileZip, err)
		return
	}

	fmt.Println("Getting token")
	token, err := auth.GetToken(cosmosHost, cosmosTokenApi, cosmosUser, cosmosPassword)
	if err != nil {
		fmt.Printf("error getting token. error=%s\n", err.Error())
		return
	}

	fmt.Println("Transforming CSV...")
	//err = csv.TransformCSV("PPV_SMALL.TXT", cosmosHost, cosmosEventDetailApi, token, outputFailureFile)
	err = csv.TransformCSV(fileInsideZip, cosmosHost, cosmosEventDetailApi, token, outputFailureFile)
	if err != nil {
		fmt.Printf("\nerror transforming csv=%s\n", err.Error())
		return
	}

	fmt.Println("Creating zip file")
	err = ziputil.ZipSource(fileInsideZip, outputSuccessPPVZipFile)
	if err != nil {
		fmt.Printf("\nerror creating zip file=%s. error=%s\n", outputSuccessPPVZipFile, err.Error())
		return
	}

	fmt.Println("--------------->SUCCESS<----------------------------\n")

}
