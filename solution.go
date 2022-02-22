package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesType float64

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum SidesType) float64 {
	var sideValuePow = math.Pow(sideLen, 2)
	switch sidesNum {
	case SidesCircle:
		return sideValuePow * math.Pi
	case SidesTriangle:
		return sideValuePow * 1 / 2 * math.Sin(60*math.Pi/180)
	case SidesSquare:
		return sideValuePow
	default:
		return 0
	}
}
