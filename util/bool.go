package util

import (
	"fmt"
	"math/rand"
)

func GetBoolFromUint16(val, offset uint16) bool {
	if offset > 15 {
		panic(fmt.Sprintf("invalid offset: %d", offset))
	}

	return val&(1<<offset) > 0
}

func RandomBool() bool {
	return rand.Uint32()%2 == 0
}
