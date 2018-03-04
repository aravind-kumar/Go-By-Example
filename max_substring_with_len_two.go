package main

func max(a,b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
    
    sMap := make(map[byte]int,len(s))
    
    start,end,targetCount,maxLen := 0,0,0,math.MinInt32
    
    for end < len(s) {
        
        if sMap[s[end]] == 0 {
            targetCount+=1
        }
        
        sMap[s[end]] +=1
        
        for targetCount > 2 && start<len(s) {
            
            sMap[s[start]]-=1
            
            if sMap[s[start]]==0 {
                targetCount-=1
            }            
            start+=1
        }
        
        maxLen = max(maxLen,end-start+1)
        
        end+=1
    }
    
    if maxLen == math.MinInt32 {
        return 0
    } else {
        return maxLen
    }
    
}


func main() {



}
