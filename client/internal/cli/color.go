package cli

import (
	"fmt"
	"math/rand"
)

type Color string

const (
	Reset          = "\x1b[0m"
	Red      Color = "\x1b[31m"
	Green    Color = "\x1b[32m"
	Yellow   Color = "\x1b[33m"
	Blue     Color = "\x1b[34m"
	Magenta  Color = "\x1b[35m"
	Cyan     Color = "\x1b[36m"
	GrayS    Color = "\x1b[90m"
	RedS     Color = "\x1b[91m"
	GreenS   Color = "\x1b[92m"
	YellowS  Color = "\x1b[93m"
	BlueS    Color = "\x1b[94m"
	MagentaS Color = "\x1b[95m"
	CyanS    Color = "\x1b[96m"
	WhiteS   Color = "\x1b[97m"
)

var ListOfColors []Color = []Color{Red, RedS, Green, GreenS, Yellow, YellowS, Blue, BlueS, Magenta, MagentaS, Cyan, CyanS, WhiteS, GrayS}

// Clear Color list via console
func clearsColorList() {
	fmt.Print("\x1b[H\x1b[J")
}

// Colorise String
func StringColorize(str string, c Color) string {
	return string(c) + str + Reset
}

// Generate Random Color
func RandColor() Color {
	return ListOfColors[rand.Intn(len(ListOfColors))]
}
