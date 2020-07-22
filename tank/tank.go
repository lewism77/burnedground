package tank

import (
	"github.com/hajimehoshi/ebiten"
)

// Tank - Holds data for a tank
type Tank struct {
	Player int
	Angle  float64
	Power  float64
	LocX   int
	LocY   int
	Sprite *ebiten.Image
}
