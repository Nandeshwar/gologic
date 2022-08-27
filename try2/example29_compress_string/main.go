package main

import "fmt"

func main() {
	fmt.Println(compressStr("aabbcc"))
}

// example: abbb
// output: ab3
func compressStr(str string) string {
	a := []rune(str)

	index := 0

	for i := 0; i < len(a); {
		// first time j, i = 0
		// 2nd time j and i will be 1
		j := i

		// first time j = 1
		// 2nd time j will be 4
		for j < len(a) && a[i] == a[j] {
			j++
		}

		// first time a[0] will be "a"
		// 2nd time a[1] will be "b"
		a[index] = a[i]
		index++

		// first time not true
		// 2nd time 4 > 1
		// a[2]  will be 4
		if j-i > 1 {
			counterStr := fmt.Sprintf("%d", j-i)

			for _, v := range counterStr {
				a[index] = rune(v)
				index++
			}
		}

		// i and j will be 1 now
		// now i will be 4 and then do not come to loop again
		i = j

	}

	return string(a[0:index])
}
