package main

import (
	"fmt"
)


type City struct {
    city string
    population float64
    zipcode string
    surface float64
    
}
func city(){

    var city City
    fmt.Println("Enter city name")
    fmt.Scanln(&city.city)
    fmt.Println("Enter population")
    fmt.Scanln(&city.population)
    fmt.Println("Enter zipcode")
    fmt.Scanln(&city.zipcode)
    fmt.Println("Enter surface")
    fmt.Scanln(&city.surface)
    fmt.Println(city)
       
}

func main(){
    
    fmt.Println("Hello World")
    city()
}