package main

import "fmt"

func twoSum(nums []int, target int) []int {
    
    var indexMap map[int][] int = make(map[int][]int,len(nums))

    for index,val := range(nums) {
      indexMap[val] = append(indexMap[val],index)
    }

    for k,v := range(indexMap) {
        
        val,ok := indexMap[target-k];
        
        if ok && target-k != k {
            return []int {v[0],val[0]}
        } else if ok && len(val)>1 {
            fmt.Println("entering else")
            return []int {val[0],val[1]}
        }
    }
    return []int{-1,-1}
}

func main() {

   fmt.Println(twoSums([]int{3,2,4},6))

}

