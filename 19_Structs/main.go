package main

import "fmt"

type Person struct{
	name string
	age int
	place string
}

type Login struct{
	userName string
	password string
}

func printPerson(p Person){
	fmt.Println(p.age)
}

func main(){
    var person1 Person

	person1.name = "Rahul"
    person1.age = 18
	person1.place = "gurgam"

	fmt.Println(person1)
	fmt.Println(person1.age)

	user1 := Login{
		userName: "hello",
		password :"Hello123",
	}
	fmt.Println(user1)

	printPerson(person1)


}