func plusOne(digits []int) []int {
    var n = len(digits)
    for i:=n-1;i>=0;i-=1 {
        if digits[i] < 9 {
            digits[i]+=1
            return digits
        }
        digits[i]=0
    }
    newArray := make([]int,n+1,n+1)
    newArray[0] = 1
    return newArray
    
}
