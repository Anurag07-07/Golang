//////////////Structs///////////////////////////

package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////

//Struct 
type car struct{
	Make string
	Model string
	Height int
	Width int
	FrontWheel Wheel
	BackWheel Wheel
}

type Wheel struct{
	Radius int
	Material string
}

// func main(){
// 	Car := car{}
//   Car.BackWheel.Material = "Steel"
// 	fmt.Println(" the wheel material",Car.BackWheel.Material) 
// }

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Anonymous Struct

type car1 struct{
	name string
	brand struct{
		name string
		model string
	}
}

// func main(){
// 	Carr := struct{
// 		Make string
// 		Model string
// 	}{
// 		Make: "Maclaren",
// 		Model: "P1",
// 	}

// 	fmt.Println(Carr)
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////////

//Embedded Struct

type car2 struct{
	make string
	model string
} 

type truck struct{
	car2
	bedsize int
}

// func main(){
// 	Car := truck{}
// 	Car.make = "Lamborghini"
// 	fmt.Println(Car.make)
// }

////////////////////////////////////////////////////////////////////////////////////////

// Struct methods

type rect struct{
	height int
	width int 
}

func (r rect) area() int {
	return r.width*r.height 
}

// func main(){
// 	r := rect{
// 		width: 10,
// 		height: 10,
// 	}

// 	fmt.Println(`The Area is `,r.area())
// }

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Assignment 

type authenticationInfo struct{
	username string
	password string
}

func (u authenticationInfo) getBasicAuth() string {
	message:= fmt.Sprintf("Basic %s:%s ",u.username,u.password)
	return message 
}

func main(){
	authinfo := authenticationInfo{
		username: "Anurag",
		password: "772002",
	} 
	fmt.Println(authinfo.getBasicAuth())
}