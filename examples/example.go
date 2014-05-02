package main

import (
	. "../"
	"fmt"
)

func main() {
	a := NewSet("colorsa")
	a.AddString("Red")
	a.AddString("Orange")
	a.AddString("Yellow")
	a.AddString("Green")
	a.AddString("Indigo")
	a.AddString("Blue")

	b := NewSet("colorsb")
	b.AddString("Purple")
	b.AddString("Red")
	b.AddString("Brown")
	b.AddString("Black")
	b.AddString("Yellow")

	fmt.Println(a.StringContains("Orange"))

	fmt.Println(a.UnionString(b, "a U b"))
	fmt.Println(a.IntersectString(b, "a X b"))

}
