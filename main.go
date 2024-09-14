package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
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
	screen.Fill(color.Black)

	ebitenutil.DebugPrintAt(screen, g.displayText, 50, 100)

	if g.currentMessageIndex >= len(scenario) {
		ebitenutil.DebugPrintAt(screen, "All messages had been displayed", 50, 200)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Nobel Game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
