// Exercise 1.5 & 1.6 :- Change lissajous to use different colours
// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// rgba(140, 20, 252, 1)
var palette = []color.Color{color.Black, color.RGBA{46, 204, 113, 1}, color.RGBA{140, 20, 252, 1}}

const (
	greenIndex  = 1 // next colour in the palette
	purpleIndex = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			var colourToUse uint8 = purpleIndex
			if i%2 == 0 {
				colourToUse = greenIndex
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colourToUse)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE:- ignoring encoding errors
}
