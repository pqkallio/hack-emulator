package graphics

type Interface interface {
	Draw()
	ShouldClose() bool
	Terminate()
}
