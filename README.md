go-cp


Copy file function in golang. Returns nil on success otherwise an error if there were any issues.

```go
import "github.com/n-marshall/go-cp"

func main() {
  err := cp.Copy("file.ext", "file.copy.ext")
  if err != nil {
    panic(err)
  }
}
```
