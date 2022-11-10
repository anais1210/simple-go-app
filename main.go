package main

import (
	"fmt"
	"sort"
)
type City struct {
    cityname string
    population uint
    zipcode uint
    surface uint
}

func city(){
    var mapping = map[string]City{}
    var numberofcities int
    fmt.Println("How many cities do you want to enter?")
    fmt.Scanln(&numberofcities)
    fmt.Println("Number of cities: ", numberofcities)

    for i := 0; i < numberofcities; i++ {
        var cityname string
        var population uint
        var zipcode uint
        var surface uint
        fmt.Println("Enter the name of the city: ")
        fmt.Scanln(&cityname)
        fmt.Println("Enter the population of the city: ")
        fmt.Scanln(&population)
        fmt.Println("Enter the zipcode of the city: ")
        fmt.Scanln(&zipcode)
        fmt.Println("Enter the surface of the city: ")
        fmt.Scanln(&surface)
        mapping[cityname] = City{cityname, population, zipcode, surface}
    }
    fmt.Println("Mapping: ", mapping)
}

func transaction(){
    var transactionNumber int
    fmt.Println("Transaction number: ")
    fmt.Scanln(&transactionNumber)


    var choice int
    fmt.Println("1. Seconde")
    fmt.Println("2. Minute")
    fmt.Println("3. Hour")
    fmt.Scanf("%d", &choice)


    switch choice {
    case 1:
        var second int
        fmt.Scanln("How many seconds?: ")
        fmt.Scanf("%d", &second)
        fmt.Printf("%d transactions per seconds", transactionNumber/second) 
    case 2:
        var minute int
        fmt.Scanln("How many minutes?: ")
        fmt.Scanf("%d", &minute)
        fmt.Printf("%d transactions by minute ", transactionNumber/minute) 
        fmt.Println("transaction by minute.")
    case 3:
        var hour int
        fmt.Scanln("How many hours?: ")
        fmt.Scanf("%d", &hour)
        fmt.Printf("%d transactions by hour ", transactionNumber/hour) 
    default:
        fmt.Println("Invalid choice")
    }

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
    city()
    transaction()
    sortName()

}