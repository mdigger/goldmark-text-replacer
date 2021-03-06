// Package replacer is a extension for the goldmark
// (http://github.com/yuin/goldmark).
//
// This extension adds support for authomaticaly replacing text in markdowns.
package replacer

import (
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// replacer replaces a list of strings with replacements in markdown text.
type replacer struct {
	html.Config
	*strings.Replacer
}

// New returns a new Replacer from a list of old, new string pairs.
// Replacements are performed in the order they appear in the target string,
// without overlapping matches. The old string comparisons are done in argument
// order.
//
// It's panics if given an odd number of arguments.
func New(oldnew ...string) goldmark.Extender {
	return &replacer{
		Config:   html.NewConfig(),
		Replacer: strings.NewReplacer(oldnew...),
	}
}

func (r *replacer) replace(source []byte) []byte {
	return util.StringToReadOnlyBytes(
		r.Replacer.Replace(util.BytesToReadOnlyString(source)))
}

func (r *replacer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	var (
		n    = node.(*ast.Text)
		text = r.replace(n.Text(source))
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

func (r *replacer) renderString(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
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
func (r *replacer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindText, r.renderText)
	reg.Register(ast.KindString, r.renderString)
}

// Extend implement goldmark.Extender interface.
func (r *replacer) Extend(m goldmark.Markdown) {
	if r.Replacer == nil {
		return
	}
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(r, 500),
	))
}

// Options return initialized text replacer goldmark.Option.
func Options(oldnew ...string) goldmark.Option {
	return goldmark.WithExtensions(New(oldnew...))
}
