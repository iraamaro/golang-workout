package main

import(
    "fmt"
    "os"
    "strings"
    "testing"
)

func BenchmarkArguments(b *testing.B) {
    for i:= 0; i < b.N; i++ {
        fmt.Println(strings.Join(os.Args[:], " "))
    }
}
