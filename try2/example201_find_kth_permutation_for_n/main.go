package main

import (
	"fmt"
)

/*
  Explanation:
    n = 4
    4! = 24 total combination
    fix:
      1  { 2, 3, 4} 3! = 6
      2  { 1, 3, 4} 3! = 6
      3  { 1, 2, 4} 3! = 6
      4  { 1, 2, 3} 3! = 6

      k =17 : want to know where 17 lies. 0 index base
      so k-1 = 16

    and 16 falls under 3 1st portion of answer.
    how to get this:
        numbers: {1, 2, 3, 4}
        factorial = 6
        k / 6 = 16 / 6 = 2(index)
        2nd index in {1, 2, 3, 4} is 3

        Add 3 in the result and remove it from numbers so remaning: 1, 2, 4
    ---------------------
    Next iteration:
       calculate k and then factorial

        next item lies in {1, 2, 4}
        fix:
        1  { 2, 4} = 2! = 2
        2  {1, 4}  = 2!= 2
        4  { 1, 2} = 2! = 2


      calculate k at end of 1st iteration:
       k = k % fact = 16 % 6 = 4
       fact = fact / 3(result in 1st iteration) = 6 / 3 = 2

      k / fact = 4 / 2 = 2(index)
        {1 , 2, 4} == 4 is at 2nd index:
            4 will be the 2nd item in the list
                 This will be continued until numbers list length is 0





*/

func main() {
	n := 4
	k := 17
	// output: 3412

	fmt.Println(findKthPermutation(n, k))
}

func findKthPermutation(n, k int) string {

	fact := 1
	var numbers []int
	for i := 1; i < n; i++ {
		fact *= i
		numbers = append(numbers, i)
	}
	numbers = append(numbers, n)

	fmt.Println("fact=", fact)       // 6
	fmt.Println("numbers=", numbers) // [1, 2, 3, 4]

	k = k - 1

	var result string
	for {
		var ind int
		if fact > 0 {
			ind = k / fact
		}

		if len(numbers) < 1 {
			break
		}
		resultItem := numbers[ind]
		result += fmt.Sprintf("%d", resultItem)

		a1 := numbers[0:ind]
		a2 := numbers[ind+1:]

		numbers = a1
		numbers = append(numbers, a2...)

		if fact > 0 {
			k = k % fact
		}
		fact = fact / resultItem
	}
	return result
}
