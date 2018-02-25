package main
import "fmt"


func indexEqual(input []int) int {

   low,high := 0,len(input)-1
   for low<=high {
       mid := low +(high-low)/2
       switch {
           case mid == input[mid]:
                return mid 
           case mid < input[mid]:
                high = mid-1
           case mid > input[mid]:
                fmt.Println("Right")
                low = mid+1
       }
   } 
   return -1
}

func main() {

   num:=0
   fmt.Println("Enter the number of elements")
   _,err := fmt.Scanf("%d",&num)
   if err == nil {
       fmt.Println("Enter the elements")
       input := make([]int,0)
       inputNum :=0
       for num>0  {
           _,err := fmt.Scanf("%d",&inputNum)
           if err == nil {
              input = append(input,inputNum)
           }
           num-=1
       }
       fmt.Println("The index whoes value is equal to index is %d",indexEqual(input)) 
   }

}
