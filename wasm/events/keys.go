package events

import "syscall/js"
import "fmt"

var Focus string = "div"

type Editor struct {
	InsertMode bool
	Div        *Div
	End        *CommandBar
	Space      bool
	Cursor     *Cursor
	Selected   *Div
}

func NewEditor() *Editor {
	e := Editor{}
	d := js.Global().Get("document").Call("getElementById", "w")
	e.Div = NewDiv(d, nil)
	e.Div.BgColor = "bg-pink-600"
	e.Div.FindChildren()
	//d = js.Global().Get("document").Call("getElementById", "cursor")
	//e.Cursor = NewCursor(d)
	d = js.Global().Get("document").Call("getElementById", "end")
	e.End = &CommandBar{d}
	e.Selected = e.Div
	return &e
}

func (e *Editor) Key(this js.Value, p []js.Value) any {
	k := p[0].String()

	e.HandleKey(k)

	return js.Undefined()
}

func (e *Editor) HandleDivInsert(k string) {
	if k == "Escape" || k == "Enter" {
		return
	}
	fmt.Println(k)
	text := e.Div.Get("innerText")
	if k == "Backspace" {
		e.Div.Set("innerText", text[0:len(text)-1])
		return
	}
	if e.Space {
		k = " " + k
		e.Space = false
	}
	if k == " " {
		e.Space = true
	}
	e.Div.Set("innerText", text+k)
}

func (e *Editor) HandleDivKey(k string) {
	//fmt.Println("e HandleDivKey", k)
	if e.InsertMode {
		e.HandleDivInsert(k)
	}

	if k == "i" && e.InsertMode == false {
		e.InsertMode = true
	} else if k == "s" && e.InsertMode == false {
		removeClass(e.Selected.Div, "bg-pink-600")
		content := e.Div.Get("innerHTML")
		addClass(e.Selected.Div, "bg-pink-600")
		fmt.Println(content)
		js.Global().Call("sendFormWithWasm", content)
	} else if k == "ArrowRight" && e.InsertMode == false {
		//e.Cursor.SetLocation("right")
	} else if k == "ArrowLeft" && e.InsertMode == false {
		//e.Cursor.SetLocation("left")
	} else if k == "ArrowUp" && e.InsertMode == false {
		//e.Cursor.SetLocation("up")
		removeClass(e.Selected.Div, "bg-pink-600")
		if len(e.Selected.Children) == 0 {
			parent := e.Selected.Parent
			parent.Index--
			if parent.Index < 0 {
				parent.Index = 0
				e.Selected = parent
				addClass(e.Selected.Div, "bg-pink-600")
			} else {
				e.Selected = parent.Children[parent.Index]
				addClass(e.Selected.Div, "bg-pink-600")
			}
		} else {
			parent := e.Selected.Parent
			if parent == nil {
				addClass(e.Selected.Div, "bg-pink-600")
			} else {
				parent.Index = 0
				e.Selected = parent
				addClass(e.Selected.Div, "bg-pink-600")
			}
		}
		e.Selected.Div.Call("scrollIntoView", `{ behavior: "smooth" }`)
	} else if k == "ArrowDown" && e.InsertMode == false {
		//e.Cursor.SetLocation("down")
		removeClass(e.Selected.Div, "bg-pink-600")
		if len(e.Selected.Children) == 0 {
			parent := e.Selected.Parent
			parent.Index++
			if parent.Index >= len(parent.Children) {
				parent.Index = 0
				if parent.Parent != nil {
					e.Selected = parent.Parent
					e.Selected.Index++
				}
				addClass(e.Selected.Div, "bg-pink-600")
			} else {
				e.Selected = parent.Children[parent.Index]
				addClass(e.Selected.Div, "bg-pink-600")
			}
		} else {
			e.Selected = e.Selected.Children[e.Selected.Index]
			addClass(e.Selected.Div, "bg-pink-600")
		}
		e.Selected.Div.Call("scrollIntoView", `{ behavior: "smooth" }`)
	} else if k == "f" && e.InsertMode == false {
		e.Selected.ReadClasses()
		if e.Selected.HasClass("flex") {
			e.Selected.RemoveClass("flex")
		} else {
			e.Selected.AddClass("flex")
		}
		e.Selected.SetClasses()
	} else if k == "c" && e.InsertMode == false {
		//e.Div.ReadClasses()
		//e.Div.BgColor = randomColor()
		//e.Div.SetClasses()
		d := js.Global().Get("document").Call("createElement", "div")
		child := NewDiv(d, e.Selected)
		child.AddClass("p-3")
		child.AddClass("border border-black")
		//child.BgColor = randomColor()
		child.SetClasses()
		e.Selected.AppendChild(child)
	} else if k == ":" && e.InsertMode == false {
		Focus = "bar"
		e.End.Show()
	} else if k == "Escape" {
		e.InsertMode = false
	}
}

func (e *Editor) HandleKey(k string) {
	if Focus == "div" {
		e.HandleDivKey(k)
	} else if Focus == "bar" {
		e.End.HandleKey(k, e.Selected)
	}
}
