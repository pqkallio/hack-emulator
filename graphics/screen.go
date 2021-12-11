package graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pqkallio/hack-emulator/hack/components/sequential/word"
)

const (
	scale = 2
	rows  = 256
	cols  = 512
)

var (
	x      = uint16(0)
	scrVal = uint16(0b11111111_11111111)
)

type Screen struct {
	window  *glfw.Window
	program uint32
	pixels  [][]*pixel
	mem     *word.ScreenMem
}

func NewScreen(mem *word.ScreenMem) *Screen {
	window := initGlfw()
	program := initOpenGL()
	pixels := makePixels()

	return &Screen{
		window:  window,
		program: program,
		pixels:  pixels,
		mem:     mem,
	}
}

func (s *Screen) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	if x == 0 {
		s.mem.Update(0, 8191, true)
	} else {
		s.mem.Update(0, x-1, true)
	}

	s.mem.Update(scrVal, x, true)
	s.mem.Tick()

	screenMem := s.mem.GetMem()

	for i, val := range screenMem {
		mask := uint16(0b10000000_00000000)
		for j := 0; j < 16; j++ {
			if val&mask != 0 {
				cellPtr := i*16 + j
				y := cellPtr / cols
				x := cellPtr % cols
				s.pixels[y][x].draw()
			}
			mask >>= 1
		}
	}

	x++
	if x == 8192 {
		x = 0
	}

	glfw.PollEvents()
	s.window.SwapBuffers()
}

func (s *Screen) ShouldClose() bool {
	return s.window.ShouldClose()
}

func (s *Screen) Terminate() {
	glfw.Terminate()
}

func makePixels() [][]*pixel {
	cells := make([][]*pixel, rows)
	for y := 0; y < rows; y++ {
		row := make([]*pixel, cols)
		for x := 0; x < cols; x++ {
			c := newPixel(x, y)
			row[x] = c
		}
		cells[y] = row
	}

	return cells
}
