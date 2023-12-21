package 类库

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

func main33() {
	fmt.Println(111)

	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	saveImage(m, "img403.png")
}

//func TestHandler1002(t *testing.T) {
//	fmt.Println(123)
//
//}
