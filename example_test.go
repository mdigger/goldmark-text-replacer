package replacer_test

import (
	"log"
	"os"

	replacer "github.com/mdigger/goldmark-text-replacer"
	"github.com/yuin/goldmark"
)

func Example() {
	repl := replacer.New(
		"(c)", "&copy;",
		"(r)", "&reg;",
		"...", "&hellip;",
		":)", "&#9786;",
	)
	md := goldmark.New(
		goldmark.WithExtensions(repl),
	)
	var source = []byte("(c)Dmitry Sedykh")
	err := md.Convert(source, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	// output: <p>©Dmitry Sedykh</p>
}
