package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"burnedground/tank"
)

const (
	screenWidth  = 600
	screenHeight = 600
	screenScale  = 1
	windowTitle  = "Burned Ground"
)

var (
	tank1 tank.Tank
)

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyKP4) {
		tank1.LocX--
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP6) {
		tank1.LocX++
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP8) {
		tank1.LocY--
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP2) {
		tank1.LocY++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		tank1.Power++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		tank1.Power--
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		tank1.Angle--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		tank1.Angle++
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	angleText := "Angle: " + fmt.Sprintf("%.2f", tank1.Angle)
	powerText := "Power: " + fmt.Sprintf("%.2f", tank1.Power)

	ebitenutil.DebugPrintAt(screen, angleText, tank1.LocX, tank1.LocY)
	ebitenutil.DebugPrintAt(screen, powerText, tank1.LocX, tank1.LocY+10)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, screenScale, windowTitle); err != nil {
		log.Fatal(err)
	}
}
