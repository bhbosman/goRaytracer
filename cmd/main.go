package main

import (
	"fmt"
	"runtime"
	"time"

	internal "github.com/bhbosman/goRaytracer/internal"
	processing03 "github.com/bhbosman/goRaytracer/internal/processing/processing03"
	"github.com/go-gl/mathgl/mgl64"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2) // Use all the machine's cores
	scene := internal.CreateDefaultScene()
	scene00002(scene)
	camera := internal.DefaultCamera()
	picture := internal.CreatePicture(camera, 800, 600)
	ss := processing03.CreateImageProcessor(picture, scene)
	t0 := time.Now()
	ss.Process(8, 12)
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
	ss.Save("dddd.jpg")
}

func scene00002(scene *internal.Scene) {
	scene.AddObject(&internal.BlackAndWhiteChessBoard{
		Plane: internal.Plane{
			Origin: mgl64.Vec3{0, -25, 0},
			Normal: mgl64.Vec3{0, -1, 0}},
		BlockSize: mgl64.Vec2{50, 50},
		Color01:   nil,
		Color02:   nil})

	scene.AddObject(&internal.StandardSphere{
		Sphere: internal.Sphere{
			Origin: mgl64.Vec3{0, 0, 0},
			Radius: 25},
		StandardMaterial: internal.MaterialParams{
			RefractionValue: 1.5,
			Pd:              0.3,
			Ps:              0.0,
			Color:           mgl64.Vec3{0, 0, 1}}})

	scene.AddObject(&internal.StandardSphere{
		Sphere: internal.Sphere{
			Origin: mgl64.Vec3{-50, 0, 0},
			Radius: 25},
		StandardMaterial: internal.MaterialParams{
			RefractionValue: 1.5,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&internal.StandardSphere{
		Sphere: internal.Sphere{
			Origin: mgl64.Vec3{50, 0, 0},
			Radius: 25},
		StandardMaterial: internal.MaterialParams{
			RefractionValue: 1.5,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&internal.StandardSphere{
		Sphere: internal.Sphere{
			Origin: mgl64.Vec3{25, 50, 0},
			Radius: 25},
		StandardMaterial: internal.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&internal.StandardSphere{
		Sphere: internal.Sphere{
			Origin: mgl64.Vec3{-25, 50, 0},
			Radius: 25},
		StandardMaterial: internal.MaterialParams{
			RefractionValue: 1.0,
			Pd:              0.3,
			Ps:              0.0,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&internal.StandardSphere{
		Sphere: internal.Sphere{
			Origin: mgl64.Vec3{25, 50, -50},
			Radius: 25},
		StandardMaterial: internal.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&internal.StandardSphere{
		Sphere: internal.Sphere{
			Origin: mgl64.Vec3{-25, 50, -50},
			Radius: 25},
		StandardMaterial: internal.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{0, 1, 0}}})
}
