# mctext

Random-text generator using Markov chain from an user-provided sample.

Example:

```go

import "github.com/zgiber/mctext"

// provide an io.Reader with the sample. Ideally UTF encoded text.
// ...
var r io.Reader

mc := mct.New()
mc.Parse(r)

output := mc.Generate(100) // Generates text with 100 words.

```
