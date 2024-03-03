package main

import (
	"fmt"
)

const (
	up     byte = 65
	kUp    byte = 107 // for vimers
	wUp    byte = 119 // for wasd
	down   byte = 66
	jDown  byte = 106 // for vimers
	sDown  byte = 115 // for wasd
	escape byte = 27
	enter  byte = 13
)

type Menu struct {
	Prompt    string
	CursorPos int
	Items     []*MenuItem
}

type MenuItem struct {
	Label      string
	ID         string
	Selectable bool
	SubMenu    *Menu
}

func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt: prompt,
		Items:  make([]*MenuItem, 0),
	}
}

func (m *Menu) Render() {
	setColor(ColorCyan)
	setBold()
	fmt.Print(m.Prompt, ": \n")
	clearTextStyle()
	for i, menuItem := range m.Items {
		prefix := " "
		if i == m.CursorPos {
			prefix = ">"
			setColor(ColorYellow)
		}
		fmt.Println(prefix, menuItem.Label)
		setColor("")
	}
	moveCursorUp(len(m.Items) + 1)
}

func (m *Menu) Load() string {
	m.Render()
	cursorOff()
	defer cursorOn()
	var key byte = 0
	for {
		key = getInput()
		switch key {
		case enter:
			return m.Items[m.CursorPos].ID
		case escape:
			return ""
		case up, kUp, wUp:
			m.CursorPos--
			if m.CursorPos < 0 {
				m.CursorPos = 0
			}
		case down, jDown, sDown:
			m.CursorPos++
			if m.CursorPos >= len(m.Items) {
				m.CursorPos = len(m.Items) - 1
			}
		}
		m.Render()
	}
}

var menuItems []*MenuItem = []*MenuItem{
	{Label: "First", ID: "1st"},
	{Label: "Second", ID: "2nd"},
	{Label: "Third", ID: "3rd"},
	{Label: "Fours", ID: "4rs"},
	{Label: "Fifs", ID: "5fs"},
}

func main() {
	m := NewMenu("General")
	m.Items = menuItems
	sel := m.Load()
	if sel != "" {
		fmt.Println("Selected menu id =", sel)
	}
}
