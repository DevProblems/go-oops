package main

import (
	"example.com/devproblems/go-oops/structs"
	"fmt"
)

func main() {
	p := structs.NewPerson("sarang", "devproblms", 12, []string{"a"})
	p1 := structs.NewPerson("sarang", "devproblems", 12, []string{"a"})
	r := p.Equals(p1)
	if r {
		fmt.Println("both the structs are equal")
	}

}
