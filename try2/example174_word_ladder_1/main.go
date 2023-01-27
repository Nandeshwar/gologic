package main

import (
	"container/list"
	"fmt"
)

type Pair struct {
	word string
	step int
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	/*
		   find shortest transformation from begin to end
		keep changing letter and check if it exists in wordList
		   hit -> hot -> dot -> dog -> cog

		it takes minimum 5 steps to chnage word from "hit" to "cog"

		ouput: 5
	*/

	result := findMinStepsForTransformation(beginWord, endWord, wordList)
	fmt.Println(result)
}

func findMinStepsForTransformation(beginWord, endWord string, wordList []string) int {
	queue := list.New()
	m := make(map[string]struct{})
	for _, v := range wordList {
		m[v] = struct{}{}
	}

	// init queue with beginWord and frequency
	queue.PushBack(Pair{beginWord, 1})

	// Remove begin word from map if exists
	delete(m, beginWord)

	for queue.Len() != 0 {
		element := queue.Remove(queue.Front())
		item := element.(Pair)
		word := item.word
		steps := item.step

		wordBytes := []byte(word)

		if word == endWord {
			return steps
		}

		for i, _ := range wordBytes {
			original := wordBytes[i]
			for ch := 'a'; ch <= 'z'; ch++ {
				wordBytes[i] = byte(ch)

				_, ok := m[string(wordBytes)]
				if ok {
					delete(m, string(wordBytes))
					queue.PushBack(Pair{string(wordBytes), steps + 1})
				}
			}
			wordBytes[i] = original
		}
	}

	return 0
}
