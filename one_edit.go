package main
import "fmt"

func areEqual(input1,input2 string) bool {
    for i:=0 ; i<len(input1) ; i+=1 {
        if input1[i] != input2[i] {
            return false
        }
    }
    return true
} 

func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

func min(a,b int) int {
    if a<b {
        return a
    } else {
        return b
    }
} 

func isOneEditDistance(s string, t string) bool {
    sLen,tLen := len(s),len(t)
        
    if abs(sLen-tLen) >=2 {
        return false
    }
    
    
    for i:=0; i<min(sLen,tLen) ; i+=1 {
        
        if s[i]!=t[i] {
            switch {
            case sLen == tLen :
                return areEqual(s[i+1:],t[i+1:])
            case sLen < tLen:
                return areEqual(s[i:],t[i+1:])
            case sLen > tLen:
                return areEqual(t[i:],s[i+1:])
            }
        }
    }
    
    return abs(sLen-tLen) == 1
}

func main() {



}
