package main

import "fmt"

type Person struct{
	name string
	age int
	
}
func (p Person) greet(city string){
	fmt.Println("Welcome to you", p.name, "Your age:",p.age, "place: ", city)

}
func(p *Person) birthday(){
	p.age++
}

func main(){
	person1 := Person{
		name: "Rahul",
		age: 19,
	}
	person1.greet("mumbai")
person1.birthday()
fmt.Println(person1.age)	

}