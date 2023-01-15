package main

import "github.com/quasilyte/gmath"

type tileKind int

const (
	tileEmpty tileKind = iota
	tileSpikeTrap
	tileFireTrap
	tileBearTrap
	tilePitTrap
	tileSugar
	tilePlayerSpawn
	numTileKinds
)

type levelMap struct {
	NumRows int
	NumCols int

	origin gmath.Vec

	tiles []tileKind
}

type tilePos struct {
	row int
	col int
}

func newLevelMap(numRows, numCols int) *levelMap {
	return &levelMap{
		NumRows: numRows,
		NumCols: numCols,
		tiles:   make([]tileKind, numRows*numCols),
	}
}

func (m *levelMap) Foreach(f func(row, col int) bool) {
	for row := 0; row < m.NumRows; row++ {
		for col := 0; col < m.NumCols; col++ {
			if !f(row, col) {
				return
			}
		}
	}
}

func (m *levelMap) GetTilePos(pos gmath.Vec) tilePos {
	realPos := pos.Sub(m.origin)
	return tilePos{
		col: int(realPos.X) / 32,
		row: int(realPos.Y) / 32,
	}
}

func (m *levelMap) GetPos(row, col int) gmath.Vec {
	pos := gmath.Vec{
		X: 32 * float64(col),
		Y: 32 * float64(row),
	}
	return pos.Add(m.origin)
}

func (m *levelMap) GetTile(row, col int) tileKind {
	return m.tiles[m.NumCols*row+col]
}

func (m *levelMap) SetTile(row, col int, t tileKind) {
	m.tiles[m.NumCols*row+col] = t
}
