package main

// keep storing currSum-> index to map
// currSum == sum return from index 0 to i
// otherwise : if currSum - sum is in map, get index from map + 1 to i(current index)
func subArrayBySum(a []int, sum int) []int {
	currSum := 0
	m := map[int]int{}
	for i := 0; i < len(a); i++ {
		currSum += a[i]
		if currSum == sum {
			return a[0 : i+1]
		}

		index, ok := m[currSum-sum]
		if ok {
			return a[index+1 : i+1]
		}
		m[currSum] = i
	}
	return []int{}
}
