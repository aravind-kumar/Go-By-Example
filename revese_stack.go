package main
import "fmt"


func reverse(stack *[]int) {
    if len(*stack) > 0 {
       top:=(*stack)[len(*stack)-1] 
       *stack=(*stack)[:len(*stack)-1]
       reverse(stack)
       *stack=append(*stack,top) 
    }
}

func main() {

   num:=0
   fmt.Println("Enter the size of the stack")
   _,err := fmt.Scanf("%d",&num) 
   if err == nil {
      inputNum:=0

      stack:=make([]int,0)

      for num>0  {
          fmt.Println("Enter the number")
          _,err := fmt.Scanf("%d",&inputNum)
          if err == nil {
              stack=append(stack,inputNum)              
          }
          num-=1
      }
      fmt.Printf("The stack is %v",stack)
      reverse(&stack)
      for num := range stack {
          fmt.Println(num) 
      } 
   }

}
