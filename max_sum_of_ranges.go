package main
import "fmt"
import "math"

func max(a,b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}

func main() {

    maxNum,numOfElements := 0,0
    _,err := fmt.Scanf("%d %d",&maxNum,&numOfElements)
    if err == nil {
        input:=make([]int,maxNum)
        for numOfElements>0 {
            start,end,weight := 0,0,0
            _,err := fmt.Scanf("%d %d %d",&start,&end,&weight)
            if err == nil {
                input[start]+=weight
                if end+1 < maxNum {
                    input[end+1]-=weight
                }
            }
            numOfElements-=1
        }
        sum,maxNum := 0,math.MinInt32
        for i := range input {
            sum+=input[i]
            maxNum = max(maxNum,sum)
        }
        
        fmt.Println(maxNum)
        
    }
}
