package main

import "fmt"

func main(){
    person := map[string] string{
		"name":"rahul",
		"place":"mysore",
	}
	fmt.Println(person)
	fmt.Println(person["place"])
}
