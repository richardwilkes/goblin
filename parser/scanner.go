// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package parser

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/richardwilkes/goblin/ast"
)

const (
	// EOF - end of file
	EOF = -1
	// EOL - end of line
	EOL = '\n'
)

var opName = map[string]int{
	"func":     FUNC,
	"return":   RETURN,
	"var":      VAR,
	"throw":    THROW,
	"if":       IF,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"in":       IN,
	"else":     ELSE,
	"new":      NEW,
	"true":     TRUE,
	"false":    FALSE,
	"nil":      NIL,
	"module":   MODULE,
	"try":      TRY,
	"catch":    CATCH,
	"finally":  FINALLY,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"make":     MAKE,
}

// Scanner holds informations for the lexer.
type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

// Init resets code to scan.
func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

// Scan for the next token.
func (s *Scanner) Scan() (tok int, lit string, pos ast.Position, err error) {
retry:
	s.skipBlank()
	pos = s.pos()
	switch ch := s.peek(); {
	case isLetter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			return
		}
		if name, ok := opName[lit]; ok {
			tok = name
		} else {
			tok = IDENT
		}
	case isDigit(ch):
		tok = NUMBER
		lit, err = s.scanNumber()
		if err != nil {
			return
		}
	case ch == '"':
		tok = STRING
		lit, err = s.scanString('"')
		if err != nil {
			return
		}
	case ch == '\'':
		tok = STRING
		lit, err = s.scanString('\'')
		if err != nil {
			return
		}
	case ch == '`':
		tok = STRING
		lit, err = s.scanRawString()
		if err != nil {
			return
		}
	default:
		switch ch {
		case EOF:
			tok = EOF
		case '#':
			for !isEOL(s.peek()) {
				s.next()
			}
			goto retry
		case '!':
			s.next()
			switch s.peek() {
			case '=':
				tok = NEQ
				lit = "!="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '=':
			s.next()
			switch s.peek() {
			case '=':
				tok = EQEQ
				lit = "=="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '+':
			s.next()
			switch s.peek() {
			case '+':
				tok = PLUSPLUS
				lit = "++"
			case '=':
				tok = PLUSEQ
				lit = "+="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '-':
			s.next()
			switch s.peek() {
			case '-':
				tok = MINUSMINUS
				lit = "--"
			case '=':
				tok = MINUSEQ
				lit = "-="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '*':
			s.next()
			switch s.peek() {
			case '*':
				tok = POW
				lit = "**"
			case '=':
				tok = MULEQ
				lit = "*="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '/':
			s.next()
			switch s.peek() {
			case '/':
				s.next()
				for !isEOL(s.peek()) {
					s.next()
				}
				goto retry
			case '*':
				s.next()
				for {
					ch = s.peek()
					if ch == EOF {
						goto retry
					}
					s.next()
					if ch == '*' && s.peek() == '/' {
						s.next()
						goto retry
					}
				}
			case '=':
				tok = DIVEQ
				lit = "/="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '>':
			s.next()
			switch s.peek() {
			case '=':
				tok = GE
				lit = ">="
			case '>':
				tok = SHIFTRIGHT
				lit = ">>"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '<':
			s.next()
			switch s.peek() {
			case '=':
				tok = LE
				lit = "<="
			case '<':
				tok = SHIFTLEFT
				lit = "<<"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '|':
			s.next()
			switch s.peek() {
			case '|':
				tok = OROR
				lit = "||"
			case '=':
				tok = OREQ
				lit = "|="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '&':
			s.next()
			switch s.peek() {
			case '&':
				tok = ANDAND
				lit = "&&"
			case '=':
				tok = ANDEQ
				lit = "&="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '.':
			s.next()
			if s.peek() == '.' {
				s.next()
				if s.peek() == '.' {
					tok = VARARG
				} else {
					err = fmt.Errorf(`syntax error "%s"`, "..")
					return
				}
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '\n':
			tok = int(ch)
			lit = string(ch)
		case '(', ')', ':', ';', '%', '?', '{', '}', ',', '[', ']', '^':
			s.next()
			if ch == '[' && s.peek() == ']' {
				s.next()
				if isLetter(s.peek()) {
					s.back()
					tok = ARRAYLIT
					lit = "[]"
				} else {
					s.back()
					s.back()
					tok = int(ch)
					lit = string(ch)
				}
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		default:
			err = fmt.Errorf(`syntax error "%s"`, string(ch))
			tok = int(ch)
			lit = string(ch)
			return
		}
		s.next()
	}
	return
}

func (s *Scanner) peek() rune {
	if s.reachEOF() {
		return EOF
	}
	return s.src[s.offset]
}

func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

func (s *Scanner) back() {
	s.offset--
}

func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

func (s *Scanner) pos() ast.Position {
	return ast.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *Scanner) skipBlank() {
	for isBlank(s.peek()) {
		s.next()
	}
}

func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for isLetter(s.peek()) || isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

func (s *Scanner) scanNumber() (string, error) {
	var ret []rune
	ch := s.peek()
	ret = append(ret, ch)
	s.next()
	if ch == '0' && s.peek() == 'x' {
		ret = append(ret, s.peek())
		s.next()
		for isHex(s.peek()) {
			ret = append(ret, s.peek())
			s.next()
		}
	} else {
		for isDigit(s.peek()) || s.peek() == '.' {
			ret = append(ret, s.peek())
			s.next()
		}
		if s.peek() == 'e' {
			ret = append(ret, s.peek())
			s.next()
			if isDigit(s.peek()) || s.peek() == '+' || s.peek() == '-' {
				ret = append(ret, s.peek())
				s.next()
				for isDigit(s.peek()) || s.peek() == '.' {
					ret = append(ret, s.peek())
					s.next()
				}
			}
			for isDigit(s.peek()) || s.peek() == '.' {
				ret = append(ret, s.peek())
				s.next()
			}
		}
		if isLetter(s.peek()) {
			return "", errors.New("identifier starts immediately after numeric literal")
		}
	}
	return string(ret), nil
}

func (s *Scanner) scanRawString() (string, error) {
	var ret []rune
	for {
		s.next()
		if s.peek() == EOF {
			return "", errors.New("unexpected EOF")
		}
		if s.peek() == '`' {
			s.next()
			break
		}
		ret = append(ret, s.peek())
	}
	return string(ret), nil
}

func (s *Scanner) scanString(l rune) (string, error) {
	var ret []rune
eos:
	for {
		s.next()
		switch s.peek() {
		case EOL:
			return "", errors.New("unexpected EOL")
		case EOF:
			return "", errors.New("unexpected EOF")
		case l:
			s.next()
			break eos
		case '\\':
			s.next()
			switch s.peek() {
			case 'b':
				ret = append(ret, '\b')
				continue
			case 'f':
				ret = append(ret, '\f')
				continue
			case 'r':
				ret = append(ret, '\r')
				continue
			case 'n':
				ret = append(ret, '\n')
				continue
			case 't':
				ret = append(ret, '\t')
				continue
			}
			ret = append(ret, s.peek())
			continue
		default:
			ret = append(ret, s.peek())
		}
	}
	return string(ret), nil
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isHex(ch rune) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}

func isEOL(ch rune) bool {
	return ch == EOL || ch == EOF
}

func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}
