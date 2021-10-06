// https://www.geeksforgeeks.org/arrange-numbers-to-form-a-valid-sequence/

/*
Given an array arr[] with N distinct numbers and another array arr1[] with N-1 operators (either < or >), the task is to organize the numbers to form a valid sequence which obeys relational operator rules with respect to provided operators.

Examples: 

Input: arr[] = {3, 12, 7, 8, 5}; arr1= {‘<‘, ‘>’, ‘>’, ‘<‘} 
Output: {3, 12, 8, 5, 7} 
Explanation: 
3 < 12 > 8 > 5 < 7 
There can be more such combinations. The task is to return one of the combinations.

Input: arr[] = {8, 2, 7, 1, 5, 9}; arr1[] = {‘>’, ‘>’, ‘<‘, ‘>’, ‘<‘} 
Output:{9, 8, 1, 7, 2, 5} 
Explanation: 
9 > 8 > 1 < 7 > 2 < 5 
*/
package main

import (
	"fmt"
	"sort"
)
func main() {
	fmt.Println("what is RAM?")

}

func newSequece(numList []int, opList []string) []int {
	newList := make([]int, len(numList))

	sort.Ints(numList)

	i := 0;
	j := len(numList) - 1;
	k := 0;

	for ; i <= j && k <= len(numList) -2; {
		if opList[k] == "<" {
			newList[k] = numList[i];
			i++;
		} else  {
			newList[k] = numList[j]
			j--
		}
		k++
	}
	newList[len(numList)-1] = numList[i] 
	return newList;
}

func test(a int) int {
	return a;
}