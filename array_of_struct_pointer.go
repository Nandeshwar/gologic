package main

import "fmt"

type Employee struct {
  id int
  name string
}

func main() {
  emp := Employee {
    id: 1, 
    name: "nks",
  }
  fmt.Println(emp)
  fmt.Println("Hello World")
  m := myFun()
  
  fmt.Println(*m[0])
  fmt.Println(*m[1])
}

func myFun() []*Employee {
  /*
  employees := []*Employee {
    {id: 1, name: "nks"},
    {id: 2, name: "nks2"},
  }
  
  return employees
  */
  
  // The given below block till return is samve as above
  emp1 := &Employee {
    id : 1, 
    name: "nks",
  }
  
  emp2 := &Employee {
    id : 2, 
    name: "nks2",
  }
  
  employees2 := []*Employee {emp1, emp2}
  return employees2
  
}
