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
	//Class    string
	Attr map[string]string
	//Parent *Tag
}

var validTagMap = map[string]int{"div": 2, "img": 3, "root": 1, "a": 2}

func (t *Tag) MakeAttr() string {
	buffer := ""

	for key, value := range t.Attr {
		buffer += fmt.Sprintf(`%s="%s" `, key, value)
	}

	return buffer
}

func fixValueForTag(name, key, value string) string {
	if name == "img" && key == "src" {
		value = fmt.Sprintf("/assets/images/%s", value)
	}
	return value
}

func getKeyValue(s string) (string, string) {
	tokens := strings.Split(s, "=")
	if len(tokens) == 2 {
		return tokens[0], tokens[1]
	}
	return "", ""
}

func makeClassAndAttrMap(name string, tokens []string) map[string]string {
	m := map[string]string{}

	class := ""
	for _, item := range tokens {
		if strings.Contains(item, "=") {
			key, value := getKeyValue(item)
			value = fixValueForTag(name, key, value)
			m[key] = value
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
	t.Attr = makeClassAndAttrMap(name, tokens[index+1:len(tokens)])
	//t.Class = classAndAttrMap["class"]
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
		delta := spaces - lastSpaces
		//fmt.Println(delta, line)
		if delta < 0 {
			delta = delta * -1
			delta = delta / 2
			offset := 2
			if delta > 2 {
				offset = (delta * 2) - 2
			}
			//fmt.Println("f", delta, offset, line)
			stack = stack[0 : len(stack)-(offset)]
		}

		tag := NewTag(spaces, tokens)
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
		html += fmt.Sprintf(` %s `, tag.MakeAttr())
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
