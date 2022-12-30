package main

// input: [["JFK", "SFO"], ["JFK", "ATL"], ["SFO", "ATL"],["ATL", "JFK"], ["ATL", "SFO"]]
// output: ["JFK", "ATL", "JFK", "SFO", "ATL", "SFO"]

/*
	JFK ---------> SFO
	 |  ^\  |--->  /
	 |	  |	|	  /
	 |	  |	|	 /
	 |	  |	|	/
	  ->	 ATL <-

	solution:
	1. store all vertex in map
	2. key: vertex, value: sorted list of destination
	3. dfs() reach to end, delete edge and add to stack
	4. pop from stack

*/

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {

	flights := [][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}

	origin := "JFK"
	graph := createGraph(flights)
	fmt.Println("graph=", graph)
	result := findItinerary(graph, origin)
	fmt.Println("Result=", result)

	/*
		output:
		graph= map[ATL:[JFK SFO] JFK:[ATL SFO] SFO:[ATL]]
		Result= [JFK ATL JFK SFO ATL SFO]
	*/
}

func createGraph(flights [][]string) map[string][]string {
	graph := make(map[string][]string)

	for i := 0; i < len(flights); i++ {
		v, ok := graph[flights[i][0]]
		if ok {
			v = append(v, flights[i][1])
			sort.Strings(v)
			graph[flights[i][0]] = v

		} else {
			graph[flights[i][0]] = []string{flights[i][1]}
		}
	}
	return graph
}

func findItinerary(graph map[string][]string, origin string) []string {
	var result []string
	stack := list.New()
	dfs(graph, origin, stack)

	for stack.Len() != 0 {
		element := stack.Remove(stack.Back())
		result = append(result, element.(string))
	}
	return result
}

func dfs(graph map[string][]string, origin string, stack *list.List) {
	l, ok := graph[origin]
	if ok {
		v := l[0]
		var newList []string
		if len(l) > 1 {
			newList = l[1:len(l)]
		}
		if len(newList) >= 1 {
			graph[origin] = newList
		} else {
			delete(graph, origin)
		}
		dfs(graph, v, stack)
	}

	stack.PushBack(origin)
}
