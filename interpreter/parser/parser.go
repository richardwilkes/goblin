package parser

import __yyfmt__ "fmt"

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/interpreter/expression"
	"github.com/richardwilkes/goblin/interpreter/statement"
)

type yySymType struct {
	yys         int
	compstmt    []interpreter.Stmt
	stmtIf      interpreter.Stmt
	stmtDefault interpreter.Stmt
	stmtCase    interpreter.Stmt
	stmtCases   []interpreter.Stmt
	stmts       []interpreter.Stmt
	stmt        interpreter.Stmt
	typ         interpreter.Type
	expr        interpreter.Expr
	exprs       []interpreter.Expr
	exprMany    []interpreter.Expr
	exprLets    interpreter.Expr
	exprPair    interpreter.Expr
	exprPairs   []interpreter.Expr
	exprIdents  []string
	tok         interpreter.Token
	term        interpreter.Token
	terms       interpreter.Token
	optTerms    interpreter.Token
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
	-2, 120,
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
	-2, 121,
	-1, 82,
	62, 3,
	-2, 120,
	-1, 85,
	52, 49,
	-2, 43,
	-1, 87,
	62, 3,
	-2, 120,
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
	-1, 141,
	52, 49,
	-2, 43,
	-1, 143,
	62, 3,
	-2, 120,
	-1, 145,
	62, 3,
	-2, 120,
	-1, 147,
	62, 1,
	-2, 36,
	-1, 150,
	62, 3,
	-2, 120,
	-1, 172,
	62, 3,
	-2, 120,
	-1, 216,
	52, 50,
	-2, 44,
	-1, 217,
	1, 45,
	45, 45,
	46, 45,
	49, 45,
	52, 51,
	62, 45,
	63, 45,
	73, 45,
	-2, 52,
	-1, 226,
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
	-1, 228,
	62, 3,
	-2, 120,
	-1, 230,
	62, 3,
	-2, 120,
	-1, 245,
	62, 3,
	-2, 120,
	-1, 262,
	62, 3,
	-2, 120,
	-1, 267,
	62, 3,
	-2, 120,
	-1, 268,
	62, 3,
	-2, 120,
	-1, 276,
	62, 3,
	-2, 120,
	-1, 277,
	62, 3,
	-2, 120,
	-1, 280,
	45, 3,
	46, 3,
	62, 3,
	-2, 120,
	-1, 284,
	62, 3,
	-2, 120,
	-1, 290,
	45, 3,
	46, 3,
	62, 3,
	-2, 120,
	-1, 303,
	62, 3,
	-2, 120,
	-1, 304,
	62, 3,
	-2, 120,
}

const yyPrivate = 57344

const yyLast = 2323

var yyAct = [...]int{

	78, 164, 235, 10, 43, 236, 167, 38, 6, 209,
	11, 89, 1, 90, 99, 250, 79, 269, 7, 85,
	6, 88, 77, 207, 91, 92, 93, 95, 97, 86,
	7, 148, 246, 212, 6, 195, 102, 176, 214, 241,
	106, 221, 10, 212, 7, 161, 110, 111, 213, 113,
	114, 115, 116, 117, 118, 119, 120, 121, 122, 123,
	124, 125, 126, 127, 128, 129, 130, 131, 132, 109,
	104, 133, 134, 135, 136, 103, 138, 141, 98, 142,
	165, 142, 137, 89, 109, 90, 247, 140, 2, 149,
	255, 155, 40, 237, 238, 146, 196, 154, 177, 308,
	152, 203, 212, 159, 307, 300, 297, 296, 162, 141,
	234, 169, 293, 292, 289, 281, 275, 274, 261, 173,
	256, 252, 100, 101, 232, 229, 227, 192, 62, 63,
	64, 65, 66, 67, 186, 264, 272, 304, 53, 144,
	180, 303, 284, 277, 268, 184, 267, 245, 10, 188,
	189, 143, 141, 87, 142, 105, 183, 199, 185, 211,
	151, 47, 191, 190, 72, 74, 205, 75, 76, 70,
	108, 147, 216, 109, 168, 262, 220, 215, 237, 238,
	222, 302, 298, 225, 5, 218, 8, 145, 233, 42,
	81, 248, 204, 165, 171, 254, 219, 239, 174, 242,
	240, 168, 226, 22, 28, 206, 202, 32, 201, 160,
	112, 80, 253, 44, 163, 4, 84, 193, 170, 41,
	17, 36, 29, 30, 31, 3, 42, 0, 107, 0,
	0, 182, 260, 0, 0, 0, 0, 0, 263, 0,
	0, 258, 194, 259, 0, 37, 0, 0, 0, 0,
	225, 208, 210, 271, 23, 27, 273, 0, 266, 34,
	0, 0, 0, 24, 25, 26, 35, 0, 33, 270,
	0, 0, 0, 0, 0, 278, 0, 0, 0, 0,
	282, 283, 0, 0, 0, 0, 0, 295, 0, 287,
	288, 0, 0, 291, 0, 0, 249, 294, 251, 0,
	0, 0, 0, 299, 56, 57, 59, 61, 71, 73,
	0, 0, 0, 0, 0, 0, 305, 306, 62, 63,
	64, 65, 66, 67, 0, 0, 68, 69, 53, 54,
	55, 0, 0, 0, 0, 0, 0, 46, 0, 286,
	58, 60, 48, 49, 50, 51, 52, 0, 0, 0,
	0, 47, 0, 280, 72, 74, 285, 75, 0, 70,
	56, 57, 59, 61, 71, 73, 0, 0, 290, 0,
	0, 0, 0, 0, 62, 63, 64, 65, 66, 67,
	0, 0, 68, 69, 53, 54, 55, 0, 0, 0,
	0, 0, 0, 46, 198, 0, 58, 60, 48, 49,
	50, 51, 52, 0, 0, 0, 0, 47, 0, 0,
	72, 74, 0, 75, 197, 70, 56, 57, 59, 61,
	71, 73, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 63, 64, 65, 66, 67, 0, 0, 68, 69,
	53, 54, 55, 0, 0, 0, 0, 0, 0, 46,
	179, 0, 58, 60, 48, 49, 50, 51, 52, 0,
	0, 0, 0, 47, 0, 0, 72, 74, 0, 75,
	178, 70, 56, 57, 59, 61, 71, 73, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 63, 64, 65,
	66, 67, 0, 0, 68, 69, 53, 54, 55, 0,
	0, 0, 0, 0, 0, 46, 0, 0, 58, 60,
	48, 49, 50, 51, 52, 0, 0, 0, 0, 47,
	0, 0, 72, 74, 301, 75, 0, 70, 56, 57,
	59, 61, 71, 73, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 63, 64, 65, 66, 67, 0, 0,
	68, 69, 53, 54, 55, 0, 0, 0, 0, 0,
	0, 46, 279, 0, 58, 60, 48, 49, 50, 51,
	52, 0, 0, 0, 0, 47, 0, 0, 72, 74,
	0, 75, 0, 70, 56, 57, 59, 61, 71, 73,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 63,
	64, 65, 66, 67, 0, 0, 68, 69, 53, 54,
	55, 0, 0, 0, 0, 0, 0, 46, 0, 0,
	58, 60, 48, 49, 50, 51, 52, 0, 276, 0,
	0, 47, 0, 0, 72, 74, 0, 75, 0, 70,
	56, 57, 59, 61, 71, 73, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 63, 64, 65, 66, 67,
	0, 0, 68, 69, 53, 54, 55, 0, 0, 0,
	0, 0, 0, 46, 0, 0, 58, 60, 48, 49,
	50, 51, 52, 0, 0, 0, 0, 47, 0, 0,
	72, 74, 0, 75, 265, 70, 56, 57, 59, 61,
	71, 73, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 63, 64, 65, 66, 67, 0, 0, 68, 69,
	53, 54, 55, 0, 0, 0, 0, 0, 0, 46,
	0, 0, 58, 60, 48, 49, 50, 51, 52, 0,
	0, 0, 0, 47, 0, 0, 72, 74, 0, 75,
	257, 70, 56, 57, 59, 61, 71, 73, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 63, 64, 65,
	66, 67, 0, 0, 68, 69, 53, 54, 55, 0,
	0, 0, 0, 0, 0, 46, 0, 0, 58, 60,
	48, 49, 50, 51, 52, 0, 0, 0, 0, 47,
	0, 0, 72, 74, 0, 75, 244, 70, 56, 57,
	59, 61, 71, 73, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 63, 64, 65, 66, 67, 0, 0,
	68, 69, 53, 54, 55, 0, 0, 0, 0, 0,
	0, 46, 0, 0, 58, 60, 48, 49, 50, 51,
	52, 0, 0, 0, 231, 47, 0, 0, 72, 74,
	0, 75, 0, 70, 56, 57, 59, 61, 71, 73,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 63,
	64, 65, 66, 67, 0, 0, 68, 69, 53, 54,
	55, 0, 0, 0, 0, 0, 0, 46, 0, 0,
	58, 60, 48, 49, 50, 51, 52, 0, 230, 0,
	0, 47, 0, 0, 72, 74, 0, 75, 0, 70,
	56, 57, 59, 61, 71, 73, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 63, 64, 65, 66, 67,
	0, 0, 68, 69, 53, 54, 55, 0, 0, 0,
	0, 0, 0, 46, 0, 0, 58, 60, 48, 49,
	50, 51, 52, 0, 228, 0, 0, 47, 0, 0,
	72, 74, 0, 75, 0, 70, 56, 57, 59, 61,
	71, 73, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 63, 64, 65, 66, 67, 0, 0, 68, 69,
	53, 54, 55, 0, 0, 0, 0, 0, 0, 46,
	0, 0, 58, 60, 48, 49, 50, 51, 52, 0,
	0, 0, 0, 47, 0, 0, 72, 74, 0, 75,
	224, 70, 56, 57, 59, 61, 71, 73, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 63, 64, 65,
	66, 67, 0, 0, 68, 69, 53, 54, 55, 0,
	0, 0, 0, 0, 0, 46, 175, 0, 58, 60,
	48, 49, 50, 51, 52, 0, 0, 0, 0, 47,
	0, 0, 72, 74, 0, 75, 0, 70, 56, 57,
	59, 61, 71, 73, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 63, 64, 65, 66, 67, 0, 0,
	68, 69, 53, 54, 55, 0, 0, 0, 0, 0,
	0, 46, 0, 0, 58, 60, 48, 49, 50, 51,
	52, 0, 172, 0, 0, 47, 0, 0, 72, 74,
	0, 75, 0, 70, 56, 57, 59, 61, 71, 73,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 63,
	64, 65, 66, 67, 0, 0, 68, 69, 53, 54,
	55, 0, 0, 0, 0, 0, 0, 46, 0, 0,
	58, 60, 48, 49, 50, 51, 52, 0, 0, 0,
	0, 47, 0, 0, 72, 74, 166, 75, 0, 70,
	56, 57, 59, 61, 71, 73, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 63, 64, 65, 66, 67,
	0, 0, 68, 69, 53, 54, 55, 0, 0, 0,
	0, 0, 0, 46, 0, 0, 58, 60, 48, 49,
	50, 51, 52, 0, 153, 0, 0, 47, 0, 0,
	72, 74, 0, 75, 0, 70, 56, 57, 59, 61,
	71, 73, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 63, 64, 65, 66, 67, 0, 0, 68, 69,
	53, 54, 55, 0, 0, 0, 0, 0, 0, 46,
	0, 0, 58, 60, 48, 49, 50, 51, 52, 0,
	150, 0, 0, 47, 0, 0, 72, 74, 0, 75,
	0, 70, 56, 57, 59, 61, 71, 73, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 63, 64, 65,
	66, 67, 0, 0, 68, 69, 53, 54, 55, 0,
	0, 0, 0, 0, 45, 46, 0, 0, 58, 60,
	48, 49, 50, 51, 52, 0, 0, 0, 0, 47,
	0, 0, 72, 74, 0, 75, 0, 70, 56, 57,
	59, 61, 71, 73, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 63, 64, 65, 66, 67, 0, 0,
	68, 69, 53, 54, 55, 0, 0, 0, 0, 0,
	0, 46, 0, 0, 58, 60, 48, 49, 50, 51,
	52, 0, 0, 0, 0, 47, 0, 0, 72, 74,
	0, 75, 0, 70, 56, 57, 59, 61, 71, 73,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 63,
	64, 65, 66, 67, 0, 0, 68, 69, 53, 54,
	55, 0, 0, 0, 0, 0, 0, 46, 0, 0,
	58, 60, 48, 49, 50, 51, 52, 0, 0, 0,
	0, 158, 0, 0, 72, 74, 0, 75, 0, 70,
	56, 57, 59, 61, 71, 73, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 63, 64, 65, 66, 67,
	0, 0, 68, 69, 53, 54, 55, 0, 0, 0,
	0, 0, 0, 46, 0, 0, 58, 60, 48, 49,
	50, 51, 52, 0, 0, 0, 0, 157, 0, 0,
	72, 74, 0, 75, 0, 70, 21, 22, 28, 0,
	0, 32, 14, 9, 15, 39, 0, 18, 0, 0,
	0, 0, 0, 0, 0, 36, 29, 30, 31, 16,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 0, 0, 0, 0, 20, 0, 0, 37,
	0, 0, 0, 0, 0, 0, 0, 0, 23, 27,
	0, 0, 0, 34, 0, 6, 0, 24, 25, 26,
	35, 0, 33, 0, 0, 7, 56, 57, 59, 61,
	0, 73, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 63, 64, 65, 66, 67, 0, 0, 68, 69,
	53, 54, 55, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 58, 60, 48, 49, 50, 51, 52, 56,
	57, 59, 61, 47, 0, 0, 72, 74, 0, 75,
	0, 70, 0, 62, 63, 64, 65, 66, 67, 0,
	0, 68, 69, 53, 54, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 58, 60, 48, 49, 50,
	51, 52, 0, 0, 0, 0, 47, 0, 0, 72,
	74, 0, 75, 0, 70, 21, 22, 187, 0, 0,
	32, 14, 9, 15, 39, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 36, 29, 30, 31, 16, 19,
	0, 0, 0, 0, 0, 0, 0, 0, 12, 13,
	0, 0, 0, 0, 0, 20, 0, 0, 37, 0,
	0, 0, 0, 0, 0, 0, 0, 23, 27, 0,
	59, 61, 34, 0, 0, 0, 24, 25, 26, 35,
	0, 33, 62, 63, 64, 65, 66, 67, 0, 0,
	68, 69, 53, 54, 55, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 60, 48, 49, 50, 51,
	52, 0, 0, 0, 0, 47, 0, 0, 72, 74,
	0, 75, 0, 70, 21, 22, 28, 0, 0, 32,
	14, 9, 15, 39, 0, 18, 0, 0, 0, 0,
	0, 0, 0, 36, 29, 30, 31, 16, 19, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 13, 0,
	0, 0, 0, 0, 20, 0, 0, 37, 0, 0,
	62, 63, 64, 65, 66, 67, 23, 27, 68, 69,
	53, 34, 0, 0, 0, 24, 25, 26, 35, 0,
	33, 0, 0, 0, 48, 49, 50, 51, 52, 21,
	22, 28, 0, 47, 32, 0, 72, 74, 0, 75,
	0, 70, 0, 0, 0, 0, 0, 0, 36, 29,
	30, 31, 0, 0, 21, 22, 28, 0, 0, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 37, 36, 29, 30, 31, 0, 0, 0,
	0, 23, 27, 21, 22, 28, 34, 0, 32, 0,
	24, 25, 26, 35, 0, 33, 243, 37, 0, 0,
	0, 0, 36, 29, 30, 31, 23, 27, 21, 22,
	28, 34, 0, 32, 0, 24, 25, 26, 35, 0,
	33, 223, 0, 0, 0, 0, 37, 36, 29, 30,
	31, 0, 0, 0, 0, 23, 27, 21, 22, 28,
	34, 0, 32, 0, 24, 25, 26, 35, 0, 33,
	200, 37, 0, 0, 0, 0, 36, 29, 30, 31,
	23, 27, 0, 0, 0, 34, 0, 0, 0, 24,
	25, 26, 35, 0, 33, 181, 0, 0, 0, 0,
	37, 0, 0, 0, 156, 0, 21, 22, 28, 23,
	27, 32, 0, 0, 34, 0, 0, 0, 24, 25,
	26, 35, 0, 33, 0, 36, 29, 30, 31, 0,
	0, 0, 0, 0, 0, 21, 22, 28, 0, 0,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 37,
	0, 0, 0, 139, 36, 29, 30, 31, 23, 27,
	226, 22, 28, 34, 0, 32, 0, 24, 25, 26,
	35, 0, 33, 0, 0, 0, 0, 0, 37, 36,
	29, 30, 31, 0, 0, 0, 0, 23, 27, 217,
	22, 28, 34, 0, 32, 0, 24, 25, 26, 35,
	0, 33, 0, 37, 0, 0, 0, 0, 36, 29,
	30, 31, 23, 27, 0, 0, 0, 34, 0, 0,
	0, 24, 25, 26, 35, 0, 33, 0, 0, 0,
	0, 0, 37, 62, 63, 64, 65, 66, 67, 0,
	0, 23, 27, 53, 0, 0, 34, 0, 0, 0,
	24, 25, 26, 35, 0, 33, 0, 0, 0, 50,
	51, 52, 96, 22, 28, 0, 47, 32, 0, 72,
	74, 0, 75, 0, 70, 0, 0, 0, 0, 0,
	0, 36, 29, 30, 31, 0, 0, 94, 22, 28,
	0, 0, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 37, 36, 29, 30, 31,
	0, 0, 0, 0, 23, 27, 83, 22, 28, 34,
	0, 32, 0, 24, 25, 26, 35, 0, 33, 0,
	37, 0, 0, 0, 0, 36, 29, 30, 31, 23,
	27, 0, 0, 0, 34, 0, 0, 0, 24, 25,
	26, 35, 0, 33, 0, 0, 0, 0, 0, 37,
	0, 0, 0, 0, 0, 0, 0, 0, 23, 27,
	0, 0, 0, 82, 0, 0, 0, 24, 25, 26,
	35, 0, 33,
}
var yyPact = [...]int{

	-55, -1000, 1810, -55, -55, -1000, -1000, -1000, -1000, 209,
	1295, 119, -1000, -1000, 2071, 2071, 207, 176, 2252, 92,
	2071, -57, -1000, 2071, 2071, 2071, 2223, 2198, -1000, -1000,
	-1000, -1000, 10, -55, -55, 2071, 7, 2, 103, 2071,
	-1000, 1532, -1000, 121, -1000, 2071, 2071, 206, 2071, 2071,
	2071, 2071, 2071, 2071, 2071, 2071, 2071, 2071, 2071, 2071,
	2071, 2071, 2071, 2071, 2071, 2071, 2071, 2071, -1000, -1000,
	2071, 2071, 2071, 2071, 2071, 2042, 2071, 102, 1351, 1351,
	90, 126, -55, 15, 26, 1239, 111, -55, 1183, 2071,
	1993, 97, 97, 97, -57, 1463, -57, 1407, 205, -23,
	2071, 187, 1127, 197, 170, -55, 1071, -1000, 2071, -55,
	1351, 1015, -1000, 2142, 2142, 97, 97, 97, 1351, 1829,
	1829, 1741, 1741, 1829, 1829, 1829, 1829, 1351, 1351, 1351,
	1351, 1351, 1351, 1351, 1589, 1351, 1632, 29, 399, 1964,
	-1000, 1351, -55, -55, 2071, -55, 72, 1701, 2071, 2071,
	-55, 2071, 65, -55, 27, 343, 1939, 204, 202, 32,
	184, 201, -29, -43, -1000, 108, -1000, -21, -1000, -31,
	197, 2125, -55, -1000, 192, 2071, -28, -1000, -1000, 1910,
	959, -1000, 2096, 64, 903, 63, -1000, 108, 847, 791,
	62, -1000, 159, 48, 133, -30, -1000, -1000, 1885, 735,
	-1000, -1000, -1000, 86, -37, 17, 183, -55, -56, -55,
	59, 2071, 191, -1000, -1000, 38, 1351, -57, 58, -1000,
	1351, -1000, 679, -1000, -1000, 1351, -57, -1000, -55, -1000,
	-55, 2071, -1000, 114, -1000, -1000, -1000, 2071, 84, -1000,
	-1000, -1000, 623, -1000, -1000, -55, 85, 83, -52, 198,
	-1000, 74, -1000, 1351, -1000, 2071, -1000, -1000, 55, 54,
	567, 82, -55, 511, -55, -1000, 53, -55, -55, 81,
	-1000, -1000, -1000, 287, -1000, -1000, -55, -55, 52, -55,
	-55, -1000, 51, 50, -55, -1000, 2071, 45, 44, 152,
	-55, -1000, -1000, -1000, 43, 455, -1000, 151, 80, -1000,
	-1000, -1000, 76, -55, -55, 42, 37, -1000, -1000,
}
var yyPgo = [...]int{

	0, 12, 225, 186, 220, 5, 2, 217, 6, 0,
	7, 10, 216, 1, 214, 4, 88, 215, 184,
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
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	16, 16, 17, 17, 18, 18,
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
	6, 5, 5, 5, 5, 4, 4, 4, 7, 9,
	0, 1, 1, 2, 1, 1,
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
	-9, -9, -9, -9, -9, -9, -9, -10, -9, 51,
	-11, -9, 52, 61, 13, 61, -1, -16, 16, 63,
	61, 49, -1, 61, -10, -9, 51, 64, 64, -15,
	4, 68, -10, -14, -13, 6, 69, -8, 4, -8,
	48, -16, 61, -11, -16, 51, 8, 69, 71, 51,
	-9, 71, -16, -1, -9, -1, 62, 6, -9, -9,
	-1, -11, 62, -7, -16, 8, 69, 71, 51, -9,
	71, 4, 4, 69, 8, -15, 4, 52, -16, 52,
	-16, 51, 64, 69, 69, -8, -9, 4, -1, 4,
	-9, 69, -9, 71, 71, -9, 4, 62, 61, 62,
	61, 63, 62, 29, 62, -6, -5, 45, 46, -6,
	-5, 69, -9, 71, 71, 61, 69, 69, 8, -16,
	71, -16, 62, -9, 4, 52, 62, 71, -1, -1,
	-9, 4, 61, -9, 51, 71, -1, 61, 61, 69,
	71, -13, 62, -9, 62, 62, 61, 61, -1, 51,
	-16, 62, -1, -1, 61, 69, 52, -1, -1, 62,
	-16, -1, 62, 62, -1, -9, 62, 62, 30, -1,
	62, 69, 30, 61, 61, -1, -1, 62, 62,
}
var yyDef = [...]int{

	-2, -2, -2, 120, 121, 122, 124, 125, 4, 39,
	-2, 0, 9, 10, 48, 0, 0, 14, 48, 0,
	0, 52, 53, 0, 0, 0, 0, 0, 61, 62,
	63, 64, 0, 120, 120, 0, 0, 0, 0, 0,
	2, -2, 123, 0, 40, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 97, 98,
	0, 0, 0, 0, 48, 0, 48, 11, 49, 12,
	0, 0, -2, 52, 0, -2, 0, -2, 0, 48,
	0, 54, 55, 56, -2, 0, -2, 0, 39, 0,
	48, 36, 0, 0, 0, 120, 0, 5, 48, 120,
	7, 0, 66, 77, 78, 79, 80, 81, 82, 83,
	84, -2, -2, 87, 88, 89, 90, 91, 92, 93,
	94, 95, 96, 99, 100, 101, 102, 0, 0, 0,
	8, -2, 120, -2, 0, -2, 0, -2, 0, 0,
	-2, 48, 0, 28, 0, 0, 0, 0, 0, 0,
	40, 39, 120, 120, 37, 0, 75, 0, 46, 0,
	0, 0, -2, 6, 0, 0, 0, 106, 108, 0,
	0, 116, 0, 0, 0, 0, 15, 61, 0, 0,
	0, 42, 0, 0, 0, 0, 104, 107, 0, 0,
	115, 58, 60, 0, 0, 0, 40, 120, 0, 120,
	0, 0, 0, 76, 117, 0, -2, -2, 0, 41,
	65, 105, 0, 114, 112, 50, -2, 13, -2, 26,
	-2, 0, 18, 0, 23, 31, 32, 0, 0, 29,
	30, 103, 0, 113, 111, -2, 0, 0, 0, 0,
	71, 0, 73, 35, 47, 0, 27, 110, 0, 0,
	0, 0, -2, 0, 120, 109, 0, -2, -2, 0,
	72, 38, 74, 0, 25, 16, -2, -2, 0, 120,
	-2, 67, 0, 0, -2, 118, 0, 0, 0, 22,
	-2, 34, 68, 69, 0, 0, 17, 21, 0, 33,
	70, 119, 0, -2, -2, 0, 0, 20, 19,
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
			yyVAL.stmts = []interpreter.Stmt{yyDollar[2].stmt}
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
			yyVAL.stmt = &statement.Variable{Names: yyDollar[2].exprIdents, Exprs: yyDollar[4].exprMany}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &statement.Variables{Left: []interpreter.Expr{yyDollar[1].expr}, Operator: "=", Right: []interpreter.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &statement.Variables{Left: yyDollar[1].exprMany, Operator: "=", Right: yyDollar[3].exprMany}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &statement.Break{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &statement.Continue{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &statement.Return{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &statement.Throw{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &statement.Module{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
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
			yyVAL.stmt = &statement.Loop{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &statement.For{Var: yyDollar[2].tok.Lit, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &statement.CFor{Expr1: yyDollar[2].exprLets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &statement.Loop{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		{
			yyVAL.stmt = &statement.Try{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		{
			yyVAL.stmt = &statement.Try{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &statement.Try{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &statement.Try{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &statement.Switch{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmtCases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &statement.Expression{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyDollar[1].stmtIf.(*statement.If).ElseIf = append(yyDollar[1].stmtIf.(*statement.If).ElseIf, &statement.If{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmtIf.SetPosition(yyDollar[1].stmtIf.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if yyVAL.stmtIf.(*statement.If).Else != nil {
				yylex.Error("multiple else statements")
			} else {
				yyVAL.stmtIf.(*statement.If).Else = append(yyVAL.stmtIf.(*statement.If).Else, yyDollar[4].compstmt...)
			}
			yyVAL.stmtIf.SetPosition(yyDollar[1].stmtIf.Position())
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmtIf = &statement.If{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmtIf.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.stmtCases = []interpreter.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmtCases = []interpreter.Stmt{yyDollar[2].stmtCase}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmtCases = []interpreter.Stmt{yyDollar[2].stmtDefault}
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
				if _, ok := stmt.(*statement.Default); ok {
					yylex.Error("multiple default statements")
				}
			}
			yyVAL.stmtCases = append(yyDollar[1].stmtCases, yyDollar[2].stmtDefault)
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmtCase = &statement.Case{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmtDefault = &statement.Default{Stmts: yyDollar[4].compstmt}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.exprPair = &expression.Pair{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprPairs = []interpreter.Expr{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprPairs = []interpreter.Expr{yyDollar[1].exprPair}
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
			yyVAL.exprLets = &expression.Vars{Left: yyDollar[1].exprMany, Operator: "=", Right: yyDollar[3].exprMany}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprMany = []interpreter.Expr{yyDollar[1].expr}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprMany = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprMany = append(yyDollar[1].exprs, &expression.Ident{Lit: yyDollar[4].tok.Lit})
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.typ = interpreter.Type{Name: yyDollar[1].tok.Lit}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.typ = interpreter.Type{Name: yyDollar[1].typ.Name + "." + yyDollar[3].tok.Lit}
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprs = nil
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []interpreter.Expr{yyDollar[1].expr}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &expression.Ident{Lit: yyDollar[4].tok.Lit})
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &expression.Ident{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			var err error
			if strings.Contains(yyDollar[1].tok.Lit, ".") || strings.Contains(yyDollar[1].tok.Lit, "e") {
				var f float64
				f, err = strconv.ParseFloat(yyDollar[1].tok.Lit, 64)
				if err == nil {
					yyVAL.expr = &expression.Number{Value: reflect.ValueOf(f)}
				}
			} else {
				var i int64
				if strings.HasPrefix(yyDollar[1].tok.Lit, "0x") {
					i, err = strconv.ParseInt(yyDollar[1].tok.Lit[2:], 16, 64)
				} else {
					i, err = strconv.ParseInt(yyDollar[1].tok.Lit, 10, 64)
				}
				if err == nil {
					yyVAL.expr = &expression.Number{Value: reflect.ValueOf(i)}
				}
			}
			if err != nil {
				tmp := &expression.Number{Value: interpreter.NilValue}
				yyVAL.expr = tmp
				tmp.Err = interpreter.NewError(yyVAL.expr, err)
			}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &expression.Unary{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &expression.Unary{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &expression.Unary{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &expression.Addr{Expr: &expression.Ident{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.Addr{Expr: &expression.Member{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &expression.Deref{Expr: &expression.Ident{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.Deref{Expr: &expression.Member{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &expression.String{Value: reflect.ValueOf(yyDollar[1].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &expression.Const{Value: interpreter.TrueValue}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &expression.Const{Value: interpreter.FalseValue}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &expression.Const{Value: reflect.ValueOf(nil)}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &expression.TernaryOp{Expr: yyDollar[1].expr, Left: yyDollar[3].expr, Right: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.Member{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.expr = &expression.Func{Args: yyDollar[3].exprIdents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.expr = &expression.Func{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.expr = &expression.Func{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].exprIdents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.expr = &expression.Func{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &expression.Array{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.expr = &expression.Array{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			mapExpr := make(map[string]interpreter.Expr)
			for _, v := range yyDollar[3].exprPairs {
				mapExpr[v.(*expression.Pair).Key] = v.(*expression.Pair).Value
			}
			yyVAL.expr = &expression.Map{Map: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 74:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			mapExpr := make(map[string]interpreter.Expr)
			for _, v := range yyDollar[3].exprPairs {
				mapExpr[v.(*expression.Pair).Key] = v.(*expression.Pair).Value
			}
			yyVAL.expr = &expression.Map{Map: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.Paren{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.New{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "**", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "<<", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: ">>", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.Assoc{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.Assoc{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.Assoc{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.Assoc{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.Assoc{Left: yyDollar[1].expr, Operator: "&=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.Assoc{Left: yyDollar[1].expr, Operator: "|=", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &expression.Assoc{Left: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &expression.Assoc{Left: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "|", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "&", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &expression.BinOp{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &expression.Call{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.Call{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &expression.AnonCall{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.AnonCall{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.Item{Value: &expression.Ident{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.Item{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.expr = &expression.Slice{Value: &expression.Ident{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.expr = &expression.Slice{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &expression.Slice{Value: &expression.Ident{Lit: yyDollar[1].tok.Lit}, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &expression.Slice{Value: yyDollar[1].expr, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &expression.Slice{Value: &expression.Ident{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &expression.Slice{Value: yyDollar[1].expr, Begin: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.Slice{Value: &expression.Ident{Lit: yyDollar[1].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.Slice{Value: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &expression.Make{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.expr = &expression.MakeArray{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.expr = &expression.MakeArray{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	}
	goto yystack /* stack new state and value */
}
