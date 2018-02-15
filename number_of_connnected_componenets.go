func findRoot(components []int,i int) int {
    for i != components[i] {
        i = components[i]
    }
    return i
}

func connected(components []int ,i,j int) bool {
    return findRoot(components,i) == findRoot(components,j)
}

func connect(components []int,i,j int) {
    var rootI,rootJ = findRoot(components,i),findRoot(components,j)
    components[rootJ] = components[rootI]
}

func countComponents(n int, edges [][]int) int {
    var max int = 0
    for _,num := range edges {
        for _,val := range num {
           if val > max {
                max = val
            } 
        }
    }
    
    var components []int = make([]int,max+1,max+1)

    for index,_ := range components {
        components[index] = index    
    }
    
    for _,num := range edges {
        fmt.Println(components)
        if !connected(components,num[0],num[1]){
            connect(components,num[0],num[1])
        }
    }
    
    var count int = 0
    
    for index,num := range components {
        if index == num {
            count+=1
        }
    }
    return count
}
