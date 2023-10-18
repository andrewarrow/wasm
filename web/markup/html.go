package markup

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ToHTML(filename string) string {
	buffer := []string{}

	asBytes, _ := ioutil.ReadFile("markup/" + filename)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")
	for _, line := range asLines {
		tokens := strings.Split(line, " ")
		spaces := countSpaces(tokens)
		if spaces == 0 {
			tag := tokens[0]
			fmt.Println(tag)
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
