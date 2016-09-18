# [Bleve](https://github.com/blevesearch/bleve) additions

> Some useful tools for Bleve, modern text indexing for go

## `asciifolding` token filter

- package `github.com/jakubkulhan/bleve-additions/analysis/token_filters/ascii_folding_filter`
- removes diacritics from utf8 terms
- based on `transformer_filter`

## `czech_stem` token filter

- package `github.com/jakubkulhan/bleve-additions/analysis/token_filters/czech_stem_filter`
- based on [`CzechStemmer`](https://lucene.apache.org/core/6_1_0/analyzers-common/org/apache/lucene/analysis/cz/CzechStemmer.html) from Lucene

## `golang.org/x/text/transform.Transfomer` token filter

- package `github.com/jakubkulhan/bleve-additions/analysis/token_filters/transformer_filter`
- allows to apply custom `golang.org/x/text/transform.Transfomer` to tokens
