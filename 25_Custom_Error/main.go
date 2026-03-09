package main 

import "fmt"

type Myerror struct{
	message string
}

func (e Myerror) Error() string{
	return e.message
}

func checkAge(age int) error{
	if age < 18{
		return Myerror{"age is less than 18"}
	}
	return nil
}


func main(){
	err := checkAge(16)

	if err != nil{
         fmt.Println(err)
	}
}