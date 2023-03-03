package main

import (
	"image/color"
	"math"
)

type Particle struct {
	pos    Vector2
	v      Vector2
	mass   float64
	radius int
	color  color.NRGBA
	id     int
}

func newParticle(x, y, vX, vY, mass float64, radius int, color color.NRGBA, id int) *Particle {
	return &Particle{
		pos:    NewVector2(x, y),
		v:      NewVector2(vX, vY),
		mass:   mass,
		radius: radius,
		color:  color,
		id:     id,
	}
}

func (p *Particle) update(particles []*Particle) {
	//Moving the particle
	p.pos.x += p.v.x
	p.pos.y += p.v.y

	go p.detectEdges()
	isCollided, p2 := p.checkCollision(particles)
	if isCollided {
		p.CalculateNewVelocity(p2)
	}
}

func (P *Particle) checkCollision(particles []*Particle) (bool, *Particle) {

	for _, p := range particles {
		if p.id == P.id {
			continue
		}

		dx := (p.pos.x + float64(p.radius)) - (P.pos.x + float64(P.radius))
		dy := (p.pos.y + float64(p.radius)) - (P.pos.y + float64(P.radius))
		distance := math.Sqrt(dx*dx + dy*dy)

		if distance <= float64(p.radius)+float64(P.radius) {
			return true, p
		}
	}
	return false, nil
}

// Calculate the velocities after collision
func (p1 *Particle) CalculateNewVelocity(p2 *Particle) {

	tmp1 := p1.v
	tmp2 := p2.v

	p1.v = tmp1.Substract(p1.pos.Substract(p2.pos).Product((2 * p2.mass / (p1.mass + p2.mass) * (Dot(tmp1.Substract(tmp2), p1.pos.Substract(p2.pos)) / (p1.pos.Substract(p2.pos).Magnitude() * p1.pos.Substract(p2.pos).Magnitude())))))
	p2.v = tmp2.Substract(p2.pos.Substract(p1.pos).Product((2 * p1.mass / (p1.mass + p2.mass) * (Dot(tmp2.Substract(tmp1), p2.pos.Substract(p1.pos)) / (p2.pos.Substract(p1.pos).Magnitude() * p2.pos.Substract(p1.pos).Magnitude())))))

}

func (p *Particle) detectEdges() {

	// Left and right edges
	if p.pos.x <= 0 || p.pos.x+2*float64(p.radius) >= WINDOW_WIDTH*2 {
		p.v.x = -p.v.x
	}
	// Up and down edges
	if p.pos.y <= 0 || p.pos.y+2*float64(p.radius) >= WINDOW_HEIGHT*2 {
		p.v.y = -p.v.y
	}
}
