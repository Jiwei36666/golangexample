package main

import (
    "fmt"
	"text/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t, err := template.ParseFiles("./t.json")
    if err != nil {
        fmt.Printf("error %v", err)
        return
    }

    conf := make(map[string]interface{})
	conf["ip"]="127.11.1.1"
    conf["port"]=80
    fmt.Println(t.Name())
	t.Execute(os.Stdout, conf)
}
