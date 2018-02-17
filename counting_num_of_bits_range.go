func countBits(num int) []int {
    
    var numOfBits []int = make([]int,num+1,num+1)
    
    numOfBits[0]=0
    for i:=1;i<=num;i+=1 {
        numOfBits[i] = numOfBits[(i&(i-1))]+1
    }
    return numOfBits
    
}


