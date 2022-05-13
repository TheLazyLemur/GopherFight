package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type cube struct {
	boxPos  rl.Vector3
	boxSize rl.Vector3
}

var cubes []cube

func spawnCubes() {
	min := 1
	max := 30
	for i := 1; i <= 5; i++ {
		rX := rand.Intn(max-min) + min
		rY := rand.Intn(max-min) + min
		rZ := rand.Intn(max-min) + min
		cubes = append(cubes, cube{
			boxPos:  rl.NewVector3(float32(rX), float32(rY), float32(rZ)),
			boxSize: rl.NewVector3(3.0, 3.0, 3.0),
		})
	}
}

func drawCubes() {
	for _, element := range cubes {
		enemyBoxPos := element.boxPos
		enemyBoxSize := element.boxSize
		rl.DrawCube(enemyBoxPos, enemyBoxSize.X, enemyBoxSize.Y, enemyBoxSize.Z, rl.Gray)
		rl.DrawCubeWires(enemyBoxPos, enemyBoxSize.X, enemyBoxSize.Y, enemyBoxSize.Z, rl.DarkGray)
	}
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - box collisions")

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(0.0, 10.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	rl.SetTargetFPS(60)
	spawnCubes()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		if rl.IsKeyDown(rl.KeyRight) {
			camera.Position.X += 0.2
		} else if rl.IsKeyDown(rl.KeyLeft) {
			camera.Position.X -= 0.2
		} else if rl.IsKeyDown(rl.KeyDown) {
			camera.Position.Z += 0.2
		} else if rl.IsKeyDown(rl.KeyUp) {
			camera.Position.Z -= 0.2
		}
		drawCubes()

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
