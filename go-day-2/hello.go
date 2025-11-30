package main

import (
	"errors"
	"fmt"
)

func sum(a int,b int) int {  
	return a+b
}

func sum1(a,b int) int {  //if both have same argument type give type to the last 
	return a+b
}


func incrementSends(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}


//Ignore a return value
func getNames()(string,string){
	return  "Anurag","Raj"
}


//Named Return Type//////////////////////////////////////////////////////

func getCoords()(int , int){
	var x int
	var y int
	return x,y
}

// Another way to write same function 

func getCoords1()(x,y int){
	return 
}

//It is good if you use like that

func getCoords2()(x int,y int){
	return x,y
}

func divide(dividend,divisor int) (int,error){
	if divisor == 0 {
		return  0, errors.New("Can't Divide")
	}

	return  dividend/divisor , nil
} 


func main() {
		// sendsSoFar := 430
		// sendsToAdd := 25

		// sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)

		// fmt.Println("You've sent", sendsSoFar, "messages")

		firstname,_ := getNames() 
		fmt.Println(firstname)

		ans,err := divide(45,0);
		if ans == 0 {
			fmt.Println(err)
		}else{
			fmt.Println(ans)
		}
}

//go run filename to run the file
//go mod init dirname to intialize the file
