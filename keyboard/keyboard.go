package main

import (
	"github.com/rostislavjadavan/gofb"
)

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

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor(84, 197, 211, 255))

		star.Draw(starPosX, starPosY)

		speed := int(w.GetDeltaTimeMs())
		if w.IsKey(gofb.KeyShift) {
			speed = int(w.GetDeltaTimeMs()) / 5
		}

		if w.IsKey(gofb.KeyUp) {
			if starPosY > 0 {
				starPosY -= speed
			}
		}
		if w.IsKey(gofb.KeyDown) {
			if starPosY < w.Height - star.Height {
				starPosY += speed
			}
		}
		if w.IsKey(gofb.KeyLeft) {
			if starPosX > 0 {
				starPosX -= speed
			}
		}
		if w.IsKey(gofb.KeyRight) {
			if starPosX < w.Width - star.Width {
				starPosX += speed
			}
		}
		if w.IsKey(gofb.KeyEscape) {
			w.Stop()
		}

		text.Draw("Use cursor keys to move", 100, 720, gofb.NewColor(0, 0, 0, 222))
		text.Draw("Hold shift to move slower", 100, 760, gofb.NewColor(0, 0, 0, 222))
		text.Draw("Press escape to exit", 100, 800, gofb.NewColor(0, 0, 0, 222))

		w.FinalizeFrame()
	}

	defer star.Release()
	defer w.Destroy()
}
