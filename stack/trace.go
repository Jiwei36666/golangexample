package main

import (
    "runtime"
    "fmt"
)

func main() {
    outer()
}

func outer() {
    inner()
}

func inner() {

    defer func() {
        if err := recover(); err != nil {
            trace := make([]byte, 1024)
            count := runtime.Stack(trace, true)
            fmt.Printf("Recover from panic: %s\n", err)
            fmt.Printf("Stack of %d bytes: %s\n", count, trace)
        }
    }()

    panic("Fake error!")


}

