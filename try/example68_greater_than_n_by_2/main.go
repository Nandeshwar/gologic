package main

func findRepeatedNumberGreaterThanNBy2(numList []int) int {
	// 1 . two loop time complexity o(n2) and space complexity o(1)
	// 2. sort number and 1 loop : o(nlogn) and space complexity 0(1)
	// 3. 1 loop and hashmap o(n) but space complexity o(n)
	// 4th startegy is o(n) and space complexity o(1)
	// get first number and 1 to the count, if same number inrement otherwise decrement,
	// repeated number greater than n/2 will have always count greater 0

	p := numList[0]
	cnt := 0
	currentNumber := numList[0]

	for _, v := range numList {
		if v == p {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			currentNumber = v
			cnt = 1
		}

		p = v
	}

	return currentNumber
}
