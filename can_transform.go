package main

func canTransform(start string, end string) bool {
    
    var pos1,pos2 int = 0,0
    
    for pos1<len(start) && pos2<len(end) {
        
        for pos1<len(start) && start[pos1] == 'X' {
            pos1+=1
        }
        
        for pos2<len(start) && end[pos2] == 'X' {
            pos2+=1
        }
        
        if pos1 == len(start) && pos2 == len(end) {
            return true
        }
        if pos1 == len(start) || pos2 == len(end) {
            return false
        }
        if start[pos1] != end[pos2] {
            return false
        }
        if start[pos1] == 'L' && pos2>pos1 {
            return false
        }
        if start[pos1] == 'R' && pos1>pos2 {
            return false
        }
        pos1+=1;
        pos2+=1;
    }
    return true;
    
}

func main() {

}
