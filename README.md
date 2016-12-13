# mctext

Random-text generator using Markov chain from an user-provided sample. It's useful in cases where generated text must be relatively random,
but expected to contain a good average distribution and repetition of words.

Example:

```go

import "github.com/zgiber/mctext"

// provide an io.Reader with the sample. Ideally UTF encoded text.
// ...
var r io.Reader

mc := mct.New()
mc.Parse(r)

output := mc.Generate("",100) // Generate text with 100 words. Without defining any starting word.

```
