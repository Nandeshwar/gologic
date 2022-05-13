package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func CallCosmosEventDetailAPI(baseUrl, api, method, accessToken, programId string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := baseUrl + api + "/" + programId

	bearer := "Bearer " + accessToken
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("cosmos api=%s, Got error %s", err.Error(), url)
	}
	//req.SetBasicAuth(username, password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("cosmos api=%s. Got error while Do call %s", url, err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("cosmos api=%s. error while reading the response bytes. error=%s", url, err.Error())
	}

	defer response.Body.Close()

	result := string([]byte(body))

	if result == `{"detail":"Not found"}` {
		return "", fmt.Errorf("data not found for id=%s", programId)
	}

	return result, nil
}
