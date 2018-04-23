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

	t.Log(canvas)

	copyRect(canvas, w, 3, 3, quad, quadW, quadH)

	t.Log(canvas)

	AssertEqual(canvas[0].Ch, rune(0), t)

	quad0InCanvas := 3 + 3 * w
	AssertEqual(canvas[quad0InCanvas].Ch, 'a', t)

	quad1InCanvas := 4 + 3 * w
	AssertEqual(canvas[quad1InCanvas].Ch, 'b', t)

	quad2InCanvas := 3 + 4 * w
	AssertEqual(canvas[quad2InCanvas].Ch, 'c', t)
}

