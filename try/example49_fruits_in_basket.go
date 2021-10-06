/*
  input:[3, 3, 3, 1,  2, 1, 1, 2, 3, 3, 4]
  output: 5 

  a. find two different number pattern together 
  b. then find max

  3 3 3 1     - 4 
  1 2 1 1 2   - 5
  3 3 4       - 3

  Algorith: 
     - store each item in map
	 - if map size is  3 i.e. 3rd pattern found
	     calaculate 
		           max
				   initialize cnt with prviousItemCount + current item
				   
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	max := 0
	cnt := 0
	previousItemCnt := 0
	previousItemCntBackup := 0

	m := map[int]struct{}{}
	list := []int{3, 3, 3, 1,  2, 1, 1, 2, 3, 3, 4}
	for _, v := range list {
		_, ok := m[v]
		if !ok {
			previousItemCntBackup = previousItemCnt;
			previousItemCnt = 0
		}
		m[v] = struct{}{} 
		cnt++
		previousItemCnt++;

		

		if len(m) == 3 {
			cnt--;
			max = int(math.Max(float64(max), float64(cnt)))
			cnt = previousItemCntBackup + 1
			m = map[int]struct{}{}
			m[v] = struct{}{} 
		}
	}
	fmt.Println(max)
}