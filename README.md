#go-cp
Copying files made easy. Golang library.

# Installation

Run `go get github.com/n-marshall/go-cp`

# Usage

Simply use `cp.Copy(sourceFile, destinationFile)`
Full documentation found [here](https://godoc.org/github.com/nmrshll/go-cp)

## Example

```go
import "github.com/n-marshall/go-cp"

func main() {
  err := cp.Copy("file.ext", "file.copy.ext")
  if err != nil {
    log.Fatal(err)
  }
}
```
