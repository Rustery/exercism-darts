package darts

import "math"

type Circle struct {
	radius float64
	points int
}

var availableCircles = []Circle{
	{radius: 1, points: 10},
	{radius: 5, points: 5},
	{radius: 10, points: 1},
}

func isInCircle(x, y float64, c Circle) bool {
	return math.Sqrt(x*x+y*y) <= c.radius
}

func Score(x, y float64) int {
	for _, c := range availableCircles {
		if isInCircle(x, y, c) {
			return c.points
		}
	}
	return 0
}
