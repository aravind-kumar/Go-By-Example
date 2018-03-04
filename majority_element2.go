package main

func majorityElement(nums []int) []int {
    num1,count1,num2,count2 := -1,0,-1,0
    
    for _,val := range nums {
        switch {
        case val == num1:
            count1+=1
        case val == num2 :
            count2+=1
        case count1 == 0 :
            num1=val
            count1=1
        case count2 == 0 :
            num2=val
            count2=1
        default : 
            count1-=1
            count2-=1
        }
    }
    
    count1,count2 = 0,0
    
    for _,val := range nums {
        if val == num1 {
            count1+=1
        } else if val == num2 {
            count2+=1
        }
    }
    
    ans := make([]int,0)
    if len(nums) > 0 {
        if count1 > len(nums)/3 {
            ans = append(ans,num1)
        }
        if count2 > len(nums)/3 {
            ans = append(ans,num2)
        }
    }
    return ans
    
}


func main() {



}
