package topo

// Buffer is the ETOPO buffer
type Buffer struct {
	Width  int
	Height int
	Data   []int16
	Max    int16
	Min    int16
	Diff   int16
}
