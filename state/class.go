package state

import "syscall/js"
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
