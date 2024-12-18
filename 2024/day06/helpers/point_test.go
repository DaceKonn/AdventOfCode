package helpers

import "testing"

func TestPointShouldCreateAndUpdate(t *testing.T) {
	wantH, wantW := -1, -1
	var point Point = NewDefaultPoint(wantH, wantW)
	if point.GetH() != wantH || point.GetW() != wantW {
		t.Errorf("Fresh point doesn't match, wanted(%d,%d), got(%d,%d)", wantH, wantW, point.GetH(), point.GetW())
	}

	wantH, wantW = 2, 3
	point.SetH(wantH)
	point.SetW(wantW)
	if point.GetH() != wantH || point.GetW() != wantW {
		t.Errorf("Updated point doesn't match, wanted(%d,%d), got(%d,%d)", wantH, wantW, point.GetH(), point.GetW())
	}
}

func TestPointsShouldAffectEachOther(t *testing.T) {
	wantH, wantW := -1, -1
	var pointA Point = NewDefaultPoint(wantH, wantW)
	pointB := pointA.Copy()

	pointA.SetW(10)
	pointA.SetH(20)

	if pointB.GetH() != wantH || pointB.GetW() != wantW {
		t.Errorf("PointB got updated, wanted(%d,%d), got(%d,%d)", wantH, wantW, pointB.GetH(), pointB.GetW())
	}
}
