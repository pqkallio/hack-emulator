package util

import "fmt"

func GetBoolFromUint16(val, offset uint16) bool {
	if offset > 15 {
		panic(fmt.Sprintf("invalid offset: %d", offset))
	}

	return val&(1<<offset) > 0
}
