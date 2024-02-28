package main

import "fmt"


// function takes a slice and finds the standard deviation of the numbers
func find_standard_deviation(numbers []int) int{
	mean := 0
	for _, number := range numbers{
		mean += number
	}
	mean = mean/len(numbers)

	deviation := 0
	for _, number := range numbers{
		deviation += (number - mean) * (number - mean)
	}
	deviation = deviation/len(numbers)
	return deviation
}



// Multiple return Values
func name_and_score(name string) (string,int){
	exams_score := make(map[string]int)
	exams_score["english"]=57
	exams_score["maths"]=90
	exams_score["science"]=88
	exams_score["social"]=70

	score := exams_score[name]

	return name,score

}

func main(){
	scores := []int{80,30,60,50,88,29,91,17,85,68}
	sd := find_standard_deviation(scores)
	fmt.Println("The final solution is ",sd)


	subject,score := name_and_score("maths")
	fmt.Printf("In %s the student scored %d\n",subject,score)
}