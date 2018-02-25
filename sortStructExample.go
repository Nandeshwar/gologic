package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	id int
	name string
	salary float32
}

type byName []Employee

func main() {
	empList := []Employee{
		Employee{
			id: 1,
			name: "Nks",
			salary: 200,
		},
		Employee{
			id: 2,
			name: "Sneha",
			salary: 300,
		},
		Employee{
			id: 4,
			name: "Spandan",
			salary: 400,
		},
		Employee{
			id: 5,
			name: "Shruti",
			salary: 500,
		},
		Employee{
			id: 6,
			name: "Krishna",
			salary: 600,
		},
		Employee{
			id: 7,
			name: "Ram",
			salary: 700,
		},
	}

	fmt.Println("Before Sorting")
	fmt.Println(empList)
	fmt.Println("After Sorrting")
	sort.Sort(byName(empList))
	fmt.Println(empList)
}

func (b byName) Len() int           { return len(b) }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byName) Less(i, j int) bool {
	if b[i].name < b[j].name {
		return true
	}
	if b[i].name > b[j].name {
		return false
	}
	return b[i].name < b[j].name
}



