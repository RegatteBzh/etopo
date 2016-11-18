package topo

import (
	"errors"
	"image"
	"math"
)

// Map is the ETOPO buffer
type Map struct {
	Width  int // Map width
	Height int // Map Height
	CellW  int // Cell Width in minutes
	CellH  int // Cell Height in minutes
	Data   []int16
	Max    int16
	Min    int16
	Diff   int16
}

// ComputeParameters compute diff, min and max
func (buffer Map) ComputeParameters() {
	for _, alt := range buffer.Data {
		if alt > buffer.Max {
			buffer.Max = alt
		}
		if alt < buffer.Min {
			buffer.Min = alt
		}
	}
	buffer.Diff = buffer.Max - buffer.Min
}

// GetAltitude get an altitude in the buffer
func (buffer Map) GetAltitude(loc image.Point) int16 {
	return buffer.Data[loc.Y*buffer.Width+loc.X]
}

// SetAltitude set an altitude in the buffer
func (buffer Map) SetAltitude(loc image.Point, alt int16) {
	buffer.Data[loc.Y*buffer.Width+loc.X] = alt
}

//Extract extract a rectangle in the buffer
func (buffer Map) Extract(start image.Point, end image.Point) (Map, error) {
	x0 := int(math.Min(float64(start.X), float64(end.X)))
	x1 := int(math.Max(float64(start.X), float64(end.X)))
	y0 := int(math.Min(float64(start.Y), float64(end.Y)))
	y1 := int(math.Max(float64(start.Y), float64(end.Y)))

	newBuffer := Map{
		Width:  x1 - x0,
		Height: y1 - y0,
		CellH:  buffer.CellH,
		CellW:  buffer.CellW,
		Data:   make([]int16, (x1-x0)*(y1-y0)),
	}

	for y := 0; y < buffer.Height; y++ {
		for x := 0; x < buffer.Width; x++ {
			newBuffer.SetAltitude(image.Point{x, y}, buffer.GetAltitude(image.Point{x + x0, y + y0}))
		}
	}

	newBuffer.ComputeParameters()

	return newBuffer, nil
}

// SetScale apply a scale on a buffer
func (buffer Map) SetScale(scale float32) (Map, error) {
	if scale > 1 {
		return Map{}, errors.New("scale must be lower than 1")
	}
	width := int(float32(buffer.Width) * scale)
	height := int(float32(buffer.Height) * scale)
	newBuffer := Map{
		Width:  width,
		Height: height,
		CellH:  buffer.CellH,
		CellW:  buffer.CellW,
		Data:   make([]int16, width*height),
		Max:    buffer.Max,
		Min:    buffer.Min,
		Diff:   buffer.Diff,
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newBuffer.Data[y*newBuffer.Width+x] = buffer.GetAltitude(image.Point{int(float32(x) / scale), int(float32(y) / scale)})
		}
	}

	return newBuffer, nil
}
