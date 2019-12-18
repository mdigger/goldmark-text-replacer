# goldmark-text-replacer

[![GoDoc](https://godoc.org/github.com/mdigger/goldmark-text-replacer?status.svg)](https://godoc.org/github.com/mdigger/goldmark-text-replacer)

[Goldmark](https://github.com/yuin/goldmark) text replacer extension.

```go
repl := repl.NewReplacer(
    "(c)", "&copy;",
    "(r)", "&reg;",
    "...", "&helip;",
    ":)", "&#9786;",
)
md := goldmark.New(
    goldmark.WithExtensions(repl),
)
var source = []byte("(c) Dmitry Sedykh")
err := md.Convert(source, os.Stdout)
if err != nil {
    log.Fatal(err)
}
```
