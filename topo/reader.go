package topo

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

//ReadEtopo read ETOPO binary
func ReadEtopo(file io.Reader, width int, height int) (buffer Buffer, err error) {
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

	buffer.ComputeParameters()

	return
}
