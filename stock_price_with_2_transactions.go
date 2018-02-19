package main

func max(a,b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}

func maxProfit(prices []int) int {
    
    var p10,p11,p20,p21 int
    p10,p20 = 0,0
    p11,p21 = math.MinInt32,math.MinInt32
    
    for _,price := range prices {
        p20 = max(p20,p21+price)
        p21 = max(p21,p10-price)
        p10 = max(p10,p11+price)
        p11 = max(p11,-price)
    }
    return p20
}

func main() {


}
