package gomenu

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
	Prompt         string
	CursorPos      int
	Items          []*MenuItem
	PrimaryColor   string
	SecondaryColor string
}

type MenuItem struct {
	Label      string
	ID         string
	Unpickable bool
	SubMenu    *Menu
}

func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt:         prompt,
		Items:          make([]*MenuItem, 0),
		PrimaryColor:   ColorCyan,
		SecondaryColor: ColorYellow,
	}
}

func (m *Menu) Up() {
	m.CursorPos--
	if m.CursorPos < 0 {
		m.CursorPos = 0
	}
	for m.CursorPos > 0 && m.Items[m.CursorPos].Unpickable {
		m.CursorPos--
	}
	if m.Items[m.CursorPos].Unpickable {
		m.Down()
	}
	m.Render()
}

func (m *Menu) Down() {
	m.CursorPos++
	if m.CursorPos >= len(m.Items) {
		m.CursorPos = len(m.Items) - 1
	}
	for m.CursorPos < len(m.Items)-1 && m.Items[m.CursorPos].Unpickable {
		m.CursorPos++
	}
	if m.Items[m.CursorPos].Unpickable {
		m.Up()
	}
	m.Render()
}

func (m *Menu) Render() {
	setColor(m.PrimaryColor)
	setBold()
	fmt.Print(m.Prompt, ": \n")
	clearTextStyle()
	for i, menuItem := range m.Items {
		if menuItem.Unpickable {
			setColor(m.PrimaryColor)
			fmt.Println(menuItem.Label)
			clearTextStyle()
		} else {
			prefix := " "
			if i == m.CursorPos {
				prefix = ">"
				setColor(m.SecondaryColor)
			}
			fmt.Println(prefix, menuItem.Label)
			setColor("")
		}
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
			m.Up()
		case down, jDown, sDown:
			m.Down()
		}
		m.Render()
	}
}
