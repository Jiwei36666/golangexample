package main

import (
"fmt"
"path/filepath"
)

func main(){
    fmt.Printf("%s\n", filepath.Base("/rootfs/testdir"))
}
