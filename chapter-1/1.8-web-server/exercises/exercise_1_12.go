package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	blackIndex = 1 //next colour in the palette
)

func main() {
	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()

		cycles, err := strconv.Atoi(keys.Get("cycles"))

		if err != nil {
			log.Print("Failed to get cycles")
			cycles = 5
		}

		size, err := strconv.Atoi(keys.Get("size"))
		if err != nil {
			log.Print("Failed to get size")
			size = 100
		}


		lissajous(w, float64(cycles), size)
	})
	log.Fatal(http.ListenAndServe("localhost:8006", nil))
}

func lissajous(out io.Writer, cycles float64, size int) {

	const (
		res     = 0.001 // angular resolution
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
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE:- ignoring encoding errors
}