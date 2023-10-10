// darts package is a game where the target rewards 4 different amounts of points,
// depending on where the dart lands
package darts

import "math"

// Circle is a struct that represents a circle of the darts target
type Circle struct {
	radius float64
	points int
}

// Constants for the 3 circles of the darts target
var BIG_CIRCLE = Circle{radius: 1, points: 10}
var MIDDLE_CIRCLE = Circle{radius: 5, points: 5}
var SMALL_CIRCLE = Circle{radius: 10, points: 1}

// isInCircle checks if a point is inside a circle
func isInCircle(x, y float64, c Circle) bool {
	return math.Sqrt(x*x+y*y) <= c.radius
}

// Score calculates the score of a dart throw
func Score(x, y float64) int {
	switch {
	case isInCircle(x, y, BIG_CIRCLE):
		return BIG_CIRCLE.points
	case isInCircle(x, y, MIDDLE_CIRCLE):
		return MIDDLE_CIRCLE.points
	case isInCircle(x, y, SMALL_CIRCLE):
		return SMALL_CIRCLE.points
	default:
		return 0
	}
}
