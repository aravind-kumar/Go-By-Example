package main

import "fmt"

func main()  {

	i:=1
	sum:=0
	// single line statements also need braces dont no why.
	for i<10 {
		sum+=i
		i=i+1
	}

	fmt.Println("The sum of the first 10 numbers is ", sum)

	sum=0
	for i=0 ; i<10; i++ {
		sum+=i
	}

	fmt.Println("The sum of the first 10 numbers using a different loop is ", sum)

	sum=0
	for i=0 ; i<10 ; {
		sum+=i
		i++
	}
	fmt.Println("The sum of the first 10 numbers using another type is ", sum)

}
