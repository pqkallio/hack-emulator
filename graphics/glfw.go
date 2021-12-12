package graphics

import "github.com/go-gl/glfw/v3.3/glfw"

type glfwWindow struct {
	window *glfw.Window
}

func newGlfwWindow(nRows, nCols, scale int) *glfwWindow {
	window := initGlfw(nRows*scale, nCols*scale)

	return &glfwWindow{window}
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
