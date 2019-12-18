package main

import (
    "fmt"
    "os"

    "github.com/mackerelio/go-osstat/memory"
    "github.com/mackerelio/go-osstat/cpu"
    "github.com/mackerelio/go-osstat/disk"
    "time"
)

func main() {
    memory, err := memory.Get()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }
    fmt.Printf("memory total: %d bytes\n", bToMb(memory.Total))
    fmt.Printf("memory used: %d bytes\n", bToMb(memory.Used))
    fmt.Printf("memory cached: %d bytes\n", bToMb(memory.Cached))
    fmt.Printf("memory free: %d bytes\n", bToMb(memory.Free))


    before, err := cpu.Get()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }
    time.Sleep(time.Duration(1) * time.Second)
    after, err := cpu.Get()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }
    total := float64(after.Total - before.Total)
    fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
    fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
    fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)

    disks, err := disk.Get()
    if err != nil {
        fmt.Printf("error should be nil but got: %v", err)
    }
    fmt.Printf("disks value: %+v", disks)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
