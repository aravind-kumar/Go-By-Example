package main

func isValidSudoku(board [][]byte) bool {
    
    for i:=0 ; i<9 ; i+=1 {
        
        var rowMap map[byte]bool = make(map[byte]bool,len(board))
        var colMap map[byte]bool = make(map[byte]bool,len(board))
        var blockMap map[byte]bool = make(map[byte]bool,len(board))
        for j:=0 ; j<9 ; j+=1 {
            
            val := board[i][j]
            if val != '.' {
                if _,ok := rowMap[val] ; ok {
                    return false
                } else {
                    rowMap[val] = true
                }
            }
            
            val = board[j][i]
            if val != '.' {
                if _,ok := colMap[val] ; ok {
                    return false
                } else {
                    colMap[val] = true
                }
            }
            
            
            blockRowIndex := 3*(i/3) + j/3 
            blockColIndex := 3*(i%3) + j%3
                
            val = board[blockRowIndex][blockColIndex]
                
            if val != '.' {
                if _,ok := blockMap[val] ; ok {
                    return false
                } else {
                    blockMap[val] = true
                }
            }
        }
    }
    return true
}

func main() {


}
