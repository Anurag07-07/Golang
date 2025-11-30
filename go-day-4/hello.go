////////////////////////////Interfaces/////////////////////////////

// Interafaces are collections of methods signatures.A type "implements" an interface
// if it has all of the methods of the given interface defined on it

package main

import (
	"fmt"
)

type shape interface{
	area() float64
	pari() float64
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

// func main(){
// 		Rect := rect{}
// 		Rect.width = 2.56
// 		Rect.height = 5.63

// 		//Using interaface
// 		var s shape = Rect    
// 		fmt.Println("The Area is ",s.area())
// 		fmt.Println("The Perimeter is",s.pari())
// }

/////////////////////////////////////////////Type Assertions/////////////////////////////////////////////////////////////////////////

type shape1 interface {
	area1() float64
}

type circle struct {
	radius float64
}

// Method for circle
func (c circle) area1() float64 {
	return 3.14 * c.radius * c.radius  
}

// func main() {
// 	Circle := circle{radius: 5}

// 	var s shape1 = Circle

// 	fmt.Println("Area:", s.area1())

// 	// Type Assertion (interface -> concrete type)
// 	c, ok := s.(circle)
// 	if ok {
// 		fmt.Println("Radius is:", c.radius)
// 	} else {
// 		fmt.Println("Type assertion failed")
// 	}
// }

/////////////////////////////////////Switch Type///////////////////////////////////////////

// A type switch makes it easy to do several type assertions in a series 


func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("Type: %T, Value: %v\n", v, v)

	case string:
		fmt.Printf("Type: %T, Value: %v\n", v, v)

	case float64:
		fmt.Printf("Type: %T, Value: %v\n", v, v)

	default:
		fmt.Printf("Unknown Type: %T, Value: %v\n", v, v)
	}
}

func main() {
	printNumericValue(1)
	printNumericValue("Hello")
	printNumericValue(3.14)
	printNumericValue(true)
}
