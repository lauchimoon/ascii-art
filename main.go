package main

import (
    "fmt"
    "log"
    "image"
    "image/color"
    "image/png"
    "os"
)

const (
    PROGRAM_NAME = "asciiart"
)

func main() {
    log.SetFlags(0)

    args := os.Args
    if len(args) < 2 {
        log.Fatalf("%s: png file is missing\n", PROGRAM_NAME)
    }

    img, err := DecodeImage(args[1])
    if err != nil {
        log.Fatalf("%v: %v\n", PROGRAM_NAME, err)
    }

    ToAscii(img)
}

func DecodeImage(path string) (image.Image, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    img, err := png.Decode(f)
    if err != nil {
        return nil, err
    }

    return ToGrayscale(img), nil

}

func ToGrayscale(img image.Image) image.Image {
    newImg := image.NewRGBA(img.Bounds())
    for y := 0; y < img.Bounds().Dy(); y++ {
        for x := 0; x < img.Bounds().Dx(); x++ {
            r, g, b, _ := img.At(x, y).RGBA()
            luminosity := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b) 

            clr := color.RGBA{uint8(luminosity/256), uint8(luminosity/256), uint8(luminosity/256), 255}
            newImg.Set(x, y, clr)
        }
    }

    return newImg
}

func ToAscii(img image.Image) {
    chars := " `'.,:;i+o*%&$#@"
    lenChars := len(chars)

    for y := 0; y < img.Bounds().Dy(); y++ {
        for x := 0; x < img.Bounds().Dx(); x++ {
            r, g, b, _ := img.At(x, y).RGBA()
            var intensity int = int(r + g + b)

            intensity = intensity*lenChars / 768
            fmt.Printf("%c", chars[iabs(intensity - lenChars) % lenChars])
        }
        fmt.Println("")
    }
}

func iabs(x int) int {
    if (x < 0) {
        return -x
    }

    return x
}
