package main

import (
	"fmt"
)

/*
input: 	a := []int{1, 7, 8, 4, 5, 6, -1, 9}
-- create tmpArr with ascending order 1, 7, 8 
-- if any number less than end (here 4 < 8)
-- replace 1st item greater than 4 with 4 . here 7 is > 4 
-- then 1st item greater than 5....here 8 now 8 and replace it with 5 and since it last item, make it max


output:
_____________-
max= 1
tmpArr.....= [1]
_____________-
max= 7
tmpArr.....= [1 7]
_____________-
max= 8
tmpArr.....= [1 7 8]
item= 4
_____________-
max= 8
tmpArr.....= [1 4 8]
item= 5
_____________-
max= 5
tmpArr.....= [1 4 5]
_____________-
max= 6
tmpArr.....= [1 4 5 6]
item= -1
_____________-
max= 6
tmpArr.....= [-1 4 5 6]
tmpArr= [-1 4 5 6 9]
5


*/

func main() {
	a := []int{1, 7, 8, 4, 5, 6, -1, 9}
	fmt.Println(lis(a))
}

var max = 0

func lis(a []int) int {
	var tmpArr []int
	tmpArr = append(tmpArr, a[0])
	max = a[0]
	for i := 1; i < len(a); i++ {
		fmt.Println("_____________-")
		fmt.Println("max=", max)
		fmt.Println("tmpArr.....=", tmpArr)
		if a[i] > max {
			tmpArr = append(tmpArr, a[i])
			max = a[i]
		} else {
			fmt.Println("item=", a[i])
			putItemInSamePlaceOrRightAfterSmallerItem(tmpArr, a[i])
		}
	}
	fmt.Println("tmpArr=", tmpArr)
	return len(tmpArr)
}

func putItemInSamePlaceOrRightAfterSmallerItem(tmpArr []int, item int) {

	for i := 0; i < len(tmpArr); i++ {
		if item == tmpArr[i] {
			return
		}

		// reach to last in arr that is greater than item
		for item > tmpArr[i] {
			i++
		}
		tmpArr[i] = item

		if i == len(tmpArr)-1 {
			max = item
		}
		break
	}
}
