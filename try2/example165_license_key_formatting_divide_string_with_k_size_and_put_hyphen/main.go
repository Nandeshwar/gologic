package main

import (
	"fmt"
)

func main() {
	/*
			Input: S = “5F3Z-2e-9-w”, K = 4
			Output: “5F3Z-2E9W”


			logic:
			  iterate from end
		      add 4 letters in new string then put hyphen
	*/

	str := "5F3Z-2e-9-w"
	k := 4

	fmt.Println(putDashAfterKLen(str, k))
}

func putDashAfterKLen(str string, k int) string {
	var nwStr string

	cnt := 0
	var tmpStr string
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == "-" {
			continue
		}
		cnt++
		tmpStr = string(str[i]) + tmpStr

		if cnt == 4 {
			if len(nwStr) == 0 {
				nwStr += tmpStr
			} else {
				nwStr = tmpStr + "-" + nwStr
			}
			cnt = 0
			tmpStr = ""
		}
	}
	if cnt > 0 && cnt < 4 {
		nwStr = tmpStr + "-" + nwStr
	}
	return nwStr
}
