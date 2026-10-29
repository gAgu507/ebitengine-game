package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	x, y  float64
	actor *ebiten.Image
	coins []Coin
	score int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if g.x > 0 {
			g.x -= 3
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if g.x < screenWidth-float64(g.actor.Bounds().Dx()) {
			g.x += 3
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.y > 0 {
			g.y -= 3
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if g.y < screenHeight-float64(g.actor.Bounds().Dy()) {
			g.y += 3
		}
	}

	for i, coin := range slices.Backward(g.coins) {
		if g.touches(coin) {
			// on collision
			g.coins = append(g.coins[:i], g.coins[i+1:]...)
			g.score++
			g.generateCoin(screenWidth, screenHeight)
		}
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	screen.DrawImage(g.actor, op)

	for _, coin := range g.coins {
		opC := &ebiten.DrawImageOptions{}
		opC.GeoM.Translate(coin.x, coin.y)
		screen.DrawImage(coin.body, opC)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %v", g.score))

}

func (g *Game) generateCoin(MaxX int, MaxY int) {
	coinImage := ebiten.NewImage(10, 10)
	coinImage.Fill(color.RGBA{255, 255, 0, 255})
	c := &Coin{float64(rand.Intn(MaxX - 20)), float64(rand.Intn(MaxY - 20)), coinImage}
	g.coins = append(g.coins, *c)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) touches(c Coin) bool {
	return (c.x+float64(c.body.Bounds().Dx())) >= g.x && (c.x <= g.x+float64(g.actor.Bounds().Dx())) &&
		c.y+float64(c.body.Bounds().Dy()) >= g.y && c.y <= g.y+float64(g.actor.Bounds().Dy())
}

type Coin struct {
	x, y float64
	body *ebiten.Image
}

func main() {
	// actor
	square := ebiten.NewImage(32, 32)
	square.Fill(color.White)
	// coins
	coinImage := ebiten.NewImage(10, 10)
	coinImage.Fill(color.RGBA{255, 255, 0, 255})
	coins := []Coin{{211, 300, coinImage},
		{120, 90, coinImage},
		{500, 250, coinImage},
	}

	// window
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("SQUARE game")
	if err := ebiten.RunGame(&Game{300, 220, square, coins, 0}); err != nil {
		log.Fatal(err)
	}
}
