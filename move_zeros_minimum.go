func moveZeroes(nums []int)  {
    zeroIndex := 0
    for _,num := range nums {
        if num!=0 {
            nums[zeroIndex] = num
            zeroIndex+=1
        }
    }
    
    for i:=zeroIndex;i<len(nums);i+=1 {
        nums[i]=0    
    }
    
}
