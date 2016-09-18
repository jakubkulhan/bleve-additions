package ascii_folding_filter

import (
	"testing"
	"github.com/blevesearch/bleve/analysis"
	"reflect"
)

func TestASCIIFoldingFilter_Filter(t *testing.T) {

	tests := []struct {
		input  analysis.TokenStream
		output analysis.TokenStream
	}{
		{analysis.TokenStream{&analysis.Token{Term: []byte("ě")}}, analysis.TokenStream{&analysis.Token{Term: []byte("e")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("š")}}, analysis.TokenStream{&analysis.Token{Term: []byte("s")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("č")}}, analysis.TokenStream{&analysis.Token{Term: []byte("c")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ř")}}, analysis.TokenStream{&analysis.Token{Term: []byte("r")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ž")}}, analysis.TokenStream{&analysis.Token{Term: []byte("z")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ý")}}, analysis.TokenStream{&analysis.Token{Term: []byte("y")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("á")}}, analysis.TokenStream{&analysis.Token{Term: []byte("a")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("í")}}, analysis.TokenStream{&analysis.Token{Term: []byte("i")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("é")}}, analysis.TokenStream{&analysis.Token{Term: []byte("e")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ú")}}, analysis.TokenStream{&analysis.Token{Term: []byte("u")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ů")}}, analysis.TokenStream{&analysis.Token{Term: []byte("u")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ď")}}, analysis.TokenStream{&analysis.Token{Term: []byte("d")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ť")}}, analysis.TokenStream{&analysis.Token{Term: []byte("t")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ň")}}, analysis.TokenStream{&analysis.Token{Term: []byte("n")}}},
	}

	for _, test := range tests {
		filter := NewASCIIFoldingFilter()
		actual := filter.Filter(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected %s, got %s", test.output, actual)
		}
	}
}
