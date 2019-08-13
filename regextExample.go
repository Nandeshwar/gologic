package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	r, _ := regexp.Compile("(\\\\\"|\")password(\\\\\"|\"):(.*\\\\\"|\")\\w+(\\\\\"|\")")

	str := `{\"host\":\"11.100.112.23\",\"port\":\"5672\",\"username\":\"nandeshwar\",\"password\":\"abc\"}`
	//str := `{\"host\":\"11.100.112.23\",\"port\":\"5672\",\"username\":\"nandeshwar\","password":"abc"}`
	foundItems := r.FindAllString(str, -1)
	str = strings.ReplaceAll(str, foundItems[0], "\"password\": \"*****\"")
	fmt.Println(str)
}
