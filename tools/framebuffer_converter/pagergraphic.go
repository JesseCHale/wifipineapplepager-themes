// Pagergraphic converter
// (c) 2025 Hak5
//
// Code licensed under GPLv2+
//
// mike@hak5.org

package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"reflect"
	"unsafe"

	"pagergraphic/rgb565"
)

func spinImageToBuffer(img *rgb565.Image) []rgb565.Color {
	var ret = make([]rgb565.Color, 480*222)

	for ix := 0; ix < 480; ix++ {
		for iy := 0; iy < 222; iy++ {
			fy := ix
			fx := 222 - iy - 1

			ret[fy*222+fx] = img.RGB565At(ix, iy)
		}
	}

	return ret
}

func writeRotatedImg(img *rgb565.Image, fb *os.File) {
	buf := spinImageToBuffer(img)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&(buf)))
	data := (*[480 * 222 * 2]byte)(unsafe.Pointer(hdr.Data))[:]
	_, _ = fb.Write(data[:])
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("ERROR: Expected ", os.Args[0], " [input png] [output fb]")
		return
	}

	pngfile, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("ERROR: Could not open ", os.Args[1], "for reading: ", err.Error())
		return
	}
	defer func() {
		_ = pngfile.Close()
	}()

	pngimg, err := png.Decode(pngfile)
	if err != nil {
		fmt.Println("ERROR: Could not decode PNG ", os.Args[1], ": ", err.Error())
		return
	}

	fb, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("ERROR: Could not open ", os.Args[2], " for writing: ", err.Error())
	}
	defer func() {
		_ = fb.Close()
	}()

	img := rgb565.New(image.Rect(0, 0, 480, 222))

	draw.Draw(img, pngimg.Bounds(), pngimg, image.Point{0, 0}, draw.Over)

	writeRotatedImg(img, fb)
}
