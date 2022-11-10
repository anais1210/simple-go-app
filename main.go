package main

import (
	"fmt"
	"sort"
)

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