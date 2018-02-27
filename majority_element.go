package main
import "fmt"

func getMajority(nums []int) int {

    majority,count := 0,0
    for _,num := range nums {
        if num == majority {
           count+=1 
        } else if count == 0 {
           majority = num 
           count = 1
        } else {
           count-=1 
        }
    }
    return majority
}

func main() {
   num:=0
   fmt.Println("Enter the size of the array")
   _,err := fmt.Scanf("%d",&num)
   if err == nil {
      input := make([]int,0)
      inputNum := 0
      for num>0 {
         if _,err := fmt.Scanf("%d",&inputNum) ; err == nil {
              input=append(input,inputNum)
         }
         num-=1 
      }
      fmt.Printf("The majority element is %d",getMajority(input)) 
   }
}
