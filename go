package main

import "fmt"
import "strings"
import "strconv"

func main() {
    var text = "120 м/с"
    editText := strings.Split(text, " ")
    number, _:= strconv.Atoi(editText[0])
    fNumber := float64(number)/1000*60
    fmt.Println(fNumber, "км/ч") 
}
