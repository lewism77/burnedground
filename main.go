package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/lewism77/burnedground/keymap"
	"github.com/lewism77/burnedground/projectile"
	"github.com/lewism77/burnedground/tank"
)

const (
	screenWidth  = 800
	screenHeight = 600
	screenScale  = 1
	windowTitle  = "Burned Ground"
	tankWidth    = 30
	tankHeight   = 10
	turretWidth  = 4
	turretHeight = 20
)

var (
	currentTank      *tank.Tank
	projectileImage  *ebiten.Image
	tankImage        *ebiten.Image
	turretImage      *ebiten.Image
	tanks            []*tank.Tank
	projectiles      []*projectile.Projectile
	currentTankIndex int
)

func init() {
	projectileImage, _ = ebiten.NewImage(4, 4, ebiten.FilterDefault)
	projectileImage.Fill(color.White)

	tankImage, _ = ebiten.NewImage(tankWidth, tankHeight, ebiten.FilterDefault)
	tankImage.Fill(color.White)

	turretImage, _ = ebiten.NewImage(turretWidth, turretHeight, ebiten.FilterDefault)
	turretImage.Fill(color.White)

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

func update(screen *ebiten.Image) error {
	keymap.Logic(&currentTank, &currentTankIndex, &tanks, &projectiles)

	debugFps := fmt.Sprintf(`TPS: %0.2f
	FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, debugFps)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for ix := range projectiles {
		draw(screen, projectileImage, int(projectiles[ix].Position.X), int(projectiles[ix].Position.Y))
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
		draw(screen, tankImage, tank.LocX, tank.LocY)
		drawTurret(screen, turretImage, tank.LocX, tank.LocY, tank.Angle)

		angleText := "Angle: " + fmt.Sprintf("%.2f", tank.Angle)
		powerText := "Power: " + fmt.Sprintf("%.2f", tank.Power)

		ebitenutil.DebugPrintAt(screen, angleText, tank.LocX, tank.LocY)
		ebitenutil.DebugPrintAt(screen, powerText, tank.LocX, tank.LocY+10)

		ebitenutil.DebugPrintAt(screen, "Projectile Count: "+strconv.Itoa(len(projectiles)), tank.LocX, tank.LocY+20)
	}

	return nil
}

func draw(screen, brushImage *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(brushImage, op)
}

func drawTurret(screen, brushImage *ebiten.Image, x, y int, angle float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate((angle - 90) * (math.Pi / 180))
	op.GeoM.Translate(float64(x)+(tankWidth/2), float64(y))
	screen.DrawImage(brushImage, op)
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, screenScale, windowTitle); err != nil {
		log.Fatal(err)
	}
}
