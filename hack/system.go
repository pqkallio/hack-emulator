package hack

import (
	"runtime"
	"time"

	combWord "github.com/pqkallio/hack-emulator/hack/components/combinational/word"
	seqWord "github.com/pqkallio/hack-emulator/hack/components/sequential/word"
	"github.com/pqkallio/hack-emulator/io"
)

const (
	fps = 30
)

var ()

func Run(instr []uint16) {
	runtime.LockOSThread()

	screenMem := seqWord.NewScreenMem()
	kbdMem := combWord.NewKeyboardMem()
	ram := seqWord.NewRam16kFlat()
	rom := combWord.NewROM32KFlat()

	ram.Write(0, 16)
	rom.Flash(instr)

	peripheral := io.NewScreenAndKeyboard(256, 512, 2, screenMem, kbdMem)
	defer peripheral.Terminate()

	reset := false
	computer := NewComputer(ram, screenMem, kbdMem, rom, &reset)
	go computer.Run()

	for !peripheral.ShouldClose() {
		t := time.Now()
		peripheral.Process()
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}
