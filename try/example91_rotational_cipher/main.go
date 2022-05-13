// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	str := "zZ"

	strArr := []rune(str)
	var newStrArr []rune

	for _, v := range strArr {
		if v >= 65 && v <= 90 {
			nv := v + 10
			var nvr rune
			if nv > 90 {
				nvr = 65 + nv - 91
				newStrArr = append(newStrArr, nvr)
			} else {
				newStrArr = append(newStrArr, nv)
			}
		}

		if v >= 97 && v <= 122 {
			nv := v + 10
			var nvr rune
			if nv > 122 {
				nvr = 97 + nv - 123
				newStrArr = append(newStrArr, nvr)
			} else {
				newStrArr = append(newStrArr, nv)
			}
		}

	}

	fmt.Println(string(newStrArr))
}
