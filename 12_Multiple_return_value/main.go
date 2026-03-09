package main

import "fmt"

func division(a int, b int)( remainder int, quoetient int){
	remainder = a%b
	quoetient = a/b
	return 
}

func main(){
  remainder, quotient := division(10,3)
  fmt.Println("Remainder: ", remainder)
  fmt.Println("Quotient: ", quotient)

}