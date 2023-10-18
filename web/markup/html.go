package markup

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tag struct {
	Name     string
	Children []*Tag
	//Parent *Tag
}

func NewTag(name string) *Tag {
	t := Tag{}
	t.Name = name
	t.Children = []*Tag{}
	//t.Parent = parent
	return &t
}

func ToHTML(filename string) string {
	buffer := []string{}

	asBytes, _ := ioutil.ReadFile("markup/" + filename)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")
	root := NewTag("root")

	stack := []*Tag{root}
	var lastSpaces int

	for _, line := range asLines {
		tokens := strings.Split(line, " ")

		spaces := countSpaces(tokens)
		fmt.Println(spaces)
		var tag *Tag
		if spaces == 0 {
			tag = NewTag(tokens[0])
		} else if spaces == 2 && lastSpaces == 0 {
			tag = NewTag(tokens[2])
		} else if spaces == 4 && lastSpaces == 2 {
			tag = NewTag(tokens[4])
		} else if spaces == 4 && lastSpaces == 6 {
			stack = stack[0 : len(stack)-1]
			tag = NewTag(tokens[4])
		} else if spaces == 6 && lastSpaces == 4 {
			tag = NewTag(tokens[6])
		}

		parent := stack[len(stack)-1]
		parent.Children = append(parent.Children, tag)
		stack = append(stack, tag)
		lastSpaces = spaces
	}

	buffer = append(buffer, `<div>hi2</div>`)
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
