package main

import (
	"fmt"
)


type Employee struct {
	id int
	name string
}


func (e Employee) String() string {
        return fmt.Sprintf("%d %s", e.id, e.name)
}



func main() {
	fmt.Println("Hello, playground")
	e1 := getAllEmployee1()
	fmt.Println(*e1)
	
	fmt.Println("________")
	fmt.Println(getAllEmployee())
}

func getAllEmployee1() *Employee {
	e1 := &Employee {
	id : 1, 
	name: "Ram",
	}
	
	return e1
}


func getAllEmployee() []*Employee {
	e1 := &Employee {
	id : 1, 
	name: "Ram",
	}
	
	e2 := &Employee {
	id : 2, 
	name: "Krishna",
	}
	
	employees := []*Employee{e1, e2}
	return employees
}

/*  
   output: 
   Hello, playground
1 Ram
________
[1 Ram 2 Krishna]
*/
