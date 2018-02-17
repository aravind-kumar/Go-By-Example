package main

func twoSum(nums []int, target int) []int {
    sort.Ints(nums)
    
    var i,j int = 0,len(nums)-1
    
    for i<j {
        var current int = nums[i]+nums[j]
        switch {
        case current==target :
            ansArray := []int{i,j}
            return ansArray
        case current<target :
            i+=1
        case current>target:
            j-=1
        }
    }
    return nil
}

func main() {


}
