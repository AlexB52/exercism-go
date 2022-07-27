package bowling

import (
	"errors"
)

type Game struct {
	CurrentFrameIndex int
	Pins              []int
	Frames            []*Frame
}

func NewGame() *Game {
	game := &Game{}
	game.Frames = make([]*Frame, 11)
	for i := 0; i <= 10; i++ {
		game.Frames[i] = &Frame{}
	}
	return game
}

func (g *Game) Roll(pins int) error {
	current := g.CurrentFrame()

	if g.IsFinished() || current.IsInvalidFrameRoll(pins) {
		return errors.New("invalid roll")
	}

	g.Pins = append(g.Pins, pins)
	current.Rolls = append(current.Rolls, pins)

	if current.Type() != Incomplete && g.CurrentFrameIndex < 10 {
		g.CurrentFrameIndex++
		g.CurrentFrame().PinsIndex = len(g.Pins)
	}

	return nil
}

func (g *Game) Score() (result int, err error) {
	if !g.IsFinished() {
		return 0, errors.New("unfinished game")
	}

	for i := 0; i < 10; i++ {
		idx := g.Frames[i].PinsIndex
		switch g.Frames[i].Type() {
		case Spare, Strike:
			result += g.Pins[idx] + g.Pins[idx+1] + g.Pins[idx+2]
		case Open:
			result += g.Pins[idx] + g.Pins[idx+1]

		}
	}
	return result, err
}

func (g *Game) IsFinished() bool {
	switch g.Frames[9].Type() {
	case Open:
		return len(g.Frames[9].Rolls) >= 2
	case Spare:
		return len(g.Frames[10].Rolls) >= 1
	case Strike:
		return len(g.Frames[10].Rolls) >= 2
	default:
		return false
	}
}

func (g *Game) CurrentFrame() *Frame {
	return g.Frames[g.CurrentFrameIndex]
}
