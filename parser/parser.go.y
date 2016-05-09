%{
package parser

import(
        "github.com/soh335/mtexport/ast"
)
%}

%union {
        token  ast.Token
        expr   ast.Expr
        stmt   ast.Stmt
        stmts  []ast.Stmt
};

%type<expr> line multiline_body
%type<stmt> entry section multiline_section field
%type<stmts> entries sections fields

%token<token> IDENTIFIER STRING VALUE END_OF_SECTION END_OF_ENTRY FIELD_KEY NL

%%

entries:
        entry END_OF_ENTRY NL
        {
                $$ = []ast.Stmt{$1}
                yylex.(*Lexer).stmts = $$
        }
        | entry END_OF_ENTRY NL entries
        {
                $$ = append([]ast.Stmt{$1}, $4...)
                yylex.(*Lexer).stmts = $$
        }

entry:
        sections
        {
                $$ = &ast.EntryStmt{
                        SectionStmts: $1,
                }
        }

sections:
        section END_OF_SECTION NL
        {
                $$ = []ast.Stmt{$1}
        }
        | section END_OF_SECTION NL sections
        {
                $$ = append([]ast.Stmt{$1}, $4...)
        }

section: fields
       {
                $$ = &ast.NormalSectionStmt{
                        FieldStmts: $1,
                }
       }
       | multiline_section
       {
                $$ = $1
       }

multiline_section:
        FIELD_KEY ':' NL fields multiline_body
        {
                $$ = &ast.MultilineSectionStmt{
                        Key: $1.Literal,
                        FieldStmts: $4,
                        Body: string($5.(ast.StringExpr)),
                }
        }
        | FIELD_KEY ':' NL multiline_body
        {
                $$ = &ast.MultilineSectionStmt{
                        Key: $1.Literal,
                        Body: string($4.(ast.StringExpr)),
                }
        }

multiline_body: line
        {
                $$ = $1
        }
        | line multiline_body
        {
                $$ = ast.StringExpr( string($1.(ast.StringExpr)) + string($2.(ast.StringExpr)) )
        }

fields:
      field
      {
                $$ = []ast.Stmt{$1}
      }
      | field fields
      {
                $$ = append([]ast.Stmt{$1}, $2...)
      }

field: FIELD_KEY ':' VALUE NL
      {
                $$ = &ast.FieldStmt{Key: $1.Literal, Value: $3.Literal}
                yylex.(*Lexer).line = $1.Line
      }


line: STRING NL
        {
                $$ = ast.StringExpr( $1.Literal + $2.Literal )
                yylex.(*Lexer).line = $1.Line
        }
        | NL
        {
                $$ = ast.StringExpr( $1.Literal )
                yylex.(*Lexer).line = $1.Line
        }

%%
