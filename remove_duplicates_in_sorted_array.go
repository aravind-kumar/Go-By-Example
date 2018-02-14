func removeDuplicates(nums []int) int {
    
    var j int
    if len(nums) > 0 {
        j=1
    } else {
        j=0
    }
    for _,num := range nums {
        if num > nums[j-1] {
            nums[j] = num
            j+=1
        }
    }
    return j;
    
}

func main() {


}
