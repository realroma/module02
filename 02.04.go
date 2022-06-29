package main

import "fmt"

func main () {
	f := []string{"zachem eto?"}
	var client map[string]map[string]string = map[string]map[string]string{
	"Anatoliy Fosjv": {
		"Geroi ne ricari i magi": "1",
		"ne moloko:sostav": "1",
		"nas ne dogonyat": f[0]},
	"fostr agilw": {
		"fogot": "1",},
	}
	fmt.Println(client)
	books := make([]string, 0, 0)
	for t := range client{
		fmt.Println(t)
		books = make ([]string, 0, len(client[t]))
		for k := range client[t]{
		books = append(books, k)
		}
	fmt.Println(books)
	}
}
