package main

import "fmt"

type Parham struct{}

func (*Parham) String() string {
	return "パルハム"
}

func funPrinting() {
	fmt.Printf("Hello ye, %s\n", new(Parham))
}

func main() {
    funPrinting()
}
