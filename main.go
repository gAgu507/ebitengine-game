package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	x, y  float64
	actor *ebiten.Image
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.x -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.x += 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.y -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.y += 3
	}
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	screen.DrawImage(g.actor, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	square := ebiten.NewImage(32, 32)
	square.Fill(color.White)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("SQUARE game")
	if err := ebiten.RunGame(&Game{300, 220, square}); err != nil {
		log.Fatal(err)
	}
}
