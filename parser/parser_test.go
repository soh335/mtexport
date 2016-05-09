package parser

import (
	"os"
	"testing"
)

func TestParser1(t *testing.T) {
	// https://movabletype.org/documentation/appendices/import-export-format.html
	f, err := os.Open("./_testdata/1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	lexer := NewLexer(f, nil)
	if yyParse(lexer) != 0 {
		t.Fatal(lexer.e, lexer.line)
	}
}

func TestParserInvalidField(t *testing.T) {
	yyErrorVerbose = true
	// https://movabletype.org/documentation/appendices/import-export-format.html
	f, err := os.Open("./_testdata/invalid_field.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	lexer := NewLexer(f, nil)
	if yyParse(lexer) == 0 {
		t.Fatal("should error")
	} else {
		t.Log(lexer.e, lexer.line)
	}
}

func TestParserExtraFiekdKey(t *testing.T) {
	yyErrorVerbose = true
	// https://movabletype.org/documentation/appendices/import-export-format.html
	f, err := os.Open("./_testdata/extra_field_key.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	lexer := NewLexer(f, []string{"IMAGE"})
	if yyParse(lexer) != 0 {
		t.Fatal(lexer.e, lexer.line)
	}
}
