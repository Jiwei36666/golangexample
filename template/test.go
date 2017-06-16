package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!\n")
	p := Person{UserName: "hustcat"}
	t.Execute(os.Stdout, p)
}
