package main

import (
	"fmt"

	"github.com/glenntrewitt/reference/demo"
)

func main() {
	var d demo.S
	d.Value.Value = 10
	fmt.Printf("%#v\n", d)
}
