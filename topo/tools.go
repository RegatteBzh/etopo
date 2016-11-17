package topo

import "errors"

// GetAltitude get an altitude in the buffer
func (buffer Buffer) GetAltitude(x uint32, y uint32) int16 {
	return buffer.Data[y*buffer.Width+x]
}

// SetScale apply a scale on a buffer
func (buffer Buffer) SetScale(scale float32) (Buffer, error) {
	if scale > 1 {
		return Buffer{}, errors.New("scale must be lower than 1")
	}
	width := uint32(float32(buffer.Width) * scale)
	height := uint32(float32(buffer.Height) * scale)
	newBuffer := Buffer{
		width,
		height,
		make([]int16, width*height),
		buffer.Max,
		buffer.Min,
		buffer.Diff,
	}

	for y := uint32(0); y < height; y++ {
		for x := uint32(0); x < width; x++ {
			newBuffer.Data[y*newBuffer.Width+x] = buffer.GetAltitude(uint32(float32(x)/scale), uint32(float32(y)/scale))
		}
	}

	return newBuffer, nil
}
