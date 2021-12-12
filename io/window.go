package graphics

type window interface {
	Terminate()
	PollEvents()
	ShouldClose() bool
	SwapBuffers()
}
