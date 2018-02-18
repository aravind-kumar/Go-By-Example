package main

func permute(nums []int) [][]int {
    var result [][]int

    if len(nums) > 0 {
        generatePermutations(nums,&result,0)
    }
    return result
}


func generatePermutations(nums []int,result *[][]int,start int) {
    
    if start == len(nums)-1 {
        fmt.Println(nums)
        *result = append(*result,nums)
    } else {
        for i:=start ; i<len(nums) ; i+=1 {
            nums[i],nums[start] = nums[start],nums[i]
            generatePermutations(nums,result,start+1)
        }
    }
}

func main() {


}
