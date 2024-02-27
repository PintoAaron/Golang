package main

import "fmt"

func main(){

	my_slice := []int{1,2,3,4,5,6,7,8,9,10}
	my_map := make(map[string]int)
	my_map["one"] = 1
	my_map["two"] = 2 

	sum := 0
	//used to iterate over a slice or map
	for _, value := range my_slice{
		sum += value
	}
	fmt.Println("Sum of my_slice: ",sum)

	for k, v := range my_map{
		fmt.Printf("Key: %s, Value: %d\n",k,v)
	}

	for index,character := range "PINTOAARONTETTEH" {
		fmt.Printf("Character: %c, Index: %d\n",character,index)
	}
}