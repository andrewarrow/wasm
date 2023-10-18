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
		buffer = append(buffer, `<div>hi2</div>`)
		fmt.Println(line)
	}

	return strings.Join(buffer, "\n")
}
