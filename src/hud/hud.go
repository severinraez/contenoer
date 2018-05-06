package hud

import (
	"time"
	"github.com/nsf/termbox-go"
	"github.com/severinraez/cotenoer/inventory"
	"github.com/severinraez/cotenoer/bundle"
	"fmt"
	"errors"
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

type state struct {
	SelectedBundleName string
	BundleOverviews []bundle.Overview
}

// @return the exit code
func Run(session inventory.Inventory) int {
	state, err := initState(session)

	if err != nil {
		fmt.Printf("Could not initialize HUD.\nCause: %v\n", err)
		return 1
	}

	initTermbox()
	defer teardownTermbox()

	reallocBackBuffer(termbox.Size())

	draw(state)

	consoleEvents := make(chan consoleEvent)
	bundleOverview := make(chan []bundle.Overview)

	go pollConsole(consoleEvents)
	go pollBundleOverviews(inventory.Bundles(session), 1000 * time.Millisecond, bundleOverview)

mainloop:
	for {
		select {
		case event := <- consoleEvents:
			isExitEvent := handleConsoleEvent(event)
			if isExitEvent {
				break mainloop
			}
		case bundleOverviews := <- bundleOverview:
			 	state.BundleOverviews = bundleOverviews

			draw(state)
		}
	}

	return 0
}

func initState(session inventory.Inventory) (state, error) {
	bundles := inventory.BundleNames(session)

	if len(bundles) == 0 {
		return state{}, errors.New("No bundles configured. Did you forget to add any or is COTENOER_SESSION not set correctly?")
	}

	return state{
		BundleOverviews: []bundle.Overview{},
		SelectedBundleName: bundles[0]}, nil
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

func draw(state state) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	bundles, bundlesW := bundleView(state.BundleOverviews)
	if len(bundles) > 0 {
		copyRect(backbuf, bbw, 5, 5, bundles, bundlesW)
	}

	backbuf[0] = termbox.Cell{Ch: 'a', Fg: termbox.ColorWhite}

	copy(termbox.CellBuffer(), backbuf)
	//_, h := termbox.Size()
	termbox.Flush()
}

func bundleView(bundles []bundle.Overview) ([]termbox.Cell, int) {
	const(
		numberColumn = iota
		nameColumn = iota
		containersColumn = iota
		containerSymbolColumn = iota
		resourcesColumn = iota
		columnCount = iota
	)

	linesPerBundle := 2
	columnSpacing := 1
	view := gridMake(columnCount, len(bundles) * linesPerBundle, columnSpacing)

	for i, bundle := range bundles {
		top := i * linesPerBundle;

		gridSet(fmt.Sprintf("%d", i + 1), numberColumn, top + 1, view)
		gridSet(fmt.Sprintf("%s  ", bundle.Name), nameColumn, top + 1, view)
		gridSet(fmt.Sprintf("%d", bundle.ActiveContainers), containersColumn, top + 1, view)

		gridSet(" _ ", containerSymbolColumn, top, view)
		gridSet("[_]", containerSymbolColumn, top + 1, view)

		gridSet(" |   |", resourcesColumn, top + 1, view)
	}

	return linesToRect(gridToLines(view))
}

// @return ([]cells, width)
func linesToRect(lines []string) ([]termbox.Cell, int) {
	width := 0
	for _, line := range lines {
		lineWidth := len(line)
		if lineWidth > width {
			width = lineWidth
		}
	}

	cells := make([]termbox.Cell, len(lines) * width)
	for i, line := range lines {
		baseIndex := i * width

		for j, charCode := range line {
			cells[baseIndex + j] = termbox.Cell{Ch: rune(charCode)}
		}
	}

	return cells, width
}

func copyRect(dest []termbox.Cell, destW int, x int, y int, src []termbox.Cell, srcW int) {
	srcH := len(src) / srcW
	for i := 0; i < srcW; i++ {
		for j := 0; j < srcH; j++ {
			dest[(x + i) + (y + j) * destW] = src[i + j * srcW]
		}
	}
}

func reallocBackBuffer(w, h int) {
	bbw, bbh = w, h
	backbuf = make([]termbox.Cell, w*h)
}

func pollBundleOverviews(bundles []inventory.Bundle, interval time.Duration, channel chan<- []bundle.Overview) {
	for {
		overviews := make([]bundle.Overview, len(bundles))
		for i, b := range bundles {
			overviews[i] = bundle.GetOverview(b)
		}
		channel <- overviews

		time.Sleep(interval)
	}
}
