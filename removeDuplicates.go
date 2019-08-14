package main

import "fmt"

func main() {
	str := []byte("abbcdab")
	fmt.Println("Original string: ", string(str))

	exists := func(v byte, a []byte) bool {
		for i := 0; i < len(a); i++ {
			if str[i] == v {
				return true
			}
		}
		return false
	}
	j := 0

	for i := 0; i < len(str); i++ {
		v := str[i]
		if exists(v, str[i+1:]) {
			continue
		}
		str[j] = str[i]
		j++
	}
	fmt.Println("After Removing duplicates", string(str[0:j]))
}

/* Ouput
Original string:  abbcdab
After Removing duplicates cdab
*/
