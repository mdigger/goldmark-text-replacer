package replacer_test

import (
	"log"
	"os"

	replacer "github.com/mdigger/goldmark-text-replacer"
	"github.com/yuin/goldmark"
)

func Example() {
	md := goldmark.New(
		replacer.Options(
			"(c)", "&copy;",
			"(r)", "&reg;",
			"...", "&hellip;",
			"(tm)", "&trade;",
			"<-", "&larr;",
			"->", "&rarr;",
			"<->", "&harr;",
			"--", "&mdash;",
		),
	)
	var source = []byte("(c)Dmitry Sedykh")
	err := md.Convert(source, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	// output: <p>©Dmitry Sedykh</p>
}
