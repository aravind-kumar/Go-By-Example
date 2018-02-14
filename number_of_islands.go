func numIslands(grid [][]byte) int {

    var count int = 0    
    for i:=0;i<len(grid);i+=1 {
        for j:=0;j<len(grid[0]);j+=1 {
            if grid[i][j] == '1' {
                count+=1
                sinkIsland(i,j,len(grid),len(grid[0]),grid)
            }
        }
    }
    return count
    
}

func sinkIsland(i,j,m,n int,grid [][] byte) {
    if i>=0 && j>=0 && i<m && j<n && grid[i][j] == '1' {
        grid[i][j] = 0
        sinkIsland(i+1,j,m,n,grid)
        sinkIsland(i-1,j,m,n,grid)
        sinkIsland(i,j+1,m,n,grid)
        sinkIsland(i,j-1,m,n,grid)
    }
}
