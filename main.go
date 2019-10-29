package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/lewism77/burnedground/keymap"
	"github.com/lewism77/burnedground/projectile"
	"github.com/lewism77/burnedground/tank"
	"github.com/lewism77/burnedground/terrain"
)

const (
	screenWidth  = 800
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
	terrainColumns   []*ebiten.Image
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

	terrain.Init(screenWidth, screenHeight)
	terrain.Generate()

	var groundHeight int
	var groundColumn *ebiten.Image
	for _, pixCol := range terrain.Grid {
		for _, onoff := range pixCol {
			if onoff == 0 {
				continue
			}
			groundHeight++
		}
		groundColumn, _ = ebiten.NewImage(1, groundHeight, ebiten.FilterDefault)
		groundColumn.Fill(color.White)
		terrainColumns = append(terrainColumns, groundColumn)
		groundHeight = 0
	}
}

func update(screen *ebiten.Image) error {
	debugFps := fmt.Sprintf(`TPS: %0.2f
	FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, debugFps)

	keymap.Logic(&currentTank, &currentTankIndex, &tanks, &projectiles)

	for ix := range projectiles {
		draw(screen, brushImage, int(projectiles[ix].Position.X), int(projectiles[ix].Position.Y))
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
		draw(screen, brushImage, tank.LocX, tank.LocY)

		angleText := "Angle: " + fmt.Sprintf("%.2f", tank.Angle)
		powerText := "Power: " + fmt.Sprintf("%.2f", tank.Power)

		ebitenutil.DebugPrintAt(screen, angleText, tank.LocX, tank.LocY)
		ebitenutil.DebugPrintAt(screen, powerText, tank.LocX, tank.LocY+10)

		ebitenutil.DebugPrintAt(screen, "Projectile Count: "+strconv.Itoa(len(projectiles)), tank.LocX, tank.LocY+20)
	}

	for pixX, groundColumn := range terrainColumns {
		_, pixY := groundColumn.Size()
		draw(screen, groundColumn, pixX, screenHeight-pixY)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}

func draw(screen, image *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(image, op)
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, screenScale, windowTitle); err != nil {
		log.Fatal(err)
	}
}
