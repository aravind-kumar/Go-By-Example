package main
import "fmt"
import "math"


func min(a,b int) int {
    if a<b {
       return a 
    } else {
       return b 
    }
}

func minSubArrayLen(s int, nums []int) int {
    
    start,end,currentSum,minLen := 0,0,0,math.MaxInt32
    
    for end < len(nums) {
        
        currentSum += nums[end]
        
        for currentSum >= s {
            minLen = min(minLen,end-start+1)
            currentSum -= nums[start]
            start+=1
        }
        end+=1
    }
    
    if minLen == math.MaxInt32 {
        return 0
    }
    return minLen
}


func main() {



}
