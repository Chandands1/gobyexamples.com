package main

import "fmt"

type Animal interface{
	speak()
}

type Dog struct{}
type Cat struct{}

func (d Dog)speak(){
	fmt.Println("Dog Barks")

}
func (c Cat)speak(){
	fmt.Println("Cat Meow")
}

func main(){
	var a Animal
	
	a = Dog{}
	a.speak()

	a = Cat{}
	a.speak()

}