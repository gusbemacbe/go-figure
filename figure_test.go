package figure

import (
	"os"
	"testing"
)

func TestExampleAlphabet(t *testing.T) {
	myFigure := NewFigure("Hello world", "alphabet", true)
	t.Logf("\n%v", myFigure)
	// Output:
	// H  H     l l                     l    d
	// H  H     l l                     l    d
	// HHHH eee l l ooo   w   w ooo rrr l  ddd
	// H  H e e l l o o   w w w o o r   l d  d
	// H  H ee  l l ooo    w w  ooo r   l  ddd
}

func TestExampleDefault(t *testing.T) {
	myFigure := NewFigure("Hello world", "", true)
	t.Logf("\n%v", myFigure)
	// Output:
	//   _   _          _   _                                       _       _
	//  | | | |   ___  | | | |   ___     __      __   ___    _ __  | |   __| |
	//  | |_| |  / _ \ | | | |  / _ \    \ \ /\ / /  / _ \  | '__| | |  / _` |
	//  |  _  | |  __/ | | | | | (_) |    \ V  V /  | (_) | | |    | | | (_| |
	//  |_| |_|  \___| |_| |_|  \___/      \_/\_/    \___/  |_|    |_|  \__,_|
}

func TestExampleColor(t *testing.T) {
	NewColorFigure("Hello World", "", "green", false).Print()
}

func TestExampleScroll(t *testing.T) {
	figure := NewFigure("Hello world", "", true)
	figure.Scroll(5000, 200, "right")
	figure.Scroll(5000, 100, "left")
}

func TestExampleBlink(t *testing.T) {
	NewFigure("Give your reasons", "doom", true).Blink(10000, 500, -1)
}

func TestExampleDance(t *testing.T) {
	NewFigure("It's been waiting for you", "larry3d", true).Dance(10000, 500)
}

func TestExampleNewFigureWithFont(t *testing.T) {
	f, _ := os.OpenFile("./fonts/doom.flf", os.O_RDONLY, 0)
	NewFigureWithFont("Hello World!", f, true).Print()
}

func TestExampleWrite(t *testing.T) {
	figure := NewColorFigure("Hello World", "", "green", false)
	Write(os.Stdout, figure)
}
