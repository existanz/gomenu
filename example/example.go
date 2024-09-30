package main

import (
	"fmt"

	"github.com/existanz/gomenu"
)

func main() {
	m := gomenu.NewMenu("Choose your destiny")
	m.AddItem("1", "Novice")
	m.AddItem("2", "Veteran")
	m.AddItem("3", "Evil")
	m.AddUnpickableItem("", "First Delimiter")
	m.AddItem("4", "Options")
	m.AddItem("5", "Killer Code")
	m.AddUnpickableItem("", "---------------")
	m.AddItem("6", "Exit")

	sm1 := gomenu.NewMenu("Novice")
	m.GetMenuItem("1").SubMenu = sm1
	sm1.AddItem("1_1", "Easy")
	sm1.AddItem("1_2", "Rookie")
	sm1.AddItem("1_3", "Beginner")

	sm2 := gomenu.NewMenu("Veteran")
	m.GetMenuItem("2").SubMenu = sm2
	sm2.AddItem("2_1", "Medium")
	sm2.AddItem("2_2", "Normal")
	sm2.AddItem("2_3", "Average")

	sm3 := gomenu.NewMenu("Evil")
	m.GetMenuItem("3").SubMenu = sm3
	sm3.AddItem("3_1", "Hard")
	sm3.AddItem("3_2", "Expert")
	sm3.AddItem("3_3", "Difficult")

	sel := m.Load()
	fmt.Println("Selected menu ID", sel)
}
