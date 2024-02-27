package main

import "fmt"


func main(){
	exams_scores := make(map[string]int)
	exams_scores["maths"] = 90
	exams_scores["science"] = 80
	exams_scores["english"] = 70
	fmt.Println(len(exams_scores))
	fmt.Println(exams_scores)
	

	delete(exams_scores,"english")
	fmt.Println(exams_scores)

	clear(exams_scores)
	fmt.Println(exams_scores)

	fmt.Println(len(exams_scores))
}