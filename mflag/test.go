//ref: https://lawlessguy.wordpress.com/2013/07/23/filling-a-slice-using-command-line-flags-in-go-golang/
package main
 
import (
    "flag"
    "fmt"
    "strconv"
)
 
// Define a type named "intslice" as a slice of ints
type intslice []int
 
// Now, for our new type, implement the two methods of
// the flag.Value interface...
// The first method is String() string
func (i *intslice) String() string {
    return fmt.Sprintf("%d", *i)
}
 
// The second method is Set(value string) error
func (i *intslice) Set(value string) error {
    fmt.Printf("%s\n", value)
    tmp, err := strconv.Atoi(value)
    if err != nil {
        *i = append(*i, -1)
    } else {
        *i = append(*i, tmp)
    }
    return nil
}
 
var myints intslice
 
func main() {
    flag.Var(&myints, "i", "List of integers")
    flag.Parse()
    if flag.NFlag() == 0 {
        flag.PrintDefaults()
    } else {
        fmt.Println("Here are the values in 'myints'")
        for i := 0; i < len(myints); i++ {
            fmt.Printf("%d\n", myints[i])
        }
    }
}
