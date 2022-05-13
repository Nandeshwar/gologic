package cosmosapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"tba/pkg/models"
)

type CosmosApi struct {
	BaseUrl      string
	Api          string
	AccessToken  string
	IngestSource string
}

func New(baseUrl, api, accessToken, ingestSource string) *CosmosApi {
	return &CosmosApi{
		BaseUrl:      baseUrl,
		Api:          api,
		AccessToken:  accessToken,
		IngestSource: ingestSource,
	}
}

func (c *CosmosApi) CheckStartDateExist(sourceId int64, startDateTime string) (*models.Schedule, bool) {
	found := false

	var nextId int64
	for {
		jsonStr, err := c.callCosmosScheduleAPI(sourceId, nextId)
		if err != nil {
			fmt.Println("error calling cosmos api=", err)
			return nil, false
		}

		cosmosModel := &models.CosmosApiModel{}
		if len(jsonStr) > 1 {
			err = json.Unmarshal([]byte(jsonStr), cosmosModel)
			if err != nil {
				fmt.Printf("\nerror Unmarsalling comsos schedule api json. error=%s, jsonStr=%s", err, jsonStr)
				return nil, false
			}
		}

		/*
		   sourceId=15377, beginTimeUtcStr=2022-05-19 23:00:00

		   "startTime": "21:00:00",
		        "programId": "SH022888590000",
		        "startDate": "2022-05-09",

		*/

		dtArr := strings.Split(startDateTime, " ")
		startDate := dtArr[0]
		startTime := dtArr[1]

		for _, schedule := range cosmosModel.Schedules {
			if schedule.StartDate == startDate && schedule.StartTime == startTime {
				return &schedule, true
			}
		}

		if cosmosModel.Response.Header.NextId == 0 {
			break
		}
		nextId = cosmosModel.Response.Header.NextId
	}

	return nil, found
}

/*
curl -X 'GET' \
  'https://cosmos-api-p.supair.io/v1_0/schedule/?source_id=76950&next_id=0&ingest_source=gn&load_type=full' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJuYW5kZXNod2FyLnNhaEBkaXNoLmNvbSIsImV4cCI6MTY4Mzc0ODc2N30.iHNqQmSVWczCmg273zJDRIzLFroc6qHU4DDmq3mQxEA'
*/
func (c *CosmosApi) callCosmosScheduleAPI(sourceId int64, nextId int64) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	baseUrl, err := url.Parse(c.BaseUrl + c.Api)
	if err != nil {
		return "", fmt.Errorf("Malformed cosmos schedule URL=", err.Error())
	}

	params := url.Values{}
	params.Add("source_id", fmt.Sprintf("%d", sourceId))
	params.Add("next_id", fmt.Sprintf("%d", nextId))
	params.Add("ingest_source", c.IngestSource)
	params.Add("load_type", "full")

	baseUrl.RawQuery = params.Encode()
	cosmosUrl := baseUrl.String()

	bearer := "Bearer " + c.AccessToken
	req, err := http.NewRequest("GET", cosmosUrl, nil)
	if err != nil {
		return "", fmt.Errorf("cosmos api=%s, Got error %s", err.Error(), cosmosUrl)
	}
	//req.SetBasicAuth(username, password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("cosmos api=%s. Got error while Do call %s", baseUrl, err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("cosmos api=%s. error while reading the response bytes. error=%s", baseUrl, err.Error())
	}
	//fmt.Println("response.StatusCode=", response.StatusCode)

	defer response.Body.Close()

	result := string([]byte(body))

	// TODO: check right error message
	if result == `{"detail":"Ingest source and source ID mapping Not found"}` {
		return "", fmt.Errorf("data not found for id=%s", sourceId)
	}

	return result, nil
}
