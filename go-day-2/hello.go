package main

import "fmt"

func sum(a int,b int) int {
	return a+b
}

func main(){
	var ans int = sum(45,56)
	fmt.Println(ans)
}

//go run filename to run the file
//go mod init dirname to intialize the file