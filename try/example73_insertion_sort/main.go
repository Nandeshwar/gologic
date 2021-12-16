package main

// 8 1 5 3 10
// logically divide this to two list - sorted(8) and unsorted(1, 5, 3, 10). To achieve this I am going to start outer loop from i = 1
// store ith value to tmp
// and j will be i -1

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		tmp := slice[i]
		j := i - 1

		for j >= 0 && slice[j] > tmp {
			slice[j+1] = slice[j] // first time 8 will move next place
			slice[j] = tmp
			j--
		}
		slice[j+1] = tmp
	}
}
