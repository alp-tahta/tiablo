package main

import (
	"github.com/alp-tahta/tiablo.git/internal/logger"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	l := logger.Init()
	g := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		l.Error(err.Error())
	}
}
