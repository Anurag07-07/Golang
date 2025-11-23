package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello This is Golang")

	var name string = "Anurag"
	var position string = "CEO"

	var companysize int = 4

	fmt.Println(`Hello myself `+name+" i am a "+position+" of Edunax company and our company size is ",companysize) //int can't concat with string


	// Declaring a Variable
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermisson bool
	var username string
	fmt.Println("%v %f %v %q\n",smsSendingLimit,costPerSMS,hasPermisson,username)


	// Short Variable Declaration 
	congrats := "Best day is Coming"
	fmt.Println(congrats)

	pennisPerText:=45.56
	fmt.Println("The type of pennisPerText is ",pennisPerText)

	//Declaring multiple value in same line
	username,role:="Anurag","Full stack Developer"
	fmt.Println("Hello myself "+username+" and i am a "+role)

	//Typecasting
	accountFloat:=94565.456
	accountInt:=int(accountFloat)
	fmt.Println("The Account balance is ",accountInt,"$")

	//Print number of second in hrs
	const secondsInMinute = 60
	const	minutesInHour = 60
	const secondsInHour = secondsInMinute*minutesInHour
	fmt.Println("Number of seconds in an hour:",secondsInHour)

	//Interplotation
	const name1 = "Anurag Raj"
	const openrate = 30.5
	msg:=fmt.Sprintf("Hi %s, your open rate is %.2f percent",name1,openrate)

	fmt.Println(msg)

	//Conditionals
	const height int = 5
	if height > 6{
		fmt.Println("You are super tall!")
	}else if height > 4 {
		fmt.Println("You are tall Enough")
	}else{
			fmt.Println("You are not tall Enough")
	}

	//Shorten Code in if
  // email := "anurag07raj@gmail.com"
  email := ""
	if length:= len(email);length<1 {  //This type of syntax is used to make code concise and the scope limited
		fmt.Println("Email is invalid")
		}else{
		fmt.Println("Email is Good")
	}
}

//go run filename.go to run the code
//go build to build the app
//go is strongly and static type language
//Nibbles is 4 bits
//Byte is 8 bits