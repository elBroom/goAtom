package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "абв"
    fmt.Println(len(s))

    /*for _, r := range s {
        fmt.Printf("%c\n", r)
    }*/

    fmt.Println(strings.HasPrefix(s, "а"))
}
