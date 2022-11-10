package main

import (
	"fmt"
	"sort"
)
func printSlice(s []string) {
    fmt.Printf("len=%d cap=%d %v", len(s),cap(s), s)
}

func sortName(){
    var name string
    var nameList []string
    for {
        fmt.Println("Enter a name:")
        fmt.Scanf("%s", &name)
        nameList = append(nameList, name)
        if(name == "X"){
            break
        }
    }
    sort.Strings(nameList)
    fmt.Printf("%v",nameList)
}

func main(){
    sortName()
}