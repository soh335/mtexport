package parser

//go:generate go tool yacc -o parser.go parser.go.y

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/soh335/mtexport/ast"
)

type ErrorType int

const (
	EOF            ErrorType = -1
	UNEXPECTED_ERR           = -2
)

type Error struct {
	typ  ErrorType
	err  error
	line int
}

func (e *Error) Error() string {
	return e.err.Error()
}

var fieldKeys = []string{
	"AUTHOR",
	"TITLE",
	"BASENAME",
	"STATUS",
	"ALLOW COMMENTS",
	"ALLOW PINGS",
	"CONVERT BREAKS",
	"PRIMARY CATEGORY",
	"CATEGORY",
	"DATE",
	"TAGS",
	"NO ENTRY",

	"BODY",
	"EXTENDED BODY",
	"EXCERPT",
	"KEYWORDS",
	"COMMENT",

	"EMAIL",
	"URL",
	"IP",

	"PING",

	"BLOG NAME",
}

var tokenName = map[int]string{
	END_OF_ENTRY:   "--------",
	END_OF_SECTION: "-----",
	NL:             "\n",
}

type Item interface{}

type Scanner struct {
	r        *bufio.Reader
	itemChan chan Item
	line     int

	fieldKeyMap map[string]bool
}

func NewScanner(r io.Reader, extraFieldKeys []string) *Scanner {

	fieldKeyMap := map[string]bool{}

	for _, fieldKey := range fieldKeys {
		fieldKeyMap[fieldKey] = true
	}

	for _, fieldKey := range extraFieldKeys {
		fieldKeyMap[fieldKey] = true
	}

	return &Scanner{
		r:           bufio.NewReader(r),
		itemChan:    make(chan Item, 1),
		fieldKeyMap: fieldKeyMap,
	}
}

func (s *Scanner) Scan() {
	for {
		line, err := s.r.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		s.line += 1

		if err != nil {
			if err == io.EOF {
				s.itemChan <- &Error{EOF, nil, s.line}
			} else {
				s.itemChan <- &Error{UNEXPECTED_ERR, err, s.line}
			}
			return
		}

		switch line {
		case tokenName[END_OF_ENTRY]:
			s.itemChan <- ast.Token{END_OF_ENTRY, line, s.line}
		case tokenName[END_OF_SECTION]:
			s.itemChan <- ast.Token{END_OF_SECTION, line, s.line}
		default:
			strs := strings.SplitN(line, ":", 2)
			if len(strs) >= 1 && s.fieldKeyMap[strs[0]] {
				s.itemChan <- ast.Token{FIELD_KEY, strs[0], s.line}
				s.itemChan <- ast.Token{':', ":", s.line}
				if len(strs) == 2 && strs[1] != "" {
					s.itemChan <- ast.Token{VALUE, strs[1], s.line}
				}
			} else {
				s.itemChan <- ast.Token{STRING, line, s.line}
			}
		}
		s.itemChan <- ast.Token{NL, tokenName[NL], s.line}
	}
}

type Lexer struct {
	s     *Scanner
	e     error
	stmts []ast.Stmt
	line  int
}

func NewLexer(r io.Reader, extraKeywords []string) *Lexer {
	l := &Lexer{}
	l.s = NewScanner(r, extraKeywords)
	go l.s.Scan()
	return l
}

func (l *Lexer) Lex(lval *yySymType) int {
	switch item := <-l.s.itemChan; item.(type) {
	case ast.Token:
		lval.token = item.(ast.Token)
		return lval.token.Token
	case *Error:
		e := item.(*Error)
		switch e.typ {
		case EOF:
			return 0
		case UNEXPECTED_ERR:
			l.e = e
			return 0
		default:
			panic("not reach")
		}
	default:
		panic("not reach")
	}
}

func (l *Lexer) Error(e string) {
	l.e = &Error{UNEXPECTED_ERR, fmt.Errorf("%s", e), l.line}
}

func Parse(r io.Reader, extraKeywords []string) ([]ast.Stmt, error) {
	l := NewLexer(r, extraKeywords)
	if yyParse(l) != 0 {
		return nil, l.e
	}
	return l.stmts, nil
}
