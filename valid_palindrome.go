func isPalindrome(s string) bool {
    
    for i,j := 0,len(s)-1;i<j;i,j=i+1,j-1 {
        for !isLetter(s[i]) && i<j {
            i+=1
        }
        for !isLetter(s[j]) && i<j {
            j-=1
        }
        if toLower(s[i]) != toLower(s[j]) {
            return false;
        }
    }
    
    return true;
    
}

func isLetter(c byte) bool {
    return (c>='a' && c<='z') ||
            (c>='A' && c<='Z') ||
    (c>='0' && c<='9')
}


func toLower(c byte) byte {
    if c>='A' && c<='Z' {
        return c + ('a' - 'A')
    }
    return c
}
