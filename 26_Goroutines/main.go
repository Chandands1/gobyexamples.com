package main 


import ("fmt"
"time")

func printMessage(){
	fmt.Println("This is from the goroutien")
}
func task(name string){
	fmt.Println("Name: ", name)
}
func main(){
	go printMessage()
	time.Sleep(time.Second)
	fmt.Println("This is from the main")

  go	task("A")
go 	task("B")
	go task("C")
go 	task("D")

	time.Sleep(time.Second)
}