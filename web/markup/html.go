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
	Close    bool
	Class    string
	//Parent *Tag
}

var validTagMap = map[string]int{"div": 2, "img": 3, "root": 1}

func makeClassAndAttrMap(tokens []string) map[string]string {
	m := map[string]string{}

	class := ""
	for _, item := range tokens {
		if strings.Contains(item, "=") {
		} else {
			class += item + " "
		}
	}
	m["class"] = class

	return m
}

func NewTag(index int, tokens []string) *Tag {
	t := Tag{}
	name := tokens[index]
	classAndAttrMap := makeClassAndAttrMap(tokens[index+1 : len(tokens)])
	t.Class = classAndAttrMap["class"]
	flavor := validTagMap[name]
	if flavor > 0 {
		t.Close = flavor == 2
		t.Name = name
	} else {
		t.Text = name
	}
	t.Children = []*Tag{}
	//t.Parent = parent
	return &t
}

func ToHTML(filename string) string {
	asBytes, _ := ioutil.ReadFile("markup/" + filename)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")
	root := NewTag(0, []string{"root"})

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
			tag = NewTag(0, tokens)
		} else if spaces == 2 && lastSpaces == 0 {
			tag = NewTag(2, tokens)
		} else if spaces == 4 && lastSpaces == 2 {
			tag = NewTag(4, tokens)
		} else if spaces == 4 && lastSpaces == 6 {
			stack = stack[0 : len(stack)-1]
			stack = stack[0 : len(stack)-1]
			tag = NewTag(4, tokens)
		} else if spaces == 6 && lastSpaces == 4 {
			tag = NewTag(6, tokens)
		}

		parent := stack[len(stack)-1]
		parent.Children = append(parent.Children, tag)
		stack = append(stack, tag)
		lastSpaces = spaces
	}

	final := renderHTML(root)
	fmt.Println(final)
	return final
}

func renderHTML(tag *Tag) string {
	html := ""

	if tag.Name != "root" && tag.Name != "" {
		html += "<" + tag.Name
		html += fmt.Sprintf(` class="%s" `, tag.Class)
		if tag.Close == false {
			html += "/>"
		} else {
			html += ">"
		}
		html += "\n"
	}

	for _, child := range tag.Children {
		html += renderHTML(child)
	}

	if tag.Name != "root" && tag.Name != "" && tag.Close {
		html += "</" + tag.Name + ">"
		html += "\n"
	}

	if tag.Text != "" {
		html += tag.Text
		html += "\n"
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
