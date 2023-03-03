package main

import "math"

type Vector2 struct {
	x float64
	y float64
}

func (v *Vector2) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Dot(v1 Vector2, v2 Vector2) float64 {
	return v1.x*v2.x + v1.y*v2.y
}

func Angle(v1 Vector2, v2 Vector2) float64 {
	return math.Acos(Dot(v1, v2) / (v1.Magnitude() * v2.Magnitude()))
}
