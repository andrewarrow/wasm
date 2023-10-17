package events

import "syscall/js"
import "fmt"

type Cursor struct {
	X   int
	Y   int
	Div js.Value
}

func NewCursor(div js.Value) *Cursor {
	d := Cursor{}
	d.Div = div
	d.X = 0
	d.Y = 0
	return &d
}

func (d *Cursor) SetLocation(dir string) {
	if dir == "up" {
		d.Y -= 1
	} else if dir == "down" {
		d.Y += 1
	} else if dir == "left" {
		d.X -= 1
	} else if dir == "right" {
		d.X += 1
	}
	d.Div.Set("className", fmt.Sprintf(" absolute bg-white text-black top-%d left-%d", d.Y, d.X))
}
