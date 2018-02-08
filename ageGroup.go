package main

import "fmt"

func main(){
	fmt.Println(ageGroup(44))
}

func ageGroup(age int) (ageGroup string){
	if age > 0 && age < 13 {
		ageGroup = "Child"
	} else if age > 12 && age < 20 {
		ageGroup = "teen"
	} else if age > 19 && age < 45 {
		ageGroup = "Adult"
	} else if age > 44 && age < 125 {
		ageGroup = "old"
	} else {
		ageGroup = "Invalid"
	}
	return ageGroup
}
