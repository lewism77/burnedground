package keymap

import (
	"github.com/lewism77/burnedground/projectile"
	"github.com/lewism77/burnedground/tank"
	"github.com/lewism77/burnedground/vector"

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
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			(*currentTank).LocY--
		} else {
			if (*currentTank).Power < 100 {
				(*currentTank).Power++
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			(*currentTank).LocY++
		} else {
			if ((*currentTank).Power) > 0 {
				(*currentTank).Power--
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			(*currentTank).LocX--
		} else {
			if ((*currentTank).Angle) > 0 {
				(*currentTank).Angle--
			} else {
				(*currentTank).Angle = 359
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			(*currentTank).LocX++
		} else {
			if ((*currentTank).Angle) < 360 {
				(*currentTank).Angle++
			} else {
				(*currentTank).Angle = 0
			}
		}
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
			Y: float64((*currentTank).LocY) - 5.0,
		}
		*projectiles = append(*projectiles, &newProjectile)
	}
}
