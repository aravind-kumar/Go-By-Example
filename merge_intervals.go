/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Intervals []Interval

func (input Intervals) Len() int {
    return len(input)
}

func (input Intervals) Swap(i,j int) {
    input[i],input[j] = input[j],input[i]
}

func (input Intervals) Less(i,j int) bool {
    return input[i].Start < input[j].Start || 
    ((input[i].Start==input[j].Start) && (input[i].End<input[j].End))
}

func max(a,b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func merge(intervals []Interval) []Interval {
    
    ans := make([]Interval,0)
    
    if len(intervals) > 0 {
        
        sort.Sort(Intervals(intervals))
        
        ans = append(ans,intervals[0])
        for i:=2 ; i<=len(intervals) ; i+=1 {
            if intervals[i-1].Start <= ans[len(ans)-1].End {
                ans[len(ans)-1].End = max(ans[len(ans)-1].End,intervals[i-1].End)
            } else {
                ans = append(ans,intervals[i-1])
            }
        }
        
    }
    return ans
    
}

func main() {



}

