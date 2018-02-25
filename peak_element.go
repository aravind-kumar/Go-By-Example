package main
import "fmt"

func peak(input []int) int {

    low,high := 0,len(input)
    for low <= high {
        mid := low + (high-low)/2
        if mid==0 || input[mid-1]<=input[mid] || mid == high || input[mid+1] <= input[mid] {
            return mid 
        } else if mid < high && input[mid] <= input[mid+1] { 
            low = mid+1 
        } else {
            high = mid-1 
        }
    } 
    return -1   
}



func main() {
   size:=0
   fmt.Println("Enter the size of the array")
   _,err := fmt.Scanf("%d",&size)
   if err == nil {
      fmt.Println("Enter the array");
      input := make([]int,0)
      inputNum:=0
      for size>0 {
          _,err := fmt.Scanf("%d",&inputNum) 
          if err == nil {
              input=append(input,inputNum) 
          }
          size-=1 
      }
      k:=0
      fmt.Println("Enter the k to be seached")
      _,err := fmt.Scanf("%d",&k)
      if err == nil {
         fmt.Printf("The peak index of k is %d",peak(input))
      }
   }
 

}
