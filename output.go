package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"time"
)

func TimeStamp() string {
	timestamp := time.Now()
	hr, mn, sc := timestamp.Clock()
	yr, mo, dy := timestamp.Date()
	return fmt.Sprintf("otp-s%vm%vh%v-%v%v%v.png", sc, mn, hr, dy, mo, yr)
}

func BinaryFilter(lattice []string) *[]string {
	out := []string{}
	for _, s := range lattice {

		newS := ""
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '0':
				newS += "0"
			case '1':
				newS += "1"
			case '2':
				newS += "1"
			case '3':
				newS += "1"
			case '4':
				newS += "0"
			case '5':
				newS += "1"
			case '6':
				newS += "1"
			case '7':
				newS += "0"
			default:
				newS += "-"
			}
		}
		out = append(out, newS)
	}

	return &out
}

func PrintTimeStamp(flagOut *string,
	flagLat *string,
	flagGens *int,
	flagIc *string) {

	fmt.Println("Program Arguments")
	fmt.Println("-out: " + *flagOut)
	fmt.Println("-lat: " + *flagLat)
	fmt.Println("-gens: " + strconv.Itoa(*flagGens))
	fmt.Println("IC: " + *flagIc)
	fmt.Println(time.Now())
}

func PrintBinaryFilter(lattice *[]string) {
	binlat := BinaryFilter(*lattice)
	for i := 0; i < len(*binlat); i++ {
		fmt.Println((*binlat)[i])
	}
}

func PrintStateFilter(lattice *[]string) {
	for i := 0; i < len(*lattice); i++ {
		fmt.Println((*lattice)[i])
	}
}

func PngBinaryFilter(lattice *[]string) {
	height := len(*lattice)
	width := len((*lattice)[height-1])

	upleft := image.Point{0, 0}
	lowright := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upleft, lowright})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch (*lattice)[y][x] {
			case '0':
				img.Set(x, y, color.White)
			case '1':
				img.Set(x, y, color.Black)
			case '2':
				img.Set(x, y, color.Black)
			case '3':
				img.Set(x, y, color.Black)
			case '4':
				img.Set(x, y, color.White)
			case '5':
				img.Set(x, y, color.Black)
			case '6':
				img.Set(x, y, color.Black)
			case '7':
				img.Set(x, y, color.White)
			}
		}
	}

	f, err := os.Create(TimeStamp())
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)
}

func PngStateFilter(lattice *[]string) {
	height := len(*lattice)
	width := len((*lattice)[height-1])

	upleft := image.Point{0, 0}
	lowright := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upleft, lowright})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch (*lattice)[y][x] {
			case '0':
				img.Set(x, y, color.White)
			case '1':
				img.Set(x, y, color.RGBA{255, 0, 0, 0xff})
			case '2':
				img.Set(x, y, color.RGBA{255, 127, 0, 0xff})
			case '3':
				img.Set(x, y, color.RGBA{255, 255, 0, 0xff})
			case '4':
				img.Set(x, y, color.RGBA{0, 255, 0, 0xff})
			case '5':
				img.Set(x, y, color.RGBA{0, 0, 255, 0xff})
			case '6':
				img.Set(x, y, color.RGBA{75, 0, 130, 0xff})
			case '7':
				img.Set(x, y, color.RGBA{148, 0, 211, 0xff})
			}
		}
	}

	f, err := os.Create(TimeStamp())
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)
}
