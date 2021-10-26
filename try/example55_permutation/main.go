package main

func permutation(s string, left, right int, allStrList []string) []string {
	if left == right {
		allStrList = append(allStrList, s)
		return allStrList
	}
	for i := left; i <= right; i++ {
		strArr := []rune(s)
		strArr[left], strArr[i] = strArr[i], strArr[left]
		allStrList = permutation(string(strArr), left+1, right, allStrList)
		strArr[left], strArr[i] = strArr[i], strArr[left]
	}
	return allStrList
}

func permutation2(s string, resultStr string, allStrList []string) []string {
	if len(s) == 0 {
		allStrList = append(allStrList, resultStr)
		return allStrList
	}

	for i := 0; i < len(s); i++ {
		leftPortionOfString := string(s[0:i])
		rightPortionOfString := string(s[i+1:])

		newStr := leftPortionOfString + rightPortionOfString

		allStrList = permutation2(newStr, resultStr+string(s[i]), allStrList)

	}
	return allStrList
}
