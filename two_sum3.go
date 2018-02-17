package main

func twoSum(nums []int, target int) []int {
    
    var charIndexMap map[int]int = make(map[int]int,len(nums))
    
    for index,val := range nums {
        searchNum := target-val
        
        if valIndex,ok := charIndexMap[searchNum] ; ok {
            return ([]int{index,valIndex})
        } else {
            charIndexMap[val]=index    
        }
    }
    
    return nil
}

func main() {


}
