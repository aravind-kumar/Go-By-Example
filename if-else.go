package main

import "fmt"

func main()  {

	if num:=0 ; num == 0 {
		fmt.Println("The number is equal to 0")
	}

	value := 100

	if value> 100 {
		fmt.Println("The value is greater than 100" ,value)
	}else if value < 10 {
		fmt.Println("The value is a single digit ", value)
	} else {
		fmt.Println("The value is between 10-100 ", value)
	}
}
