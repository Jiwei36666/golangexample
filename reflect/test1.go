package main

import (
	"fmt"
	"reflect"
    "io"
    "os"
)

func main() {
/*
	t := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)
*/
var w io.Writer = os.Stdout
fmt.Println(reflect.ValueOf(w)) // "*os.File"
}
