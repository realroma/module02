package main

import "fmt"

func main() {
    var a *int
    var b int = 2
    a = &b
    fmt.Println(*a) 
    *a = 3
    fmt.Println(b)
}
