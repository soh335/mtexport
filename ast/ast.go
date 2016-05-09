package ast

type Token struct {
	Token   int
	Literal string
	Line    int
}

type StringExpr string

type EntryStmt struct {
	SectionStmts []Stmt
}

type MultilineSectionStmt struct {
	Key        string
	FieldStmts []Stmt
	Body       string
}

type NormalSectionStmt struct {
	FieldStmts []Stmt
}

type FieldStmt struct {
	Key   string
	Value string
}

type Expr interface{}

type Stmt interface{}
