package topo

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

// Buffer is the ETOPO buffer
type Buffer struct {
	Width  uint32
	Height uint32
	Data   []int16
	Max    int16
	Min    int16
	Diff   int16
}

//ReadEtopo read ETOPO binary
func ReadEtopo(file io.Reader, width uint32, height uint32) (buffer Buffer, err error) {
	buffer = Buffer{
		width,
		height,
		make([]int16, width*height),
		0,
		0,
		0,
	}
	preData := make([]byte, width*height*2)
	if _, err = file.Read(preData); err != nil {
		log.Fatal("file.Read failed (ReadEtopo)\n", err)
	}

	dataBuf := bytes.NewReader(preData)
	if err = binary.Read(dataBuf, binary.LittleEndian, &buffer.Data); err != nil {
		log.Fatal("Byte to uint32 failed\n", err)
		return
	}

	for _, alt := range buffer.Data {
		if alt > buffer.Max {
			buffer.Max = alt
		}
		if alt < buffer.Min {
			buffer.Min = alt
		}
	}

	buffer.Diff = buffer.Max - buffer.Min

	return
}
