%{
package goblin

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
	compstmt    []Stmt
	stmtIf      Stmt
	stmtDefault Stmt
	stmtCase    Stmt
	stmtCases   []Stmt
	stmts       []Stmt
	stmt        Stmt
	typ         Type
	expr        Expr
	exprs       []Expr
	exprMany    []Expr
	exprLets    Expr
	exprPair    Expr
	exprPairs   []Expr
	exprIdents  []string
	tok         Token
	term        Token
	terms       Token
	optTerms    Token
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
		$$ = []Stmt{$2}
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
		$$ = &VarStmt{Names: $2, Exprs: $4}
		$$.SetPosition($1.Position())
	}
	| expr '=' expr
	{
		$$ = &LetsStmt{LHSS: []Expr{$1}, Operator: "=", RHSS: []Expr{$3}}
	}
	| exprMany '=' exprMany
	{
		$$ = &LetsStmt{LHSS: $1, Operator: "=", RHSS: $3}
	}
	| BREAK
	{
		$$ = &BreakStmt{}
		$$.SetPosition($1.Position())
	}
	| CONTINUE
	{
		$$ = &ContinueStmt{}
		$$.SetPosition($1.Position())
	}
	| RETURN exprs
	{
		$$ = &ReturnStmt{Exprs: $2}
		$$.SetPosition($1.Position())
	}
	| THROW expr
	{
		$$ = &ThrowStmt{Expr: $2}
		$$.SetPosition($1.Position())
	}
	| MODULE IDENT '{' compstmt '}'
	{
		$$ = &ModuleStmt{Name: $2.Lit, Stmts: $4}
		$$.SetPosition($1.Position())
	}
	| stmtIf
	{
		$$ = $1
		$$.SetPosition($1.Position())
	}
	| FOR '{' compstmt '}'
	{
		$$ = &LoopStmt{Stmts: $3}
		$$.SetPosition($1.Position())
	}
	| FOR IDENT IN expr '{' compstmt '}'
	{
		$$ = &ForStmt{Var: $2.Lit, Value: $4, Stmts: $6}
		$$.SetPosition($1.Position())
	}
	| FOR exprLets ';' expr ';' expr '{' compstmt '}'
	{
		$$ = &CForStmt{Expr1: $2, Expr2: $4, Expr3: $6, Stmts: $8}
		$$.SetPosition($1.Position())
	}
	| FOR expr '{' compstmt '}'
	{
		$$ = &LoopStmt{Expr: $2, Stmts: $4}
		$$.SetPosition($1.Position())
	}
	| TRY '{' compstmt '}' CATCH IDENT '{' compstmt '}' FINALLY '{' compstmt '}'
	{
		$$ = &TryStmt{Try: $3, Var: $6.Lit, Catch: $8, Finally: $12}
		$$.SetPosition($1.Position())
	}
	| TRY '{' compstmt '}' CATCH '{' compstmt '}' FINALLY '{' compstmt '}'
	{
		$$ = &TryStmt{Try: $3, Catch: $7, Finally: $11}
		$$.SetPosition($1.Position())
	}
	| TRY '{' compstmt '}' CATCH IDENT '{' compstmt '}'
	{
		$$ = &TryStmt{Try: $3, Var: $6.Lit, Catch: $8}
		$$.SetPosition($1.Position())
	}
	| TRY '{' compstmt '}' CATCH '{' compstmt '}'
	{
		$$ = &TryStmt{Try: $3, Catch: $7}
		$$.SetPosition($1.Position())
	}
	| SWITCH expr '{' stmtCases '}'
	{
		$$ = &SwitchStmt{Expr: $2, Cases: $4}
		$$.SetPosition($1.Position())
	}
	| expr
	{
		$$ = &ExprStmt{Expr: $1}
		$$.SetPosition($1.Position())
	}


stmtIf :
	stmtIf ELSE IF expr '{' compstmt '}'
	{
		$1.(*IfStmt).ElseIf = append($1.(*IfStmt).ElseIf, &IfStmt{If: $4, Then: $6})
		$$.SetPosition($1.Position())
	}
	| stmtIf ELSE '{' compstmt '}'
	{
		if $$.(*IfStmt).Else != nil {
			yylex.Error("multiple else statement")
		} else {
			$$.(*IfStmt).Else = append($$.(*IfStmt).Else, $4...)
		}
		$$.SetPosition($1.Position())
	}
	| IF expr '{' compstmt '}'
	{
		$$ = &IfStmt{If: $2, Then: $4, Else: nil}
		$$.SetPosition($1.Position())
	}

stmtCases :
	{
		$$ = []Stmt{}
	}
	| optTerms stmtCase
	{
		$$ = []Stmt{$2}
	}
	| optTerms stmtDefault
	{
		$$ = []Stmt{$2}
	}
	| stmtCases stmtCase
	{
		$$ = append($1, $2)
	}
	| stmtCases stmtDefault
	{
		for _, stmt := range $1 {
			if _, ok := stmt.(*DefaultStmt); ok {
				yylex.Error("multiple default statement")
			}
		}
		$$ = append($1, $2)
	}

stmtCase :
	CASE expr ':' optTerms compstmt
	{
		$$ = &CaseStmt{Expr: $2, Stmts: $5}
	}

stmtDefault :
	DEFAULT ':' optTerms compstmt
	{
		$$ = &DefaultStmt{Stmts: $4}
	}

exprPair :
	STRING ':' expr
	{
		$$ = &PairExpr{Key: $1.Lit, Value: $3}
	}

exprPairs :
	{
		$$ = []Expr{}
	}
	| exprPair
	{
		$$ = []Expr{$1}
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
		$$ = &LetsExpr{LHSS: $1, Operator: "=", RHSS: $3}
	}

exprMany :
	expr
	{
		$$ = []Expr{$1}
	}
	| exprs ',' optTerms expr
	{
		$$ = append($1, $4)
	}
	| exprs ',' optTerms IDENT
	{
		$$ = append($1, &IdentExpr{Lit: $4.Lit})
	}

typ : IDENT
	{
		$$ = Type{Name: $1.Lit}
	}
	| typ '.' IDENT
	{
		$$ = Type{Name: $1.Name + "." + $3.Lit}
	}

exprs :
	{
		$$ = nil
	}
	| expr 
	{
		$$ = []Expr{$1}
	}
	| exprs ',' optTerms expr
	{
		$$ = append($1, $4)
	}
	| exprs ',' optTerms IDENT
	{
		$$ = append($1, &IdentExpr{Lit: $4.Lit})
	}

expr :
	IDENT
	{
		$$ = &IdentExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| NUMBER
	{
		$$ = &NumberExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| '-' expr %prec UNARY
	{
		$$ = &UnaryExpr{Operator: "-", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '!' expr %prec UNARY
	{
		$$ = &UnaryExpr{Operator: "!", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '^' expr %prec UNARY
	{
		$$ = &UnaryExpr{Operator: "^", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '&' IDENT %prec UNARY
	{
		$$ = &AddrExpr{Expr: &IdentExpr{Lit: $2.Lit}}
		$$.SetPosition($2.Position())
	}
	| '&' expr '.' IDENT %prec UNARY
	{
		$$ = &AddrExpr{Expr: &MemberExpr{Expr: $2, Name: $4.Lit}}
		$$.SetPosition($2.Position())
	}
	| '*' IDENT %prec UNARY
	{
		$$ = &DerefExpr{Expr: &IdentExpr{Lit: $2.Lit}}
		$$.SetPosition($2.Position())
	}
	| '*' expr '.' IDENT %prec UNARY
	{
		$$ = &DerefExpr{Expr: &MemberExpr{Expr: $2, Name: $4.Lit}}
		$$.SetPosition($2.Position())
	}
	| STRING
	{
		$$ = &StringExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| TRUE
	{
		$$ = &ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| FALSE
	{
		$$ = &ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| NIL
	{
		$$ = &ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| expr '?' expr ':' expr
	{
		$$ = &TernaryOpExpr{Expr: $1, LHS: $3, RHS: $5}
		$$.SetPosition($1.Position())
	}
	| expr '.' IDENT
	{
		$$ = &MemberExpr{Expr: $1, Name: $3.Lit}
		$$.SetPosition($1.Position())
	}
	| FUNC '(' exprIdents ')' '{' compstmt '}'
	{
		$$ = &FuncExpr{Args: $3, Stmts: $6}
		$$.SetPosition($1.Position())
	}
	| FUNC '(' IDENT VARARG ')' '{' compstmt '}'
	{
		$$ = &FuncExpr{Args: []string{$3.Lit}, Stmts: $7, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| FUNC IDENT '(' exprIdents ')' '{' compstmt '}'
	{
		$$ = &FuncExpr{Name: $2.Lit, Args: $4, Stmts: $7}
		$$.SetPosition($1.Position())
	}
	| FUNC IDENT '(' IDENT VARARG ')' '{' compstmt '}'
	{
		$$ = &FuncExpr{Name: $2.Lit, Args: []string{$4.Lit}, Stmts: $8, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| '[' optTerms exprs optTerms ']'
	{
		$$ = &ArrayExpr{Exprs: $3}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '[' optTerms exprs ',' optTerms ']'
	{
		$$ = &ArrayExpr{Exprs: $3}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '{' optTerms exprPairs optTerms '}'
	{
		mapExpr := make(map[string]Expr)
		for _, v := range $3 {
			mapExpr[v.(*PairExpr).Key] = v.(*PairExpr).Value
		}
		$$ = &MapExpr{MapExpr: mapExpr}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '{' optTerms exprPairs ',' optTerms '}'
	{
		mapExpr := make(map[string]Expr)
		for _, v := range $3 {
			mapExpr[v.(*PairExpr).Key] = v.(*PairExpr).Value
		}
		$$ = &MapExpr{MapExpr: mapExpr}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '(' expr ')'
	{
		$$ = &ParenExpr{SubExpr: $2}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| NEW '(' typ ')'
	{
		$$ = &NewExpr{Type: $3.Name}
		$$.SetPosition($1.Position())
	}
	| expr '+' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "+", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '-' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "-", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '*' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "*", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '/' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "/", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '%' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "%", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr POW expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "**", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr SHIFTLEFT expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "<<", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr SHIFTRIGHT expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: ">>", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr EQEQ expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "==", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr NEQ expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "!=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '>' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: ">", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr GE expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: ">=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '<' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "<", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr LE expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "<=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr PLUSEQ expr
	{
		$$ = &AssocExpr{LHS: $1, Operator: "+=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr MINUSEQ expr
	{
		$$ = &AssocExpr{LHS: $1, Operator: "-=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr MULEQ expr
	{
		$$ = &AssocExpr{LHS: $1, Operator: "*=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr DIVEQ expr
	{
		$$ = &AssocExpr{LHS: $1, Operator: "/=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr ANDEQ expr
	{
		$$ = &AssocExpr{LHS: $1, Operator: "&=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr OREQ expr
	{
		$$ = &AssocExpr{LHS: $1, Operator: "|=", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr PLUSPLUS
	{
		$$ = &AssocExpr{LHS: $1, Operator: "++"}
		$$.SetPosition($1.Position())
	}
	| expr MINUSMINUS
	{
		$$ = &AssocExpr{LHS: $1, Operator: "--"}
		$$.SetPosition($1.Position())
	}
	| expr '|' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "|", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr OROR expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "||", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr '&' expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "&", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| expr ANDAND expr
	{
		$$ = &BinOpExpr{LHS: $1, Operator: "&&", RHS: $3}
		$$.SetPosition($1.Position())
	}
	| IDENT '(' exprs VARARG ')'
	{
		$$ = &CallExpr{Name: $1.Lit, SubExprs: $3, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| IDENT '(' exprs ')'
	{
		$$ = &CallExpr{Name: $1.Lit, SubExprs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '(' exprs VARARG ')'
	{
		$$ = &AnonCallExpr{Expr: $1, SubExprs: $3, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| expr '(' exprs ')'
	{
		$$ = &AnonCallExpr{Expr: $1, SubExprs: $3}
		$$.SetPosition($1.Position())
	}
	| IDENT '[' expr ']'
	{
		$$ = &ItemExpr{Value: &IdentExpr{Lit: $1.Lit}, Index: $3}
		$$.SetPosition($1.Position())
	}
	| expr '[' expr ']'
	{
		$$ = &ItemExpr{Value: $1, Index: $3}
		$$.SetPosition($1.Position())
	}
	| IDENT '[' expr ':' expr ']'
	{
		$$ = &SliceExpr{Value: &IdentExpr{Lit: $1.Lit}, Begin: $3, End: $5}
		$$.SetPosition($1.Position())
	}
	| expr '[' expr ':' expr ']'
	{
		$$ = &SliceExpr{Value: $1, Begin: $3, End: $5}
		$$.SetPosition($1.Position())
	}
	| MAKE '(' typ ')'
	{
		$$ = &MakeExpr{Type: $3.Name}
		$$.SetPosition($1.Position())
	}
	| MAKE '(' ARRAYLIT typ ',' expr ')'
	{
		$$ = &MakeArrayExpr{Type: $4.Name, LenExpr: $6}
		$$.SetPosition($1.Position())
	}
	| MAKE '(' ARRAYLIT typ ',' expr ',' expr ')'
	{
		$$ = &MakeArrayExpr{Type: $4.Name, LenExpr: $6, CapExpr: $8}
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
