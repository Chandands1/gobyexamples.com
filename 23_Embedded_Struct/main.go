package main

import "fmt"

type Person struct{
	name string
	age int
}

type Employee struct{
	Person
	id int
}


func main(){

	e := Employee{
		Person: Person{name: "KIshor", age: 19},
		id:28,
	}
	
	fmt.Println(e)
	
}
