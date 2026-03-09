package main

import "fmt"

func main(){

	numbers := map[string]string{"A":"apple", "B":"ball"}

	

	for key, value := range numbers{
		fmt.Println(key, value)
	}

}
