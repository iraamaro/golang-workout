package main

import(
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    start := time.Now()
    fmt.Println(start)
    fmt.Println(strings.Join(os.Args[:], " "))
    fmt.Println("%.2fs elapsed\n", time.Since(start).Seconds())
}
