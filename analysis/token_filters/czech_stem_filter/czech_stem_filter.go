package czech_stem_filter

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
)

const (
	Name = "czech_stem"
)

type CzechStemFilter struct {
}

func NewCzechStemFilter() *CzechStemFilter {
	return &CzechStemFilter{}
}

// see https://github.com/apache/lucene-solr/blob/master/lucene/analysis/common/src/java/org/apache/lucene/analysis/cz/CzechStemmer.java
// and also see http://portal.acm.org/citation.cfm?id=1598600
func (f *CzechStemFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	for _, token := range input {
		if token.Type != analysis.AlphaNumeric {
			continue
		}

		term := []rune(string(token.Term))
		term = removeCase(term)
		term = removePossesives(term)
		if len(token.Term) > 0 {
			term = normalize(term)
		}
		token.Term = []byte(string(term))
	}

	return input
}

func removeCase(term []rune) []rune {
	if len(term) > 7 && endsWith(term, []rune("atech")) {
		return term[:len(term)-5]
	}

	if len(term) > 6 &&
		(endsWith(term, []rune("ětem")) ||
			endsWith(term, []rune("etem")) ||
			endsWith(term, []rune("atům"))) {
		return term[:len(term)-4]
	}

	if len(term) > 5 &&
		(endsWith(term, []rune("ech")) ||
			endsWith(term, []rune("ich")) ||
			endsWith(term, []rune("ích")) ||
			endsWith(term, []rune("ého")) ||
			endsWith(term, []rune("ěmi")) ||
			endsWith(term, []rune("emi")) ||
			endsWith(term, []rune("ému")) ||
			endsWith(term, []rune("ěte")) ||
			endsWith(term, []rune("ete")) ||
			endsWith(term, []rune("ěti")) ||
			endsWith(term, []rune("eti")) ||
			endsWith(term, []rune("ího")) ||
			endsWith(term, []rune("iho")) ||
			endsWith(term, []rune("ími")) ||
			endsWith(term, []rune("ímu")) ||
			endsWith(term, []rune("imu")) ||
			endsWith(term, []rune("ách")) ||
			endsWith(term, []rune("ata")) ||
			endsWith(term, []rune("aty")) ||
			endsWith(term, []rune("ých")) ||
			endsWith(term, []rune("ama")) ||
			endsWith(term, []rune("ami")) ||
			endsWith(term, []rune("ové")) ||
			endsWith(term, []rune("ovi")) ||
			endsWith(term, []rune("ými"))) {
		return term[:len(term)-3]
	}

	if len(term) > 4 &&
		(endsWith(term, []rune("em")) ||
			endsWith(term, []rune("es")) ||
			endsWith(term, []rune("ém")) ||
			endsWith(term, []rune("ím")) ||
			endsWith(term, []rune("ům")) ||
			endsWith(term, []rune("at")) ||
			endsWith(term, []rune("ám")) ||
			endsWith(term, []rune("os")) ||
			endsWith(term, []rune("us")) ||
			endsWith(term, []rune("ým")) ||
			endsWith(term, []rune("mi")) ||
			endsWith(term, []rune("ou"))) {
		return term[:len(term)-2]
	}

	if len(term) > 3 &&
		(endsWith(term, []rune("a")) ||
			endsWith(term, []rune("e")) ||
			endsWith(term, []rune("i")) ||
			endsWith(term, []rune("o")) ||
			endsWith(term, []rune("u")) ||
			endsWith(term, []rune("ů")) ||
			endsWith(term, []rune("y")) ||
			endsWith(term, []rune("á")) ||
			endsWith(term, []rune("é")) ||
			endsWith(term, []rune("í")) ||
			endsWith(term, []rune("ý")) ||
			endsWith(term, []rune("ě"))) {

		return term[:len(term)-1]
	}

	return term
}

func removePossesives(term []rune) []rune {
	if len(term) > 5 &&
		(endsWith(term, []rune("ov")) ||
			endsWith(term, []rune("in")) ||
			endsWith(term, []rune("ův"))) {
		return term[:len(term)-2]
	}

	return term
}

func normalize(term []rune) []rune {
	if endsWith(term, []rune("čt")) {
		term[len(term)-2] = 'c'
		term[len(term)-1] = 'k'
		return term
	}

	if endsWith(term, []rune("št")) {
		term[len(term)-2] = 's'
		term[len(term)-1] = 'k'
		return term
	}

	if endsWith(term, []rune("c")) || endsWith(term, []rune("č")) {
		term[len(term)-1] = 'k'
		return term
	}

	if endsWith(term, []rune("z")) || endsWith(term, []rune("ž")) {
		term[len(term)-1] = 'h'
		return term
	}

	if len(term) > 1 && term[len(term)-2] == 'e' {
		term[len(term)-2] = term[len(term)-1]
		return term[:len(term)-1]
	}

	if len(term) > 1 && term[len(term)-2] == 'ů' {
		term[len(term)-2] = 'o'
		return term
	}

	return term
}

func endsWith(term []rune, suffix []rune) bool {
	if len(term) < len(suffix) {
		return false
	}

	for i, j, l := len(term)-len(suffix), 0, len(suffix); j < l; i, j = i+1, j+1 {
		if term[i] != suffix[j] {
			return false
		}
	}

	return true
}

func CzechStemFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	return NewCzechStemFilter(), nil
}

func init() {
	registry.RegisterTokenFilter(Name, CzechStemFilterConstructor)
}
