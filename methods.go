package main


import "fmt"



type boy struct{
	name string
	age int
}


func (b *boy) check_validity() bool {

	if b.age < 18 {
		return false
	} else {
		return true
	
	}
}

func main(){

	var name string
	var age int

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	s := boy{name,age}
	is_valid := s.check_validity()
	if is_valid {
		fmt.Println("Valid, can vote")
	} else {
		fmt.Println("Not valid, cannot vote")
	}


}