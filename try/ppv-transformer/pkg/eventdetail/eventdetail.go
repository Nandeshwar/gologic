package eventdetail

import (
	"encoding/json"
	"fmt"

	"ppv-transformer/pkg/api"
)

type EventDetail struct {
	EchostarId         int64              `json:"echostar_id"`
	ProgramId          string             `json:"source_id"`
	EchostarSeriesInfo EchostarSeriesInfo `json:"echostar_series_info"`
}

type EchostarSeriesInfo struct {
	EchostarEventString string `json:"echostar_event_string"`
}

func (e *EventDetail) parseEventDetail(jsonStr string) error {

	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		return fmt.Errorf("error parsing json=%s. error=%s", jsonStr, err)
	}
	return nil
}

func GetEchostar3Id(programId, cosmosHost, cosmosEventDetailApi, token string) (string, error) {

	result, err := api.CallCosmosEventDetailAPI(cosmosHost, cosmosEventDetailApi, "GET", token, programId)
	if err != nil {
		return "", fmt.Errorf("\nerror calling cosmos api. error=%s", err.Error())
	}

	e := &EventDetail{}
	err = e.parseEventDetail(result)
	if err != nil {
		return "", err
	}

	return e.EchostarSeriesInfo.EchostarEventString, nil
}
