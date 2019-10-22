package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"burnedground/projectile"
	"burnedground/tank"
	"burnedground/vector"
)

const (
	screenWidth  = 600
	screenHeight = 600
	screenScale  = 1
	windowTitle  = "Burned Ground"
)

var (
	tank1       tank.Tank
	brushImage  *ebiten.Image
	projectiles []*projectile.Projectile
)

func init() {
	brushImage, _ = ebiten.NewImage(4, 4, ebiten.FilterDefault)
	brushImage.Fill(color.White)
	projectiles = make([]*projectile.Projectile, 0)
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
	if repeatingKeyPressed(ebiten.KeySpace) {
		newProjectile := projectile.New(tank1.Power, tank1.Angle)

		newProjectile.Owner = tank1.Player
		newProjectile.Position = vector.Vector{
			X: float64(tank1.LocX),
			Y: float64(tank1.LocY),
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
		if proj.Position.X < -10 ||
			proj.Position.X > screenWidth+10 ||
			proj.Position.Y < -10 ||
			proj.Position.Y > screenHeight+10 {
		} else {
			projectiles[ix] = proj
			ix++
		}
	}
	projectiles = projectiles[:ix]

	draw(screen, tank1.LocX, tank1.LocY)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	angleText := "Angle: " + fmt.Sprintf("%.2f", tank1.Angle)
	powerText := "Power: " + fmt.Sprintf("%.2f", tank1.Power)

	ebitenutil.DebugPrintAt(screen, angleText, tank1.LocX, tank1.LocY)
	ebitenutil.DebugPrintAt(screen, powerText, tank1.LocX, tank1.LocY+10)

	ebitenutil.DebugPrintAt(screen, "Projectile Count: "+strconv.Itoa(len(projectiles)), tank1.LocX, tank1.LocY+20)

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
