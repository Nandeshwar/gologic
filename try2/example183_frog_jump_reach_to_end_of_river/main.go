package main

import (
	"fmt"
)

func main() {
	/*


		  stones    : 0  1  3  5  6  8  12  17
		         from stone 0, there can be 1 jump and can reach to stone 1
				  m[0] = 1
				 from stone 1, there can be following ways
				                1 - 1 step = 0  Not possible, so don't store in map
								   1  --------------> from 1 to steps 1 -- it will be jump to stone 2 and there is no stone2,
								                       this will be stored in map and rejected from loop
								1 + 1 step = 2 so from current position(1) + 2 = 3
				m[0] = 1 (jump)
				m[1] = 1, 2 (jump)

				--- from 1 + 2 will jump to 3






	*/

	stones := []int{0, 1, 3, 5, 6, 8, 12, 17} // expectation true
	// stones := []int{0, 1, 5} // expectation false
	// from 0 step to 1
	// from 1 three options 1-1 = 0 not possbile
	//                      1 i.e. current position(1) + 1 = 2 there is no stone 2
	//                      1 + 1 = 2 current position(1) + 2  = 3 there is stone 3
	fmt.Println(canFrogReachToEnd(stones))
}

func canFrogReachToEnd(a []int) bool {
	// key -> sets
	m := map[int]map[int]struct{}{}

	// assign key with empty sets
	for _, v := range a {
		m[v] = make(map[int]struct{})
	}

	// Stone 0 - sets value 1
	// there is 1 jump possible from 0
	set := m[0]
	set[1] = struct{}{}
	m[0] = set

	for i := 0; i < len(a); i++ {
		currentStone := a[i]

		jumpOptionsMap := m[currentStone]
		for jumpOption, _ := range jumpOptionsMap {

			// if jump and end stone reaches, return true
			if currentStone+jumpOption == a[len(a)-1] {
				return true
			}

			// find next stones sets and add all possbile jumps from that stone
			sets, ok := m[currentStone+jumpOption]
			if ok {
				newJumpOption1 := jumpOption - 1
				if newJumpOption1 > 0 {
					sets[newJumpOption1] = struct{}{}
				}
				newJumpOption2 := jumpOption
				sets[newJumpOption2] = struct{}{}

				newJumpOption3 := jumpOption + 1
				sets[newJumpOption3] = struct{}{}

				m[jumpOption] = sets
			}
		}
	}
	return false
}
