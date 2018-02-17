package student

import "fmt"

type Student struct {
	id int
	name string
}

func( student Student) GetInfo(id int, name string){
	student.id = id
	student.name = name
}

func( student Student) DisplayInfo(){
	fmt.Println(student.id)
	fmt.Println(student.name)
}

func New(id int, name string) Student{
	return Student{
		id : id,
		name :name,
	}
}
