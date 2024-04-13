package main

import (
	"strconv"
	"strings"
)

func main() {

}

func restoreIpAddresses(s string) []string {

	result := []string{}
	i := 0
	dot := 0
	currIp := ""
	generateIpAddress(i, dot, currIp, s, &result)
	return result
}

func generateIpAddress(i, dot int, currIp, s string, result *[]string) {
	if dot == 4 && i == len(s) {
		lastIndexOfDot := strings.LastIndex(currIp, ".")
		*result = append(*result, currIp[0:lastIndexOfDot])
		return
	}

	if dot > 4 {
		return
	}

	for j := i; j < Min(i+3, len(s)); j++ {
		eachIpSectionStr := s[i : j+1]
		eachIpSectionInt, _ := strconv.Atoi(eachIpSectionStr)
		if eachIpSectionInt <= 255 && (len(eachIpSectionStr) == 1 || s[i] != '0') {
			generateIpAddress(j+1, dot+1, currIp+eachIpSectionStr+".", s, result)
		}
	}

}

func generateIpAddress2(i, dot int, currIp, s string, result *[]string) {
	if dot == 4 && i == len(s) {
		lastIndexOfDot := strings.LastIndex(currIp, ".")
		*result = append(*result, currIp[0:lastIndexOfDot])
		return
	}

	if dot > 4 {
		return
	}

	for j := i; j < Min(i+3, len(s)); j++ {
		eachIpSectionStr := s[i : j+1]
		eachIpSectionInt, _ := strconv.Atoi(eachIpSectionStr)
		if eachIpSectionInt > 255 {
			return
		}
		if len(eachIpSectionStr) > 1 && eachIpSectionStr[0] == '0' {
			return
		}

		generateIpAddress(j+1, dot+1, currIp+eachIpSectionStr+".", s, result)

	}

}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
