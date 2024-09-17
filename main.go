package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	en_scenario = []string{
		"Here is a my house.",
		"I don't know what to do...",
		"next day...",
	}
	ja_scenario = []string{
		"ここは私の家です。",
		"何をしようかな...",
		"翌日...",
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
	if g.currentMessageIndex < len(ja_scenario) {
		//if g.currentMessageIndex < len(en_scenario) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if time.Since(g.lastClickTime) > 200*time.Millisecond {
				g.lastClickTime = time.Now()
				// 日本語を扱う場合
				runes := []rune(ja_scenario[g.currentMessageIndex])
				if g.currentCharIndex >= len(runes) {
					g.currentMessageIndex++
					g.displayText = ""
					g.currentCharIndex = 0
					// 英語を扱う場合
					//if g.currentCharIndex >= len(en_scenario[g.currentMessageIndex]) {
					//	g.currentMessageIndex++ // 次の文章へ
					//	g.displayText = ""
					//	g.currentCharIndex = 0
					//}
				}
			}
			// 日本語を扱う場合
			if g.currentMessageIndex < len(ja_scenario) {
				runes := []rune(ja_scenario[g.currentMessageIndex])
				if g.currentCharIndex < len(runes) {
					g.displayText += string(runes[g.currentCharIndex])
					g.currentCharIndex++
				}
			}
			// messageIndexをチェックすればscenario内の要素が最後まで来たのかわかる
			// 英語を扱う場合
			//if g.currentMessageIndex < len(en_scenario) {
			//	if g.currentCharIndex < len(en_scenario[g.currentMessageIndex]) {
			//		g.displayText += string(en_scenario[g.currentMessageIndex][g.currentCharIndex])
			//		g.currentCharIndex++ // 次の文字へ
			//	}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	ebitenutil.DebugPrintAt(screen, g.displayText, 50, 100)

	if g.currentMessageIndex >= len(ja_scenario) {
		ebitenutil.DebugPrintAt(screen, "All messages had been displayed", 50, 200)
	}
	op := &text.DrawOptions{}
	op.ColorScale.Scale(1, 1, 1, 1)
	op.LineSpacing = 48 * 1.5
	text.Draw(screen, g.displayText, fontFace, op)
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
