package main

import "fmt"

func main () {
	day := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	workday := make ([]string, 5, 5)
	a := copy(workday, day)
	fmt.Println(a, workday)
	a = len(workday)
	freeday := []string{day[a], day[a+1]}
	fmt.Println(freeday)
	week := append(workday, freeday...)
	fmt.Println(week)
}
