package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	x, y  float64
	actor *ebiten.Image
	coins *[]Coin
	score int
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

	for i, coin := range *g.coins {
		if (coin.x+float64(coin.body.Bounds().Dx())) >= g.x && (coin.x <= g.x+float64(g.actor.Bounds().Dx())) {
			if (coin.y+float64(coin.body.Bounds().Dy())) >= g.y && (coin.y <= (g.y + float64(g.actor.Bounds().Dy()))) {
				// on collision
				*g.coins = append((*g.coins)[:i], (*g.coins)[i+1:]...)
				g.score++
				g.generateCoin(ScreenWidth, ScreenHeight)
			}
		}
		op_C := &ebiten.DrawImageOptions{}
		op_C.GeoM.Translate(coin.x, coin.y)
		screen.DrawImage(coin.body, op_C)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %v", g.score))

}

func (g *Game) generateCoin(MaxX int, MaxY int) {
	c := &Coin{float64(rand.Intn(MaxX)), float64(rand.Intn(MaxY)), ebiten.NewImage(10, 10)}
	c.body.Fill(color.RGBA{255, 255, 0, 255})
	*g.coins = append(*g.coins, *c)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
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
	coins := []Coin{{211, 300, ebiten.NewImage(10, 10)},
		{120, 90, ebiten.NewImage(10, 10)},
		{500, 250, ebiten.NewImage(10, 10)}}
	for _, coin := range coins {
		coin.body.Fill(color.RGBA{255, 255, 0, 255})
	}

	// window
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("SQUARE game")
	if err := ebiten.RunGame(&Game{300, 220, square, &coins, 0}); err != nil {
		log.Fatal(err)
	}
}
