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

	var frameUpdateTimeMs int64 = 0          // how much time elapsed
	var frame = 0                            // current animation frame frame
	var frameSize = gofb.NewPoint2(128, 170) // atlas frame region

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor(120, 220, 230, 255))

		// calculate frame position in atlas
		r := gofb.NewRegion(
			frame*int(frameSize.X), 0,
			int(frameSize.X), int(frameSize.Y))

		// draw atlas region
		atlas.DrawRegion(400, 250, r)

		frameUpdateTimeMs += w.GetDeltaTimeMs()
		// if 100ms elapsed switch to next frame from atlas (and reset timer)
		if frameUpdateTimeMs > 100 {
			frame++
			frameUpdateTimeMs = 0
			// there are 8 frames in atlas
			if frame > 8 {
				frame = 0
			}
		}

		w.FinalizeFrame()
	}

	defer atlas.Release()
	defer w.Destroy()
}
