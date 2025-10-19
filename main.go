package main

import (
    "fmt"
    "log"
    "image/png"
    "os"
)

func main() {
    chars := " `'.,:;i+o*%&$#@"
    lenChars := len(chars)

    f, err := os.Open("sample2.png")
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
            r, g, b, _ := clr.RGBA()
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
