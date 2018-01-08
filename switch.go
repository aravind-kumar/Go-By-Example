package main

import "fmt"

func main() {
	num:=4

	switch num {
	case 1 : fmt.Println("The number is 1")
	case 2 : fmt.Println("The number is 2")
	case 3 : fmt.Println("The number is 3")
	case 4 : fmt.Println("The number is 4")
	default:
		fmt.Println("The number is unknown")
	}

	//One or more cases having the same case;

	num1:=10

	switch num1 {
	case 1,10 : fmt.Println("The number is either 1 or 10")
	default:
		fmt.Println("The number is other than 1 or 10")
	}


	//Switch can have conditionals as well unlike c or c++
	//But care must be taken to ensure that switch is not bound to a variable

	num2:=100

	switch {

	case num2==100 : fmt.Println("The number equals 100")

	case num1+num2<10 : fmt.Print("The sum num1 and num2 is less than 10")

	case num2 < 10 : fmt.Println("The number is a single digit number")

	case num2 > 100 : fmt.Println("The number is greate than 100")

	default:
		fmt.Println("Unkown number")
	}

	//variables can take in functions just like values
	// Values can also be initialized in the switch statment.

	whatis:=func(i interface{}) {
		switch t:=i.(type) {

		case bool:
			fmt.Println("It has boolean type")
		case int:
			fmt.Println("It has integer type")
		default:
			fmt.Println("It it an unknown of type ", t)
		}
	}

	whatis(true)
	whatis("Aravind")
	whatis(5)
}
