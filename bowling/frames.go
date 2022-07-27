package bowling

type FrameType int

const (
	Incomplete FrameType = iota
	Spare
	Open
	Strike
)

type Frame struct {
	PinsIndex int
	Rolls     []int
}

func (f *Frame) Type() FrameType {
	switch {
	case len(f.Rolls) == 2 && f.Throws() == 10:
		return Spare
	case len(f.Rolls) == 1 && f.Throws() == 10:
		return Strike
	case len(f.Rolls) == 2:
		return Open
	default:
		return Incomplete
	}
}

func (f *Frame) Throws() int {
	var result int
	for _, r := range f.Rolls {
		result += r
	}
	return result
}

func (f *Frame) IsInvalidFrameRoll(pins int) bool {
	return pins < 0 || pins > 10 || f.Throws() < 10 && f.Throws()+pins > 10
}
