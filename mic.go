package main

import (
    "fmt"
    "mic"
)

func main() {
    var arr [2]float64
    arr[0] = 0.5
    arr[1] = 0.5
    H := mic.Entropy(arr[:])
    fmt.Println(H)
}