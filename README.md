# mctext
Random-text generator using Markov chain from an user-provided sample.

Example:
```go

import "github.com/zgiber/mctext"

var r io.Reader
// provide an io.Reader with the sample. Ideally UTF encoded text.
// ...

mc := &mct.MChain{Nodes: map[string]*mctext.Node{})
mc.Parse(r)

output := mc.Generate(100) // Generates text with 100 words.

```
