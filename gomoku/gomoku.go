package main

import (
	"github.com/nsf/termbox-go"
)

var width = 10
var height = 10

func draw_strings(x, y int, horizontal bool, output string, fg, bg termbox.Attribute) {
	for _, c := range output {
		termbox.SetCell(x, y, c, fg, bg)
		if horizontal {
			y++
		} else {
			x++
		}
	}
	termbox.Flush()
}

func draw_board() {
	fg := termbox.ColorWhite
	bg := termbox.ColorBlack

	tline := "+"
	mline := "|"
	for w := 0; w < width; w++ {
		tline = tline + "---+"
		mline = mline + "   |"
	}

	draw_strings(0, 0, false, tline, fg, bg)

	for h := 1; h < height*2; h += 2 {
		draw_strings(0, h, false, mline, fg, bg)
		draw_strings(0, h+1, false, tline, fg, bg)
	}
}

func draw_cursor() {
	fg := termbox.ColorWhite
	bg := termbox.ColorCyan

	draw_strings(1, 1, false, "   ", fg, bg)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)

	draw_board()
	draw_cursor()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlC {
				break loop
			}
		}
	}
}
