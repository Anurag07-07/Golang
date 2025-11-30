////////////////////////////Interfaces/////////////////////////////

// Interafaces are collections of methods signatures.A type "implements" an interface
// if it has all of the methods of the given interface defined on it 

package main

import "fmt"

type shape interface{
	area() float64
	peri() float64
}

type rect struct{
	height,width float64
}

func (r rect) area() float64 {
	return  r.width*r.height
}

func (r rect) pari() float64{
	return  2*r.width+2*r.height
}

func main(){
		Rect := rect{}
		Rect.width = 2.56
		Rect.height = 5.63
		fmt.Println("The area is ",Rect.area())	
		fmt.Println("The Perimeter is ",Rect.pari())	
}