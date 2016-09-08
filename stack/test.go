package main

import "fmt"

type D struct{
    a int
    b string
}

type StartCallback func(*D, int)

type trace struct{}

func main() {
	//slice := make([]string, 2, 4)

	var t trace
    d := &D{1, "aaa"}
	t.Example(d, "hello", f)
}

func f(d *D, a int){

}

//func (t *trace) Example(slice []string, str string, i int) {
func (t *trace) Example(d *D, str string, f StartCallback)(int, error) {
	fmt.Printf("Receiver Address: %p\n", t)
	panic("Want stack trace")
}
