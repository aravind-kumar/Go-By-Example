package main

func isValid(i,j,m,n int) bool {
    return i>=0 && j>=0 && i<m && j<n
}

func getNumOfLives(i,j int,board [][]int) (count int) {

    coord := [][]int {{1,0},{0,1},{-1,0},{0,-1},{1,1},{-1,1},{1,-1},{-1,-1}}   
    
    for _,val := range coord {
        var newRow,newCol int = i+val[0],j+val[1]
        if isValid(newRow,newCol,len(board),len(board[0])) {
            count += (board[newRow][newCol] & 1)       
        }
    }
    
    return 
    
}

func gameOfLife(board [][]int)  {
    if(len(board) > 0 ) {
    
        for i,val := range board {
            for j,value := range val {
                numOfLives := getNumOfLives(i,j,board)
                
                if value == 1 && (numOfLives ==2 || numOfLives ==3) {
                    board[i][j] = 3
                }
                if value == 0 && numOfLives == 3 {
                    board[i][j] = 2
                }
            }
        }
        
        for i,val := range board {
            for j,_ := range val {
                board[i][j]>>=1
            }
        }
    }
}

func main() {



}
