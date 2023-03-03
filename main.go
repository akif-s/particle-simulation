package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

const (
	WINDOW_WIDTH  = 1300
	WINDOW_HEIGHT = 700
)

func main() {
	rand.Seed(time.Now().UnixNano()) // set the seed
	go func() {
		w := app.NewWindow(
			app.Title("Collision"),
			app.Size(unit.Dp(WINDOW_WIDTH), unit.Dp(WINDOW_HEIGHT)),
		)

		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var particles []*Particle

func loop(w *app.Window) error {
	var ops op.Ops

	particles = createParticles(500, particles)
	fmt.Printf("Total %d particles corrected", count)

	var keytag struct{}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			fillBackground(color.NRGBA{R: 220, G: 220, B: 220, A: 0xff}, gtx)

			//Keyboard inputs
			key.InputOp{
				Tag:  &keytag,
				Keys: key.NameEscape,
			}.Add(gtx.Ops)

			for _, ev := range gtx.Queue.Events(&keytag) {
				if e, ok := ev.(key.Event); ok {
					if e.State == key.Press {
						if e.Name == key.NameEscape {
							w.Perform(system.ActionClose)
						}
					}
				}
			}

			//Draw and update the particles
			for _, p := range particles {
				a := clip.Ellipse{Min: image.Pt(int(p.x), int(p.y)), Max: image.Pt(int(p.x)+p.radius*2, int(p.y)+p.radius*2)}.Push(gtx.Ops)
				paint.ColorOp{Color: p.color}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				a.Pop()
				p.update(particles)
			}

			op.InvalidateOp{}.Add(gtx.Ops)
			e.Frame(gtx.Ops)

		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}

var count int

func createParticles(n int, particles []*Particle) []*Particle {
	for i := 0; i < n; i++ {
		var x, y float64
		radius := 10

		x = float64(rand.Intn(WINDOW_WIDTH*2 - radius*2))
		y = float64(rand.Intn(WINDOW_HEIGHT*2 - radius*2))

		//Randomizing velocities
		vy := 0. //float64(rand.Intn(20) - 10)
		vx := 0. //float64(rand.Intn(20) - 10)
		p := newParticle(x, y, vx, vy, 1, radius, color.NRGBA{R: 95, G: 190, B: 190, A: 0xff}, i)

		// Making sure two particle do not spawn at same space.
		for {
			if isColliding, _ := p.checkCollision(particles); !isColliding {
				break
			}
			p.x = float64(rand.Intn(WINDOW_WIDTH*2 - radius*2))
			p.y = float64(rand.Intn(WINDOW_HEIGHT*2 - radius*2))
			fmt.Println("Corrected: ", p.id)
			count++
		}
		particles = append(particles, p)
	}
	return particles
}

func fillBackground(color color.NRGBA, gtx layout.Context) {
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
}
