package markup

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tag struct {
	Name     string
	Text     string
	Children []*Tag
	//Parent *Tag
}

var validTagMap = map[string]bool{"div": true, "img": true, "root": true}

func NewTag(name string) *Tag {
	t := Tag{}
	if validTagMap[name] {
		t.Name = name
	} else {
		t.Text = name
	}
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
		if len(tokens) == 1 {
			continue
		}

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

	final := renderHTML(root)
	fmt.Println(final)
	buffer = append(buffer, `<div>hi2</div>`)
	return strings.Join(buffer, "\n")
}

func renderHTML(tag *Tag) string {
	html := ""

	if tag.Name != "root" && tag.Name != "" {
		html += "<" + tag.Name
		html += ">"
	}

	for _, child := range tag.Children {
		html += renderHTML(child)
	}

	if tag.Name != "root" && tag.Name != "" {
		html += "</" + tag.Name + ">"
	}

	if tag.Text != "" {
		html += tag.Text
	}

	return html
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
