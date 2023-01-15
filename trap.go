package main

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/resource"
	"github.com/quasilyte/gmath"
)

type trap struct {
	scene *ge.Scene

	pos gmath.Vec

	kind         tileKind
	triggerSound resource.AudioID

	sprite      *ge.Sprite
	animation   *ge.Animation
	playingAnim bool
	enabled     bool
}

func newTrap(pos gmath.Vec, t tileKind) *trap {
	return &trap{
		pos:     pos,
		kind:    t,
		enabled: true,
	}
}

func (t *trap) Init(scene *ge.Scene) {
	t.scene = scene

	var img resource.ImageID
	switch t.kind {
	case tileSpikeTrap:
		img = ImageSpikeTrap
		t.triggerSound = AudioSpikesTrap
	case tileFireTrap:
		img = ImageFireTrap
		t.triggerSound = AudioFireTrap
	case tileBearTrap:
		img = ImageBearTrap
		t.triggerSound = AudioBearTrap
	default:
		panic("unexpected tile kind used for trap")
	}
	t.sprite = scene.NewSprite(img)
	t.sprite.Pos.Base = &t.pos
	t.sprite.Pos.Offset.Y = -8
	t.sprite.Centered = false
	scene.AddGraphics(t.sprite)

	t.animation = ge.NewAnimation(t.sprite, -1)
	t.animation.SetSecondsPerFrame(0.03)
}

func (t *trap) Trigger() {
	t.playingAnim = true
	t.animation.Rewind()
	t.scene.Audio().PlaySound(t.triggerSound)
	if t.kind == tileBearTrap {
		t.enabled = false
	}
}

func (t *trap) SetVisibility(visibile bool) {
	t.sprite.Visible = visibile
}

func (t *trap) Update(delta float64) {
	if t.playingAnim && t.animation.Tick(delta) {
		if t.kind != tileBearTrap {
			t.animation.Rewind()
			t.sprite.FrameOffset.X = 0
		}
		t.playingAnim = false
	}
}

func (t *trap) IsDisposed() bool { return false }
