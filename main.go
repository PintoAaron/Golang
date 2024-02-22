package main

import "fmt"

func main() {
	var username string
	var email string
	var food string
	//unit stands for unsigned integer, so in case of integer it will be just int
	var banku uint = 15
	var rice int = 6
	var kenkey int = 20
	fmt.Println("Welcome To Quantum Kitchen")
	fmt.Println("Hello, I will like to know some few info about you")
	fmt.Println("Please me I know your name: ")
	fmt.Scan(&username)
	fmt.Println("How about your email: ")
	fmt.Scan(&email)
	fmt.Println("We Only Have these foods available")
	fmt.Printf("Banku   available: %v\n", banku)
	fmt.Printf("Rice   available: %v\n", rice)
	fmt.Printf("Kenkey   available: %v\n", kenkey)
	fmt.Println("What will you like to eat today: ")
	fmt.Scan(&food)
	fmt.Printf("Thank you %v, we are processing your food. We will notify you when your %v is ready.\n",username,food)

}
