package network

import "fmt"

func Save() {
	s, code := GetTo("/markup/list.html", "")
	fmt.Println(s, code)
}
