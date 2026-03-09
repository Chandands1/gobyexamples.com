package main

import "fmt"

type Day int

type Status int

const(
	PENDING Status = iota
	APPROVED
	ACCEPTED


)

const(
	Monday Day = iota
	Tuesday
	Wednesday

	Thursday
	Friday
	Saturday
	Sunday

)

func main(){
	today := Wednesday
	fmt.Println(today)

	status := PENDING
	fmt.Println(status)


}
