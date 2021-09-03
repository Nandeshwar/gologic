/*
Algorithm: https://www.youtube.com/watch?v=dqHw4vOwXn8
  1. if 1st digit in string is greater than next digit remove first digit
  2. if digits are in ascending order, remove digit from end
  3. Remove leading zero 

	Example1:  
    	Input: num="1432219", k = 3
		Ouput: "1219"
	Example2:
		Input: num = "10200", k = 1
		Output: "200"

	Example3: 
		Input: num = "10", k = 2
		Output: "0"

*/
package main
import (
	"fmt"
	"container/list"
	
)

func main() {
	input := "1432219"
	k := 3

	// input := "12345"
	// k := 3

	result := removeKDigit(input, k)
	fmt.Println(result);

	result2 := removeKDigit2(input, k)
	fmt.Println(result2);
}

func removeKDigit(input string, k int) string {
	if k == 0 {
		return input
	}

	if len(input) == k {
		return ""
	}

	strArr := []rune(input)
	var strResult string
	for i := 0; i < len(strArr) - 1 ; i++ {
		if strArr[i] > strArr[i+1] && k > 0{
			k--
			continue
		}
		strResult += string(strArr[i])
	}
	// Adding last digit
	strResult += string(strArr[len(strArr) - 1])

	// Remove k digits from end
	if k > 0 {
		strResult = strResult[0:len(strResult) - k]
	}

	return strResult
}

// approach2 : using stack
func removeKDigit2(input string, k int) string {
	if k == 0 {
		return input
	}

	if len(input) == k {
		return ""
	}

	stack := list.New()
	strArr := []rune(input)
	var strResult string
	for i := 0; i < len(strArr); i++ {
		if k > 0 && stack.Len() > 0 {
			element := stack.Back().Value
			if element.(rune) > strArr[i] {
				stack.Remove(stack.Back())
				stack.PushBack(strArr[i])
				k--
				continue
			}
		}
		stack.PushBack(strArr[i])
	}
	stackLen := stack.Len()
	for i := 0; i < stackLen; i++ {
		removedElement := string(stack.Remove(stack.Back()).(rune))
		strResult += removedElement
	}
	strResult = reverse(strResult);
	
	// Remove k digits from end
	if k > 0 {
		strResult = strResult[0:len(strResult) - k]
	}

	return strResult
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}