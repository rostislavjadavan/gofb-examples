package main

import (
	"github.com/rostislavjadavan/gofb"
	"strconv"
)

// Linear interpolation between points p, q and time t=[0, 1]
func lerp(t float32, px, py, qx, qy int) (x, y int) {
	return int(float32(px) + t*float32(qx-px)), int(float32(py) + t*float32(qy-py))
}

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)

	text, err := gofb.NewFont("../assets/uni0553-webfont.ttf", 40)
	if err != nil {
		panic(err)
	}

	star, err := gofb.NewSurfaceFromFile("../assets/pixel_star.png")
	if err != nil {
		panic(err)
	}

	starPosX := 500
	starPosY := 400
	sourceStarPosX := starPosX
	sourceStarPosY := starPosY
	targetStarPosX := starPosX
	targetStarPosY := starPosY

	var t float32 = 0.0

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor(84, 197, 211, 255))

		star.Draw(starPosX-star.Width/2, starPosY-star.Height/2)

		speed := float32(w.GetDeltaTimeMs()) / 1000
		cursorPos := w.GetCursorPos()

		if w.IsInput(gofb.KeyEscape) {
			w.Stop()
		}
		if w.IsInput(gofb.MouseButtonLeft) {
			sourceStarPosX = starPosX
			sourceStarPosY = starPosY
			targetStarPosX = int(cursorPos.X)
			targetStarPosY = int(cursorPos.Y)
			t = 0
		}

		if starPosX != targetStarPosX && starPosY != targetStarPosY {
			starPosX, starPosY = lerp(t, sourceStarPosX, sourceStarPosY, targetStarPosX, targetStarPosY)
			t += speed
			if t > 1 {
				starPosX = targetStarPosX
				starPosY = targetStarPosY
			}
			text.Draw("X", targetStarPosX - 15, targetStarPosY - 30, gofb.NewColor(0, 0, 0, 120))
		}

		posAsString := strconv.FormatFloat(float64(cursorPos.X), 'f', 0, 32) + ", " + strconv.FormatFloat(float64(cursorPos.Y), 'f', 0, 32)
		text.Draw("Mouse is at "+posAsString, 100, 760, gofb.NewColor(0, 0, 0, 222))
		text.Draw("Press escape to exit", 100, 800, gofb.NewColor(0, 0, 0, 222))

		w.FinalizeFrame()
	}

	defer star.Release()
	defer w.Destroy()
}
