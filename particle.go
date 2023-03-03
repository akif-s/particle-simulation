package main

import (
	"image/color"
	"math"
)

type Particle struct {
	x, y   float64
	vX, vY float64
	mass   float64
	radius int
	color  color.NRGBA
	id     int
}

func newParticle(x, y, vX, vY, mass float64, radius int, color color.NRGBA, id int) *Particle {
	return &Particle{
		x:      x,
		y:      y,
		vX:     vX,
		vY:     vY,
		mass:   mass,
		radius: radius,
		color:  color,
		id:     id,
	}
}

func (p *Particle) update(particles []*Particle) {
	//Moving the particle
	p.x += p.vX
	p.y += p.vY

	go p.detectEdges()
	p.checkCollision(particles)
}

func (P *Particle) checkCollision(particles []*Particle) {

	for _, p := range particles {
		if p.id == P.id {
			continue
		}

		dx := p.x - P.x
		dy := p.y - P.y
		distance := math.Sqrt(dx*dx + dy*dy)

		if distance <= float64(p.radius)+float64(P.radius) {
			// X-Axis Calculations
			tmPX := P.vX
			tmpX := p.vX

			P.vX = tmPX*(P.mass-p.mass)/(P.mass+p.mass) + tmpX*(2*p.mass)/(P.mass+p.mass)
			p.vX = tmPX*(2*P.mass)/(p.mass+P.mass) + tmpX*(p.mass-P.mass)/(p.mass+P.mass)

			// Y-Axis Calculations
			tmPY := P.vY
			tmpY := p.vY
			P.vY = tmPY*(P.mass-p.mass)/(P.mass+p.mass) + tmpY*(2*p.mass)/(P.mass+p.mass)
			p.vY = tmPY*(2*P.mass)/(p.mass+P.mass) + tmpY*(p.mass-P.mass)/(p.mass+P.mass)
		}
	}

}

func (p *Particle) detectEdges() {

	// Left and right edges
	if p.x <= 0 || p.x+2*float64(p.radius) >= WINDOW_WIDTH*2 {
		p.vX = -p.vX
	}
	// Up and down edges
	if p.y <= 0 || p.y+2*float64(p.radius) >= WINDOW_HEIGHT*2 {
		p.vY = -p.vY
	}
}
