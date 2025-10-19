package main

import (
    "fmt"
    "log"
    "image/png"
    "image/color"
    "os"
)

const (
    BLACKCHAR = '\''
    WHITECHAR = 'M'
)

var (
    BLACK = color.RGBA{0, 0, 0, 255}
)

func main() {
    f, err := os.Open("sample.png")
    if err != nil {
        log.Fatalf("%v\n", err)
    }
    defer f.Close()

    sample, err := png.Decode(f)
    if err != nil {
        log.Fatalf("%v\n", err)
    }

    for y := 0; y < sample.Bounds().Dy(); y++ {
        for x := 0; x < sample.Bounds().Dx(); x++ {
            clr := sample.At(x, y)
            if clr == BLACK {
                fmt.Printf("%c", BLACKCHAR)
            } else {
                fmt.Printf("%c", WHITECHAR)
            }
        }
        fmt.Println("")
    }
}
