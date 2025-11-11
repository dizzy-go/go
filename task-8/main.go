package main

import "fmt"

func main() {
    var N int
    total := 0
    fmt.Scan(&N)
    for i := 0; i <= N; i++ {
        total += i
    }
    fmt.Println(total)
}
