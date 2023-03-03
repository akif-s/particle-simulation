package main

import (
	"math"
)

type Vector2 struct {
	x float64
	y float64
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{
		x: x,
		y: y,
	}
}

// Multiplies the each element of vector2 with a
func (v Vector2) Product(a float64) Vector2 {
	return NewVector2(v.x*a, v.y*a)
}

// Adds two vector
func (v1 Vector2) Sum(v2 Vector2) Vector2 {
	return NewVector2(v1.x+v2.x, v1.y+v2.y)
}

// Substracts v2 from v1
func (v1 Vector2) Substract(v2 Vector2) Vector2 {
	return NewVector2(v1.x-v2.x, v1.y-v2.y)
}

// Calculates the magnitude of the vector
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Calculates the dot product of two vectors
func Dot(v1 Vector2, v2 Vector2) float64 {
	return v1.x*v2.x + v1.y*v2.y
}

// Calculates the angle between two vectors
func Angle(v1 Vector2, v2 Vector2) float64 {
	return math.Acos(Dot(v1, v2) / (v1.Magnitude() * v2.Magnitude()))
}
