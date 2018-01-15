package main

import (
	"log"

	cp "github.com/nmrshll/go-cp"
)

func main() {
	err := cp.CopyFile("examples/examplefile.ext", "examples/examplefile.copy.ext")
	if err != nil {
		log.Fatal(err)
	}
}
