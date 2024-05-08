package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

var blankColor = color.RGBA{110, 110, 110, 255}
var treeColor = color.RGBA{63, 171, 51, 255}
var fireColor = color.RGBA{255, 69, 41, 255}
var palette = color.Palette{blankColor, treeColor, fireColor}

var upLeft = image.Point{0, 0}

// var resultChan = make(chan *image.Paletted)

func (f *Forest) generateArtPalletted() *image.Paletted {
	img := image.NewPaletted(image.Rectangle{upLeft, point(f.width, f.height)}, palette)
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			v, ok := f.forestMap.Load(coords{x, y})
			if ok && v == 1 {
				img.Set(x, y, treeColor)
			} else if ok && v == 2 {
				img.Set(x, y, fireColor)
			} else {
				img.Set(x, y, blankColor)
			}

		}
	}
	return img
	// f.gifArray = append(f.gifArray, img)
}

func (f *Forest) saveCurrentFramePaletted() {
	img := image.NewPaletted(image.Rectangle{upLeft, point(f.width, f.height)}, palette)
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			v, ok := f.forestMap.Load(coords{x, y})
			if ok && v == 1 {
				img.Set(x, y, treeColor)
			} else if ok && v == 2 {
				img.Set(x, y, fireColor)
			} else {
				img.Set(x, y, blankColor)
			}

		}
	}
	//return img
	f.gifArray = append(f.gifArray, img)
}

func filterGifArray(gifArray []*image.Paletted) []*image.Paletted {
	var filtered []*image.Paletted
	for i, value := range gifArray {
		if i%20 == 0 {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func saveGif(imgs []*image.Paletted, filename string) {
	f, _ := os.Create(filename + ".gif")
	if e := gif.EncodeAll(f, &gif.GIF{
		Image: imgs,
		Delay: make([]int, len(imgs)),
	}); e != nil {
		fmt.Println(e)
	}
}

func point(w, h int) image.Point {
	return image.Point{w, h}
}

func (f *Forest) gifWorker() {
	//for img, ok := <-f.frameChannel; ok; img, ok = <-f.frameChannel {
	//	f.gifArray = append(f.gifArray, img)
	//}
	for img := range f.frameChannel {
		// fmt.Println("zapisuję")
		f.gifArray = append(f.gifArray, img)
	}
	f.endChannel <- true
	//f.wggif.Done()
}

// func GenerateArtRGBA(worldMap *sync.Map) *image.RGBA {
// 	img := image.NewRGBA(image.Rectangle{upLeft, lowRight()})
// 	for x := 0; x < globalWidth; x++ {
// 		for y := 0; y < globalHeight; y++ {
// 			v, ok := worldMap.Load(coords{x, y})
// 			if ok && v == 1 {
// 				img.Set(x, y, treeColor)
// 			} else if ok && v == 2 {
// 				img.Set(x, y, fireColor)
// 			} else {
// 				img.Set(x, y, blankColor)
// 			}

// 		}
// 	}
// 	return img
// }
