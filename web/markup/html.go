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
	//t.Children = []*Tag{}
	t.Parent = parent
	return &t
}

func ToHTML(filename string) string {
	all := []*Tag{}
	buffer := []string{}

	asBytes, _ := ioutil.ReadFile("markup/" + filename)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")
	var root *Tag
	var current *Tag
	var lastSpaces int
	for _, line := range asLines {
		tokens := strings.Split(line, " ")
		spaces := countSpaces(tokens)
		fmt.Println(spaces)
		if spaces == 0 {
			root = NewTag(tokens[0], nil)
			current = root
		} else if spaces == 2 && lastSpaces == 0 {
			current = NewTag(tokens[2], current)
			all = append(all, current)
		} else if spaces == 4 && lastSpaces == 2 {
			current = NewTag(tokens[4], current)
			all = append(all, current)
		} else if spaces == 4 && lastSpaces == 6 {
			current = current.Parent
			current = NewTag(tokens[4], current)
			all = append(all, current)
		} else if spaces == 6 && lastSpaces == 4 {
			current = NewTag(tokens[6], current)
			all = append(all, current)
		}
		lastSpaces = spaces
	}

	tags := findChildren(root, all)
	fmt.Println(tags)

	buffer = append(buffer, `<div>hi2</div>`)
	return strings.Join(buffer, "\n")
}

func findChildren(root *Tag, all []*Tag) []*Tag {
	some := []*Tag{}
	for _, item := range all {
		if item.Parent == root {
			fmt.Println(item.Name)
			some = append(some, item)
		}
	}
	return some
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
