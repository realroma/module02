package main

import "fmt"

func contains (a []string, x string) bool{
	var res bool
	for _, q :=  range a{
		if q == x {
			return true
		}
	}
	return res
}

func getMax (x ...int)  int {
	var getOut int
	for a, b := range x{
		if a == 0 {
			getOut = b
		} else if b > getOut {
			getOut = b
		}
	}
	return getOut
}

func main(){
	var a []string
	a = []string{"i", "know", "where"}
	x := "you"
	fmt.Println(contains(a, x))
	a = []string{"i", "love", "you"}
	fmt.Println(contains(a, x))
	fmt.Println(getMax(1, 2, 3, 4, 5, 6))
	fmt.Println(getMax(-1, -3, 2, -4))
}
