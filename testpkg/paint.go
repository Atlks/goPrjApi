package testpkg

import (
	"bufio"
	"fmt"
	"goapiPrj/lib"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"testing"
)

func TestHandler100(t *testing.T) {

	f := "c:/w/egls850wd.txt"

	//	f2 := "egls850wdAbbr.txt"
	lib.ReadLine(f, func(line string) {
		fmt.Println(line)

		//firstCh := lib.Left(line, 1)
		//other := lib.Substr(line, 1)
		//lib.Write(f2, line+"\n")
	})

	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	saveImage(m, "img403.png")
}

func saveImage(img image.Image, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = jpeg.Encode(b, img, nil)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}
