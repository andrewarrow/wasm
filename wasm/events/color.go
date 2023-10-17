package events

import "syscall/js"
import "math/rand"
import "fmt"
import "strings"

func addClass(w js.Value, className string) {
	currentClass := w.Get("className").String()
	newClass := currentClass + " " + className
	w.Set("className", newClass)
}

func removeClass(w js.Value, className string) {
	currentClass := w.Get("className").String()
	tokens := strings.Split(currentClass, " ")
	buffer := []string{}
	for _, item := range tokens {
		if item == className {
			continue
		}
		buffer = append(buffer, item)
	}
	w.Set("className", strings.Join(buffer, " "))
}

var colors = []string{"gray",
	"red",
	"yellow",
	"green",
	"blue",
	"indigo",
	"purple",
	"pink",
	"rose",
	"teal",
}

func randomColor() string {
	randInt := rand.Intn(8) + 1
	return fmt.Sprintf("bg-%s-%d00", colors[rand.Intn(len(colors))], randInt)
}
