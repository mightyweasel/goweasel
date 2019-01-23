package weaselmodels

import (
    "fmt"
)

func init() {
	// we will use this later
	fmt.Println("Init: Weasel Models Initialized!")
}
// dummy function to show usage of package compilation
func NoOp() {
    i := "no-op success"
    fmt.Println("weaselmodels: " + i)        
}
