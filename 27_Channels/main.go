package main

import "fmt"


func main(){

	ch := make(chan string)
	cha := make(chan int,2 )

	go func(){
		ch <- "this is from the channel"
	}()

	message := <- ch

	fmt.Println(message)
	cha <- 1
	cha <- 2
	fmt.Println(<- cha)
	fmt.Println(<- cha)
}
