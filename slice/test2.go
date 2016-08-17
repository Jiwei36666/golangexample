package main

import "fmt"
import "strings"

func main() {
	args := [][2]string{{"IP", "10.55.206.46"}, {"arg2", "value2"}}
    fmt.Println("Elements of args: ")
	for _, v := range args {
		fmt.Println(strings.Join(v[:], "="))
	}
	fmt.Println()
}
