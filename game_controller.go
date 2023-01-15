package main

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/gmath"
)

type gameController struct {
	scene *ge.Scene

	levelMap *levelMap

	trapDamage [numTileKinds]int
	tiles      [][]*roomTile

	sugar        *ge.Sprite
	sugarTilePos tilePos

	input *input.Handler
}

func newGameController(h *input.Handler) *gameController {
	c := &gameController{
		input: h,
	}
	for i := range c.trapDamage {
		c.trapDamage[i] = 1
	}
	return c
}

func (c *gameController) Init(scene *ge.Scene) {
	c.scene = scene

	playerSpawnPos := gmath.Vec{
		X: (scene.Context().WindowWidth / 2) - 16,
		Y: scene.Context().WindowHeight / 2,
	}

	c.initLevel(playerSpawnPos)

	p := newPlayer(playerSpawnPos, c.input)
	scene.AddObject(p)
	p.EventMoved.Connect(nil, c.onPlayerMoved)

	c.updateTilesVisibility(c.levelMap.GetTilePos(playerSpawnPos))
}

func (c *gameController) Update(delta float64) {}

func (c *gameController) onPlayerMoved(p *player) {
	tpos := c.levelMap.GetTilePos(p.pos)

	switch k := c.levelMap.GetTile(tpos.row, tpos.col); k {
	case tileSpikeTrap, tileFireTrap, tileBearTrap:
		trap := c.tiles[tpos.row][tpos.col].trap
		if trap.enabled {
			trap.Trigger()
			p.hp -= c.trapDamage[k]
			c.trapDamage[k] *= 2
		}
	case tileSugar:
		c.onVictory()
	}

	c.updateTilesVisibility(tpos)

	if p.hp <= 0 {
		c.onDefeat()
	}
}

func (c *gameController) updateTilesVisibility(tpos tilePos) {
	for deltaRow := -1; deltaRow <= 1; deltaRow++ {
		for deltaCol := -1; deltaCol <= 1; deltaCol++ {
			row := gmath.Clamp(tpos.row+deltaRow, 0, c.levelMap.NumRows-1)
			col := gmath.Clamp(tpos.col+deltaCol, 0, c.levelMap.NumCols-1)
			tile := c.tiles[row][col]
			tile.SetVisibility(true)
			if c.sugarTilePos.row == row && c.sugarTilePos.col == col {
				c.sugar.Visible = true
			}
		}
	}
}

func (c *gameController) onDefeat() {
	// Restarts the game from the beginning.
	c.scene.Context().ChangeScene(newGameController(c.input))
}

func (c *gameController) onVictory() {
	c.scene.Audio().PlaySound(AudioVictorySound)
	c.onDefeat()
}

func (c *gameController) initLevel(playerSpawnPos gmath.Vec) {
	windowHeight := c.scene.Context().WindowHeight
	windowWidth := c.scene.Context().WindowWidth

	numRows := int(windowHeight / 32)
	numCols := int(windowWidth / 32)
	c.levelMap = newLevelMap(numRows, numCols)

	c.tiles = make([][]*roomTile, numRows)
	for i := range c.tiles {
		c.tiles[i] = make([]*roomTile, numCols)
	}

	for offsetY := 0.0; offsetY < windowHeight; offsetY += 32 {
		for offsetX := 0.0; offsetX < windowWidth; offsetX += 32 {
			pos := gmath.Vec{X: offsetX, Y: offsetY}
			tile := newRoomTile(pos)
			c.scene.AddObject(tile)
			c.tiles[int(pos.Y)/32][int(pos.X)/32] = tile
		}
	}

	playerTilePos := c.levelMap.GetTilePos(playerSpawnPos)
	c.levelMap.SetTile(playerTilePos.row, playerTilePos.col, tilePlayerSpawn)

	picker := gmath.NewRandPicker[tileKind](c.scene.Rand())
	picker.AddOption(tileEmpty, 5)
	picker.AddOption(tileSpikeTrap, 1)
	picker.AddOption(tileFireTrap, 1)
	picker.AddOption(tileBearTrap, 1)
	c.levelMap.Foreach(func(row, col int) bool {
		if row == playerTilePos.row && col == playerTilePos.col {
			return true
		}
		kind := picker.Pick()
		c.levelMap.SetTile(row, col, kind)
		if kind == tileEmpty {
			return true
		}
		pos := c.levelMap.GetPos(row, col)
		trap := newTrap(pos, kind)
		c.scene.AddObject(trap)
		c.tiles[row][col].trap = trap
		return true
	})

	emptyTiles := make([]tilePos, 0, 32)
	c.levelMap.Foreach(func(row, col int) bool {
		kind := c.levelMap.GetTile(row, col)
		if kind == tileEmpty {
			emptyTiles = append(emptyTiles, tilePos{row: row, col: col})
		}
		return true
	})

	// Try to pick a pos that is not too close to the player starting location.
	var sugarTilePos tilePos
	for i := 0; i < 5; i++ {
		sugarTilePos = gmath.RandElem(c.scene.Rand(), emptyTiles)
		dist := c.levelMap.GetPos(sugarTilePos.row, sugarTilePos.col).DistanceTo(playerSpawnPos)
		if dist >= 128 {
			break
		}
	}
	sugarSprite := c.scene.NewSprite(ImageSugar)
	sugarSprite.Pos.Offset = c.levelMap.GetPos(sugarTilePos.row, sugarTilePos.col)
	sugarSprite.Centered = false
	sugarSprite.Visible = false
	c.sugar = sugarSprite
	c.sugarTilePos = sugarTilePos
	c.scene.AddGraphics(sugarSprite)
	c.levelMap.SetTile(sugarTilePos.row, sugarTilePos.col, tileSugar)

	for _, rowTiles := range c.tiles {
		for _, tile := range rowTiles {
			tile.SetVisibility(false)
		}
	}
}
