package main

import "fmt"

func main () { 
	f := []string{"zachem eto?"} 
	var client map[string]map[string]string = map[string]map[string]string{
	"Anatoliy Fosjv": {
		"Geroi ne ricari i magi": "1",
		"ne moloko:sostav": "1",
		"nas ne dogonyat": f[0]},
	}
	fmt.Println(client)
	books := make([]string, 0, len(client["Anatoliy Fosjv"]))
	for k := range client["Anatoliy Fosjv"]{
		books = append(books, k)
	}
	
	fmt.Println(len(books))
}
