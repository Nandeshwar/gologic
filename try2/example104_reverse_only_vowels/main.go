package main

import (
	"fmt"
)

func main() {
	a := "Shruti"
	b := "Shruti"
	reverseVowels(&a)
	fmt.Println(a)
	reverseVowels2(&b)
	fmt.Println("b=", b)
}

func reverseVowels2(a *string) {
	aArr := []byte(*a)
	vowels := map[byte]struct{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
		'A': struct{}{},
		'E': struct{}{},
		'I': struct{}{},
		'O': struct{}{},
		'U': struct{}{},
	}
	i := 0
	j := len(aArr) - 1
	fmt.Println(aArr)

	for i < j {
		_, ok := vowels[(aArr)[i]]
		if !ok {
			i++
			continue
		}

		_, ok = vowels[(aArr)[j]]
		if !ok {
			j--
			continue
		}

		(aArr)[i], (aArr)[j] = (aArr)[j], (aArr[i])
		i++
		j--
	}
	*a = string(aArr)
}

func reverseVowels(a *string) {
	aArr := []byte(*a)
	vowels := map[byte]struct{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
		'A': struct{}{},
		'E': struct{}{},
		'I': struct{}{},
		'O': struct{}{},
		'U': struct{}{},
	}
	i := 0
	j := len(aArr) - 1
	fmt.Println(aArr)

	for i < j {
		for {
			_, ok := vowels[aArr[i]]
			if !ok {
				i++
			}
			if ok || i >= j {
				break
			}
		}

		for {
			_, ok := vowels[aArr[j]]
			if !ok {
				j--
			}
			if ok || j <= i {
				break
			}
		}

		if i < j {
			aArr[i], aArr[j] = aArr[j], aArr[i]
		}

		i++
		j--

	}
	*a = string(aArr)
}
