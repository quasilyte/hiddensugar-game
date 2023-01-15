package main

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type roomTile struct {
	pos gmath.Vec

	sprite *ge.Sprite

	trap *trap
}

func newRoomTile(pos gmath.Vec) *roomTile {
	return &roomTile{
		pos: pos,
	}
}

func (t *roomTile) Init(scene *ge.Scene) {
	t.sprite = scene.NewSprite(ImageRoomTile)
	t.sprite.Pos.Base = &t.pos
	t.sprite.Centered = false
	t.sprite.FlipHorizontal = scene.Rand().Bool()
	t.sprite.FlipVertical = scene.Rand().Bool()
	scene.AddGraphics(t.sprite)
}

func (t *roomTile) Update(delta float64) {
}

func (t *roomTile) IsDisposed() bool { return false }

func (t *roomTile) SetVisibility(visible bool) {
	t.sprite.Visible = visible
	if t.trap != nil {
		t.trap.SetVisibility(visible)
	}
}
