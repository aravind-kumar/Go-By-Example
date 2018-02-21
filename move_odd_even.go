package main

import "fmt"
import "strconv"

func swap(num1,num2 *int) {
    var temp int = *num1
    *num1 = *num2
    *num2 = temp
}

func partition(inputArray *[]int) {

     wall:=-1
     for i:=0 ; i<len(*inputArray) ; i+=1 {
         if (*inputArray)[i] % 2 == 0 {
             wall+=1 
             swap(&((*inputArray)[wall]),&((*inputArray)[i]))    
         }
     }
}

func printArray(inputArray []int) {
    
    for _,num := range inputArray {
        fmt.Println(num)
    }
}

func main() {
   
    var num int
    fmt.Println("Enter the size of the input")
    _,err := fmt.Scanf("%d",&num)
    if err == nil {
        inputArray := make([]int,0,num)
        var inputNum int
        fmt.Println("Enter the array")

        for i:=0 ; i<num ; i+=1 {
            fmt.Scanf("%d",&inputNum)
            inputArray = append(inputArray,inputNum) 
        }
        partition(&inputArray)
        printArray(inputArray)
    } 


}
