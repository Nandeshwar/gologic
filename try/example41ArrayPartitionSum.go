/*
	Check if the list of integer contains items that can sum to given number
*/
package main
import (
	"fmt"
)

func main() {
	list := []int{10, 20, 2, 4, 8}
	result := 12

	m := map[string]bool{}

	fmt.Println(checkSum(list, 0, 0, result, m))
}

func checkSum(list []int, index int, sum int, result int, m map[string]bool) bool {
	state :=fmt.Sprintf("%d%d", index, sum)

	v, ok := m[state]
	if ok {
		return v
	}

	if sum == result {
		return true
	}

	if index >= len(list) {
		return false
	}

	r :=  checkSum(list, index + 1, sum, result, m) || checkSum(list, index + 1, sum + list[index], result, m)
	m[state] = r
	return r

}
