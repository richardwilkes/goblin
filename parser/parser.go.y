%{
package parser

import (
	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/statement"
)

%}

%type<compstmt> compstmt
%type<stmts> stmts
%type<stmt> stmt
%type<stmtIf> stmtIf
%type<stmtDefault> stmtDefault
%type<stmtCase> stmtCase
%type<stmtCases> stmtCases
%type<typ> typ
%type<expr> expr
%type<exprs> exprs
%type<exprMany> exprMany
%type<exprLets> exprLets
%type<exprPair> exprPair
%type<exprPairs> exprPairs
%type<exprIdents> exprIdents

%union{
	compstmt    []goblin.Stmt
	stmtIf      goblin.Stmt
	stmtDefault goblin.Stmt
	stmtCase    goblin.Stmt
	stmtCases   []goblin.Stmt
	stmts       []goblin.Stmt
	stmt        goblin.Stmt
	typ         goblin.Type
	expr        goblin.Expr
	exprs       []goblin.Expr
	exprMany    []goblin.Expr
	exprLets    goblin.Expr
	exprPair    goblin.Expr
	exprPairs   []goblin.Expr
	exprIdents  []string
	tok         goblin.Token
	term        goblin.Token
	terms       goblin.Token
	optTerms    goblin.Token
}

%token<tok> IDENT NUMBER STRING ARRAY VARARG FUNC RETURN VAR THROW IF ELSE FOR IN EQEQ NEQ GE LE OROR ANDAND NEW TRUE FALSE NIL MODULE TRY CATCH FINALLY PLUSEQ MINUSEQ MULEQ DIVEQ ANDEQ OREQ BREAK CONTINUE PLUSPLUS MINUSMINUS POW SHIFTLEFT SHIFTRIGHT SWITCH CASE DEFAULT MAKE ARRAYLIT

%right '='
%right '?' ':'
%left OROR
%left ANDAND
%left IDENT
%nonassoc EQEQ NEQ ','
%left '>' GE '<' LE SHIFTLEFT SHIFTRIGHT

%left '+' '-' PLUSPLUS MINUSMINUS
%left '*' '/' '%'
%right UNARY

%%

compstmt : optTerms
	{
		$$ = nil
	}
	| stmts optTerms
	{
		$$ = $1
	}

stmts :
	{
		$$ = nil
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| optTerms stmt
	{
		$$ = []goblin.Stmt{$2}
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| stmts terms stmt
	{
		if $3 != nil {
			$$ = append($1, $3)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = $$
			}
		}
	}

stmt :
	VAR exprIdents '=' exprMany
	{
		$$ = &statement.VarStmt{Names: $2, Exprs: $4}
		$$.SetPosition($1.Position())
	}
	| expr '=' expr
	{
		$$ = &statement.LetsStmt{Left: []goblin.Expr{$1}, Operator: "=", Right: []goblin.Expr{$3}}
	}
	| exprMany '=' exprMany
	{
		$$ = &statement.LetsStmt{Left: $1, Operator: "=", Right: $3}
	}
	| BREAK
	{
		$$ = &statement.BreakStmt{}
		$$.SetPosition($1.Position())
	}
	| CONTINUE
	{
		$$ = &statement.ContinueStmt{}
		$$.SetPosition($1.Position())
	}
	| RETURN exprs
	{
		$$ = &statement.ReturnStmt{Exprs: $2}
		$$.SetPosition($1.Position())
	}
	| THROW expr
	{
		$$ = &statement.ThrowStmt{Expr: $2}
		$$.SetPosition($1.Position())
	}
	| MODULE IDENT '{' compstmt '}'
	{
		$$ = &statement.ModuleStmt{Name: $2.Lit, Stmts: $4}
		$$.SetPosition($1.Position())
	}
	| stmtIf
	{
		$$ = $1
		$$.SetPosition($1.Position())
	}
	| FOR '{' compstmt '}'
	{
		$$ = &statement.LoopStmt{Stmts: $3}
		$$.SetPosition($1.Position())
	}
	| FOR IDENT IN expr '{' compstmt '}'
	{
		$$ = &statement.ForStmt{Var: $2.Lit, Value: $4, Stmts: $6}
		$$.SetPosition($1.Position())
	}
	| FOR exprLets ';' expr ';' expr '{' compstmt '}'
	{
		$$ = &statement.CForStmt{Expr1: $2, Expr2: $4, Expr3: $6, Stmts: $8}
		$$.SetPosition($1.Position())
	}
	| FOR expr '{' compstmt '}'
	{
		$$ = &statement.LoopStmt{Expr: $2, Stmts: $4}
		$$.SetPosition($1.Position())
	}
	| TRY '{' compstmt '}' CATCH IDENT '{' compstmt '}' FINALLY '{' compstmt '}'
	{
		$$ = &statement.TryStmt{Try: $3, Var: $6.Lit, Catch: $8, Finally: $12}
		$$.SetPosition($1.Position())
	}
	| TRY '{' compstmt '}' CATCH '{' compstmt '}' FINALLY '{' compstmt '}'
	{
		$$ = &statement.TryStmt{Try: $3, Catch: $7, Finally: $11}
		$$.SetPosition($1.Position())
	}
	| TRY '{' compstmt '}' CATCH IDENT '{' compstmt '}'
	{
		$$ = &statement.TryStmt{Try: $3, Var: $6.Lit, Catch: $8}
		$$.SetPosition($1.Position())
	}
	| TRY '{' compstmt '}' CATCH '{' compstmt '}'
	{
		$$ = &statement.TryStmt{Try: $3, Catch: $7}
		$$.SetPosition($1.Position())
	}
	| SWITCH expr '{' stmtCases '}'
	{
		$$ = &statement.SwitchStmt{Expr: $2, Cases: $4}
		$$.SetPosition($1.Position())
	}
	| expr
	{
		$$ = &statement.ExprStmt{Expr: $1}
		$$.SetPosition($1.Position())
	}


stmtIf :
	stmtIf ELSE IF expr '{' compstmt '}'
	{
		$1.(*statement.IfStmt).ElseIf = append($1.(*statement.IfStmt).ElseIf, &statement.IfStmt{If: $4, Then: $6})
		$$.SetPosition($1.Position())
	}
	| stmtIf ELSE '{' compstmt '}'
	{
		if $$.(*statement.IfStmt).Else != nil {
			yylex.Error("multiple else statement")
		} else {
			$$.(*statement.IfStmt).Else = append($$.(*statement.IfStmt).Else, $4...)
		}
		$$.SetPosition($1.Position())
	}
	| IF expr '{' compstmt '}'
	{
		$$ = &statement.IfStmt{If: $2, Then: $4, Else: nil}
		$$.SetPosition($1.Position())
	}

stmtCases :
	{
		$$ = []goblin.Stmt{}
	}
	| optTerms stmtCase
	{
		$$ = []goblin.Stmt{$2}
	}
	| optTerms stmtDefault
	{
		$$ = []goblin.Stmt{$2}
	}
	| stmtCases stmtCase
	{
		$$ = append($1, $2)
	}
	| stmtCases stmtDefault
	{
		for _, stmt := range $1 {
			if _, ok := stmt.(*statement.DefaultStmt); ok {
				yylex.Error("multiple default statement")
			}
		}
		$$ = append($1, $2)
	}

stmtCase :
	CASE expr ':' optTerms compstmt
	{
		$$ = &statement.CaseStmt{Expr: $2, Stmts: $5}
	}

stmtDefault :
	DEFAULT ':' optTerms compstmt
	{
		$$ = &statement.DefaultStmt{Stmts: $4}
	}

exprPair :
	STRING ':' expr
	{
		$$ = &goblin.PairExpr{Key: $1.Lit, Value: $3}
	}

exprPairs :
	{
		$$ = []goblin.Expr{}
	}
	| exprPair
	{
		$$ = []goblin.Expr{$1}
	}
	| exprPairs ',' optTerms exprPair
	{
		$$ = append($1, $4)
	}

exprIdents :
	{
		$$ = []string{}
	}
	| IDENT
	{
		$$ = []string{$1.Lit}
	}
	| exprIdents ',' optTerms IDENT
	{
		$$ = append($1, $4.Lit)
	}

exprLets : exprMany '=' exprMany
	{
		$$ = &goblin.LetsExpr{LHSS: $1, Operator: "=", RHSS: $3}
	}

exprMany :
	expr
	{
		$$ = []goblin.Expr{$1}
	}
	| exprs ',' optTerms expr
	{
		$$ = append($1, $4)
	}
	| exprs ',' optTerms IDENT
	{
		$$ = append($1, &goblin.IdentExpr{Lit: $4.Lit})
	}

typ : IDENT
	{
		$$ = goblin.Type{Name: $1.Lit}
	}
	| typ '.' IDENT
	{
		$$ = goblin.Type{Name: $1.Name + "." + $3.Lit}
	}

exprs :
	{
		$$ = nil
	}
	| expr 
	{
		$$ = []goblin.Expr{$1}
	}
	| exprs ',' optTerms expr
	{
		$$ = append($1, $4)
	}
	| exprs ',' optTerms IDENT
	{
		$$ = append($1, &goblin.IdentExpr{Lit: $4.Lit})
	}

expr :
	IDENT
	{
		$$ = &goblin.IdentExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| NUMBER
	{
		$$ = &goblin.NumberExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| '-' expr %prec UNARY
	{
		$$ = &goblin.UnaryExpr{Operator: "-", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '!' expr %prec UNARY
	{
		$$ = &goblin.UnaryExpr{Operator: "!", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '^' expr %prec UNARY
	{
		$$ = &goblin.UnaryExpr{Operator: "^", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '&' IDENT %prec UNARY
	{
		$$ = &goblin.AddrExpr{Expr: &goblin.IdentExpr{Lit: $2.Lit}}
		$$.SetPosition($2.Position())
	}
	| '&' expr '.' IDENT %prec UNARY
	{
		$$ = &goblin.AddrExpr{Expr: &goblin.MemberExpr{Expr: $2, Name: $4.Lit}}
		$$.SetPosition($2.Position())
	}
	| '*' IDENT %prec UNARY
	{
		$$ = &goblin.DerefExpr{Expr: &goblin.IdentExpr{Lit: $2.Lit}}
		$$.SetPosition($2.Position())
	}
	| '*' expr '.' IDENT %prec UNARY
	{
		$$ = &goblin.DerefExpr{Expr: &goblin.MemberExpr{Expr: $2, Name: $4.Lit}}
		$$.SetPosition($2.Position())
	}
	| STRING
	{
		$$ = &goblin.StringExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| TRUE
	{
		$$ = &goblin.ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| FALSE
	{
		$$ = &goblin.ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| NIL
	{
		$$ = &goblin.ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| expr '?' expr ':' expr
	{
		$$ = &goblin.TernaryOpExpr{Expr: $1, LHS: $3, RHS: $5}
		$$.SetPosition($1.Position())
	}
	| expr '.' IDENT
	{
		$$ = &goblin.MemberExpr{Expr: $1, Name: $3.Lit}
		$$.SetPosition($1.Position())
	}
	| FUNC '(' exprIdents ')' '{' compstmt '}'
	{
		$$ = &goblin.FuncExpr{Args: $3, Stmts: $6}
		$$.SetPosition($1.Position())
	}
	| FUNC '(' IDENT VARARG ')' '{' compstmt '}'
	{
		$$ = &goblin.FuncExpr{Args: []string{$3.Lit}, Stmts: $7, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| FUNC IDENT '(' exprIdents ')' '{' compstmt '}'
	{
		$$ = &goblin.FuncExpr{Name: $2.Lit, Args: $4, Stmts: $7}
		$$.SetPosition($1.Position())
	}
	| FUNC IDENT '(' IDENT VARARG ')' '{' compstmt '}'
	{
		$$ = &goblin.FuncExpr{Name: $2.Lit, Args: []string{$4.Lit}, Stmts: $8, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| '[' optTerms exprs optTerms ']'
	{
		$$ = &goblin.ArrayExpr{Exprs: $3}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '[' optTerms exprs ',' optTerms ']'
	{
		$$ = &goblin.ArrayExpr{Exprs: $3}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '{' optTerms exprPairs optTerms '}'
	{
		mapExpr := make(map[string]goblin.Expr)
		for _, v := range $3 {
			mapExpr[v.(*goblin.PairExpr).Key] = v.(*goblin.PairExpr).Value
		}
		$$ = &goblin.MapExpr{MapExpr: mapExpr}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '{' optTerms exprPairs ',' optTerms '}'
	{
		mapExpr := make(map[string]goblin.Expr)
		for _, v := range $3 {
			mapExpr[v.(*goblin.PairExpr).Key] = v.(*goblin.PairExpr).Value
		}
		$$ = &goblin.MapExpr{MapExpr: mapExpr}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '(' expr ')'
	{
		$$ = &goblin.ParenExpr{SubExpr: $2}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| NEW '(' typ ')'
	{
		$$ = &goblin.NewExpr{Type: $3.Name}
		$$.SetPosition($1.Position())
	}
	| expr '+' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "+", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '-' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "-", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '*' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "*", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '/' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "/", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '%' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "%", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr POW expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "**", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr SHIFTLEFT expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "<<", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr SHIFTRIGHT expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: ">>", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr EQEQ expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "==", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr NEQ expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "!=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '>' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: ">", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr GE expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: ">=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '<' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "<", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr LE expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "<=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr PLUSEQ expr
	{
		$$ = &goblin.AssocExpr{LHS: $1, Operator: "+=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr MINUSEQ expr
	{
		$$ = &goblin.AssocExpr{LHS: $1, Operator: "-=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr MULEQ expr
	{
		$$ = &goblin.AssocExpr{LHS: $1, Operator: "*=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr DIVEQ expr
	{
		$$ = &goblin.AssocExpr{LHS: $1, Operator: "/=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr ANDEQ expr
	{
		$$ = &goblin.AssocExpr{LHS: $1, Operator: "&=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr OREQ expr
	{
		$$ = &goblin.AssocExpr{LHS: $1, Operator: "|=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr PLUSPLUS
	{
		$$ = &goblin.AssocExpr{LHS: $1, Operator: "++"}
		$$.SetPosition($1.Position())
	}
	| expr MINUSMINUS
	{
		$$ = &goblin.AssocExpr{LHS: $1, Operator: "--"}
		$$.SetPosition($1.Position())
	}
	| expr '|' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "|", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr OROR expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "||", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '&' expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "&", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr ANDAND expr
	{
		$$ = &goblin.BinOpExpr{LHS: $1, Operator: "&&", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| IDENT '(' exprs VARARG ')'
	{
		$$ = &goblin.CallExpr{Name: $1.Lit, SubExprs: $3, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| IDENT '(' exprs ')'
	{
		$$ = &goblin.CallExpr{Name: $1.Lit, SubExprs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '(' exprs VARARG ')'
	{
		$$ = &goblin.AnonCallExpr{Expr: $1, SubExprs: $3, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| expr '(' exprs ')'
	{
		$$ = &goblin.AnonCallExpr{Expr: $1, SubExprs: $3}
		$$.SetPosition($1.Position())
	}
	| IDENT '[' expr ']'
	{
		$$ = &goblin.ItemExpr{Value: &goblin.IdentExpr{Lit: $1.Lit}, Index: $3}
		$$.SetPosition($1.Position())
	}
	| expr '[' expr ']'
	{
		$$ = &goblin.ItemExpr{Value: $1, Index: $3}
		$$.SetPosition($1.Position())
	}
	| IDENT '[' expr ':' expr ']'
	{
		$$ = &goblin.SliceExpr{Value: &goblin.IdentExpr{Lit: $1.Lit}, Begin: $3, End: $5}
		$$.SetPosition($1.Position())
	}
	| expr '[' expr ':' expr ']'
	{
		$$ = &goblin.SliceExpr{Value: $1, Begin: $3, End: $5}
		$$.SetPosition($1.Position())
	}
	| MAKE '(' typ ')'
	{
		$$ = &goblin.MakeExpr{Type: $3.Name}
		$$.SetPosition($1.Position())
	}
	| MAKE '(' ARRAYLIT typ ',' expr ')'
	{
		$$ = &goblin.MakeArrayExpr{Type: $4.Name, LenExpr: $6}
		$$.SetPosition($1.Position())
	}
	| MAKE '(' ARRAYLIT typ ',' expr ',' expr ')'
	{
		$$ = &goblin.MakeArrayExpr{Type: $4.Name, LenExpr: $6, CapExpr: $8}
		$$.SetPosition($1.Position())
	}

optTerms : /* none */
	| terms
	;


terms : term
	{
	}
	| terms term
	{
	}
	;

term : ';'
	{
	}
	| '\n'
	{
	}
	;

%%
