package main

func containsDuplicate(nums []int) bool {
    var numMap map[int] bool = make(map[int]bool,len(nums)+1)
    
    for _,num := range nums {
        if _,ok := numMap[num]; ok {
            return true
        } 
        numMap[num]=true
    }
    return false
}

func main() {


}
