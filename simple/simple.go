package main

import (
	"github.com/rostislavjadavan/gofb"
	"math/rand"
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)
	bg := gofb.NewSurface(1200, 900)

	for y := 200; y < 700; y++ {
		for x := 200; x < 1000; x++ {
			gray := uint8(rand.Intn(255))
			bg.SetPixel(x, y, gofb.NewColor(gray, gray, gray, 255))
		}
	}

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor(0, 0, 0, 255))

		bg.Draw(0, 0)

		w.FinalizeFrame()
	}

	defer bg.Release()
	defer w.Destroy()
}
