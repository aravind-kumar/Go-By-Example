package main
import "fmt"

func adder(num1 int) (func(int) int) {
    return func (num2 int) int {
           return num1+num2 
         }
}

func main() {
   add3 := adder(3)
   fmt.Println("The value of 3 + 5 is ", add3(5))
}
