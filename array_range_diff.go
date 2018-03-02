package main
import "fmt"
import "math"

func max(a,b int) int{
    if a>b {
        return a
    } else {
        return b
    }
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    maxNumRange,numOfIterations := 0,0
    
    _,err := fmt.Scanf("%d %d",&maxNumRange,&numOfIterations)
    
    if err == nil {
        start,end,sum := 0,0,0
        
        inputArray := make([]int,maxNumRange+1)
        
        for numOfIterations>0 {
            _,err := fmt.Scanf("%d %d %d",&start,&end,&sum)
            if err == nil {
                inputArray[start]+=sum
                if end+1<=maxNumRange {
                    inputArray[end+1]-=sum    
                }
                
            }
            numOfIterations-=1
        }
        
        x := 0
        maxNum := math.MinInt32
        
        for i:= 1 ; i<=maxNumRange ; i+=1 {
            x+=inputArray[i]
            maxNum = max(maxNum,x)
        }
        
        fmt.Println(maxNum)
    }
}
