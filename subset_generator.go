import "fmt"

// TODO: This line has an error the pass by reffernce overwrites previous solutions 
// HAVE TO FIX.
func subsetsHelper(nums,intermediate []int,ans *[][]int,start int) {
    

    fmt.Printf("\n Pre %v", intermediate)
    *ans = append(*ans,intermediate)
    fmt.Printf("\n Post %v", intermediate)

    for i:= start ; i<len(nums) ; i+=1 {
        intermediate = append(intermediate,nums[i])
      
        subsetsHelper(nums,intermediate,ans,i+1)

        intermediate = intermediate[0:len(intermediate)-1]
    }
}

func subsets(nums []int) [][]int {
    
    intermediate := make([]int,0)
    ans := make([][]int,0)
    start := 0
    
    subsetsHelper(nums,intermediate,&ans,start)
    return ans
    
}


func main() {

  input := 0
  fmt.Println("Enter the number of inputs")
  _,err := fmt.Scanf("%d",&input)
  if err == nil {
    
    inputArray := make([]int,input)
    
    fmt.Println("Enter the input array")
    
    for i:=0 ;i<input ; i+=1 {
      givenInput := 0
      _,err := fmt.Scanf("$d",&givenInput)
      
      if err == nil {
        input[i] = givenInput
      }
    }
    subsets(input)
  }

}
