package lib

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/jpeg"
	"os"
)

func M2() {
	fmt.Println("MM222")
}

func MainPain() {

	//f := "c:/w/egls850wd.txt"
	//
	////	f2 := "egls850wdAbbr.txt"
	//lib.ReadLine(f, func(line string) {
	//	fmt.Println(line)
	//
	//	//firstCh := lib.Left(line, 1)
	//	//other := lib.Substr(line, 1)
	//	//lib.Write(f2, line+"\n")
	//})

	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	//	dc.SetRGB(0, 0, 0)
	dc.SetRGB(0.5, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")

}

func saveImageV2(img image.Image, filename string) error {
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
