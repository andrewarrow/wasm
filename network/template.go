package network

func GetTemplate(name string) string {
	s, _ := GetTo("/markup/"+name+".html", "")
	return s
}
