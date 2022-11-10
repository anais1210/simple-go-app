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
    var seconde int
    fmt.Scanln("Enter: ")
    fmt.Scanf("%d", &seconde)
    fmt.Printf("Transaction by seconde %d for %d: ", transactionNumber, seconde) 
    case 2:
    var minute int
    fmt.Scanln("Enter: ")
    fmt.Scanf("%d", &minute)
    fmt.Printf("Transaction by minute %d for %d: ", transactionNumber, minute) 
    fmt.Println("transaction by minute.")
    case 3:
    var hour int
    fmt.Scanln("Enter: ")
    fmt.Scanf("%d", &hour)
    fmt.Printf("Transaction by hour %d for %d: ", transactionNumber, hour) 
    }

}

func main(){
    transaction()
}