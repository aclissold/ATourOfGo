package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
    XStart, YStart, XEnd, YEnd int
}

func (i *Image) Bounds() image.Rectangle {
    return image.Rect(i.XStart, i.YStart, i.XEnd, i.YEnd)
}

func (i *Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
    m := Image{0, 0, 100, 100}
    pic.ShowImage(&m)
}
