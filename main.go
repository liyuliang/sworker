package main

import (
    "fmt"
    "os"

    "github.com/mackerelio/go-osstat/memory"
)

func main() {
    memory, err := memory.Get()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }
    fmt.Printf("memory total: %d bytes\n", memory.Total)
    fmt.Printf("memory used: %d bytes\n", memory.Used)
    fmt.Printf("memory cached: %d bytes\n", memory.Cached)
    fmt.Printf("memory free: %d bytes\n", memory.Free)
}
