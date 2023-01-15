package main

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/gesignal"
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/gmath"
)

type player struct {
	ctx *ge.Context

	pos gmath.Vec

	hp int

	input *input.Handler

	animation *ge.Animation

	EventMoved gesignal.Event[*player]
}

func newPlayer(pos gmath.Vec, h *input.Handler) *player {
	return &player{
		pos:   pos,
		hp:    70,
		input: h,
	}
}

func (p *player) Init(scene *ge.Scene) {
	p.ctx = scene.Context()

	sprite := scene.NewSprite(ImageGopher)
	sprite.Pos.Base = &p.pos
	scene.AddGraphics(sprite)

	p.animation = ge.NewRepeatedAnimation(sprite, -1)
	p.animation.SetAnimationSpan(1.5)
}

func (p *player) Update(delta float64) {
	p.animation.Tick(delta)

	oldPos := p.pos
	switch {
	case p.input.ActionIsJustPressed(ActionMoveRight):
		p.pos.X += 32
	case p.input.ActionIsJustPressed(ActionMoveLeft):
		p.pos.X -= 32
	case p.input.ActionIsJustPressed(ActionMoveDown):
		p.pos.Y += 32
	case p.input.ActionIsJustPressed(ActionMoveUp):
		p.pos.Y -= 32
	}
	p.pos.X = gmath.Clamp(p.pos.X, 16, p.ctx.WindowWidth-16)
	p.pos.Y = gmath.Clamp(p.pos.Y, 16, p.ctx.WindowHeight-16)
	if oldPos != p.pos {
		p.EventMoved.Emit(p)
	}
}

func (p *player) IsDisposed() bool { return false }
