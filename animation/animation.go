package main

import (
	"github.com/rostislavjadavan/gofb"
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)

	spriteSheet, err := gofb.NewSpriteSheetFromFile("../assets/person_atlas.png")
	if err != nil {
		panic(err)
	}

	spriteSheet.Surface().Scale = 2
	spriteSheet.SetFrameRegion(128, 163)

	var frameUpdateTimeMs int64 = 0 // how much time elapsed
	var frame = 0                   // current animation frame

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor(120, 220, 230, 255))

		spriteSheet.Surface().FlipHorizontal = false
		spriteSheet.DrawFrame(200, 250, frame, 0)

		spriteSheet.Surface().FlipHorizontal = true
		spriteSheet.DrawFrame(600, 250, frame, 1)

		frameUpdateTimeMs += w.GetDeltaTimeMs()
		// when 100ms elapsed switch to next frame from sprite sheet (and reset timer)
		if frameUpdateTimeMs > 100 {
			frame++
			frameUpdateTimeMs = 0
			// there are 8 frames in sprite sheet
			if frame > 8 {
				frame = 0
			}
		}

		w.FinalizeFrame()
	}

	defer spriteSheet.Release()
	defer w.Destroy()
}
