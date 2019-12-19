# goldmark-text-replacer

[![GoDoc](https://godoc.org/github.com/mdigger/goldmark-text-replacer?status.svg)](https://godoc.org/github.com/mdigger/goldmark-text-replacer)

[Goldmark](https://github.com/yuin/goldmark) text replacer extension.

```go
md := goldmark.New(
    replacer.Options(
        "(c)", "&copy;",
        "(r)", "&reg;",
        "...", "&hellip;",
        "(tm)", "&trade;",
        "<-", "&larr;",
        "->", "&rarr;",
        "<->", "&harr;",
    ),
)
var source = []byte("(c)Dmitry Sedykh")
err := md.Convert(source, os.Stdout)
if err != nil {
    log.Fatal(err)
}
```
