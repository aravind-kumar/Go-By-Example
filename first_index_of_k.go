package main
import "fmt"

func bst(input []int,k int) int {

    low,high,res := 0,len(input),-1

    for low<=high {
        mid:= low + ((high-low)/2)
        switch {
            case input[mid] > k: 
                high=mid-1
            case input[mid] == k:
                res,high = mid,mid-1
            case input[mid] < k: 
                low=mid+1 
        }
    }
    return res
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
         fmt.Printf("The first index of k is %d",bst(input,k))
      }
   }
   

}
