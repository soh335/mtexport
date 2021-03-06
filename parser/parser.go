//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/soh335/mtexport/ast"
)

//line parser.go.y:9
type yySymType struct {
	yys   int
	token ast.Token
	expr  ast.Expr
	stmt  ast.Stmt
	stmts []ast.Stmt
}

const IDENTIFIER = 57346
const STRING = 57347
const VALUE = 57348
const END_OF_SECTION = 57349
const END_OF_ENTRY = 57350
const FIELD_KEY = 57351
const NL = 57352

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"STRING",
	"VALUE",
	"END_OF_SECTION",
	"END_OF_ENTRY",
	"FIELD_KEY",
	"NL",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:119

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 17
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 33

var yyAct = [...]int{

	22, 5, 3, 24, 16, 1, 24, 12, 25, 11,
	18, 25, 8, 13, 17, 29, 26, 15, 20, 21,
	19, 14, 27, 12, 28, 9, 10, 18, 7, 6,
	4, 2, 23,
}
var yyPact = [...]int{

	3, -1000, 17, -1000, 19, -1000, -1000, 14, 2, 11,
	7, -1000, -7, 4, 3, 3, 21, -2, 6, -1000,
	-1000, 1, -1000, 1, 5, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 32, 0, 31, 30, 29, 28, 5, 2, 1,
}
var yyR1 = [...]int{

	0, 7, 7, 3, 8, 8, 4, 4, 5, 5,
	2, 2, 9, 9, 6, 1, 1,
}
var yyR2 = [...]int{

	0, 3, 4, 1, 3, 4, 1, 1, 5, 4,
	1, 2, 1, 2, 4, 2, 1,
}
var yyChk = [...]int{

	-1000, -7, -3, -8, -4, -9, -5, -6, 9, 8,
	7, -9, 9, 11, 10, 10, 11, 10, 6, -7,
	-8, -9, -2, -1, 5, 10, 10, -2, -2, 10,
}
var yyDef = [...]int{

	0, -2, 0, 3, 0, 6, 7, 12, 0, 0,
	0, 13, 0, 0, 1, 4, 0, 0, 0, 2,
	5, 0, 9, 10, 0, 16, 14, 8, 11, 15,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 11,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:26
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[1].stmt}
			yylex.(*Lexer).stmts = yyVAL.stmts
		}
	case 2:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:31
		{
			yyVAL.stmts = append([]ast.Stmt{yyDollar[1].stmt}, yyDollar[4].stmts...)
			yylex.(*Lexer).stmts = yyVAL.stmts
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:38
		{
			yyVAL.stmt = &ast.EntryStmt{
				SectionStmts: yyDollar[1].stmts,
			}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:46
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[1].stmt}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:50
		{
			yyVAL.stmts = append([]ast.Stmt{yyDollar[1].stmt}, yyDollar[4].stmts...)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:55
		{
			yyVAL.stmt = &ast.NormalSectionStmt{
				FieldStmts: yyDollar[1].stmts,
			}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:61
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:67
		{
			yyVAL.stmt = &ast.MultilineSectionStmt{
				Key:        yyDollar[1].token.Literal,
				FieldStmts: yyDollar[4].stmts,
				Body:       string(yyDollar[5].expr.(ast.StringExpr)),
			}
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:75
		{
			yyVAL.stmt = &ast.MultilineSectionStmt{
				Key:  yyDollar[1].token.Literal,
				Body: string(yyDollar[4].expr.(ast.StringExpr)),
			}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:83
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:87
		{
			yyVAL.expr = ast.StringExpr(string(yyDollar[1].expr.(ast.StringExpr)) + string(yyDollar[2].expr.(ast.StringExpr)))
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:93
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[1].stmt}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.stmts = append([]ast.Stmt{yyDollar[1].stmt}, yyDollar[2].stmts...)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.stmt = &ast.FieldStmt{Key: yyDollar[1].token.Literal, Value: yyDollar[3].token.Literal}
			yylex.(*Lexer).line = yyDollar[1].token.Line
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.expr = ast.StringExpr(yyDollar[1].token.Literal + yyDollar[2].token.Literal)
			yylex.(*Lexer).line = yyDollar[1].token.Line
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.expr = ast.StringExpr(yyDollar[1].token.Literal)
			yylex.(*Lexer).line = yyDollar[1].token.Line
		}
	}
	goto yystack /* stack new state and value */
}
