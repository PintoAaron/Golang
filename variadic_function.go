package main 


import "fmt"


func sum_all(nums...int) int{
	total := 0
	for _,num := range nums{

		total += num
	}

	return total
}



func longest_word(words...string) (string,int){
	temp := ""
	for _,word := range words{
		if len(word) > len(temp){
			temp = word
		}
	}

	return temp,len(temp)
}


func main(){

	res := sum_all(4,5,2,1,3,3,4,21,2,5,3,2,1,3,8,9,)
	fmt.Printf("Total is %d\n",res)


	word,word_length := longest_word("Master","encyclopedia","opotptot","givatheboyback","macquena","seriousuy","yakaukukukuuuuuuuueuueueueueue")
	fmt.Printf("The longest word is %s with %d characters.",word,word_length)


	
}