package graphics

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pqkallio/hack-emulator/hack/components/sequential/word"
)

const (
	scale        = 2
	rows         = 256
	cols         = 512
	screenHeight = scale * rows
	screenWidth  = scale * cols
	pixelHeight  = (1 / float32(rows)) * scale
	pixelWidth   = (1 / float32(cols)) * scale

	vertexShaderSource = `
    #version 450
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 450
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(0, 1, 0, 1);
    }
` + "\x00"
)

var (
	square = []float32{
		0, 1, 0,
		0, 0, 0,
		1, 0, 0,

		0, 1, 0,
		1, 1, 0,
		1, 0, 0,
	}

	x      = uint16(0)
	scrVal = uint16(0b11111111_11111111)
	pixels = [][]*pixel{}
)

func Init() (window *glfw.Window, program uint32, terminate func()) {
	window = initGlfw()
	program = initOpenGL()
	pixels = makePixels()

	return window, program, func() {
		glfw.Terminate()
	}
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(screenWidth, screenHeight, "Hack emulator", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)

	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
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

func Draw(scr *word.ScreenMem, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	if x == 0 {
		scr.Update(0, 8191, true)
	} else {
		scr.Update(0, x-1, true)
	}

	scr.Update(scrVal, x, true)
	scr.Tick()

	screenMem := scr.GetMem()

	for i, val := range screenMem {
		mask := uint16(0b10000000_00000000)
		for j := 0; j < 16; j++ {
			if val&mask != 0 {
				cellPtr := i*16 + j
				y := cellPtr / cols
				x := cellPtr % cols
				pixels[y][x].draw()
			}
			mask >>= 1
		}
	}

	x++
	if x == 8192 {
		x = 0
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
