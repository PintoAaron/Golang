package main


import "fmt"


type Order struct{
	Order_id int
	Order_date string
	Order_amount float64
	Order_status string
	customer_id int
}


func create_order(order_id int,order_date string,order_amount float64,order_status string,customer_id int) Order{
	return Order{order_id,order_date,order_amount,order_status,customer_id}
}


func main(){

	o := create_order(1,"2020-01-01",1000.00,"Pending",1)
	fmt.Println(o)


}