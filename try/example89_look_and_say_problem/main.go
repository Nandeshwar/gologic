// https://www.youtube.com/watch?v=9fI_26Dl1IA
package main

import "fmt"

func main() {
	fmt.Println(lookAndSay(5))
}

/*
print frequeny and then item
   1
 1 1
2 1
1 2 1 1
1 1 1 2 2 1
*/
func lookAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	str := lookAndSay(n - 1)
	var nwStr string
	charArr := []rune(str)
	cnt := 0
	var i int
	for i = 0; i < len(charArr); i++ {
		cnt++
		if i+1 < len(charArr) && charArr[i] != charArr[i+1] {
			nwStr += fmt.Sprintf("%d%c", cnt, charArr[i])
			cnt = 0
		}
	}
	nwStr += fmt.Sprintf("%d%c", cnt, charArr[i-1])
	return nwStr

}
