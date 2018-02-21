package main
import "fmt"
import "strconv"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var input int
    _,err := fmt.Scanf("%d",&input)
    if err == nil {
        for i:=1 ; i<=input; i+=1 {
            //output:=""
            if i%3 == 0 && i%5 == 0 {
                fmt.Println("FizzBuzz")                
            } else if i%3 == 0 {
                fmt.Println("Fizz")
            } else if i%5 == 0 {
                fmt.Println("Buzz")
            } else {
                fmt.Println(strconv.Itoa(i))
            }
        }
    }
    
}
