package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Uso da variável 'i'\n")
    fmt.Println("Tamanho de 'len(os.Args[])'")
    fmt.Println(len(os.Args))
    for i := 0; i < len(os.Args); i++ {
        fmt.Println("\ni=\n")
        fmt.Println(i)
        fmt.Println("\nValor de os.Args[i]\n")
        fmt.Println(os.Args[i])
    }
}
