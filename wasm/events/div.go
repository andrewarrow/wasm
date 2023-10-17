package events

import "strings"
import "syscall/js"
import "fmt"

type Div struct {
	ClassMap map[string]bool
	BgColor  string
	Div      js.Value
	Index    int
	Children []*Div
	Parent   *Div
}

func NewDiv(div js.Value, parent *Div) *Div {
	d := Div{}
	d.Div = div
	d.Parent = parent
	d.ClassMap = map[string]bool{}
	d.Children = []*Div{}
	return &d
}

func (d *Div) AppendChild(child *Div) {
	d.Div.Call("appendChild", child.Div)
	d.Children = append(d.Children, child)
	child.Parent = d
}

func (d *Div) Classes() string {
	buffer := []string{}
	for k, _ := range d.ClassMap {
		buffer = append(buffer, k)
	}
	buffer = append(buffer, d.BgColor)

	return strings.Join(buffer, " ")
}

func (d *Div) SetClasses() {
	d.Set("className", d.Classes())
}

func (d *Div) ReadClasses() {
	currentClass := d.Get("className")
	tokens := strings.Fields(currentClass)
	for _, token := range tokens {
		d.ClassMap[token] = true
	}
}

func (d *Div) RemoveClass(c string) {
	delete(d.ClassMap, c)
}

func (d *Div) HasClass(c string) bool {
	return d.ClassMap[c]
}

func (d *Div) AddClass(c string) {
	d.ClassMap[c] = true
}
func (d *Div) Get(c string) string {
	return d.Div.Get(c).String()
}
func (d *Div) Set(k, v string) {
	d.Div.Set(k, v)
}

func (d *Div) FindChildren() {
	children := d.Div.Get("children")

	for i := 0; i < children.Length(); i++ {
		childDiv := NewDiv(children.Index(i), d)
		d.Children = append(d.Children, childDiv)
		childDiv.FindChildren()
	}
}

func (d *Div) Debug() {

	for i, item := range d.Children {
		fmt.Println("d", i, item.Get("outerHTML"))
		item.Debug()
	}

}
