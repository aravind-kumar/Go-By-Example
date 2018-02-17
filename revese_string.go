package main

func reverseString(s string) string {
    
    var sArray []byte = []byte(s)
    for i,j := 0,len(s)-1; i<j ; i,j = i+1,j-1 {
        sArray[i],sArray[j] = sArray[j],sArray[i]
    }
    var ans string = string(sArray)
    return ans
}

func main() {


}
