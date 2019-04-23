package main

import (
	"fmt"

	"github.com/jsliacan/hello/hello"
	"github.com/jsliacan/hello/wave"
)

func main() {

	// say hello
	fmt.Println(hello.Greet())

	// wave
	fmt.Println(wave.Wave())
}