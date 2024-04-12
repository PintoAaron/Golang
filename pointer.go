package main

import "fmt"

func change_pointer_value_to_10(ptr *int) {
	*ptr = 10
}

func main() {

	numb := 50
	change_pointer_value_to_10(&numb)
	fmt.Println(numb)
	fmt.Println(&numb)

}
