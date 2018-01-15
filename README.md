#go-cp
Copying files made easy. Golang library.

# Installation

Run `go get github.com/n-marshall/go-cp`

# Usage

Simply use `cp.Copy(sourceFile, destinationFile)`
Full documentation found [here](https://godoc.org/github.com/nmrshll/go-cp)

## Example

[embedmd]:# (./examples/example.go)
```go
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
```
