package topo

import "errors"

// GetAltitude get an altitude in the buffer
func (buffer Buffer) GetAltitude(x int, y int) int16 {
	return buffer.Data[y*buffer.Width+x]
}

// SetScale apply a scale on a buffer
func (buffer Buffer) SetScale(scale float32) (Buffer, error) {
	if scale > 1 {
		return Buffer{}, errors.New("scale must be lower than 1")
	}
	width := int(float32(buffer.Width) * scale)
	height := int(float32(buffer.Height) * scale)
	newBuffer := Buffer{
		width,
		height,
		make([]int16, width*height),
		buffer.Max,
		buffer.Min,
		buffer.Diff,
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newBuffer.Data[y*newBuffer.Width+x] = buffer.GetAltitude(int(float32(x)/scale), int(float32(y)/scale))
		}
	}

	return newBuffer, nil
}
