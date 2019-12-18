package replacer

import (
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// Replacer replaces a list of strings with replacements in markdown text.
type Replacer struct {
	html.Config
	r *strings.Replacer
}

// NewReplacer returns a new Replacer from a list of old, new string pairs.
// Replacements are performed in the order they appear in the target string,
// without overlapping matches. The old string comparisons are done in argument
// order.
//
// NewReplacer panics if given an odd number of arguments.
func NewReplacer(oldnew ...string) *Replacer {
	return &Replacer{
		Config: html.NewConfig(),
		r:      strings.NewReplacer(oldnew...),
	}
}

func (r *Replacer) replace(source []byte) []byte {
	return util.StringToReadOnlyBytes(
		r.r.Replace(util.BytesToReadOnlyString(source)))
}

func (r *Replacer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	var (
		n    = node.(*ast.Text)
		text = r.replace(n.Segment.Value(source))
	)
	if n.IsRaw() {
		r.Writer.RawWrite(w, text)
	} else {
		r.Writer.Write(w, text)
		if n.HardLineBreak() || (n.SoftLineBreak() && r.HardWraps) {
			if r.XHTML {
				_, _ = w.WriteString("<br />\n")
			} else {
				_, _ = w.WriteString("<br>\n")
			}
		} else if n.SoftLineBreak() {
			_ = w.WriteByte('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *Replacer) renderString(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	var (
		n    = node.(*ast.String)
		text = r.replace(n.Value)
	)
	if n.IsCode() {
		_, _ = w.Write(text)
	} else {
		if n.IsRaw() {
			r.Writer.RawWrite(w, text)
		} else {
			r.Writer.Write(w, text)
		}
	}
	return ast.WalkContinue, nil
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs interface.
func (r *Replacer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindText, r.renderText)
	reg.Register(ast.KindString, r.renderString)
}

// Extend implement goldmar.Extender interface.
func (r *Replacer) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(r, 500),
	))
}
