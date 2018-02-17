package main

func climbStairs(n int) int {
    if n==0 || n==1 {
        return n
    }
    pre,prepre,count := 1,1,0
    
    for i:=2;i<=n;i=i+1 {
        count = pre +prepre
        prepre = pre
        pre = count
    }
    return count
    
}

func main() {


}
