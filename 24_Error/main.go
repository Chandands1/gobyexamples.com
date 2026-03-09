package main

import ("fmt"
"errors")


func divide(a int, b int)(int, error){
	if b==0{
		return 0, errors.New("cannot divide by zero")
	}
	return a/b, nil
}



func main(){
	res , err := divide(10,0)
	if err != nil{
		fmt.Println("Error:", err)
		return
	}else{
		fmt.Println(res)
	}
}
