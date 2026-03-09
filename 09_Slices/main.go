package main

import "fmt"

func main(){
	numbers := []int{10,20,30}
	fmt.Println(numbers)
	numbers = append(numbers, 40)
	fmt.Println(numbers)
	for i :=0 ; i < len(numbers); i++{
		fmt.Println(numbers[i])
	}
}