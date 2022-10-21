// https://www.youtube.com/watch?v=9fI_26Dl1IA
package main

import "fmt"

func main() {
	fmt.Println(lookAndSay(3))
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
	fmt.Println(str)
	var nwStr string
	cnt := 0
	var i int
	for i = 0; i < len(str); i++ {
		cnt++

		if i == len(str)-1 || str[i] != str[i+1] {
			nwStr += fmt.Sprintf("%d%c", cnt, str[i])
			cnt = 0
		}
	}
	return nwStr

}
