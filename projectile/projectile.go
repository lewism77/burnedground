package projectile

import (
	"math"

	"github.com/lewism77/burnedground/vector"
)

// Projectile stores info on a projectile
type Projectile struct {
	Owner        int
	Acceleration vector.Vector
	Velocity     vector.Vector
	Position     vector.Vector
}

func resolution(power float64, directionInRad float64) vector.Vector {
	return vector.Vector{
		X: power * math.Cos(directionInRad),
		Y: power * math.Sin(directionInRad),
	}
}

func degreesToRadians(deg float64) float64 { return math.Pi / 180 * deg }

// New intialises a new projectiles values
func New(power float64, directionInDegrees float64) Projectile {
	gravityPower := 0.2 // in pixels
	p := Projectile{
		Velocity:     resolution(power, degreesToRadians(directionInDegrees)),
		Acceleration: resolution(gravityPower, degreesToRadians(90.0)), // 270 gravity down
	}
	return p
}

// Logic progresses the projectile along it's path
func (p *Projectile) Logic() {
	p.Velocity = vector.Add(p.Velocity, p.Acceleration)
	p.Position = vector.Add(p.Position, p.Velocity)
}
