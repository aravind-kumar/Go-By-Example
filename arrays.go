package main

import "fmt"

func main() {

	var numbers [5]int

	fmt.Println("The numbers are ",numbers)

	numbers[0]=0
	numbers[1]=1
	numbers[2]=2
	numbers[3]=3
	numbers[4]=4

	fmt.Println("The numbers are ",numbers)


	newNumbers:= [5]int {2,3,4,5,6}

	fmt.Println("the newnumber are ",newNumbers)

	twoD := [2][3] int {{1,2},{3,4}}

	fmt.Println("The 2d array is ", twoD)

}
