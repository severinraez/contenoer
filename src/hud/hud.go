package hud

import (
	"github.com/nsf/termbox-go"
	"github.com/severinraez/cotenoer/inventory"
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

type consoleEvent struct {
	Kind string
	Key string
	ResizeWidth int
	ResizeHeight int
}

func Start(session inventory.Inventory) {
	initTermbox()
	defer teardownTermbox()

	reallocBackBuffer(termbox.Size())
	updateAndRedraw(-1, -1)

	consoleEvents := make(chan consoleEvent)

	go pollConsole(consoleEvents)

mainloop:
	for {
		select {
		case event := <- consoleEvents:
			isExitEvent := handleConsoleEvent(event)
			if isExitEvent {
				break mainloop
			}
		}
	}
}

func handleConsoleEvent(event consoleEvent) bool {
	switch event.Kind {
	case "key":
		if event.Key == "esc" {
			return true
		}
	case "resize":
		reallocBackBuffer(event.ResizeWidth, event.ResizeHeight)
	}
	updateAndRedraw(-1, -1)
	return false
}

func pollConsole(channel chan<- consoleEvent) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				channel <- consoleEvent{
					Kind: "key",
					Key: "esc"}
				return
			}
		case termbox.EventResize:
			channel <- consoleEvent{
				Kind: "resize",
				ResizeWidth: ev.Width,
				ResizeHeight: ev.Height}
		}

	}
}

func initTermbox() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc)
}

func teardownTermbox() {
	termbox.Close()
}

func updateAndRedraw(mx, my int) {
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
