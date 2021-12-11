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
	scr = word.NewScreenMem()
)

func main() {
	runtime.LockOSThread()

	window, prog, terminate := graphics.Init()
	defer terminate()

	for !window.ShouldClose() {
		t := time.Now()
		graphics.Draw(scr, window, prog)
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}
