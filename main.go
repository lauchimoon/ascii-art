package main

import (
    "fmt"
    "log"
    "image"
    "image/color"
    "image/png"
    "os"
)

func main() {
    chars := " `'.,:;i+o*%&$#@"
    lenChars := len(chars)

    f, err := os.Open("sample3.png")
    if err != nil {
        log.Fatalf("%v\n", err)
    }
    defer f.Close()

    sample, err := png.Decode(f)
    if err != nil {
        log.Fatalf("%v\n", err)
    }

    imgGray := ToGrayscale(sample)

    for y := 0; y < imgGray.Bounds().Dy(); y++ {
        for x := 0; x < imgGray.Bounds().Dx(); x++ {
            r, g, b, _ := imgGray.At(x, y).RGBA()
            var intensity int = int(r + g + b)

            intensity = intensity*lenChars / 768
            fmt.Printf("%c", chars[iabs(intensity - lenChars) % lenChars])
        }
        fmt.Println("")
    }
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

func iabs(x int) int {
    if (x < 0) {
        return -x
    }

    return x
}
