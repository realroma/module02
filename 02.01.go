package main

import (
    "fmt") 
    
type americanVelocity float64
type evropianVelocity float64

var Aspeed americanVelocity
var Espeed evropianVelocity

func americanSpeed (data americanVelocity) {
    cof:= 2.237
    speed := float64(data)*cof
    Aspeed = americanVelocity(speed)
    fmt.Println(Aspeed, "миль/час")
}

func evropianSpeed (data evropianVelocity){
    var cof evropianVelocity= 3.588
     Espeed := data*cof
    fmt.Printf( "%.2f километры/час", Espeed)
}

func main() {
    var dataA americanVelocity = 130
    americanSpeed(dataA)
    var dataE evropianVelocity = 120.4
    evropianSpeed(dataE)
}
