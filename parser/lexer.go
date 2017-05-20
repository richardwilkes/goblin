package parser

//go:generate goyacc -l -v /dev/null -o parser.go parser.go.y

import (
	"github.com/richardwilkes/goblin/ast"
)

// Lexer provides scanning for tokens.
type Lexer struct {
	scanner *Scanner
	lit     string
	pos     ast.Position
	stmts   []ast.Stmt
	err     error
}

// Parse the source into statements.
func Parse(src string) ([]ast.Stmt, error) {
	lexer := Lexer{scanner: &Scanner{src: []rune(src)}}
	yyErrorVerbose = true
	if yyParse(&lexer) != 0 {
		return nil, lexer.err
	}
	return lexer.stmts, lexer.err
}

// Lex scans the token and literals.
func (lexer *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos, err := lexer.scanner.Scan()
	if err != nil {
		lexer.err = &ast.Error{Message: err.Error(), Pos: pos, Fatal: true}
	}
	lval.tok = ast.Token{Tok: tok, Lit: lit}
	lval.tok.SetPosition(pos)
	lexer.lit = lit
	lexer.pos = pos
	return tok
}

// Error sets the parse error.
func (lexer *Lexer) Error(msg string) {
	lexer.err = &ast.Error{Message: msg, Pos: lexer.pos}
}
