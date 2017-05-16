package parser

import __yyfmt__ "fmt"

import (
	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/statement"
)

type yySymType struct {
	yys         int
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

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VARARG = 57350
const FUNC = 57351
const RETURN = 57352
const VAR = 57353
const THROW = 57354
const IF = 57355
const ELSE = 57356
const FOR = 57357
const IN = 57358
const EQEQ = 57359
const NEQ = 57360
const GE = 57361
const LE = 57362
const OROR = 57363
const ANDAND = 57364
const NEW = 57365
const TRUE = 57366
const FALSE = 57367
const NIL = 57368
const MODULE = 57369
const TRY = 57370
const CATCH = 57371
const FINALLY = 57372
const PLUSEQ = 57373
const MINUSEQ = 57374
const MULEQ = 57375
const DIVEQ = 57376
const ANDEQ = 57377
const OREQ = 57378
const BREAK = 57379
const CONTINUE = 57380
const PLUSPLUS = 57381
const MINUSMINUS = 57382
const POW = 57383
const SHIFTLEFT = 57384
const SHIFTRIGHT = 57385
const SWITCH = 57386
const CASE = 57387
const DEFAULT = 57388
const MAKE = 57389
const ARRAYLIT = 57390
const UNARY = 57391

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"FUNC",
	"RETURN",
	"VAR",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"OROR",
	"ANDAND",
	"NEW",
	"TRUE",
	"FALSE",
	"NIL",
	"MODULE",
	"TRY",
	"CATCH",
	"FINALLY",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"ANDEQ",
	"OREQ",
	"BREAK",
	"CONTINUE",
	"PLUSPLUS",
	"MINUSMINUS",
	"POW",
	"SHIFTLEFT",
	"SHIFTRIGHT",
	"SWITCH",
	"CASE",
	"DEFAULT",
	"MAKE",
	"ARRAYLIT",
	"'='",
	"'?'",
	"':'",
	"','",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"'{'",
	"'}'",
	"';'",
	"'.'",
	"'!'",
	"'^'",
	"'&'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'|'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 114,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	52, 48,
	-2, 1,
	-1, 10,
	52, 49,
	-2, 24,
	-1, 41,
	52, 48,
	-2, 115,
	-1, 82,
	62, 3,
	-2, 114,
	-1, 85,
	52, 49,
	-2, 43,
	-1, 87,
	62, 3,
	-2, 114,
	-1, 94,
	1, 57,
	8, 57,
	45, 57,
	46, 57,
	49, 57,
	51, 57,
	52, 57,
	61, 57,
	62, 57,
	63, 57,
	69, 57,
	71, 57,
	73, 57,
	-2, 52,
	-1, 96,
	1, 59,
	8, 59,
	45, 59,
	46, 59,
	49, 59,
	51, 59,
	52, 59,
	61, 59,
	62, 59,
	63, 59,
	69, 59,
	71, 59,
	73, 59,
	-2, 52,
	-1, 121,
	17, 0,
	18, 0,
	-2, 85,
	-1, 122,
	17, 0,
	18, 0,
	-2, 86,
	-1, 140,
	52, 49,
	-2, 43,
	-1, 142,
	62, 3,
	-2, 114,
	-1, 144,
	62, 3,
	-2, 114,
	-1, 146,
	62, 1,
	-2, 36,
	-1, 149,
	62, 3,
	-2, 114,
	-1, 170,
	62, 3,
	-2, 114,
	-1, 210,
	52, 50,
	-2, 44,
	-1, 211,
	1, 45,
	45, 45,
	46, 45,
	49, 45,
	52, 51,
	62, 45,
	63, 45,
	73, 45,
	-2, 52,
	-1, 218,
	1, 51,
	8, 51,
	45, 51,
	46, 51,
	52, 51,
	62, 51,
	63, 51,
	69, 51,
	71, 51,
	73, 51,
	-2, 52,
	-1, 220,
	62, 3,
	-2, 114,
	-1, 222,
	62, 3,
	-2, 114,
	-1, 235,
	62, 3,
	-2, 114,
	-1, 252,
	62, 3,
	-2, 114,
	-1, 257,
	62, 3,
	-2, 114,
	-1, 258,
	62, 3,
	-2, 114,
	-1, 266,
	62, 3,
	-2, 114,
	-1, 267,
	62, 3,
	-2, 114,
	-1, 270,
	45, 3,
	46, 3,
	62, 3,
	-2, 114,
	-1, 274,
	62, 3,
	-2, 114,
	-1, 280,
	45, 3,
	46, 3,
	62, 3,
	-2, 114,
	-1, 293,
	62, 3,
	-2, 114,
	-1, 294,
	62, 3,
	-2, 114,
}

const yyPrivate = 57344

const yyLast = 1945

var yyAct = [...]int{

	78, 162, 227, 10, 43, 228, 165, 38, 6, 203,
	2, 1, 201, 240, 40, 11, 79, 259, 7, 85,
	6, 88, 77, 6, 91, 92, 93, 95, 97, 147,
	7, 109, 236, 7, 86, 89, 102, 90, 233, 215,
	106, 191, 10, 174, 100, 101, 110, 111, 237, 113,
	114, 115, 116, 117, 118, 119, 120, 121, 122, 123,
	124, 125, 126, 127, 128, 129, 130, 131, 132, 109,
	206, 133, 134, 135, 136, 208, 138, 140, 99, 206,
	159, 89, 137, 90, 207, 141, 197, 141, 104, 103,
	163, 154, 139, 146, 145, 245, 298, 153, 297, 151,
	229, 230, 192, 157, 175, 148, 290, 206, 160, 140,
	287, 167, 286, 283, 282, 279, 169, 226, 271, 265,
	172, 264, 246, 251, 171, 242, 224, 221, 219, 188,
	182, 294, 293, 143, 274, 141, 267, 258, 257, 235,
	142, 87, 98, 105, 180, 254, 262, 10, 184, 185,
	108, 140, 178, 109, 179, 205, 181, 150, 76, 229,
	230, 186, 292, 190, 199, 166, 187, 288, 225, 8,
	210, 202, 204, 81, 214, 209, 238, 163, 216, 217,
	252, 144, 212, 198, 62, 63, 64, 65, 66, 67,
	244, 213, 166, 231, 53, 234, 232, 200, 196, 195,
	62, 63, 64, 65, 66, 67, 243, 158, 112, 168,
	53, 107, 239, 5, 241, 80, 44, 47, 42, 161,
	72, 74, 84, 75, 250, 70, 50, 51, 52, 189,
	253, 4, 248, 47, 249, 41, 72, 74, 17, 75,
	217, 70, 3, 261, 0, 0, 263, 256, 62, 63,
	64, 65, 66, 67, 0, 42, 68, 69, 53, 0,
	0, 0, 0, 0, 268, 270, 0, 0, 0, 272,
	273, 0, 48, 49, 50, 51, 52, 285, 277, 278,
	280, 47, 281, 0, 72, 74, 284, 75, 0, 70,
	0, 0, 289, 56, 57, 59, 61, 71, 73, 0,
	0, 0, 0, 0, 0, 295, 296, 62, 63, 64,
	65, 66, 67, 0, 0, 68, 69, 53, 54, 55,
	0, 0, 0, 0, 0, 0, 46, 0, 276, 58,
	60, 48, 49, 50, 51, 52, 0, 0, 0, 0,
	47, 0, 0, 72, 74, 275, 75, 0, 70, 56,
	57, 59, 61, 71, 73, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 63, 64, 65, 66, 67, 0,
	0, 68, 69, 53, 54, 55, 0, 0, 0, 0,
	0, 0, 46, 194, 0, 58, 60, 48, 49, 50,
	51, 52, 0, 0, 0, 0, 47, 0, 0, 72,
	74, 0, 75, 193, 70, 56, 57, 59, 61, 71,
	73, 0, 0, 0, 0, 0, 0, 0, 0, 62,
	63, 64, 65, 66, 67, 0, 0, 68, 69, 53,
	54, 55, 0, 0, 0, 0, 0, 0, 46, 177,
	0, 58, 60, 48, 49, 50, 51, 52, 0, 0,
	0, 0, 47, 0, 0, 72, 74, 0, 75, 176,
	70, 56, 57, 59, 61, 71, 73, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 63, 64, 65, 66,
	67, 0, 0, 68, 69, 53, 54, 55, 0, 0,
	0, 0, 0, 0, 46, 0, 0, 58, 60, 48,
	49, 50, 51, 52, 0, 0, 0, 0, 47, 0,
	0, 72, 74, 291, 75, 0, 70, 56, 57, 59,
	61, 71, 73, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 63, 64, 65, 66, 67, 0, 0, 68,
	69, 53, 54, 55, 0, 0, 0, 0, 0, 0,
	46, 269, 0, 58, 60, 48, 49, 50, 51, 52,
	0, 0, 0, 0, 47, 0, 0, 72, 74, 0,
	75, 0, 70, 56, 57, 59, 61, 71, 73, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 63, 64,
	65, 66, 67, 0, 0, 68, 69, 53, 54, 55,
	0, 0, 0, 0, 0, 0, 46, 0, 0, 58,
	60, 48, 49, 50, 51, 52, 0, 266, 0, 0,
	47, 0, 0, 72, 74, 0, 75, 0, 70, 56,
	57, 59, 61, 71, 73, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 63, 64, 65, 66, 67, 0,
	0, 68, 69, 53, 54, 55, 0, 0, 0, 0,
	0, 0, 46, 0, 0, 58, 60, 48, 49, 50,
	51, 52, 0, 0, 0, 0, 47, 0, 0, 72,
	74, 0, 75, 255, 70, 56, 57, 59, 61, 71,
	73, 0, 0, 0, 0, 0, 0, 0, 0, 62,
	63, 64, 65, 66, 67, 0, 0, 68, 69, 53,
	54, 55, 0, 0, 0, 0, 0, 0, 46, 0,
	0, 58, 60, 48, 49, 50, 51, 52, 0, 0,
	0, 0, 47, 0, 0, 72, 74, 0, 75, 247,
	70, 56, 57, 59, 61, 71, 73, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 63, 64, 65, 66,
	67, 0, 0, 68, 69, 53, 54, 55, 0, 0,
	0, 0, 0, 0, 46, 0, 0, 58, 60, 48,
	49, 50, 51, 52, 0, 0, 0, 223, 47, 0,
	0, 72, 74, 0, 75, 0, 70, 56, 57, 59,
	61, 71, 73, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 63, 64, 65, 66, 67, 0, 0, 68,
	69, 53, 54, 55, 0, 0, 0, 0, 0, 0,
	46, 0, 0, 58, 60, 48, 49, 50, 51, 52,
	0, 222, 0, 0, 47, 0, 0, 72, 74, 0,
	75, 0, 70, 56, 57, 59, 61, 71, 73, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 63, 64,
	65, 66, 67, 0, 0, 68, 69, 53, 54, 55,
	0, 0, 0, 0, 0, 0, 46, 0, 0, 58,
	60, 48, 49, 50, 51, 52, 0, 220, 0, 0,
	47, 0, 0, 72, 74, 0, 75, 0, 70, 56,
	57, 59, 61, 71, 73, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 63, 64, 65, 66, 67, 0,
	0, 68, 69, 53, 54, 55, 0, 0, 0, 0,
	0, 0, 46, 173, 0, 58, 60, 48, 49, 50,
	51, 52, 0, 0, 0, 0, 47, 0, 0, 72,
	74, 0, 75, 0, 70, 56, 57, 59, 61, 71,
	73, 0, 0, 0, 0, 0, 0, 0, 0, 62,
	63, 64, 65, 66, 67, 0, 0, 68, 69, 53,
	54, 55, 0, 0, 0, 0, 0, 0, 46, 0,
	0, 58, 60, 48, 49, 50, 51, 52, 0, 170,
	0, 0, 47, 0, 0, 72, 74, 0, 75, 0,
	70, 56, 57, 59, 61, 71, 73, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 63, 64, 65, 66,
	67, 0, 0, 68, 69, 53, 54, 55, 0, 0,
	0, 0, 0, 0, 46, 0, 0, 58, 60, 48,
	49, 50, 51, 52, 0, 0, 0, 0, 47, 0,
	0, 72, 74, 164, 75, 0, 70, 56, 57, 59,
	61, 71, 73, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 63, 64, 65, 66, 67, 0, 0, 68,
	69, 53, 54, 55, 0, 0, 0, 0, 0, 0,
	46, 0, 0, 58, 60, 48, 49, 50, 51, 52,
	0, 152, 0, 0, 47, 0, 0, 72, 74, 0,
	75, 0, 70, 56, 57, 59, 61, 71, 73, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 63, 64,
	65, 66, 67, 0, 0, 68, 69, 53, 54, 55,
	0, 0, 0, 0, 0, 0, 46, 0, 0, 58,
	60, 48, 49, 50, 51, 52, 0, 149, 0, 0,
	47, 0, 0, 72, 74, 0, 75, 0, 70, 56,
	57, 59, 61, 71, 73, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 63, 64, 65, 66, 67, 0,
	0, 68, 69, 53, 54, 55, 0, 0, 0, 0,
	0, 45, 46, 0, 0, 58, 60, 48, 49, 50,
	51, 52, 0, 0, 0, 0, 47, 0, 0, 72,
	74, 0, 75, 0, 70, 56, 57, 59, 61, 71,
	73, 0, 0, 0, 0, 0, 0, 0, 0, 62,
	63, 64, 65, 66, 67, 0, 0, 68, 69, 53,
	54, 55, 0, 0, 0, 0, 0, 0, 46, 0,
	0, 58, 60, 48, 49, 50, 51, 52, 0, 0,
	0, 0, 47, 0, 0, 72, 74, 0, 75, 0,
	70, 56, 57, 59, 61, 71, 73, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 63, 64, 65, 66,
	67, 0, 0, 68, 69, 53, 54, 55, 0, 0,
	0, 0, 0, 0, 46, 0, 0, 58, 60, 48,
	49, 50, 51, 52, 0, 0, 0, 0, 156, 0,
	0, 72, 74, 0, 75, 0, 70, 56, 57, 59,
	61, 71, 73, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 63, 64, 65, 66, 67, 0, 0, 68,
	69, 53, 54, 55, 0, 0, 0, 0, 0, 0,
	46, 0, 0, 58, 60, 48, 49, 50, 51, 52,
	0, 0, 0, 0, 155, 0, 0, 72, 74, 0,
	75, 0, 70, 21, 22, 28, 0, 0, 32, 14,
	9, 15, 39, 0, 18, 0, 0, 0, 0, 0,
	0, 0, 36, 29, 30, 31, 16, 19, 0, 0,
	0, 0, 0, 0, 0, 0, 12, 13, 0, 0,
	0, 0, 0, 20, 0, 0, 37, 0, 0, 0,
	0, 0, 0, 0, 0, 23, 27, 0, 0, 0,
	34, 0, 6, 0, 24, 25, 26, 35, 0, 33,
	0, 0, 7, 56, 57, 59, 61, 0, 73, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 63, 64,
	65, 66, 67, 0, 0, 68, 69, 53, 54, 55,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 58,
	60, 48, 49, 50, 51, 52, 56, 57, 59, 61,
	47, 0, 0, 72, 74, 0, 75, 0, 70, 0,
	62, 63, 64, 65, 66, 67, 0, 0, 68, 69,
	53, 54, 55, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 58, 60, 48, 49, 50, 51, 52, 0,
	0, 0, 0, 47, 0, 0, 72, 74, 0, 75,
	0, 70, 21, 22, 183, 0, 0, 32, 14, 9,
	15, 39, 0, 18, 0, 0, 0, 0, 0, 0,
	0, 36, 29, 30, 31, 16, 19, 0, 0, 0,
	0, 0, 0, 0, 0, 12, 13, 0, 0, 0,
	0, 0, 20, 0, 0, 37, 0, 0, 0, 0,
	0, 0, 0, 0, 23, 27, 0, 59, 61, 34,
	0, 0, 0, 24, 25, 26, 35, 0, 33, 62,
	63, 64, 65, 66, 67, 0, 0, 68, 69, 53,
	54, 55, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 58, 60, 48, 49, 50, 51, 52, 0, 0,
	0, 0, 47, 0, 0, 72, 74, 0, 75, 0,
	70, 21, 22, 28, 0, 0, 32, 14, 9, 15,
	39, 0, 18, 0, 0, 0, 0, 0, 0, 0,
	36, 29, 30, 31, 16, 19, 218, 22, 28, 0,
	0, 32, 0, 0, 12, 13, 0, 0, 0, 0,
	0, 20, 0, 0, 37, 36, 29, 30, 31, 0,
	0, 0, 0, 23, 27, 21, 22, 28, 34, 0,
	32, 0, 24, 25, 26, 35, 0, 33, 0, 37,
	0, 0, 0, 0, 36, 29, 30, 31, 23, 27,
	218, 22, 28, 34, 0, 32, 0, 24, 25, 26,
	35, 0, 33, 260, 0, 0, 0, 0, 37, 36,
	29, 30, 31, 0, 0, 0, 0, 23, 27, 211,
	22, 28, 34, 0, 32, 0, 24, 25, 26, 35,
	0, 33, 0, 37, 0, 0, 0, 0, 36, 29,
	30, 31, 23, 27, 96, 22, 28, 34, 0, 32,
	0, 24, 25, 26, 35, 0, 33, 0, 0, 0,
	0, 0, 37, 36, 29, 30, 31, 0, 0, 0,
	0, 23, 27, 94, 22, 28, 34, 0, 32, 0,
	24, 25, 26, 35, 0, 33, 0, 37, 0, 0,
	0, 0, 36, 29, 30, 31, 23, 27, 83, 22,
	28, 34, 0, 32, 0, 24, 25, 26, 35, 0,
	33, 0, 0, 0, 0, 0, 37, 36, 29, 30,
	31, 0, 0, 0, 0, 23, 27, 0, 0, 0,
	34, 0, 0, 0, 24, 25, 26, 35, 0, 33,
	0, 37, 0, 0, 0, 0, 0, 0, 0, 0,
	23, 27, 0, 0, 0, 82, 0, 0, 0, 24,
	25, 26, 35, 0, 33,
}
var yyPact = [...]int{

	-55, -1000, 1687, -55, -55, -1000, -1000, -1000, -1000, 212,
	1172, 109, -1000, -1000, 1741, 1741, 211, 159, 1874, 80,
	1741, -33, -1000, 1741, 1741, 1741, 1849, 1820, -1000, -1000,
	-1000, -1000, 74, -55, -55, 1741, 21, 20, 91, 1741,
	-1000, 1409, -1000, 101, -1000, 1741, 1741, 204, 1741, 1741,
	1741, 1741, 1741, 1741, 1741, 1741, 1741, 1741, 1741, 1741,
	1741, 1741, 1741, 1741, 1741, 1741, 1741, 1741, -1000, -1000,
	1741, 1741, 1741, 1741, 1741, 1741, 1741, 83, 1228, 1228,
	79, 120, -55, 13, 42, 1116, 108, -55, 1060, 1741,
	1741, 153, 153, 153, -33, 1340, -33, 1284, 203, 12,
	1741, 171, 1004, 188, 161, -55, 948, -1000, 1741, -55,
	1228, 892, -1000, 169, 169, 153, 153, 153, 1228, 217,
	217, 1618, 1618, 217, 217, 217, 217, 1228, 1228, 1228,
	1228, 1228, 1228, 1228, 1466, 1228, 1509, 35, 388, -1000,
	1228, -55, -55, 1741, -55, 68, 1578, 1741, 1741, -55,
	1741, 67, -55, 33, 332, 195, 194, 17, 175, 193,
	-40, -43, -1000, 104, -1000, 15, -1000, 6, 188, 1795,
	-55, -1000, 187, 1741, -30, -1000, -1000, 1741, 1766, 66,
	836, 65, -1000, 104, 780, 724, 64, -1000, 139, 55,
	114, -31, -1000, -1000, 1741, -1000, -1000, 78, -37, -21,
	168, -55, -58, -55, 63, 1741, 186, -1000, -1000, 43,
	1228, -33, 60, -1000, 1228, -1000, 668, 1228, -33, -1000,
	-55, -1000, -55, 1741, -1000, 119, -1000, -1000, -1000, 1741,
	94, -1000, -1000, -1000, 612, -55, 77, 76, -52, 1712,
	-1000, 84, -1000, 1228, -1000, 1741, -1000, -1000, 59, 57,
	556, 75, -55, 500, -55, -1000, 56, -55, -55, 73,
	-1000, -1000, -1000, 276, -1000, -1000, -55, -55, 53, -55,
	-55, -1000, 52, 51, -55, -1000, 1741, 50, 48, 137,
	-55, -1000, -1000, -1000, 44, 444, -1000, 132, 71, -1000,
	-1000, -1000, 70, -55, -55, 36, 34, -1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 242, 169, 238, 5, 2, 229, 6, 0,
	7, 15, 222, 1, 219, 4, 10, 231, 213,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 7, 7,
	7, 7, 7, 6, 5, 13, 14, 14, 14, 15,
	15, 15, 12, 11, 11, 11, 8, 8, 10, 10,
	10, 10, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 16, 16, 17, 17, 18, 18,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 5, 1, 7, 5, 5, 0, 2,
	2, 2, 2, 5, 4, 3, 0, 1, 4, 0,
	1, 4, 3, 1, 4, 4, 1, 3, 0, 1,
	4, 4, 1, 1, 2, 2, 2, 2, 4, 2,
	4, 1, 1, 1, 1, 5, 3, 7, 8, 8,
	9, 5, 6, 5, 6, 3, 4, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	3, 3, 3, 5, 4, 5, 4, 4, 4, 6,
	6, 4, 7, 9, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -16, -2, -17, -18, 63, 73, -3, 11,
	-9, -11, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 4, 5, 56, 65, 66, 67, 57, 6, 24,
	25, 26, 9, 70, 61, 68, 23, 47, -10, 13,
	-16, -17, -18, -15, 4, 49, 50, 64, 55, 56,
	57, 58, 59, 41, 42, 43, 17, 18, 53, 19,
	54, 20, 31, 32, 33, 34, 35, 36, 39, 40,
	72, 21, 67, 22, 68, 70, 49, -10, -9, -9,
	4, 14, 61, 4, -12, -9, -11, 61, -9, 68,
	70, -9, -9, -9, 4, -9, 4, -9, 68, 4,
	-16, -16, -9, 68, 68, 52, -9, -3, 49, 52,
	-9, -9, 4, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -10, -9, -11,
	-9, 52, 61, 13, 61, -1, -16, 16, 63, 61,
	49, -1, 61, -10, -9, 64, 64, -15, 4, 68,
	-10, -14, -13, 6, 69, -8, 4, -8, 48, -16,
	61, -11, -16, 51, 8, 69, 71, 51, -16, -1,
	-9, -1, 62, 6, -9, -9, -1, -11, 62, -7,
	-16, 8, 69, 71, 51, 4, 4, 69, 8, -15,
	4, 52, -16, 52, -16, 51, 64, 69, 69, -8,
	-9, 4, -1, 4, -9, 69, -9, -9, 4, 62,
	61, 62, 61, 63, 62, 29, 62, -6, -5, 45,
	46, -6, -5, 69, -9, 61, 69, 69, 8, -16,
	71, -16, 62, -9, 4, 52, 62, 71, -1, -1,
	-9, 4, 61, -9, 51, 71, -1, 61, 61, 69,
	71, -13, 62, -9, 62, 62, 61, 61, -1, 51,
	-16, 62, -1, -1, 61, 69, 52, -1, -1, 62,
	-16, -1, 62, 62, -1, -9, 62, 62, 30, -1,
	62, 69, 30, 61, 61, -1, -1, 62, 62,
}
var yyDef = [...]int{

	-2, -2, -2, 114, 115, 116, 118, 119, 4, 39,
	-2, 0, 9, 10, 48, 0, 0, 14, 48, 0,
	0, 52, 53, 0, 0, 0, 0, 0, 61, 62,
	63, 64, 0, 114, 114, 0, 0, 0, 0, 0,
	2, -2, 117, 0, 40, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 97, 98,
	0, 0, 0, 0, 48, 0, 48, 11, 49, 12,
	0, 0, -2, 52, 0, -2, 0, -2, 0, 48,
	0, 54, 55, 56, -2, 0, -2, 0, 39, 0,
	48, 36, 0, 0, 0, 114, 0, 5, 48, 114,
	7, 0, 66, 77, 78, 79, 80, 81, 82, 83,
	84, -2, -2, 87, 88, 89, 90, 91, 92, 93,
	94, 95, 96, 99, 100, 101, 102, 0, 0, 8,
	-2, 114, -2, 0, -2, 0, -2, 0, 0, -2,
	48, 0, 28, 0, 0, 0, 0, 0, 40, 39,
	114, 114, 37, 0, 75, 0, 46, 0, 0, 0,
	-2, 6, 0, 0, 0, 106, 108, 0, 0, 0,
	0, 0, 15, 61, 0, 0, 0, 42, 0, 0,
	0, 0, 104, 107, 0, 58, 60, 0, 0, 0,
	40, 114, 0, 114, 0, 0, 0, 76, 111, 0,
	-2, -2, 0, 41, 65, 105, 0, 50, -2, 13,
	-2, 26, -2, 0, 18, 0, 23, 31, 32, 0,
	0, 29, 30, 103, 0, -2, 0, 0, 0, 0,
	71, 0, 73, 35, 47, 0, 27, 110, 0, 0,
	0, 0, -2, 0, 114, 109, 0, -2, -2, 0,
	72, 38, 74, 0, 25, 16, -2, -2, 0, 114,
	-2, 67, 0, 0, -2, 112, 0, 0, 0, 22,
	-2, 34, 68, 69, 0, 0, 17, 21, 0, 33,
	70, 113, 0, -2, -2, 0, 0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	73, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 65, 3, 3, 3, 59, 67, 3,
	68, 69, 57, 55, 52, 56, 64, 58, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 51, 63,
	54, 49, 53, 50, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 70, 3, 71, 66, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 61, 72, 62,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 60,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

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
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
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
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			yyrcvr.char = -1
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
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []goblin.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.stmts
				}
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &statement.VarStmt{Names: yyDollar[2].exprIdents, Exprs: yyDollar[4].exprMany}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &statement.LetsStmt{Left: []goblin.Expr{yyDollar[1].expr}, Operator: "=", Right: []goblin.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &statement.LetsStmt{Left: yyDollar[1].exprMany, Operator: "=", Right: yyDollar[3].exprMany}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &statement.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &statement.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &statement.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &statement.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &statement.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmtIf
			yyVAL.stmt.SetPosition(yyDollar[1].stmtIf.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &statement.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &statement.ForStmt{Var: yyDollar[2].tok.Lit, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &statement.CForStmt{Expr1: yyDollar[2].exprLets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &statement.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		{
			yyVAL.stmt = &statement.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		{
			yyVAL.stmt = &statement.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &statement.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &statement.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &statement.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmtCases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &statement.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyDollar[1].stmtIf.(*statement.IfStmt).ElseIf = append(yyDollar[1].stmtIf.(*statement.IfStmt).ElseIf, &statement.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmtIf.SetPosition(yyDollar[1].stmtIf.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if yyVAL.stmtIf.(*statement.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmtIf.(*statement.IfStmt).Else = append(yyVAL.stmtIf.(*statement.IfStmt).Else, yyDollar[4].compstmt...)
			}
			yyVAL.stmtIf.SetPosition(yyDollar[1].stmtIf.Position())
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmtIf = &statement.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmtIf.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.stmtCases = []goblin.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmtCases = []goblin.Stmt{yyDollar[2].stmtCase}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmtCases = []goblin.Stmt{yyDollar[2].stmtDefault}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmtCases = append(yyDollar[1].stmtCases, yyDollar[2].stmtCase)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			for _, stmt := range yyDollar[1].stmtCases {
				if _, ok := stmt.(*statement.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmtCases = append(yyDollar[1].stmtCases, yyDollar[2].stmtDefault)
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmtCase = &statement.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmtDefault = &statement.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.exprPair = &goblin.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprPairs = []goblin.Expr{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprPairs = []goblin.Expr{yyDollar[1].exprPair}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprPairs = append(yyDollar[1].exprPairs, yyDollar[4].exprPair)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprIdents = []string{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprIdents = []string{yyDollar[1].tok.Lit}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprIdents = append(yyDollar[1].exprIdents, yyDollar[4].tok.Lit)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.exprLets = &goblin.LetsExpr{LHSS: yyDollar[1].exprMany, Operator: "=", RHSS: yyDollar[3].exprMany}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprMany = []goblin.Expr{yyDollar[1].expr}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprMany = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprMany = append(yyDollar[1].exprs, &goblin.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.typ = goblin.Type{Name: yyDollar[1].tok.Lit}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.typ = goblin.Type{Name: yyDollar[1].typ.Name + "." + yyDollar[3].tok.Lit}
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprs = nil
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []goblin.Expr{yyDollar[1].expr}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &goblin.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &goblin.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &goblin.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &goblin.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &goblin.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &goblin.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &goblin.AddrExpr{Expr: &goblin.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &goblin.AddrExpr{Expr: &goblin.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &goblin.DerefExpr{Expr: &goblin.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &goblin.DerefExpr{Expr: &goblin.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &goblin.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &goblin.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &goblin.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &goblin.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &goblin.TernaryOpExpr{Expr: yyDollar[1].expr, LHS: yyDollar[3].expr, RHS: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.expr = &goblin.FuncExpr{Args: yyDollar[3].exprIdents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.expr = &goblin.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.expr = &goblin.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].exprIdents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.expr = &goblin.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &goblin.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.expr = &goblin.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			mapExpr := make(map[string]goblin.Expr)
			for _, v := range yyDollar[3].exprPairs {
				mapExpr[v.(*goblin.PairExpr).Key] = v.(*goblin.PairExpr).Value
			}
			yyVAL.expr = &goblin.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 74:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			mapExpr := make(map[string]goblin.Expr)
			for _, v := range yyDollar[3].exprPairs {
				mapExpr[v.(*goblin.PairExpr).Key] = v.(*goblin.PairExpr).Value
			}
			yyVAL.expr = &goblin.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &goblin.NewExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "+", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "-", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "*", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "/", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "%", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "**", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "<<", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: ">>", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "==", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "!=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: ">", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: ">=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "<", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "<=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.AssocExpr{LHS: yyDollar[1].expr, Operator: "+=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.AssocExpr{LHS: yyDollar[1].expr, Operator: "-=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.AssocExpr{LHS: yyDollar[1].expr, Operator: "*=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.AssocExpr{LHS: yyDollar[1].expr, Operator: "/=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.AssocExpr{LHS: yyDollar[1].expr, Operator: "&=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.AssocExpr{LHS: yyDollar[1].expr, Operator: "|=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &goblin.AssocExpr{LHS: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &goblin.AssocExpr{LHS: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "|", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "||", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "&", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &goblin.BinOpExpr{LHS: yyDollar[1].expr, Operator: "&&", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &goblin.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &goblin.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &goblin.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &goblin.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &goblin.ItemExpr{Value: &goblin.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &goblin.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.expr = &goblin.SliceExpr{Value: &goblin.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.expr = &goblin.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &goblin.MakeExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.expr = &goblin.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.expr = &goblin.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	}
	goto yystack /* stack new state and value */
}
