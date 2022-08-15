package main

import (
	"os"
	"fmt"
	"time"
)
var start time.Time
var m time.Time
func logTime() {
	if start == m {
		start = time.Now()
		return
	}
	time := time.Now()
	res := time.Sub(start)
        fmt.Println(res)

}

func main() {
	logTime()
	data := []byte("hi")
	err := os.WriteFile("out.txt", data, 0600)     // создаем файл
	if err != nil {
		fmt.Println("i have been stopping of error")
	} 
	logTime()
}
