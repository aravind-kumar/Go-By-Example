package main

func getSrcIndex(nums []int) int {
    
    for i:= len(nums)-1; i>=1 ; i-=1 {
        if nums[i-1] < nums[i] {
            return i-1
        }
    }
    return -1
}

func getDstIndex(nums []int,startIndex int) int {
    
    if startIndex >= 0 {
         var dstIndex,src int = -1,nums[startIndex]
    
        for i:=startIndex ; i<len(nums) ; i+=1 {
            if nums[i]>src {
              dstIndex = i
            }
        }
        return dstIndex     
    }
    return -1
   
}

func reverse(nums *[]int,start int) {
    
    for i,j := start,len(*nums)-1 ; i<j ; i,j = i+1,j-1 {
        (*nums)[i],(*nums)[j] = (*nums)[j],(*nums)[i]
    }  
}

func nextPermutation(nums []int)  {
    
    srcIndex := getSrcIndex(nums)
    
    dstIndex := getDstIndex(nums,srcIndex)
    
    
    if srcIndex != -1 && dstIndex != -1 {
        nums[srcIndex],nums[dstIndex] = nums[dstIndex],nums[srcIndex]
        reverse(&nums,srcIndex+1)
    }else {
        reverse(&nums,0)
    }
}


func main() {



}
