package main

import (
    "fmt"
    // "slices"
	"math/rand"
    "time"
)

func main() {

	magic_scores := []int{0,10,20,30,40,50,60,70,80,90,100}
    var trials = 0
	var chances = 3
    var guess int
	rand.Seed(time.Now().UnixNano())

	for trials < chances{
		index := rand.Intn(len(magic_scores))
		random_value := magic_scores[index]
		fmt.Println("What is the magic score: ")
		fmt.Scan(&guess)
		if guess == random_value {
			fmt.Printf("Hidden value is %d\n", random_value)
			fmt.Println("Well Done,You got it right")
			break
		} else {
			fmt.Printf("Hidden value is %d\n", random_value)
			fmt.Println("Try again")
		}
		
		trials = trials + 1
	}
	fmt.Println("Game over")



}