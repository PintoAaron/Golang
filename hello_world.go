package main 


import "fmt"

func main(){
	var name string
	fmt.Println("Hello")
	fmt.Println("whats your name: ")
	fmt.Scan(&name)
	fmt.Printf("You are welcome %v\n",name)
	fmt.Printf("%v, we are about to do something craxy with golang\n",name)
}