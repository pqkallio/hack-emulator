package io

import (
	combWord "github.com/pqkallio/hack-emulator/hack/components/combinational/word"
	seqWord "github.com/pqkallio/hack-emulator/hack/components/sequential/word"
)

type ScreenAndKeyboard struct {
	window   window
	graphics graphics
	scr      *seqWord.ScreenMem
	kbd      *combWord.KeyboardMem
}

func NewScreenAndKeyboard(nRows, nCols, scale int, scr *seqWord.ScreenMem, kbd *combWord.KeyboardMem) *ScreenAndKeyboard {
	return &ScreenAndKeyboard{
		window:   newGlfwWindow(nRows, nCols, scale, kbd),
		graphics: newOpenGL(nRows, nCols, scale, scr),
		scr:      scr,
		kbd:      kbd,
	}
}

func (s *ScreenAndKeyboard) Process() {
	s.scr.Tick()

	s.graphics.Draw()

	s.window.SwapBuffers()
	s.window.PollEvents()
}

func (s *ScreenAndKeyboard) ShouldClose() bool {
	return s.window.ShouldClose()
}

func (s *ScreenAndKeyboard) Terminate() {
	s.window.Terminate()
}
