package main 

import "fmt" 

func main(){
	var maths_score int
	var science_score int
	fmt.Println("What did you score in maths: ")
	fmt.Scan(&maths_score)
	fmt.Println("What did you score in science: ")
	fmt.Scan(&science_score)
	switch maths_score + science_score {
	case 200:
		fmt.Println("This is Impossible, you got the answers")
	case 190:
		fmt.Println("Bravo, You're brilliant ")
	case 180:
		fmt.Println("You are good, keep it up")
	case 100:
		fmt.Println("Average, sit up guy")
	case 20:
		fmt.Println("Did you prepare for the exams")
	case 10:
		fmt.Println("I'm not ssure you have been coming to classes")
	case 0:
		fmt.Println("Bro, you don't exit")
	default:
		fmt.Println("Not bad")
	}
}