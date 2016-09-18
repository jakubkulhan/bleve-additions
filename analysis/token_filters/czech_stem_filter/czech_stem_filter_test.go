package czech_stem_filter

import (
	"reflect"
	"testing"

	"github.com/blevesearch/bleve/analysis"
)

func TestCzechStemFilter_Filter(t *testing.T) {

	tests := []struct {
		input  analysis.TokenStream
		output analysis.TokenStream
	}{
		{analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("páni")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("pánové")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("pána")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("pánů")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("pánovi")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("pánům")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("pány")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("páne")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("pánech")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("pánem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("pán")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hradu")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hrade")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hradem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hrady")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hradech")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hradům")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hradů")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hrad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("muž")}}, analysis.TokenStream{&analysis.Token{Term: []byte("muh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("muži")}}, analysis.TokenStream{&analysis.Token{Term: []byte("muh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("muže")}}, analysis.TokenStream{&analysis.Token{Term: []byte("muh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mužů")}}, analysis.TokenStream{&analysis.Token{Term: []byte("muh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mužům")}}, analysis.TokenStream{&analysis.Token{Term: []byte("muh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mužích")}}, analysis.TokenStream{&analysis.Token{Term: []byte("muh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mužem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("muh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("stroj")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stroj")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("stroje")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stroj")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("strojů")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stroj")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("stroji")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stroj")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("strojům")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stroj")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("strojích")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stroj")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("strojem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stroj")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předseda")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedové")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedy")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedů")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedovi")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedům")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedu")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedo")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedech")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("předsedou")}}, analysis.TokenStream{&analysis.Token{Term: []byte("předsd")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("soudce")}}, analysis.TokenStream{&analysis.Token{Term: []byte("soudk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("soudci")}}, analysis.TokenStream{&analysis.Token{Term: []byte("soudk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("soudců")}}, analysis.TokenStream{&analysis.Token{Term: []byte("soudk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("soudcům")}}, analysis.TokenStream{&analysis.Token{Term: []byte("soudk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("soudcích")}}, analysis.TokenStream{&analysis.Token{Term: []byte("soudk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("soudcem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("soudk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kost")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kost")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kosti")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kost")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kostí")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kost")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kostem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kost")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kostech")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kost")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kostmi")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kost")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("píseň")}}, analysis.TokenStream{&analysis.Token{Term: []byte("písň")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("písně")}}, analysis.TokenStream{&analysis.Token{Term: []byte("písn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("písni")}}, analysis.TokenStream{&analysis.Token{Term: []byte("písn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("písněmi")}}, analysis.TokenStream{&analysis.Token{Term: []byte("písn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("písních")}}, analysis.TokenStream{&analysis.Token{Term: []byte("písn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("písním")}}, analysis.TokenStream{&analysis.Token{Term: []byte("písn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("růže")}}, analysis.TokenStream{&analysis.Token{Term: []byte("růh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("růží")}}, analysis.TokenStream{&analysis.Token{Term: []byte("růh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("růžím")}}, analysis.TokenStream{&analysis.Token{Term: []byte("růh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("růžích")}}, analysis.TokenStream{&analysis.Token{Term: []byte("růh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("růžemi")}}, analysis.TokenStream{&analysis.Token{Term: []byte("růh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("růži")}}, analysis.TokenStream{&analysis.Token{Term: []byte("růh")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("žena")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ženy")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("žen")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ženě")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ženám")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ženu")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ženo")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ženách")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ženou")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("ženami")}}, analysis.TokenStream{&analysis.Token{Term: []byte("žn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("město")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("města")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("městu")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("městům")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("městě")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("městech")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("městem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("městy")}}, analysis.TokenStream{&analysis.Token{Term: []byte("měst")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("moře")}}, analysis.TokenStream{&analysis.Token{Term: []byte("moř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("moří")}}, analysis.TokenStream{&analysis.Token{Term: []byte("moř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mořím")}}, analysis.TokenStream{&analysis.Token{Term: []byte("moř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("moři")}}, analysis.TokenStream{&analysis.Token{Term: []byte("moř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mořích")}}, analysis.TokenStream{&analysis.Token{Term: []byte("moř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mořem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("moř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuře")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuřata")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuřete")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuřat")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuřeti")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuřatům")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuřatech")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuřetem")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kuřaty")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kuř")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("stavení")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stavn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("stavením")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stavn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("staveních")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stavn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("staveními")}}, analysis.TokenStream{&analysis.Token{Term: []byte("stavn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladý")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladí")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladého")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladých")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladému")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladým")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladé")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladém")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladými")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladá")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mladou")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mlad")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("jarní")}}, analysis.TokenStream{&analysis.Token{Term: []byte("jarn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("jarního")}}, analysis.TokenStream{&analysis.Token{Term: []byte("jarn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("jarních")}}, analysis.TokenStream{&analysis.Token{Term: []byte("jarn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("jarnímu")}}, analysis.TokenStream{&analysis.Token{Term: []byte("jarn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("jarním")}}, analysis.TokenStream{&analysis.Token{Term: []byte("jarn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("jarními")}}, analysis.TokenStream{&analysis.Token{Term: []byte("jarn")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("karlův")}}, analysis.TokenStream{&analysis.Token{Term: []byte("karl")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("jazykový")}}, analysis.TokenStream{&analysis.Token{Term: []byte("jazyk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("český")}}, analysis.TokenStream{&analysis.Token{Term: []byte("česk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("čeští")}}, analysis.TokenStream{&analysis.Token{Term: []byte("česk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("anglický")}}, analysis.TokenStream{&analysis.Token{Term: []byte("anglick")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("angličtí")}}, analysis.TokenStream{&analysis.Token{Term: []byte("anglick")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kniha")}}, analysis.TokenStream{&analysis.Token{Term: []byte("knih")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("knize")}}, analysis.TokenStream{&analysis.Token{Term: []byte("knih")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mazat")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mah")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("mažu")}}, analysis.TokenStream{&analysis.Token{Term: []byte("mah")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kluk")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kluk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("kluci")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kluk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("klucích")}}, analysis.TokenStream{&analysis.Token{Term: []byte("kluk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hezký")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hezk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hezčí")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hezk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hůl")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hol")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("hole")}}, analysis.TokenStream{&analysis.Token{Term: []byte("hol")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("deska")}}, analysis.TokenStream{&analysis.Token{Term: []byte("desk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("desek")}}, analysis.TokenStream{&analysis.Token{Term: []byte("desk")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("e")}}, analysis.TokenStream{&analysis.Token{Term: []byte("e")}}},
		{analysis.TokenStream{&analysis.Token{Term: []byte("zi")}}, analysis.TokenStream{&analysis.Token{Term: []byte("zi")}}},
	}

	for _, test := range tests {
		filter := NewCzechStemFilter()
		actual := filter.Filter(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected %s, got %s", test.output, actual)
		}
	}
}

type endsWithTestCase struct {
	term     []rune
	suffix   []rune
	expected bool
}

func TestEndsWith(t *testing.T) {
	testCases := []endsWithTestCase{
		{[]rune(""), []rune(""), true},
		{[]rune(""), []rune("a"), false},
		{[]rune("a"), []rune(""), true},
		{[]rune("ab"), []rune("a"), false},
		{[]rune("ab"), []rune("b"), true},
		{[]rune("ab"), []rune("ab"), true},
		{[]rune("ab"), []rune("abcd"), false},
		{[]rune("žůžo"), []rune("žo"), true},
	}

	for _, testCase := range testCases {
		got := endsWith(testCase.term, testCase.suffix)
		if got != testCase.expected {
			t.Errorf("endsWith(%q, %q) expected %v, got %v", string(testCase.term), string(testCase.suffix), testCase.expected, got)
		}
	}
}
