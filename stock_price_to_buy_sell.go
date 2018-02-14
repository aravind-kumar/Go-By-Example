func maxProfit(prices []int) int {
    
    var t00,t01,old_t00 int = 0,math.MinInt32,0
    
    for _,price := range prices {
        old_t00 = t00
        t00 = max(t00,t01+price)
        t01 = max(t01,old_t00-price)
    }
    return t00
    
}

func max(num1,num2 int) int {
    if num1 > num2 {
        return num1
    } else {
        return num2
    }
}
