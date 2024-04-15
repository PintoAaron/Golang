package main

import (
	"fmt"
	"math"
)


type shapes interface {
	
	area() float64
	perimeter() float64

}

type square struct {
	side float64
}

type rect struct{
	width float64
	height float64
}

type circle struct{
	radius float64
}



func (s square) area() float64{
	return s.side * s.side
}

func (s square) perimeter() float64{
	return 4 * s.side
}



func (r rect) area() float64{
	return r.height * r.width
}

func (r rect) perimeter() float64{
	return 2 * r.height + 2 * r.width
}


func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}


func measure_area_and_perimeter(s shapes){
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	fmt.Println("\n\n")
}



func main(){

	sqaure1 := square{side: 5.4}
	rect1 := rect{width: 6.6, height:4.7 }
	circle1 := circle{radius: 3.2}

	measure_area_and_perimeter(sqaure1)
	measure_area_and_perimeter(rect1)
	measure_area_and_perimeter(circle1)


}