package main
import "fmt"


func max(num1,num2 int) int {
    if num1 > num2 {
        return num1
    } else {
        return num2
    }
}

func lcs(input1,input2 []int) []int {

    matrix := make([][]int,len(input1)+1)
    
    for i,_ := range matrix {
        matrix[i] = make([]int,len(input2)+1)
    } 
    
    for i:=1 ; i<=len(input1) ; i+=1 {
        for j:=1 ; j<=len(input2) ; j+=1 {
            if input1[i-1] == input2[j-1] {
                matrix[i][j] = 1+matrix[i-1][j-1]
            } else {
                matrix[i][j] = max(matrix[i-1][j],matrix[i][j-1])
            }
        }
    }
    
    length := matrix[len(input1)][len(input2)]
    output := make([]int,length,length)
    index  := length
    
    for i,j := len(input1),len(input2) ; i > 0 && j >0 ; {
        if input1[i-1] == input2[j-1] {
           output[index-1] = input1[i-1]
           i-=1
           j-=1
           index-=1
        } else if matrix[i-1][j] > matrix[i][j-1] {
            i-=1
        } else {
            j-=1
        }
    }
    return output
}

func main() {

 var len1,len2 int
 _,err:= fmt.Scanf("%d %d",&len1,&len2)
    
 if err == nil {
    var input1,input2 []int = make([]int,0,len1),make([]int,0,len2) 
    for i:=0 ; i<len1 ; i+=1 {
        var num int
        _,err := fmt.Scanf("%d",&num)
        if err == nil {
            input1 = append(input1,num)
        }
    } 
    for i:=0 ; i<len2 ; i+=1 {
      var num int
        _,err := fmt.Scanf("%d",&num)
        if err == nil {
            input2 = append(input2,num)
        }
    }
     
     for _,num := range lcs(input1,input2) {
         fmt.Printf("%d ",num)
     }
}
 
}
