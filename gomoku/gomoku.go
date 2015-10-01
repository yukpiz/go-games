package main

import (
	"github.com/nsf/termbox-go"
)

var width = 10
var height = 10
var span_x = 4
var span_y = 2

var mes_x = 0
var mes_y = 0

var cur_x = 1
var cur_y = 1

var def_fg = termbox.ColorWhite
var def_bg = termbox.ColorBlack

var cur_fg = termbox.ColorWhite
var cur_bg = termbox.ColorCyan

var mes_fg = termbox.ColorWhite
var mes_bg = termbox.ColorBlack

var err_fg = termbox.ColorRed
var err_bg = termbox.ColorBlack

type cell struct {
	x     int
	y     int
	value string
}

var cells []cell

func init_cells() {
	for y := 1; y < height+1; y++ {
		for x := 1; x < width+1; x++ {
			cells = append(cells, cell{x, y, " "})
		}
	}
}

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

	mes_x = 0
	mes_y = height*2 + 1
}

func draw_cursor() {
	draw_strings(cur_x, cur_y, false, "   ", cur_fg, cur_bg)
}

func draw_message(value string) {
	draw_strings(mes_x, mes_y, false, value, mes_fg, mes_bg)
}

func draw_error(value string) {
	draw_strings(mes_x, mes_y, false, value, err_fg, err_bg)
}

func draw_cell(x, y int, value string) {
	draw_strings(x, y, false, value, cur_fg, cur_bg)
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
	cur_x -= 4
	draw_strings(cur_x, cur_y, false, "   ", cur_fg, cur_bg)
}

func right_cursor() {
	if cur_x == span_x*(width-1)+1 {
		return
	}

	draw_strings(cur_x, cur_y, false, "   ", def_fg, def_bg)
	cur_x += 4
	draw_strings(cur_x, cur_y, false, "   ", cur_fg, cur_bg)
}

func switch_cell(x, y int, value string) {
	x, y = to_cur_pos(x, y)
	draw_cell(x, y, " "+value+" ")
}

func to_cur_pos(board_x, board_y int) (cur_x, cur_y int) {
	cur_x = board_x*span_x - 3
	cur_y = board_y*span_y - 1
	return
}

func to_board_pos(cur_x, cur_y int) (board_x, board_y int) {
	board_x = cur_x/span_x + 1
	board_y = cur_y/span_y + 1
	return
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)

	init_cells()

	draw_board()
	draw_cursor()

	draw_message("Hi Player!! It's your turn. Please press the <Space> to select the cell.")

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
				right_cursor()
			case termbox.KeySpace:
				x, y := to_board_pos(cur_x, cur_y)
				switch_cell(x, y, "o")
			}
		}
	}
}
