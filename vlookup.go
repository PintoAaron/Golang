package main



import "fmt"


type array2D struct{
	data [][]int
}


func max(slice []int) int{
	max := slice[0]
	for _,x := range slice{
		if x > max{
			max = x
		}
	}
	return max
}



func (arr array2D) vlookup(value int,column int) int{
	var arr0 []int
	

	for _,row := range (arr.data){
		for index,number := range (row){
			if index == 0 && number <= value{
				arr0 = append(arr0, number)
			}
		}	
	}
	max_value := max(arr0)

	for _,numb := range arr.data{
		if numb[0] == max_value{
			final_value := numb[column - 1]
			return final_value
		}
	}

	return -1

}


func main(){


	arr := array2D{data: [][]int{{670,20,3},{749,77,1},{723,12,7}}}

	s := arr.vlookup(750,2)

	fmt.Println(s)

}