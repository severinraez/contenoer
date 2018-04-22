package hud

import (
	"github.com/nsf/termbox-go"
)


var backbuf []termbox.Cell
var bbw, bbh int

var runes = []rune{' ', '░', '▒', '▓', '█'}
var colors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorBlue,
	termbox.ColorMagenta,
	termbox.ColorCyan,
	termbox.ColorWhite,
}

type attrFunc func(int) (rune, termbox.Attribute, termbox.Attribute)

func Start() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	reallocBackBuffer(termbox.Size())
	update_and_redraw(-1, -1)

mainloop:
	for {
		mx, my := -1, -1
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break mainloop
			}
		case termbox.EventResize:
			reallocBackBuffer(ev.Width, ev.Height)
		}
		update_and_redraw(mx, my)
	}
}

func update_and_redraw(mx, my int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	backbuf[0] = termbox.Cell{Ch: 'a', Fg: termbox.ColorWhite}
	if mx != -1 && my != -1 {
		backbuf[bbw*my+mx] = termbox.Cell{Ch: runes[0], Fg: colors[0]}
	}
	copy(termbox.CellBuffer(), backbuf)
	//_, h := termbox.Size()
	termbox.Flush()
}

func reallocBackBuffer(w, h int) {
	bbw, bbh = w, h
	backbuf = make([]termbox.Cell, w*h)
}
