package keymap

import (
	"burnedground/projectile"
	"burnedground/tank"
	"burnedground/vector"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

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

// Logic executes any actions based on pressed keys this cycle
func Logic(currentTank **tank.Tank, currentTankIndex *int, tanks *[]*tank.Tank, projectiles *[]*projectile.Projectile) {
	if ebiten.IsKeyPressed(ebiten.KeyKP4) {
		(*currentTank).LocX--
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP6) {
		(*currentTank).LocX++
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP8) {
		(*currentTank).LocY--
	}
	if ebiten.IsKeyPressed(ebiten.KeyKP2) {
		(*currentTank).LocY++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		(*currentTank).Power++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		(*currentTank).Power--
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		(*currentTank).Angle--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		(*currentTank).Angle++
	}
	if repeatingKeyPressed(ebiten.KeyT) {
		if *currentTankIndex+1 >= len(*tanks) {
			*currentTankIndex = 0
		} else {
			*currentTankIndex++
		}

		*currentTank = (*tanks)[*currentTankIndex]
	}
	if repeatingKeyPressed(ebiten.KeySpace) {
		newProjectile := projectile.New((*currentTank).Power, (*currentTank).Angle)

		newProjectile.Owner = (*currentTank).Player
		newProjectile.Position = vector.Vector{
			X: float64((*currentTank).LocX),
			Y: float64((*currentTank).LocY),
		}
		*projectiles = append(*projectiles, &newProjectile)
	}
}
