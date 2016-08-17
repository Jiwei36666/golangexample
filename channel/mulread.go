package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("in p1")
		//c <- 1
        close(c)
        fmt.Println("p1 exit")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("in c1")
		<-c
		fmt.Println("c1 exit")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("in c2")
		<-c
		fmt.Println("c2 exit")
	}()

	wg.Wait()
    fmt.Println("main exit")
}
