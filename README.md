# mctext

Random-text generator using Markov chain from an user-provided sample. It's useful in cases where generated text must be relatively random,
but expected to contain a good average distribution and repetition of words.

Example:

```go

import "github.com/zgiber/mctext"

...

// provide an io.Reader with the sample. Ideally UTF encoded text.
// ...
var r io.Reader

...
  mc := mctext.New()
  mc.Parse(r)

  output := mc.Generate("",100) // Generate text with 100 words without defining any starting word.

```

If you don't want to provide your own io.Reader you can use the sample in the library:

```go
...
  mc := mctext.New()
  mc.Parse(Sample) // The default sample is Tom Sawyer from Mark Twain :)
...
```
