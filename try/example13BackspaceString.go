// https://www.youtube.com/watch?v=tUxW1JwEb9M
package main
import(
	"fmt"
)

func main() {
	s := "ab#c"
	t := "ad#c"

	fmt.Println(isStrSame(s, t))

	s = "abcd##"
	t = "axcd##"

	fmt.Println(isStrSame(s, t))

}
func isStrSame(s, t string) bool {


	i := len(s) - 1
	j := len(t) - 1

	for i >= 0 || j >= 0 {
		cnt := 0
		for i >= 0 && (cnt > 0 || s[i] == '#') {
			if s[i] == '#' {
				cnt++
			} else {
				cnt --
			}
			i --
		}

		for j >= 0 && (cnt > 0 || t[j] == '#') {
			if t[j] == '#' {
				cnt++
			} else {
				cnt--
			}
			j--
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			} else {
				i--
				j--
			}
		} else {
			if i >= 0 || j >= 0 {
				return false
			}
		}
	}
	return true
}