package main
import "fmt"

func minWindow(s string, t string) string {
    
    scount := make(map[int]int,0)
    
    for _,char := range s {
        scount[int(char)] = 0
    }
    
    for _,char := range t {
        scount[int(char)]+=1
    }
    
    start,end,minStart,minLen,numOfCharacters := 0,0,math.MaxInt32,math.MaxInt32,len(t)
    for end < len(s) {
        
        char := int(s[end])
        
        if val,ok := scount[char] ; ok && val > 0 {
            numOfCharacters -=1    
        }
        scount[char] -=1
        

        for numOfCharacters == 0 {
            if minLen > end-start+1 {
                minStart = start 
                minLen = end-start+1
            }
    
            char = int(s[start])
            if val,ok := scount[char] ; ok && val >= 0 {
                numOfCharacters+=1
            }
            scount[char] += 1

            start+=1
            
        }
        
        end+=1
    }


    if minLen == math.MaxInt32 {
        return ""
    } else {
        return s[minStart:minStart+minLen]
    }
    
}

func main() {



}
