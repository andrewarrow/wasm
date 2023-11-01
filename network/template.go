package network

import "fmt"

func GetTemplate(name string) string {
	fmt.Println("here4")
	s, code := GetTo("http://localhost:3000/markup/"+name, "")
	fmt.Println("here5")
	fmt.Println(code, s)
	return s
}
