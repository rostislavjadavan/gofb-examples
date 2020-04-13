package main

import (
	"github.com/rostislavjadavan/gofb"
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)

	atlas, err := gofb.NewSurfaceFromFile("../assets/person_atlas.png")
	if err != nil {
		panic(err)
	}
	atlas.Scale = 2

	var frameCounter int64 = 0
	frame := 0
	frameSize := gofb.NewPoint2(128, 170)

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor(120, 220, 230, 255))

		r := gofb.NewRegion(
			frame*int(frameSize.X), 0,
			int(frameSize.X), int(frameSize.Y))
		atlas.DrawRegion(400, 250, r)

		frameCounter += w.GetDeltaTimeMs()
		if frameCounter > 200 {
			frame++
			frameCounter = 0
			if frame > 8 {
				frame = 0
			}
		}

		w.FinalizeFrame()
	}

	defer atlas.Release()
	defer w.Destroy()
}
