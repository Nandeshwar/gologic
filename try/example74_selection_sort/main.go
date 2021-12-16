package main

// find index of smallest number
// swith smallest index number with i
func selectionSort(slice []int) {

	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if slice[minIndex] > slice[j] {
				minIndex = j
			}
		}
		if slice[minIndex] != slice[i] {
			slice[minIndex], slice[i] = slice[i], slice[minIndex]
		}
	}
}
