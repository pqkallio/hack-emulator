package io

type window interface {
	Terminate()
	PollEvents()
	ShouldClose() bool
	SwapBuffers()
}
