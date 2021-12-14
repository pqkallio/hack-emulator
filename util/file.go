package util

import (
	"errors"
	"io"
	"os"
)

func ReadRomFile(romFile string) (rom []uint16, err error) {
	f, err := os.Open(romFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 2)

	for err == nil {
		_, err = f.Read(buf)
		rom = append(rom, uint16(buf[0])<<8|uint16(buf[1]))
	}

	if !errors.Is(err, io.EOF) {
		return nil, err
	}

	return rom, nil
}
