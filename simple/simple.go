package main

import (
	"github.com/rostislavjadavan/gofb"
	"math/rand"
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)
	bg := gofb.NewSurface(600, 600)

	for y := 0; y < bg.Height; y++ {
		for x := 0; x < bg.Width; x++ {
			gray := uint8(rand.Intn(255))
			bg.SetPixel(x, y, gofb.NewColor(gray, gray, gray, 255))
		}
	}

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor(0, 0, 0, 255))

		bg.Draw(300, 150)
		bg.Rotation += float32(w.GetDeltaTimeMs() / 10)

		w.FinalizeFrame()
	}

	defer bg.Release()
	defer w.Destroy()
}
