package main

import (
	"fmt"
	"sort"
	"time"
)

type Emp struct {
	id        int
	name      string
	createdAt time.Time
}

// 3rd algorithm : custom logic

type E []Emp

func (e E) Len() int           { return len(e) }
func (e E) Less(i, j int) bool { return e[i].id < e[j].id }
func (e E) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func main() {
	empList := []Emp{
		Emp{2, "Ram", time.Now()},
		Emp{3, "Radha", time.Now()},
		Emp{1, "Sita", time.Now()},
		Emp{3, "Lalita", time.Now()},
		Emp{4, "Krishna", time.Now()},
		Emp{3, "Visakha", time.Now()},
	}

	sortEmpAlgo(empList)
	fmt.Println(empList)

	a := []int{10, 20, 30, 40, 50}
	fmt.Println("sort in reverse")
	fmt.Println(sort.IntSlice(a))
	sort.Ints(a)
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}

func sortEmpAlgo(empList []Emp) {
	// sort.Slice(empList, func(i, j int) bool {
	// 	return empList[i].id < empList[j].id
	// })

	// sort.SliceStable(empList, func(i, j int) bool {
	// 	return empList[i].id < empList[j].id
	// })

	sort.Sort(E(empList))

}
