package main

import "fmt"
import  "reflect"


func isAnagram(s string, t string) bool {
    
    var smap,tmap map[int]int = make(map[int]int),make(map[int]int)
    
    for _,char := range s {
        if value,ok := smap[int(char)]; ok {  
            smap[int(char)] = value+1
        } else {
            smap[int(char)] = 1
        }
    }
    
    for _,char := range t {
        if value,ok := tmap[int(char)]; ok {
            tmap[int(char)] = value+1
        } else {
            tmap[int(char)] = 1
        }
    }
    
    
    return reflect.DeepEqual(smap,tmap)
  
}

func main() {

  var s1,s2 string

  fmt.Println("Enter string 1 and string2 ")
  fmt.Scanln(&s1)
  fmt.Scanln(&s2)  
  fmt.Println(isAnagram(s1,s2))  

}
