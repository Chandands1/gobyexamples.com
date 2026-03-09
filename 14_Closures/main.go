package main

import "fmt"

func counter() func() int{
	i:=0

	return func()int{
		i++
		return i
	}
}

func main(){
	next := counter()
	fmt.Println(next())
	fmt.Println(next())
}