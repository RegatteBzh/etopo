package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/regattebzh/etopo/topo"
)

//NCOLS are the number of columns in the ETOPO1
const NCOLS = 21601

//NROWS are the number of rows in the ETOPO1
const NROWS = 10801

func main() {
	etopoFile, err := os.Open("./data/etopo1_ice_g_i2.bin")
	if err != nil {
		log.Fatal(err) //log.Fatal run an os.Exit
	}
	defer etopoFile.Close()

	etopoData, err := topo.ReadEtopo(etopoFile, NCOLS, NROWS)

	etopoScaled, err := etopoData.SetScale(0.2)

	myimage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{int(etopoScaled.Width), int(etopoScaled.Height)}})

	for x := uint32(0); x < etopoScaled.Width; x++ {
		for y := uint32(0); y < etopoScaled.Height; y++ {
			var c color.RGBA
			altitude := etopoScaled.GetAltitude(x, y)
			if altitude > 0 {
				c = color.RGBA{
					uint8(altitude * 255 / etopoScaled.Max),
					0,
					0,
					255,
				}
			} else {
				c = color.RGBA{
					0,
					0,
					uint8(altitude * 255 / etopoScaled.Min),
					255,
				}
			}
			myimage.Set(int(x), int(y), c)
		}
	}

	myfile, _ := os.Create("test.png")

	png.Encode(myfile, myimage)

}
