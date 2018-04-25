package hud

import (
	"testing"
	"github.com/nsf/termbox-go"
	. "github.com/severinraez/cotenoer/testhelper"
)

func TestCopyRect(t *testing.T) {
	w, h := 10, 10
	canvas := make([]termbox.Cell, w*h)

	quadW, quadH := 2, 2
	quad := make([]termbox.Cell, quadW*quadH)
	quad[0].Ch = 'a'
	quad[1].Ch = 'b'
	quad[2].Ch = 'c'
	quad[3].Ch = 'd'

	copyRect(canvas, w, 3, 3, quad, quadW)

	AssertEqual(canvas[0].Ch, rune(0), t)

	quad0InCanvas := 3 + 3 * w
	AssertEqual(canvas[quad0InCanvas].Ch, 'a', t)

	quad1InCanvas := 4 + 3 * w
	AssertEqual(canvas[quad1InCanvas].Ch, 'b', t)

	quad2InCanvas := 3 + 4 * w
	AssertEqual(canvas[quad2InCanvas].Ch, 'c', t)
}

func TestLinesToRect(t *testing.T) {
	lines := []string{"", "a", "ab"}

	cells, rectWidth := linesToRect(lines)

	AssertEqualH(2 * 3, len(cells), "Cell array length", t)
	AssertEqualH(2, rectWidth, "Rect width", t)

	empty := rune(0)
	AssertEqual(cells[0].Ch, empty, t)
	AssertEqual(cells[1].Ch, empty, t)
	AssertEqual(cells[2].Ch, 'a', t)
	AssertEqual(cells[3].Ch, empty, t)
	AssertEqual(cells[4].Ch, 'a', t)
	AssertEqual(cells[5].Ch, 'b', t)
}

