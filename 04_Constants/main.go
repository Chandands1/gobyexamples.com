package main

import "fmt"

func main(){
	const pi = 3.14
	const appName = "MyApp"

	fmt.Println(pi)
	fmt.Println(appName)

	//group constant
	const(
		name = "Virat"
		age = 32
		place = "London"
	)
	fmt.Println(name,age, place)
}