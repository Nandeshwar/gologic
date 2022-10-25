/*
find distinct element count in each window
window size 4
[]int{1, 2, 2, 3, 4, 4, 5, 4}
window1: 1, 2, 2, 3  --> unique 3
window2: 2, 2, 3, 4 --> unique 3
window3: 2, 3, 4, 4 -> unique 3
window4: 3, 4, 4, 5 - unique 3
window5: 4, 4, 5, 4 - unique 2

total 5 window:
*/
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 2, 3, 4, 4, 5, 4}
	windowSlide(a, 4)
}

func windowSlide(a []int, windowSize int) {
	m := map[int]int{}

	//Added 1st window to map
	for i := 0; i < windowSize; i++ {
		v, ok := m[a[i]]
		if ok {
			v++
			m[a[i]] = v
		} else {
			m[a[i]] = 1
		}
	}

	fmt.Println("distinct count=", len(m))

	// Slide window by 1
	// if 1stItemInWindow's count is 1, delete it from window, otherwise decrement it by 1
	for i := windowSize; i < len(a); i++ {
		firstItemInWindow := a[i-windowSize] // 1st item in window will be deleted

		// in order to find distinct element. delete it from map if item count is 1 otherwise decrase count
		v := m[firstItemInWindow]
		if v > 1 {
			v--
			m[firstItemInWindow] = v
		} else {
			delete(m, firstItemInWindow)
		}

		// sliding window in this section
		// Add next item in window
		v, ok := m[a[i]] // sliding window in this line: adding next item
		if ok {
			v++
			m[a[i]] = v // sliding window in this line: adding next item
		} else {
			m[a[i]] = 1 // sliding window in this line: adding next item
		}

		fmt.Println("distinct count=", len(m))
	}
}
