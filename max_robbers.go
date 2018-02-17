package main

func max(a,b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}

func maxRobber(nums []int, index int,cache map[int]int) int {
    
    if val,ok := cache[index]; ok {
        return val
    }
    
    if index<len(nums) {
        
        val := max(nums[index] + maxRobber(nums,index+2,cache),
                   maxRobber(nums,index+1,cache))
        cache[index] = val
        return val
    }
    return 0
    
}

func rob(nums []int) int {
    var cache map[int]int = make(map[int]int,len(nums))
    return maxRobber(nums,0,cache)
    
}

func main() {


}
