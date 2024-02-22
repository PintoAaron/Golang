package main

import "fmt"

const magic_number = 7

func main(){


	fmt.Println("YOU HAVE 3 CHANCES TO GUESS THE MAGIC NUMBER BETWEEN 1 TO 10\n")

	var trials = 0
	var guess int
	const chances = 3

	for trials < chances {
		fmt.Println("Guess: ")
		fmt.Scan(&guess)
		if guess == magic_number {

			fmt.Printf("Well done, you got it right..our magic number was %v\n",magic_number)
			break
		}
		trials = trials + 1
	}

	if trials == 3{
		fmt.Println("\nooops!, nice try...come back tomorrow and try again.\n")
	}


	
}