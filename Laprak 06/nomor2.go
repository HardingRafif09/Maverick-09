package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    var ikan [1000]float64

    for i := 0; i < x; i++ {
        fmt.Scan(&ikan[i])
    }

    var totalWadah [1000]float64
    idx := 0

    for i := 0; i < x; i += y {
        total := 0.0
        for j := i; j < i+y && j < x; j++ {
            total += ikan[j]
        }
        totalWadah[idx] = total
        idx++
    }


    for i := 0; i < idx; i++ {
        fmt.Printf("%.2f ", totalWadah[i])
    }
    fmt.Println()

    sum := 0.0
    for i := 0; i < idx; i++ {
        sum += totalWadah[i]
    }

    rata := sum / float64(idx)
    fmt.Printf("%.2f\n", rata)
}