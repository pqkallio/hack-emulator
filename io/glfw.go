package io

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pqkallio/hack-emulator/hack/components/combinational/word"
)

type glfwWindow struct {
	window *glfw.Window
	kbd    *word.KeyboardMem
}

func newGlfwWindow(nRows, nCols, scale int, kbd *word.KeyboardMem) *glfwWindow {
	window := initGlfw(nRows*scale, nCols*scale)

	glfwWindow := &glfwWindow{window, kbd}
	window.SetKeyCallback(glfwWindow.keyCallback)

	return glfwWindow
}

func (w *glfwWindow) Terminate() {
	glfw.Terminate()
}

func (w *glfwWindow) ShouldClose() bool {
	return w.window.ShouldClose()
}

func (w *glfwWindow) PollEvents() {
	glfw.PollEvents()
}

func (w *glfwWindow) SwapBuffers() {
	w.window.SwapBuffers()
}

func (w *glfwWindow) keyCallback(wnd *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		if mappedKey, ok := keyMap[keyMod{key, mods}]; ok {
			w.kbd.Update(mappedKey)
		}
	} else if action == glfw.Release {
		w.kbd.Update(0)
	}
}

func initGlfw(h, w int) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(w, h, "Hack emulator", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}
