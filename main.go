package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"log"
	"os"
	"time"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	scenario = []string{
		"Here is a my house.",
		"I don't know what to do...",
		"next day...",
	}
	fontFace *text.GoTextFace
)

type Game struct {
	currentMessageIndex int
	displayText         string
	currentCharIndex    int
	lastClickTime       time.Time
}

func NewGame() *Game {
	return &Game{
		currentMessageIndex: 0,
		displayText:         "",
		currentCharIndex:    0,
		lastClickTime:       time.Now(),
	}
}

func (g *Game) Update() error {
	if g.currentMessageIndex <= len(scenario) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if time.Since(g.lastClickTime) > 200*time.Millisecond {
				g.lastClickTime = time.Now()
				if g.currentCharIndex >= len(scenario[g.currentMessageIndex]) {
					g.currentMessageIndex++
					g.displayText = ""
					g.currentCharIndex = 0
				}
			}
		}
		if g.currentCharIndex < len(scenario[g.currentMessageIndex]) {
			g.displayText += string(scenario[g.currentMessageIndex][g.currentCharIndex])
			g.currentCharIndex++
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.Black)

	//ebitenutil.DebugPrintAt(screen, g.displayText, 50, 100)
	//
	//if g.currentMessageIndex >= len(scenario) {
	//	ebitenutil.DebugPrintAt(screen, "All messages had been displayed", 50, 200)
	//}
	screen.Fill(color.White)
	// text v2
	op := &text.DrawOptions{}
	op.ColorScale.Scale(0, 0, 0, 1)
	op.LineSpacing = 48 * 1.5
	text.Draw(screen, "こんにちは、世界!", fontFace, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func init() {
	f, err := os.Open("KiwiMaru-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	src, err := text.NewGoTextFaceSource(f)
	if err != nil {
		log.Fatal(err)
	}

	fontFace = &text.GoTextFace{Source: src, Size: 48}

}
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Nobel Game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
