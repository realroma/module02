package main

import "fmt"
import "strings"
import "strconv"

func main() {
    var text = "120.4 м/с"
    editText := strings.Split(text, " ")
    fmt.Println(editText[0])
    number, _ := strconv.ParseFloat(editText[0], 64)
    speed := 1.609 * number
    speedStr := strconv.FormatFloat(speed, 'g', -1, 64)
    speedSplit := strings.Split(speedStr, ".")
    splitSpeedFloat := len(speedSplit[1])
    edit := speedSplit[1][:splitSpeedFloat-2]
    fmt.Println(speedSplit[0], ".", edit, "км/ч") 
}
