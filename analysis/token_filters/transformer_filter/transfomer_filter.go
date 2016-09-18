package transformer_filter

import (
	"github.com/blevesearch/bleve/analysis"
	"golang.org/x/text/transform"
)

type TransformerFilter struct {
	transformer transform.Transformer
}

func NewTransformerFilter(transformer transform.Transformer) *TransformerFilter {
	return &TransformerFilter{transformer}
}

func (f *TransformerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	for _, token := range input {
		f.transformer.Reset()
		if result, _, err := transform.Bytes(f.transformer, token.Term); err == nil {
			token.Term = result
		}
	}
	return input
}
