package replacer

import (
	"log"
	"os"

	"github.com/yuin/goldmark"
)

func Example() {
	repl := NewReplacer(
		"(c)", "&copy;",
		"(r)", "&reg;",
		"...", "&hellip;",
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
	// output: <p>© Dmitry Sedykh</p>
}
