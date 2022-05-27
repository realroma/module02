package main

import "fmt"
import "strings"
import "strconv" 

func main() { 
    var numberI int
    var numberF float64
    var americanNoEdit = "130 м/с"
    var europianNoEdit = "120.4 м/с"
    var mile float64 = 1.609
    var cof float64 = 3.6
   
    editText := strings.Split(americanNoEdit, " ") 
    numberI, _= strconv.Atoi(editText[0])
    AmericanVelocity := float64(numberI)*cof/mile
    
    editText = strings.Split(europianNoEdit, " ")
    numberF, _= strconv.ParseFloat(editText[0], 64)
    v := numberF * 3.6
    speedStr := strconv.FormatFloat(v, 'g', -1, 64)
    speedSplit := strings.Split(speedStr, ".")
    splitSpeedFloat := len(speedSplit[1])
    edit := speedSplit[1][splitSpeedFloat-2]
    
    fmt.Println(speedSplit[0], ".", edit, "км/ч")
    fmt.Println(AmericanVelocity, "миль/ч") 
    }
