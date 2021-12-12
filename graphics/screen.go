package graphics

import (
	"github.com/pqkallio/hack-emulator/hack/components/sequential/word"
)

type Screen struct {
	window   window
	graphics graphics
	mem      *word.ScreenMem
}

func NewScreen(nRows, nCols, scale int, mem *word.ScreenMem) *Screen {
	return &Screen{
		window:   newGlfwWindow(nRows, nCols, scale),
		graphics: newOpenGL(nRows, nCols, scale, mem),
		mem:      mem,
	}
}

func (s *Screen) Draw() {
	s.mem.Tick()

	s.graphics.Draw()

	s.window.PollEvents()
	s.window.SwapBuffers()
}

func (s *Screen) ShouldClose() bool {
	return s.window.ShouldClose()
}

func (s *Screen) Terminate() {
	s.window.Terminate()
}
