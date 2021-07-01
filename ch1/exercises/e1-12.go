package main

import (
	"log"
	"net/http"
	"image"
        "image/color"
        "image/gif"
        "io"
        "math"
        "math/rand"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
        whiteIndex = 0
        blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	res := .001
	size := 100
	nframes := 64
	delay := 8

	for par, val := range r.Form {
		if par == "cycles" {
			i, err := strconv.Atoi(val[0])
			if err != nil {
				log.Print(err)
			} else {
				cycles = i
			}
		} else if par == "res" {
			i, err := strconv.ParseFloat(val[0], 64)
			if err != nil {
				log.Print(err)
			} else {
				res = i
			}
		} else if par == "size" {
			i, err := strconv.Atoi(val[0])
			if err != nil {
				log.Print(err)
			} else {
				size = i
			}
		} else if par == "nframes" {
			i, err := strconv.Atoi(val[0])
			if err != nil {
				log.Print(err)
			} else {
				nframes = i
			}
		} else if par == "delay" {
			i, err := strconv.Atoi(val[0])
			if err != nil {
				log.Print(err)
			} else {
				delay = i
			}
		}
	}
	lissajous(w, cycles, res, size, nframes, delay)
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
        freq := rand.Float64() * 3.0
        anim := gif.GIF{LoopCount: nframes}
        phase := 0.0
        for i:= 0; i < nframes; i++ {
                rect := image.Rect(0, 0, 2*size+1, 2*size+1)
                img := image.NewPaletted(rect, palette)
                for t:= 0.0; t < float64(cycles*2)*math.Pi; t += res{
                        x:= math.Sin(t)
                        y := math.Sin(t*freq+phase)
                        img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
                }

                phase += 0.1
                anim.Delay = append(anim.Delay, delay)
                anim.Image = append(anim.Image, img)
        }

        gif.EncodeAll(out, &anim)
}
