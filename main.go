package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/lewism77/burnedground/projectile"
	"github.com/lewism77/burnedground/tank"
	"github.com/lewism77/burnedground/vector"
)

const (
	screenWidth  = 600
	screenHeight = 600
	screenScale  = 1
	windowTitle  = "Burned Ground"
)

var (
	currentTank      *tank.Tank
	brushImage       *ebiten.Image
	tanks            []*tank.Tank
	projectiles      []*projectile.Projectile
	currentTankIndex int
)

func init() {
	brushImage, _ = ebiten.NewImage(4, 4, ebiten.FilterDefault)
	brushImage.Fill(color.White)
	projectiles = make([]*projectile.Projectile, 0)
	tanks = make([]*tank.Tank, 0)

	tank1 := tank.Tank{
		LocX: 10,
		LocY: 10,
	}

	tank2 := tank.Tank{
		LocX: 400,
		LocY: 100,
	}

	tanks = append(tanks, &tank1)
	tanks = append(tanks, &tank2)

	currentTank = &tank1
	currentTankIndex = 0
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyKP4) {
		currentTank.LocX--
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP6) {
		currentTank.LocX++
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP8) {
		currentTank.LocY--
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP2) {
		currentTank.LocY++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		currentTank.Power++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		currentTank.Power--
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		currentTank.Angle--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		currentTank.Angle++
	}
	if repeatingKeyPressed(ebiten.KeyT) {
		if currentTankIndex+1 >= len(tanks) {
			currentTankIndex = 0
		} else {
			currentTankIndex++
		}

		currentTank = tanks[currentTankIndex]
	}
	if repeatingKeyPressed(ebiten.KeySpace) {
		newProjectile := projectile.New(currentTank.Power, currentTank.Angle)

		newProjectile.Owner = currentTank.Player
		newProjectile.Position = vector.Vector{
			X: float64(currentTank.LocX),
			Y: float64(currentTank.LocY),
		}
		projectiles = append(projectiles, &newProjectile)
	}

	for ix := range projectiles {
		draw(screen, int(projectiles[ix].Position.X), int(projectiles[ix].Position.Y))
		projectiles[ix].Logic()

		velText := "Acc X: " + fmt.Sprintf("%.2f", projectiles[ix].Acceleration.X) + ", Acc Y: " + fmt.Sprintf("%.2f", projectiles[ix].Acceleration.Y)
		accText := "Vel X: " + fmt.Sprintf("%.2f", projectiles[ix].Velocity.X) + ", Vel Y: " + fmt.Sprintf("%.2f", projectiles[ix].Velocity.Y)
		posText := "Pos X: " + fmt.Sprintf("%.2f", projectiles[ix].Position.X) + ", Pos Y: " + fmt.Sprintf("%.2f", projectiles[ix].Position.Y)

		ebitenutil.DebugPrintAt(screen, velText, int(projectiles[ix].Position.X), int(projectiles[ix].Position.Y))
		ebitenutil.DebugPrintAt(screen, accText, int(projectiles[ix].Position.X), int(projectiles[ix].Position.Y)+10)
		ebitenutil.DebugPrintAt(screen, posText, int(projectiles[ix].Position.X), int(projectiles[ix].Position.Y)+40)
	}

	ix := 0
	for _, proj := range projectiles {
		if -10 <= proj.Position.X && proj.Position.X <= screenWidth+10 &&
			-10 <= proj.Position.Y && proj.Position.Y <= screenHeight+10 {
			projectiles[ix] = proj
			ix++
		}
	}
	projectiles = projectiles[:ix]

	for _, tank := range tanks {
		draw(screen, tank.LocX, tank.LocY)

		angleText := "Angle: " + fmt.Sprintf("%.2f", tank.Angle)
		powerText := "Power: " + fmt.Sprintf("%.2f", tank.Power)

		ebitenutil.DebugPrintAt(screen, angleText, tank.LocX, tank.LocY)
		ebitenutil.DebugPrintAt(screen, powerText, tank.LocX, tank.LocY+10)

		ebitenutil.DebugPrintAt(screen, "Projectile Count: "+strconv.Itoa(len(projectiles)), tank.LocX, tank.LocY+20)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}

func draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(brushImage, op)
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, screenScale, windowTitle); err != nil {
		log.Fatal(err)
	}
}
