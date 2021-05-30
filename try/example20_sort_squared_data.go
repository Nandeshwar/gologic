package main

import (
	"fmt"
)


func main() {
	list := []int{-4, -3, 0, 1, 2, 3, 5}
	newList := make([]int, len(list))

	i := 0;
	j := len(list) - 1

	k := j
	for i <= j {
		if i == j {
			newList[k] = list[i] * list[i]
			k++
			break
		}
		
		if list[i] * list[i] > list[j] * list[j] {
			fmt.Printf("\nlist i= %d, %d", list[i], list[i] * list[i])
			newList[k] = list[i] * list[i]
			i++
			k--
		} else  {
			fmt.Printf("\nlist j= %d %d", list[j], list[j] * list[j])
			newList[k] = list[j] * list[j]
			j--
			k++
		}
	}

	fmt.Println(fmt.Sprintf("original sorted list = %v", list))
	fmt.Println(fmt.Sprintf("sorted list after square= %v", newList))

}