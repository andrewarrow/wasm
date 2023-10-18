package markup

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tag struct {
	Name string
	//Children []*Tag
	Parent *Tag
}

func NewTag(name string, parent *Tag) *Tag {
	t := Tag{}
	t.Name = name
	//t.Children = []*tag{}
	t.Parent = parent
	return &t
}

func ToHTML(filename string) string {
	buffer := []string{}

	asBytes, _ := ioutil.ReadFile("markup/" + filename)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")
	var current *Tag
	for _, line := range asLines {
		tokens := strings.Split(line, " ")
		spaces := countSpaces(tokens)
		fmt.Println(spaces)
		if spaces == 0 {
			root := NewTag(tokens[0], nil)
			current = root
		} else if spaces == 2 {
			current = NewTag(tokens[2], current)
		}
		buffer = append(buffer, `<div>hi2</div>`)
	}

	return strings.Join(buffer, "\n")
}

func countSpaces(tokens []string) int {
	count := 0
	for _, item := range tokens {
		if item == "" {
			count++
		} else {
			break
		}
	}
	return count
}
