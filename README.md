# go-cp
### Golang library to make copying files easy

Copying a file in vanilla Go forces you to deal with filesystem I/O, it's a bit overcomplicated for the general case, when you just want to copy a file as easy as using `cp $source $dest`.
This tiny, one-use-case library is just for that.

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
