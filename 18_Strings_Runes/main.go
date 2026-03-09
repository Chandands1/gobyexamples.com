package main

import "fmt"

func main(){
	var s string = "Hello this is me"

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[2])

	for index, value := range s{
		fmt.Println(index, value, string(s[index]))
	}

	var r rune = 'A'

	fmt.Println(r)
	fmt.Println(string(r))
}
