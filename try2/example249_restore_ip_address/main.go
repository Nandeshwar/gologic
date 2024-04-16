package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("restoreIpAddresses(\"12345\"): ", restoreIpAddresses("12345"))
	fmt.Println("restoreIpAddresses2(\"12345\"): ", restoreIpAddresses2("12345"))
	fmt.Println("restoreIpAddresses3(\"12345\"): ", restoreIpAddresses3("12345"))
	fmt.Println("restoreIpAddresses4(\"12345\"): ", restoreIpAddresses4("12345"))
	fmt.Println("restoreIpAddresses4(\"12345\"): ", len(restoreIpAddresses4("12345")))
	m := make(map[string]struct{})
	for _, v := range restoreIpAddresses4("12345") {
		m[v] = struct{}{}
	}
	fmt.Println("restoreIpAddresses4(\"12345\") map: ", len(m))
}

func restoreIpAddresses(s string) []string {

	result := []string{}
	i := 0
	dot := 0
	currIp := ""
	generateIpAddress(i, dot, currIp, s, &result)
	return result
}

func restoreIpAddresses2(s string) []string {

	result := []string{}
	i := 0
	dot := 0
	currIp := ""
	generateIpAddress2(i, dot, currIp, s, &result)
	return result
}

func restoreIpAddresses3(s string) []string {

	result := []string{}
	i := 0
	dot := 0
	currIp := ""
	generateIpAddress3(i, dot, currIp, s, &result)
	return result
}

func restoreIpAddresses4(s string) []string {

	result := []string{}
	i := 0
	dot := 0
	currIp := ""
	generateIpAddress4(i, dot, currIp, s, &result)
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

		generateIpAddress2(j+1, dot+1, currIp+eachIpSectionStr+".", s, result)

	}

}

func generateIpAddress3(i, dot int, currIp, s string, result *[]string) {
	if dot == 4 && i == len(s) {
		lastIndexOfDot := strings.LastIndex(currIp, ".")
		*result = append(*result, currIp[0:lastIndexOfDot])
		return
	}

	if dot > 4 {
		return
	}

	if i+1 <= len(s) {
		eachIpSectionStr := s[i : i+1]
		generateIpAddress3(i+1, dot+1, currIp+eachIpSectionStr+".", s, result)
	}

	if i+2 <= len(s) {
		eachIpSectionStr := s[i : i+2]
		if eachIpSectionStr[0] != '0' {
			generateIpAddress3(i+2, dot+1, currIp+eachIpSectionStr+".", s, result)
		}
	}

	if i+3 <= len(s) {
		eachIpSectionStr := s[i : i+3]
		if eachIpSectionStr[0] != '0' {
			eachIpSectionInt, _ := strconv.Atoi(eachIpSectionStr)
			if eachIpSectionInt < 256 {
				generateIpAddress3(i+3, dot+1, currIp+eachIpSectionStr+".", s, result)
			}
		}
	}
}

func generateIpAddress4(i, dot int, currIp, s string, result *[]string) {

	if dot == 4 {
		lastIndexOfDot := strings.LastIndex(currIp, ".")
		*result = append(*result, currIp[0:lastIndexOfDot])
		return
	}

	if i >= len(s) {
		return
	}

	eachIpSectionStr := s[i : i+1]
	eachIpSectionInt, _ := strconv.Atoi(eachIpSectionStr)
	if eachIpSectionInt <= 255 {
		originalCurrIp := currIp
		currIp = currIp + eachIpSectionStr + "."
		generateIpAddress4(i+1, dot+1, currIp, s, result)
		currIp = originalCurrIp
	}
	if i+2 < len(s) {
		eachIpSectionStr = s[i : i+2]
		eachIpSectionInt, _ = strconv.Atoi(eachIpSectionStr)
		if eachIpSectionInt <= 255 {
			originalCurrIp := currIp
			currIp = currIp + eachIpSectionStr + "."
			generateIpAddress4(i+1, dot+1, currIp, s, result)
			currIp = originalCurrIp
		}
	}

	if i+3 < len(s) {
		eachIpSectionStr = s[i : i+3]
		eachIpSectionInt, _ = strconv.Atoi(eachIpSectionStr)
		if eachIpSectionInt <= 255 {
			originalCurrIp := currIp
			currIp = currIp + eachIpSectionStr + "."
			generateIpAddress4(i+1, dot+1, currIp, s, result)
			currIp = originalCurrIp
		}
	}

	generateIpAddress4(i+1, dot, currIp, s, result)
	generateIpAddress4(i+2, dot, currIp, s, result)
	generateIpAddress4(i+3, dot, currIp, s, result)

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
nandeshwar.sah@C02DX2V4MD6R example249_restore_ip_address % go run main.go
restoreIpAddresses("12345"):  [1.2.3.45 1.2.34.5 1.23.4.5 12.3.4.5]
restoreIpAddresses2("12345"):  [1.2.3.45 1.2.34.5 1.23.4.5 12.3.4.5]
restoreIpAddresses3("12345"):  [1.2.3.45 1.2.34.5 1.23.4.5 12.3.4.5]
restoreIpAddresses4("12345"):  [1.2.3.4 1.2.3.5 1.2.34.4 1.2.34.5 1.2.4.5 1.23.3.4 1.23.3.5 1.23.34.4 1.23.34.5 1.23.4.5 1.234.3.4 1.234.3.5 1.234.34.4 1.234.34.5 1.234.4.5 1.3.4.5 1.34.4.5 12.2.3.4 12.2.3.5 12.2.34.4 12.2.34.5 12.2.4.5 12.23.3.4 12.23.3.5 12.23.34.4 12.23.34.5 12.23.4.5 12.234.3.4 12.234.3.5 12.234.34.4 12.234.34.5 12.234.4.5 12.3.4.5 12.34.4.5 123.2.3.4 123.2.3.5 123.2.34.4 123.2.34.5 123.2.4.5 123.23.3.4 123.23.3.5 123.23.34.4 123.23.34.5 123.23.4.5 123.234.3.4 123.234.3.5 123.234.34.4 123.234.34.5 123.234.4.5 123.3.4.5 123.34.4.5 2.3.4.5 2.34.4.5 23.3.4.5 23.34.4.5 234.3.4.5 234.34.4.5]

*/
