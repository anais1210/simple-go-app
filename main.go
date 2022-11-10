package main

import (
	"fmt"
)

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

func main(){
    transaction()
}