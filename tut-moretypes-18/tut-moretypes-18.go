package main

import "golang.org/x/tour/pic"
//import "fmt"

func Pic(dx, dy int) [][]uint8 {
    s := make([][]uint8, dy)
    for y:=0; y<dy; y++ {
        s[y] = make([]uint8, dx)
        for x:=0; x<dx; x++ {
            s[y][x] = uint8( 256*float32(x)/(float32(y)+.0001) )
        }
    }
    return s
}

func main() {
    pic.Show(Pic)
}
