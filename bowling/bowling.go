package bowling

import (
	"errors"
)

type FrameType int

const (
	Incomplete FrameType = iota
	Spare
	Open
	Strike
)

type Game struct {
	CurrentFrameIndex int
	Pins              []int
	Frames            []*Frame
}

type Frame struct {
	PinsIndex int
	Rolls     []int
}

func (f *Frame) Type() FrameType {
	switch {
	case len(f.Rolls) == 2 && f.Pins() == 10:
		return Spare
	case len(f.Rolls) == 1 && f.Pins() == 10:
		return Strike
	case len(f.Rolls) == 2:
		return Open
	default:
		return Incomplete
	}
}

func (f *Frame) Pins() int {
	var result int
	for _, r := range f.Rolls {
		result += r
	}
	return result
}

func NewGame() *Game {
	game := &Game{}
	game.Frames = make([]*Frame, 11)
	for i := 0; i <= 10; i++ {
		game.Frames[i] = &Frame{}
	}
	return game
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

func (g *Game) Roll(pins int) error {
	if g.IsFinished() || pins < 0 || pins > 10 {
		return errors.New("invalid roll")
	}

	current := g.CurrentFrame()

	if current.Pins() < 10 && current.Pins()+pins > 10 {
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

func (g *Game) Score() (int, error) {
	if !g.IsFinished() {
		return 0, errors.New("unfinised game")
	}

	var result int
	for i := 0; i < 10; i++ {
		result += g.Frames[i].Pins()
		nextPinIdx := g.Frames[i+1].PinsIndex
		switch g.Frames[i].Type() {
		case Spare:
			result += g.Pins[nextPinIdx]
		case Strike:
			result += g.Pins[nextPinIdx] + g.Pins[nextPinIdx+1]
		}
	}
	return result, nil
}
