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

	for x := 0; x < etopoScaled.Width; x++ {
		for y := 0; y < etopoScaled.Height; y++ {
			var c color.RGBA
			altitude := etopoScaled.GetAltitude(x, y)
			if altitude > 0 {
				coef := float32(1) - float32(altitude)/float32(etopoScaled.Max)
				r, g, b := color.YCbCrToRGB(uint8(coef*128), uint8(80), uint8(80))
				c = color.RGBA{
					r,
					g,
					b,
					255,
				}
			} else {
				coef := float32(altitude) / float32(etopoScaled.Min)
				c = color.RGBA{
					0,
					0,
					uint8(coef * 255),
					255,
				}
			}
			myimage.Set(int(x), int(y), c)
		}
	}

	myfile, _ := os.Create("test.png")

	png.Encode(myfile, myimage)

}
