package graphics

import "github.com/go-gl/gl/v4.6-core/gl"

type pixel struct {
	drawable uint32

	x int
	y int
}

func (c *pixel) draw() {
	gl.BindVertexArray(c.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}

func newPixel(x, y int) *pixel {
	points := make([]float32, len(square))
	copy(points, square)

	for i, p := range points {
		if i%3 == 0 {
			points[i] = p*pixelWidth - 1 + float32(x)*pixelWidth
		} else if i%3 == 1 {
			points[i] = (p-1)*pixelHeight + 1 - float32(y)*pixelHeight
		}
	}

	return &pixel{
		drawable: makeVao(points),

		x: x,
		y: y,
	}
}

func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
