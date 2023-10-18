package events

import "syscall/js"

type CommandBar struct {
	Div js.Value
}

func (cb *CommandBar) Hide() {
	addClass(cb.Div, "invisible")
	cb.Div.Set("innerText", ": ")
}

func (cb *CommandBar) Show() {
	removeClass(cb.Div, "invisible")
}

func (cb *CommandBar) Append(s string) {
	text := cb.Div.Get("innerText").String()
	cb.Div.Set("innerText", text+s)
}

func (cb *CommandBar) RemoveLast() {
	text := cb.Div.Get("innerText").String()
	cb.Div.Set("innerText", text[0:len(text)-1])
}

func (cb *CommandBar) HandleKey(k string) {
	//fmt.Println("ecb HandleKey")
	if k == "i" {
	} else if k == "Escape" {
		Focus = "div"
		cb.Hide()
	} else if k == "Backspace" {
		cb.RemoveLast()
	} else if k == "Enter" {
		//text := cb.Div.Get("innerText").String()
		Focus = "div"
		cb.Hide()
	} else {
		cb.Append(k)
	}
}
