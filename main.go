package main

import (
	"runtime"
	"time"

	"github.com/pqkallio/hack-emulator/graphics"
	"github.com/pqkallio/hack-emulator/hack/components/sequential/word"
)

const (
	fps = 30
)

var (
	screenMem = word.NewScreenMem()
)

func main() {
	runtime.LockOSThread()

	screen := graphics.NewScreen(256, 512, 2, screenMem)
	defer screen.Terminate()

	for !screen.ShouldClose() {
		t := time.Now()
		screen.Draw()
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}
