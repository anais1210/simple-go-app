package main

import (
	"fmt"
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

func main(){
    city()
}