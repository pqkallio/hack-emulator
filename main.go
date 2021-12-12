package main

import (
	"runtime"
	"time"

	combWord "github.com/pqkallio/hack-emulator/hack/components/combinational/word"
	seqWord "github.com/pqkallio/hack-emulator/hack/components/sequential/word"
	graphics "github.com/pqkallio/hack-emulator/io"
)

const (
	fps = 30
)

var (
	screenMem = seqWord.NewScreenMem()
	kbdMem    = combWord.NewKeyboardMem()
)

func main() {
	runtime.LockOSThread()

	io := graphics.NewScreenAndKeyboard(256, 512, 2, screenMem, kbdMem)
	defer io.Terminate()

	for !io.ShouldClose() {
		t := time.Now()
		io.Process()
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}
