package main

import "fmt"


func sum(number ...int) int{
	total := 0
	for _, num := range number{
		total = total + num
	}
	return total
}

func main(){
	result1 := sum(2,3,4,52)
	result2 := sum(43,23,1)
	result3 := sum(10)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

}