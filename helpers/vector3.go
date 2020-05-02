package helpers

import "math"

// Vector3 is a 3-dimensional vector structure
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// NewVector3 creates a new 3D vector
func NewVector3(x, y, z float32) *Vector3 {
	return &Vector3{x, y, z}
}

// Dot product of two vectors
func (v *Vector3) Dot(v2 *Vector3) float64 {
	return float64(v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z)
}

// Magnitude of a single vector
func (v *Vector3) Magnitude() float64 {
	return math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2) + math.Pow(float64(v.Z), 2))
}
