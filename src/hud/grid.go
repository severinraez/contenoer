package hud

import (
	"strings"
)

type Grid struct {
	w int
	h int
	cells []string
	columnSpacing int
}

func gridMake(w int, h int, columnSpacing int) Grid {
	return Grid{
		w: w,
		h: h,
		columnSpacing: columnSpacing,
		cells: make([]string, w*h)}
}

func gridSet(value string, x int, y int, grid Grid) {
	index := y * grid.w + x

	grid.cells[index] = value
}

func gridGet(x int, y int, grid Grid) string {
	index := y * grid.w + x

	return grid.cells[index]
}

func gridToLines(grid Grid) []string {
	columnWidths := make([]int, grid.w)
	for x := 0; x < grid.w; x++ {
		maxLength := -1
		for y := 0; y < grid.h; y++ {
			lineLength := len(gridGet(x, y, grid))
			if lineLength > maxLength {
				maxLength = lineLength
			}
		}
		columnWidths[x] = maxLength
	}

	lines := make([]string, grid.h)
	for x := 0; x < grid.w; x++ {
		for y := 0; y < grid.h; y++ {
			content := gridGet(x, y, grid)

			padding := columnWidths[x] - len(content)
			spaces := strings.Repeat(" ", padding + grid.columnSpacing)

			lines[y] += content + spaces
		}
	}

	return lines
}
