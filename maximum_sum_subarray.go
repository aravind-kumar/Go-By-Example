func max(a,b int) int {
    if a>b {
        return a
    } else {
        return b
    } 
}

func maxSubArray(nums []int) int {
    
    var maxEndingHere,maxSoFar int = 0,math.MinInt32
    
    for _,num := range nums {
        maxEndingHere = max(num,maxEndingHere+num)
        maxSoFar = max(maxSoFar,maxEndingHere)
    }
    return maxSoFar
    
}
