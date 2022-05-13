package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/logic-building/functional-go/fp"
)

func GetToken(baseUrl, tokenApi, user, password string) (string, error) {

	data := url.Values{
		"grant_type": {"password"},
		"username":   {user},
		"password":   {password},
	}

	resp, err := http.PostForm(baseUrl+tokenApi, data)

	if err != nil {
		return "", fmt.Errorf("error posting url=%s. error=%s", baseUrl+tokenApi, err.Error())
	}

	var maskPassword string
	for range fp.RangeInt(0, len(password)) {
		maskPassword += "*"
	}
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	accessToken := fmt.Sprintf("%v", result["access_token"])
	if accessToken == "<nil>" {
		return "", fmt.Errorf("error getting token from cosmos. url=%s. user=%s, password=%v, error=%v", baseUrl+tokenApi, user, maskPassword, result)
	}

	return accessToken, nil

}
