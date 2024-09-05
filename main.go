package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	choiceSelected bool
	message        string
}

func NewGame() *Game {
	return &Game{
		choiceSelected: false,
		message:        "Please select a choice",
	}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.isChoiceSelected(x, y) {
			g.choiceSelected = true
			g.message = "selected！"
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.message)

	ebitenutil.DebugPrintAt(screen, "1. select A", 100, 200)
	ebitenutil.DebugPrintAt(screen, "2. select B", 100, 230)
}

func (g *Game) isChoiceSelected(x, y int) bool {
	if x >= 100 && x <= 300 && y >= 200 && y <= 220 {
		return true
	}
	return false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple game")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
