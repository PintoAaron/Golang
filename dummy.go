package main

import (
	"fmt"
	"strconv"
	"strings"
)


func odd_zeroes(numbs []int) int{
	temp := 0
	for _,x := range numbs{
		s := strconv.Itoa(x)
		if strings.Count(s,"0") % 2 != 0{
			temp ++
		}
	}
	
	return temp
} 


func main(){

	s := []int{300,4000,50}

	
	fmt.Println(odd_zeroes(s))

}