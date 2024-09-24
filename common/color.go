/*
Copyright © 2024 ktsnkmrcom
*/
package common

import (
	"fmt"
	"math/rand"
)

var (
	Red        = "ffcdd2"
	Pink       = "f8bbd0"
	Purple     = "e1bee7"
	Deeppurple = "d1c4e9"
	Indigo     = "c5cae9"
	Blue       = "bbdefb"
	Lightblue  = "b3e5fc"
	Cyan       = "b2ebf2"
	Teal       = "b2dfdb"
	Green      = "c8e6c9"
	Lightgreen = "dcedc8"
	Lime       = "f0f4c3"
	Yellow     = "fff9c4"
	Amber      = "ffecb3"
	Orange     = "ffe0b2"
	Deeporange = "ffccbc"
	Brown      = "d7ccc8"
	Gray       = "f5f5f5"
	Bluegray   = "cfd8dc"
)

func RandomColor() string {
	colors := [...]string{
		Red,
		Pink,
		Purple,
		Deeppurple,
		Indigo,
		Blue,
		Lightblue,
		Cyan,
		Teal,
		Green,
		Lightgreen,
		Lime,
		Yellow,
		Amber,
		Orange,
		Deeporange,
		Brown,
		Gray,
		Bluegray,
	}
	r := rand.Intn(19)
	fmt.Println("Selected Color:" + colors[r])
	return colors[r]
}
