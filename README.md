## go-emoji

transform emoji markup to emoji

### Example

```go
import (
  "github.com/kitagry/go-emoji"
  "golang.org/x/text/transform"
)

r := emoji.NewReplacer()
tf := transform.NewReader(strings.NewReader("Hello World:blush:"), r)
io.Copy(os.Stdout, tf)
// Hello World😊
```
