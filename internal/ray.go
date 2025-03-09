package internal

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Ray struct {
	Origin    mgl64.Vec3
	Direction mgl64.Vec3
}

func (r Ray) PositionAt(time float64) (result mgl64.Vec3) {
	return r.Origin.Add(r.Direction.Mul(time))
}
