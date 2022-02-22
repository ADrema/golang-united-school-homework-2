package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesType int

const (
	circleSides     = 0
	triangleSides   = 3
	foursquareSides = 4
)

func CalcSquare(sideLen float64, sidesNum SidesType) float64 {
	var sideValuePow = math.Pow(sideLen, 2)
	switch sidesNum {
	case circleSides:
		return sideValuePow * math.Pi
	case triangleSides:
		return sideValuePow * (1 / 2) * math.Sin(60*math.Pi/180)
	case foursquareSides:
		return sideValuePow
	default:
		return 0
	}
}
