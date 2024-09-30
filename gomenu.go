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

// NewMenu Creates a new menu with the given prompt and default colors.
//
// prompt - The text to be displayed as the menu prompt.
//
// Returns a pointer to the newly created menu.
func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt:         prompt,
		Items:          make([]*MenuItem, 0),
		PrimaryColor:   ColorCyan,
		SecondaryColor: ColorYellow,
	}
}

// AddItem adds a new menu item to the menu.
// To add unpickable items, use AddUnpickableItem instead.
//
// id - MenuItem Id.
// label - MenuItem Label.
//
// No return values.
func (m *Menu) AddItem(id, label string) {
	m.Items = append(m.Items, &MenuItem{
		Label: label,
		ID:    id,
	})
}

// AddUnpickableItem adds a new unpickable menu item to the menu.
//
// id - MenuItem Id.
// label - MenuItem Label.
//
// No return values.
func (m *Menu) AddUnpickableItem(id, label string) {
	m.Items = append(m.Items, &MenuItem{
		Label:      label,
		ID:         id,
		Unpickable: true,
	})
}

// GetMenuItem returns the menu item with the given id.
//
// id - The id of the menu item to return.
//
// Returns a pointer to the menu item, or nil if the item is not found.
func (m *Menu) GetMenuItem(id string) *MenuItem {
	for _, item := range m.Items {
		if item.ID == id {
			return item
		}
	}
	return nil
}

// Load displays the menu and waits for user input to navigate and select an item.
//
// No parameters.
// Returns the ID of the selected menu item as a string, or an empty string if the user escapes.
func (m *Menu) Load() string {
	m.render()
	cursorOff()
	defer cursorOn()
	var key byte = 0
	for {
		key = getInput()
		switch key {
		case enter:
			curItem := m.Items[m.CursorPos]
			if curItem.SubMenu != nil {
				clearScreen()
				return curItem.SubMenu.Load()
			}
			return curItem.ID
		case escape:
			return ""
		case up, kUp, wUp:
			m.up()
		case down, jDown, sDown:
			m.down()
		}
		m.render()
	}
}

// up moves the cursor up in the menu.
//
// No parameters.
// No return values.
func (m *Menu) up() {
	m.CursorPos--
	if m.CursorPos < 0 {
		m.CursorPos = 0
	}
	for m.CursorPos > 0 && m.Items[m.CursorPos].Unpickable {
		m.CursorPos--
	}
	if m.Items[m.CursorPos].Unpickable {
		m.down()
	}
	m.render()
}

// down Moves the cursor down in the menu.
//
// No parameters.
// No return values.
func (m *Menu) down() {
	m.CursorPos++
	if m.CursorPos >= len(m.Items) {
		m.CursorPos = len(m.Items) - 1
	}
	for m.CursorPos < len(m.Items)-1 && m.Items[m.CursorPos].Unpickable {
		m.CursorPos++
	}
	if m.Items[m.CursorPos].Unpickable {
		m.up()
	}
	m.render()
}

// render Displays the menu and its items with proper styling and cursor positioning.
//
// No parameters.
// No return values.
func (m *Menu) render() {
	clearLine()
	setColor(m.PrimaryColor)
	setBold()
	fmt.Print(m.Prompt, ": \n")
	clearTextStyle()
	for i, menuItem := range m.Items {
		clearLine()
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
