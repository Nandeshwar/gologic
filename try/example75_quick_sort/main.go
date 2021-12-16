package main

func quickSort(a []int) {
	quickS(a, 0, len(a)-1)
}

func quickS(a []int, begin, end int) {
	if begin < end {
		pivotIndex := parition(a, begin, end)
		quickS(a, begin, pivotIndex-1)
		quickS(a, pivotIndex+1, end)
	}
}

// take first item as pivot
// fix two pointers i to 0 and j to end
// compare from left side - if item <= pivot, incrment i
// compare from right side - if item > pivot, decrement j
// otherwise if i < j, swap itme
// repeat it using outer for loop
// once j is less than i i.e once controll is out of outer for loop
// , swap pivot with j position and return j position with is pivot index
func parition(a []int, begin, end int) int {
	pivot := a[begin]

	i := begin
	j := end

	for i < j {
		for a[i] <= pivot {
			i++
		}

		for a[j] > pivot {
			j--
		}

		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	a[begin] = a[j]
	a[j] = pivot
	return j
}
