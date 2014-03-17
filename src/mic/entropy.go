package mic

import "math"

func Entropy(pdf []float64) float64 {
    H := 0.0
    for _, p := range pdf {
        if( p > 0.0 ) {
            H += math.Log2(p) * p
        }
    }
    return -H
}