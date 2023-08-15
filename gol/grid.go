package gol

import (
	"fmt"
	"strings"
)

func NewGrid(dim uint) *Grid {
	grid := make([][]cell, dim+2)
	for i := range grid {
		grid[i] = make([]cell, dim+2)
	}
	return &Grid{grid, 1, dim}
}

func NewGridFromString(strGrid string) *Grid {
	cells := buildGridSlice(strGrid)
	grid := NewGrid(uint(len(cells)))
	for i, row := range cells {
		for j, col := range row {
			if col == 1 {
				grid.SetCellAt(uint(i+1), uint(j+1))
			}
		}
	}
	return grid
}

// Grid is the grid on which the game of life can be executed. For convenience,
// the grid is one-indexed. A new grid can be built by setting live cells after
// creating an empty grid using NewGrid or it can be created using string representation
// of the grid.
type Grid struct {
	grid      [][]cell
	low, high uint
}

func (g *Grid) String() string {
	var sb strings.Builder
	for _, row := range g.grid[g.low : g.high+1] {
		for _, col := range row[g.low : g.high+1] {
			sb.WriteString(fmt.Sprintf("%d ", col))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// SetCellAt turns the cell at (i, j) alive. Grid is 1 - indexed for both i and j.
// Invalid value for (i, j) causes silent return.
func (g *Grid) SetCellAt(row, col uint) {
	if row < g.low || row > g.high ||
		col < g.low || col > g.high {
		return
	}
	g.grid[row][col] = alive
}

// Advance initiates the creation of next generation. A new grid is returned for the new generation.
func (g *Grid) Advance() *Grid {
	ng := NewGrid(g.high)
	for i := g.low; i <= g.high; i++ {
		for j := g.low; j <= g.high; j++ {
			nc := g.neighbourCellCountAt(i, j)
			ng.grid[i][j] = g.grid[i][j].transition(nc)
		}
	}
	return ng
}

// neighbourCellCountAt counts the number of alive neighbouring cells for a
// given cell at (i, j). Grid being 1-indexed helps avoid branching.
func (g *Grid) neighbourCellCountAt(i, j uint) uint {
	return uint(
		g.grid[i-1][j-1] +
			g.grid[i-1][j] +
			g.grid[i-1][j+1] +
			g.grid[i][j-1] +
			g.grid[i][j+1] +
			g.grid[i+1][j-1] +
			g.grid[i+1][j] +
			g.grid[i+1][j+1])
}

// cell implements the cell behavior in a grid.
type cell uint

var (
	alive cell = 1
	dead  cell = 0
)

func (c cell) transition(nc uint) cell {
	switch c {
	case alive:
		if nc < 2 || nc > 3 {
			return dead
		}
	case dead:
		if nc == 3 {
			return alive
		}
	}
	return c
}

func buildGridSlice(strGrid string) [][]uint {
	var cells [][]uint
	cells = append(cells, []uint{})
	row, col, maxCol := 1, 0, 0
	for _, c := range strings.Trim(strGrid, "\n") {
		switch c {
		case ' ', '\t':
			continue
		case '\n':
			cells = append(cells, []uint{})
			row++
		case '0', '1':
			cells[row-1] = append(cells[row-1], uint(c)-'0')
		default:
			panic("bad grid character, only 0 and 1 allowed")
		}
	}
	for i, r := range cells {
		if i == 0 {
			maxCol = len(r)
			continue
		}
		col = len(r)
		if col != len(cells) || col != maxCol {
			fmt.Println("invalid grid input", col, maxCol, row)
			panic("xx")
		}
	}
	return cells
}
