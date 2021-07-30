package main

import (
	"fmt"
)

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var m2 []int

	l := 0
	r := len(m1[0]) - 1
	t := 0
	b := len(m1) - 1

	dir := 'r'

	for l <= r && t <= b {
		if dir == 'r' {
		
			for i := l; i <= r; i++ {
				m2 = append(m2, m1[t][i])
			}
			t++;
			dir = 'b'
		}

		if dir == 'b' {
			for i := t; i <= b; i++ {
				m2 = append(m2, m1[i][r])
			}
			r--;
			dir = 'l'
		}

		if dir == 'l' {
			for i := r; i >= l; i-- {
				m2 = append(m2, m1[b][i])
			}
			b--
			dir = 't'
		}

		if dir == 't' {
			for i := b; i >= t; i-- {
				m2 = append(m2, m1[i][l])
			}
			l++;
			dir = 'r'
		}
	}

	fmt.Println(m2)

}