package main

import (
	"github.com/nsf/termbox-go"
)

var width = 10
var height = 10

var cur_x = 1
var cur_y = 1

var def_fg = termbox.ColorWhite
var def_bg = termbox.ColorBlack

var cur_fg = termbox.ColorWhite
var cur_bg = termbox.ColorCyan

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
	tline := "+"
	mline := "|"
	for w := 0; w < width; w++ {
		tline = tline + "---+"
		mline = mline + "   |"
	}

	draw_strings(0, 0, false, tline, def_fg, def_bg)

	for h := 1; h < height*2; h += 2 {
		draw_strings(0, h, false, mline, def_fg, def_bg)
		draw_strings(0, h+1, false, tline, def_fg, def_bg)
	}
}

func draw_cursor() {
	draw_strings(cur_x, cur_y, false, "   ", cur_fg, cur_bg)
}

func up_cursor() {
	if cur_y == 1 {
		return
	}

	draw_strings(cur_x, cur_y, false, "   ", def_fg, def_bg)
	cur_y -= 2
	draw_strings(cur_x, cur_y, false, "   ", cur_fg, cur_bg)
}

func down_cursor() {
	if cur_y == height*2-1 {
		return
	}

	draw_strings(cur_x, cur_y, false, "   ", def_fg, def_bg)
	cur_y += 2
	draw_strings(cur_x, cur_y, false, "   ", cur_fg, cur_bg)
}

func left_cursor() {
	if cur_x == 1 {
		return
	}

	draw_strings(cur_x, cur_y, false, "   ", def_fg, def_bg)
	cur_x += 4
	draw_strings(cur_x, cur_y, false, "   ", cur_fg, cur_bg)
}

func right_cursor() {
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
			switch ev.Key {
			case termbox.KeyCtrlC:
				break loop
			case termbox.KeyArrowUp:
				//Move to up cursor
				up_cursor()
			case termbox.KeyArrowDown:
				//Move to down cursor
				down_cursor()
			case termbox.KeyArrowLeft:
				//Move to left cursor
				left_cursor()
			case termbox.KeyArrowRight:
				//Move to right cursor
			}
		}
	}
}
