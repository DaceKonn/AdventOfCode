package helpers

type Point interface {
	GetH() int
	GetW() int
	SetH(h int)
	SetW(w int)
	Copy() Point
}

type DefaultPoint struct {
	h int
	w int
}

func NewDefaultPoint(h, w int) Point {
	return &DefaultPoint{h: h, w: w}
}

func (dp *DefaultPoint) GetH() int {
	return dp.h
}

func (dp *DefaultPoint) GetW() int {
	return dp.w
}

func (dp *DefaultPoint) SetH(h int) {
	dp.h = h
}

func (dp *DefaultPoint) SetW(w int) {
	dp.w = w
}

func (dp *DefaultPoint) Copy() Point {
	return &DefaultPoint{
		h: dp.h,
		w: dp.w,
	}
}
