package io

type Interface interface {
	Process()
	ShouldClose() bool
	Terminate()
}
