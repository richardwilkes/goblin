package interpreter

// Token holds a token.
type Token struct {
	PosImpl
	Tok int
	Lit string
}
