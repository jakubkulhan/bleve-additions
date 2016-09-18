package ascii_folding_filter

import (
	"unicode"

	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
	"github.com/jakubkulhan/bleve-additions/analysis/token_filters/transformer_filter"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	Name = "asciifolding"
)

func NewASCIIFoldingFilter() *transformer_filter.TransformerFilter {
	return transformer_filter.NewTransformerFilter(transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC))
}

func ASCIIFoldingFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	return NewASCIIFoldingFilter(), nil
}

func init() {
	registry.RegisterTokenFilter(Name, ASCIIFoldingFilterConstructor)
}
