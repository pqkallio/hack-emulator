package graphics

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/pqkallio/hack-emulator/hack/components/sequential/word"
)

const (
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

type openGL struct {
	nCols  int
	pixels [][]*pixel
	mem    *word.ScreenMem
}

func newOpenGL(nRows, nCols, scale int, mem *word.ScreenMem) *openGL {
	initOpenGL()
	return &openGL{nCols, makePixels(nRows, nCols, scale), mem}
}

func (o *openGL) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	screenMem := o.mem.GetMem()

	for i, val := range screenMem {
		mask := uint16(0b10000000_00000000)

		for j := 0; j < 16; j++ {
			if val&mask != 0 {
				cellPtr := i*16 + j
				y := cellPtr / o.nCols
				x := cellPtr % o.nCols
				o.pixels[y][x].draw()
			}

			mask >>= 1
		}
	}
}

func initOpenGL() {
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
	gl.UseProgram(prog)
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

func makePixels(nRows, nCols, scale int) [][]*pixel {
	h := (1 / float32(nRows)) * float32(scale)
	w := (1 / float32(nCols)) * float32(scale)

	pixels := make([][]*pixel, nRows)

	for y := 0; y < nRows; y++ {
		row := make([]*pixel, nCols)

		for x := 0; x < nCols; x++ {
			c := newPixel(x, y, h, w)
			row[x] = c
		}

		pixels[y] = row
	}

	return pixels
}
