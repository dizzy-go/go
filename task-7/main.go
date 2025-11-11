package main

import "fmt"

func main() {
    var score int
    fmt.Scan(&score)
    switch {
        case score >= 90 && score <= 100:
            fmt.Println("A")
        case score >= 80 && score <= 89:
            fmt.Println("B")
        case score >= 70 && score <= 79:
            fmt.Println("C")
        case score >= 60 && score <= 69:
            fmt.Println("D")
        case score >= 0 && score <= 59:
            fmt.Println("F")
        default: 
            fmt.Println("invalid")
    }
}
