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

    return img, nil
}

func ToAscii(img image.Image) {
    chars := " `'.,:;i+o*%&$#@"
    lenChars := len(chars)

    for y := 0; y < img.Bounds().Dy(); y++ {
        for x := 0; x < img.Bounds().Dx(); x++ {
            intensity := GetPixelIntensity(img, x, y, lenChars)
            fmt.Printf("%c", chars[iabs(intensity - lenChars) % lenChars])
        }
        fmt.Println("")
    }
}

func GetPixelIntensity(img image.Image, x, y, lenChars int) int {
    pixel := img.At(x, y)
    r, g, b, _ := color.GrayModel.Convert(pixel).RGBA()

    intensity := int(r + g + b)
    return intensity*lenChars / 768
}

func iabs(x int) int {
    if (x < 0) {
        return -x
    }

    return x
}
