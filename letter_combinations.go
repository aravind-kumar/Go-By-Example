package main

var phoneMap []string = []string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"} 

func letterCombinationsHelper(input,intermediate string,result *[]string,index int) {
    if len(intermediate) == len(input) {
        *result = append(*result,intermediate)
    } else {
        for i:=index ; i<len(input) ; i+=1 {
            key := input[i] - '0'
            for _,val := range phoneMap[key] {
                intermediate += string(val)
                letterCombinationsHelper(input,intermediate,result,i+1)
                intermediate = intermediate[:len(intermediate)-1]
            }
        }
    }
}
func letterCombinations(digits string) []string {
    
    var result []string = make([]string,0)
    if len(digits) > 0 {
        var intermediate string = ""
        letterCombinationsHelper(digits,intermediate,&result,0)
    }
    return result
}

func main() {
   


}
