package markup

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ToHTML(m map[string]any, filename string) string {
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

	final := renderHTML(m, root)
	fmt.Println(final)
	return final
}

func renderHTML(m map[string]any, tag *Tag) string {
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
		html += renderHTML(m, child)
	}

	if tag.Name != "root" && tag.Name != "" && tag.Close {
		html += "</" + tag.Name + ">"
		html += "\n"
	}

	if tag.Text != "" {
		if strings.HasPrefix(tag.Text, "#") {
			key := tag.Text[1:len(tag.Text)]
			html += m[key].(string)
		} else {
			html += tag.Text
			html += "\n"
		}
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
