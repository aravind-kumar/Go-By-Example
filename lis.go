package main
import "fmt"
import "math"

func max(num1,num2 int) int {
    if num1 > num2 {
        return num1
    } else {
        return num2
    }
}

func sliceMax(input []int) int {
    maxNum := math.MinInt32
    
    for _,num := range input {
        maxNum = max(maxNum,num)
    }
    
    return maxNum
}

func lis(input []int) int {
    cache := make([]int,len(input))
    
    for i,_ := range cache {
        cache[i] = 1
    }
    
    for i,_ := range input {
        for j:= 0 ; j<i ; j+=1 {
            if input[i]>input[j] && cache[j]+1 > cache[i] {
                cache[i] = cache[j]+1
            }
        }
    }
    return sliceMax(cache)
}

func main() {

 var num int
 _,err := fmt.Scanf("%d",&num)
    if err == nil {
        input := make([]int,0,num)
        for num > 0 {
            var inputNum int
            _,err := fmt.Scanf("%d",&inputNum)
            if err == nil {
                input=append(input,inputNum)
            }
            num-=1
        }
        fmt.Println(lis(input))
    }
    
}
