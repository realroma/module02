package main

import "fmt"

func contains (a []string, x string) (res bool){
	for _, q :=  range a{
		if q == x {
			return true
		}
	}
	return res
}

func main(){
	var a []string
	a = []string{"i", "know", "where"}
	x := "you"
	fmt.Println(contains(a, x))
	a = []string{"i", "love", "you"}
	fmt.Println(contains(a, x))
}
