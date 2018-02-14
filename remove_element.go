func removeElement(nums []int, val int) int {
  
    valueIndex:=0
    
    for _,num := range nums {
        if num!=val {
            nums[valueIndex] = num
            valueIndex+=1
        }
    }
    return valueIndex
    
}
