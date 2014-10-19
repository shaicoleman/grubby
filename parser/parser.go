//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:3
import (
	"github.com/grubby/grubby/ast"
	"strings"
)

var Statements []ast.Node

//line parser.y:16
type RubySymType struct {
	yys          int
	operator     string
	genericValue ast.Node
	genericSlice ast.Nodes
	stringSlice  []string
}

const OPERATOR = 57346
const NODE = 57347
const REF = 57348
const SPECIAL_CHAR_REF = 57349
const CAPITAL_REF = 57350
const LPAREN = 57351
const RPAREN = 57352
const COMMA = 57353
const DO = 57354
const DEF = 57355
const END = 57356
const IF = 57357
const ELSE = 57358
const ELSIF = 57359
const UNLESS = 57360
const CLASS = 57361
const MODULE = 57362
const FOR = 57363
const WHILE = 57364
const UNTIL = 57365
const BEGIN = 57366
const RESCUE = 57367
const ENSURE = 57368
const BREAK = 57369
const NEXT = 57370
const REDO = 57371
const RETRY = 57372
const RETURN = 57373
const YIELD = 57374
const AND = 57375
const OR = 57376
const LAMBDA = 57377
const TRUE = 57378
const FALSE = 57379
const LESSTHAN = 57380
const GREATERTHAN = 57381
const EQUALTO = 57382
const BANG = 57383
const COMPLEMENT = 57384
const POSITIVE = 57385
const NEGATIVE = 57386
const STAR = 57387
const WHITESPACE = 57388
const NEWLINE = 57389
const SEMICOLON = 57390
const COLON = 57391
const DOUBLECOLON = 57392
const DOT = 57393
const PIPE = 57394
const SLASH = 57395
const AMPERSAND = 57396
const QUESTIONMARK = 57397
const CARET = 57398
const LBRACKET = 57399
const RBRACKET = 57400
const LBRACE = 57401
const RBRACE = 57402
const DOLLARSIGN = 57403
const ATSIGN = 57404
const FILE_CONST_REF = 57405
const EOF = 57406

var RubyToknames = []string{
	"OPERATOR",
	"NODE",
	"REF",
	"SPECIAL_CHAR_REF",
	"CAPITAL_REF",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"DO",
	"DEF",
	"END",
	"IF",
	"ELSE",
	"ELSIF",
	"UNLESS",
	"CLASS",
	"MODULE",
	"FOR",
	"WHILE",
	"UNTIL",
	"BEGIN",
	"RESCUE",
	"ENSURE",
	"BREAK",
	"NEXT",
	"REDO",
	"RETRY",
	"RETURN",
	"YIELD",
	"AND",
	"OR",
	"LAMBDA",
	"TRUE",
	"FALSE",
	"LESSTHAN",
	"GREATERTHAN",
	"EQUALTO",
	"BANG",
	"COMPLEMENT",
	"POSITIVE",
	"NEGATIVE",
	"STAR",
	"WHITESPACE",
	"NEWLINE",
	"SEMICOLON",
	"COLON",
	"DOUBLECOLON",
	"DOT",
	"PIPE",
	"SLASH",
	"AMPERSAND",
	"QUESTIONMARK",
	"CARET",
	"LBRACKET",
	"RBRACKET",
	"LBRACE",
	"RBRACE",
	"DOLLARSIGN",
	"ATSIGN",
	"FILE_CONST_REF",
	"EOF",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:1028

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	50, 115,
	-2, 20,
	-1, 88,
	10, 77,
	11, 77,
	-2, 176,
	-1, 98,
	50, 115,
	-2, 20,
	-1, 104,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	22, 13,
	24, 13,
	32, 13,
	41, 13,
	42, 13,
	43, 13,
	44, 13,
	48, 13,
	-2, 11,
	-1, 119,
	50, 115,
	-2, 113,
	-1, 220,
	50, 116,
	-2, 114,
	-1, 227,
	10, 77,
	11, 77,
	-2, 176,
	-1, 351,
	27, 192,
	28, 192,
	-2, 13,
	-1, 352,
	27, 192,
	28, 192,
	-2, 13,
}

const RubyNprod = 218
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2133

var RubyAct = []int{

	9, 138, 368, 35, 262, 322, 190, 253, 20, 188,
	277, 90, 89, 139, 187, 215, 174, 25, 97, 66,
	98, 88, 96, 93, 103, 2, 3, 95, 306, 173,
	72, 115, 133, 72, 134, 143, 358, 142, 111, 112,
	215, 272, 4, 245, 109, 118, 120, 68, 61, 62,
	108, 110, 92, 107, 305, 275, 87, 196, 131, 130,
	76, 103, 70, 69, 140, 351, 352, 377, 72, 91,
	76, 104, 141, 60, 59, 150, 151, 152, 153, 71,
	155, 156, 157, 158, 213, 160, 161, 162, 135, 165,
	215, 146, 167, 168, 74, 75, 274, 141, 95, 208,
	141, 76, 164, 304, 74, 75, 165, 73, 147, 72,
	122, 84, 182, 183, 141, 228, 347, 73, 341, 176,
	172, 84, 72, 339, 208, 25, 97, 66, 98, 67,
	205, 214, 194, 72, 344, 74, 75, 345, 206, 121,
	5, 201, 202, 218, 105, 165, 212, 166, 73, 215,
	95, 215, 84, 141, 350, 68, 61, 62, 221, 342,
	292, 225, 226, 258, 215, 209, 336, 230, 267, 337,
	260, 233, 205, 191, 260, 106, 205, 63, 119, 64,
	266, 60, 59, 123, 124, 125, 126, 127, 128, 231,
	243, 191, 334, 234, 189, 335, 205, 132, 205, 244,
	250, 205, 74, 75, 383, 357, 380, 379, 170, 261,
	220, 148, 192, 259, 263, 73, 72, 154, 105, 84,
	95, 193, 159, 375, 144, 242, 163, 72, 165, 271,
	192, 378, 248, 380, 379, 349, 141, 205, 331, 193,
	205, 270, 311, 213, 313, 177, 178, 179, 180, 210,
	278, 211, 184, 185, 264, 265, 247, 205, 205, 281,
	198, 65, 76, 296, 289, 246, 279, 243, 308, 310,
	113, 269, 213, 114, 329, 94, 285, 284, 317, 241,
	213, 217, 317, 205, 222, 229, 213, 283, 205, 285,
	284, 238, 205, 111, 112, 199, 74, 75, 136, 186,
	137, 77, 78, 79, 307, 117, 110, 116, 181, 73,
	82, 80, 81, 84, 169, 314, 232, 149, 237, 324,
	207, 216, 1, 320, 131, 346, 343, 145, 205, 205,
	56, 205, 55, 54, 53, 52, 51, 17, 16, 15,
	25, 219, 66, 98, 67, 14, 43, 340, 298, 205,
	299, 22, 23, 24, 319, 362, 363, 364, 12, 366,
	11, 321, 205, 296, 296, 296, 372, 21, 10, 19,
	68, 61, 62, 13, 18, 382, 36, 39, 38, 33,
	32, 34, 215, 296, 280, 387, 388, 31, 296, 296,
	296, 389, 63, 286, 64, 0, 60, 59, 0, 0,
	0, 297, 0, 374, 376, 0, 309, 0, 0, 0,
	0, 312, 0, 384, 0, 385, 318, 0, 0, 0,
	318, 0, 0, 326, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 25, 26, 66, 27, 67, 0, 332,
	333, 40, 0, 48, 338, 0, 49, 41, 42, 0,
	58, 0, 50, 0, 0, 28, 0, 0, 0, 0,
	57, 0, 0, 68, 61, 62, 0, 0, 0, 44,
	45, 46, 47, 353, 354, 355, 356, 0, 0, 0,
	0, 0, 99, 359, 360, 63, 0, 64, 0, 60,
	59, 0, 0, 0, 0, 0, 365, 0, 0, 0,
	0, 297, 297, 297, 0, 0, 0, 0, 381, 0,
	0, 0, 0, 99, 0, 0, 0, 0, 386, 99,
	0, 297, 0, 0, 0, 0, 297, 297, 297, 0,
	99, 99, 99, 99, 0, 99, 99, 99, 99, 0,
	99, 99, 99, 0, 99, 76, 0, 99, 99, 0,
	0, 0, 0, 99, 37, 0, 0, 0, 0, 0,
	0, 99, 0, 0, 0, 0, 0, 99, 99, 0,
	0, 25, 97, 66, 98, 88, 0, 0, 103, 74,
	75, 102, 0, 0, 77, 78, 79, 0, 0, 0,
	0, 0, 73, 82, 80, 81, 84, 0, 99, 0,
	99, 68, 61, 62, 0, 99, 0, 0, 0, 0,
	0, 0, 102, 0, 0, 268, 0, 0, 102, 0,
	0, 0, 99, 91, 0, 104, 99, 60, 59, 102,
	102, 102, 102, 0, 102, 102, 102, 102, 0, 102,
	102, 102, 0, 102, 0, 0, 102, 102, 0, 0,
	0, 0, 102, 29, 0, 0, 0, 0, 0, 0,
	102, 0, 0, 0, 99, 0, 102, 102, 0, 99,
	0, 0, 0, 0, 0, 99, 0, 0, 0, 0,
	100, 0, 0, 99, 99, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 102, 0, 102,
	0, 0, 0, 0, 102, 0, 0, 0, 0, 0,
	0, 100, 0, 0, 0, 0, 0, 100, 0, 0,
	0, 102, 0, 0, 99, 102, 0, 0, 100, 100,
	100, 100, 0, 100, 100, 100, 100, 0, 100, 100,
	100, 0, 100, 0, 0, 100, 100, 0, 0, 0,
	0, 100, 0, 30, 0, 0, 0, 0, 0, 100,
	0, 0, 0, 102, 0, 100, 100, 175, 102, 0,
	0, 0, 0, 0, 102, 0, 0, 0, 0, 99,
	101, 0, 102, 102, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 100, 0, 100, 0,
	0, 0, 0, 100, 0, 0, 0, 0, 0, 0,
	0, 101, 0, 0, 99, 0, 0, 101, 129, 0,
	100, 0, 0, 102, 100, 0, 0, 0, 101, 101,
	101, 101, 0, 101, 101, 101, 101, 0, 101, 101,
	101, 0, 101, 0, 0, 101, 101, 0, 0, 0,
	0, 101, 0, 25, 97, 66, 98, 67, 0, 101,
	103, 0, 100, 0, 0, 101, 101, 100, 0, 0,
	0, 171, 0, 100, 0, 0, 0, 0, 102, 0,
	0, 100, 100, 68, 61, 62, 195, 0, 197, 0,
	0, 0, 0, 0, 0, 200, 101, 0, 101, 0,
	0, 0, 0, 101, 0, 63, 0, 104, 0, 60,
	59, 0, 0, 102, 0, 0, 0, 0, 0, 0,
	101, 0, 100, 0, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 25, 97, 66, 98, 227, 0,
	236, 103, 239, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 101, 0, 68, 61, 62, 101, 0, 256,
	257, 0, 0, 101, 0, 0, 0, 100, 0, 0,
	0, 101, 101, 0, 0, 0, 63, 0, 104, 0,
	60, 59, 0, 0, 25, 26, 66, 27, 67, 0,
	0, 0, 40, 371, 300, 370, 369, 301, 41, 42,
	0, 58, 100, 50, 0, 0, 302, 303, 282, 0,
	0, 57, 101, 287, 68, 61, 62, 0, 291, 0,
	44, 45, 46, 47, 0, 76, 294, 295, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 0, 64, 83,
	60, 59, 0, 327, 328, 0, 0, 0, 0, 0,
	330, 0, 0, 0, 85, 86, 0, 0, 0, 74,
	75, 0, 0, 0, 77, 78, 79, 101, 0, 0,
	0, 0, 73, 82, 80, 81, 84, 0, 0, 0,
	0, 0, 0, 0, 348, 0, 0, 0, 0, 0,
	200, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 101, 0, 0, 0, 0, 361, 0, 256,
	257, 25, 26, 66, 27, 67, 0, 0, 0, 40,
	367, 300, 370, 369, 301, 41, 42, 0, 58, 0,
	50, 0, 0, 302, 303, 0, 0, 0, 57, 0,
	0, 68, 61, 62, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 294, 295, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 0, 64, 0, 60, 59, 25,
	26, 66, 27, 67, 0, 0, 0, 40, 373, 300,
	0, 0, 301, 41, 42, 0, 58, 0, 50, 0,
	0, 302, 303, 0, 0, 0, 57, 0, 0, 68,
	61, 62, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 294, 295, 0, 0, 0, 0, 0, 0, 0,
	0, 63, 0, 64, 0, 60, 59, 25, 26, 66,
	27, 67, 0, 0, 0, 40, 293, 300, 0, 0,
	301, 41, 42, 0, 58, 0, 50, 0, 0, 302,
	303, 0, 0, 0, 57, 0, 0, 68, 61, 62,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 294,
	295, 0, 0, 0, 0, 0, 0, 0, 0, 63,
	0, 64, 0, 60, 59, 25, 26, 66, 27, 67,
	0, 0, 0, 40, 288, 48, 255, 254, 49, 41,
	42, 0, 58, 0, 50, 0, 0, 0, 0, 0,
	0, 0, 57, 0, 0, 68, 61, 62, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 203, 204, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 0, 64,
	0, 60, 59, 25, 26, 66, 27, 67, 0, 0,
	0, 40, 252, 48, 255, 254, 49, 41, 42, 0,
	58, 0, 50, 0, 0, 0, 0, 0, 0, 0,
	57, 0, 0, 68, 61, 62, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 203, 204, 0, 25, 26,
	66, 27, 67, 0, 0, 63, 40, 64, 48, 60,
	59, 49, 41, 42, 0, 58, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 57, 0, 0, 68, 61,
	62, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 0, 64, 0, 60, 59, 0, 8, 25, 26,
	66, 27, 67, 0, 0, 0, 40, 0, 300, 0,
	0, 301, 41, 42, 0, 58, 0, 50, 0, 0,
	302, 303, 0, 0, 0, 57, 0, 0, 68, 61,
	62, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	294, 295, 0, 25, 26, 66, 27, 67, 0, 0,
	63, 40, 64, 48, 60, 59, 49, 41, 42, 0,
	58, 0, 50, 260, 0, 0, 0, 0, 0, 323,
	57, 0, 0, 68, 61, 62, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 315, 316, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 0, 64, 0, 60,
	59, 25, 26, 66, 27, 67, 0, 0, 0, 40,
	325, 48, 0, 0, 49, 41, 42, 0, 58, 0,
	50, 0, 0, 0, 0, 0, 0, 0, 57, 0,
	0, 68, 61, 62, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 203, 204, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 0, 64, 0, 60, 59, 25,
	26, 66, 27, 67, 0, 0, 0, 40, 290, 48,
	0, 0, 49, 41, 42, 0, 58, 0, 50, 0,
	0, 0, 0, 0, 0, 0, 57, 0, 0, 68,
	61, 62, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 203, 204, 0, 25, 26, 66, 27, 67, 0,
	0, 63, 40, 64, 48, 60, 59, 49, 41, 42,
	0, 58, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 57, 0, 0, 68, 61, 62, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 203, 204, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 0, 64, 276,
	60, 59, 25, 26, 66, 27, 67, 0, 0, 0,
	40, 273, 48, 0, 0, 49, 41, 42, 0, 58,
	0, 50, 0, 0, 0, 0, 0, 0, 0, 57,
	0, 0, 68, 61, 62, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 203, 204, 0, 0, 0, 0,
	0, 0, 0, 0, 63, 0, 64, 0, 60, 59,
	25, 26, 66, 27, 67, 0, 0, 0, 40, 251,
	48, 0, 0, 49, 41, 42, 0, 58, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 57, 0, 0,
	68, 61, 62, 0, 0, 0, 44, 45, 46, 47,
	0, 0, 203, 204, 0, 0, 0, 0, 0, 0,
	0, 0, 63, 0, 64, 0, 60, 59, 25, 26,
	66, 27, 67, 0, 0, 0, 40, 249, 48, 0,
	0, 49, 41, 42, 0, 58, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 57, 0, 0, 68, 61,
	62, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	203, 204, 0, 25, 26, 66, 27, 67, 0, 0,
	63, 40, 64, 48, 60, 59, 49, 41, 42, 0,
	58, 0, 50, 0, 0, 0, 0, 0, 0, 0,
	57, 0, 0, 68, 61, 62, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 203, 204, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 0, 64, 240, 60,
	59, 25, 26, 66, 27, 67, 0, 0, 0, 40,
	235, 48, 0, 0, 49, 41, 42, 0, 58, 0,
	50, 0, 0, 0, 0, 0, 0, 0, 57, 0,
	0, 68, 61, 62, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 203, 204, 0, 25, 26, 66, 27,
	67, 0, 0, 63, 40, 64, 48, 60, 59, 49,
	41, 42, 0, 58, 0, 50, 0, 0, 0, 0,
	0, 0, 0, 57, 0, 0, 68, 61, 62, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 203, 204,
	0, 25, 26, 66, 27, 67, 224, 0, 63, 40,
	64, 48, 60, 59, 49, 41, 42, 0, 58, 0,
	50, 0, 0, 0, 0, 0, 0, 0, 57, 0,
	0, 68, 61, 62, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 0, 223, 25, 97, 66, 98, 88,
	0, 0, 103, 63, 0, 64, 0, 60, 59, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 61, 62, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 91, 0, 104,
	0, 60, 59,
}
var RubyPact = []int{

	-22, 1393, -1000, -1000, -1000, 15, -1000, -1000, -1000, 1031,
	-1000, -1000, -1000, 38, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 12, 135, 13, 10,
	4, -1000, -1000, -1000, -1000, -1000, -1000, 255, -20, -1000,
	301, 170, 170, 99, 428, 428, 428, 428, 428, 428,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 120, 428, 26,
	292, -1000, -1000, 120, -1000, -15, 215, -1000, 49, -1000,
	-1000, -1000, 428, 311, 120, 120, 120, 120, 428, 120,
	120, 120, 120, 428, 120, 120, 120, 428, 120, -1000,
	136, 120, 120, 308, 197, 97, -1000, 2070, 209, -1000,
	-1000, -1000, 0, -23, -23, 120, 428, 428, 428, 428,
	302, 120, 120, 428, 428, 293, 185, 185, 19, -1000,
	-1000, 428, 289, 53, 53, 53, 53, 53, 94, 1981,
	113, 97, 118, -1000, -1000, 243, -1000, -1000, 88, 73,
	541, -1000, 335, 202, 120, 2026, -1000, -23, 53, 929,
	97, 97, 97, 97, 53, 97, 97, 97, 97, 53,
	66, 97, 97, 53, 275, 541, 848, 258, 97, -1000,
	848, 1936, -1000, 285, -1000, 1878, 269, 53, 53, 53,
	53, -1000, 97, 97, 53, 53, -1000, -1000, 179, 167,
	-1000, 3, 259, 250, -1000, 1833, 170, 1775, 53, -1000,
	1348, -1000, -1000, -1000, -1000, 1031, 53, 149, 120, -1000,
	-1000, -1000, -1000, 120, -1000, -1000, -1000, 169, 164, 566,
	-1000, 261, 53, -1000, -1000, 136, -1000, 120, 120, -1000,
	97, -1000, 1, 97, -1000, -1000, 1717, 44, -1000, 1659,
	-1000, -1000, -7, 167, 256, 428, -1000, -1000, -7, -1000,
	-1000, -1000, -1000, 273, 428, -1000, 1290, 1614, -1000, -1000,
	152, 97, 1232, 97, 43, -32, -1000, 428, 120, -1000,
	232, 97, 428, -1000, -1000, 238, -1000, 1498, -1000, -1000,
	53, 1498, 1556, -1000, 428, -1000, 53, 1981, -1000, 260,
	-1000, 1981, 234, -1000, -1000, -1000, 1031, 53, -1000, -1000,
	428, 428, 177, 151, -1000, 428, -1000, 117, 1031, 53,
	97, -1000, 53, -1000, 104, -1000, -1000, 1031, 53, -1000,
	145, 119, -1000, 120, 102, -1000, 53, 1981, 1981, -1000,
	1981, 229, 107, 18, 428, 428, 428, 428, 201, -13,
	-7, -1000, -1000, -1000, 428, 428, 113, -1000, 1981, -1000,
	-1000, -1000, -1000, 53, 53, 53, 53, 428, 120, 53,
	53, 1981, 1116, 989, 1174, 212, 56, -1000, 217, 428,
	-1000, -1000, 190, -1000, -7, -1000, -7, -1000, -1000, 428,
	-1000, 53, 1453, -1000, -7, -7, 53, 1453, 1453, 1453,
}
var RubyPgo = []int{

	0, 138, 387, 381, 22, 380, 379, 378, 753, 377,
	5, 376, 374, 373, 369, 0, 653, 554, 368, 367,
	361, 8, 360, 6, 455, 358, 3, 354, 353, 352,
	351, 350, 348, 2, 346, 345, 339, 338, 337, 336,
	335, 334, 333, 332, 330, 767, 327, 323, 12, 16,
	7, 322, 14, 321, 4, 320, 13, 10, 318, 1,
	281, 275, 11, 9, 261, 37,
}
var RubyR1 = []int{

	0, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	51, 65, 65, 45, 45, 45, 45, 45, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 21, 21, 21, 21, 21, 21, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 48, 48, 59, 59, 56, 56, 56,
	56, 62, 62, 62, 62, 61, 61, 61, 18, 18,
	27, 27, 27, 27, 57, 57, 57, 57, 57, 57,
	57, 52, 52, 23, 23, 23, 23, 63, 63, 63,
	22, 22, 25, 26, 26, 64, 64, 13, 13, 13,
	13, 13, 13, 8, 8, 24, 24, 16, 16, 34,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 2, 5, 6, 6, 3, 3, 53, 53, 53,
	53, 60, 60, 60, 4, 4, 4, 4, 49, 58,
	58, 58, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 50, 50, 50, 50, 46, 46, 46, 7,
	14, 10, 10, 10, 55, 55, 47, 47, 19, 20,
	11, 30, 54, 54, 54, 54, 54, 54, 54, 31,
	31, 31, 31, 31, 31, 31, 32, 32, 32, 32,
	32, 33, 33, 33, 33, 29, 28, 9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 4, 1, 4, 4, 2,
	3, 3, 4, 4, 3, 2, 3, 3, 3, 3,
	3, 4, 6, 3, 1, 1, 3, 0, 1, 1,
	3, 1, 1, 3, 3, 1, 3, 3, 7, 7,
	0, 1, 3, 3, 0, 2, 2, 2, 2, 2,
	3, 1, 3, 1, 2, 3, 2, 0, 1, 3,
	4, 6, 4, 1, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 4, 7,
	8, 3, 7, 8, 3, 4, 4, 3, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 3, 5,
	6, 5, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 2, 3, 5, 0, 2, 1, 2, 2, 2,
	5, 5, 0, 2, 2, 2, 2, 2, 2, 0,
	1, 3, 3, 1, 3, 3, 5, 6, 5, 6,
	5, 4, 3, 3, 2, 3, 3, 2,
}
var RubyChk = []int{

	-1000, -51, 47, 48, 64, -1, 47, 48, 64, -15,
	-18, -22, -25, -13, -35, -36, -37, -38, -12, -14,
	-21, -19, -30, -29, -28, 5, 6, 8, -24, -16,
	-8, -2, -5, -6, -3, -26, -11, -17, -7, -9,
	13, 19, 20, -34, 41, 42, 43, 44, 15, 18,
	24, -39, -40, -41, -42, -43, -44, 32, 22, 62,
	61, 36, 37, 57, 59, -64, 7, 9, 35, 48,
	47, 64, 15, 51, 38, 39, 4, 43, 44, 45,
	53, 54, 52, 18, 55, 33, 34, 18, 9, -48,
	-62, 57, 40, 11, -61, -15, -4, 6, 8, -24,
	-16, -8, -17, 12, 59, 9, 40, 40, 40, 40,
	51, 38, 39, 15, 18, 51, 6, 4, -26, 8,
	-26, 40, 11, -1, -1, -1, -1, -1, -1, -45,
	-59, -15, -1, 6, 8, 62, 6, 8, -59, -56,
	-15, -21, -65, 50, 9, -46, -4, 59, -1, 6,
	-15, -15, -15, -15, -1, -15, -15, -15, -15, -1,
	-15, -15, -15, -1, -56, -15, 11, -15, -15, 6,
	11, -45, -49, 52, -49, -45, -56, -1, -1, -1,
	-1, 6, -15, -15, -1, -1, 6, -52, -63, 9,
	-23, 6, 45, 54, -52, -45, 38, -45, -1, 6,
	-45, 47, 48, 47, 48, -15, -1, -55, 11, 47,
	6, 8, 58, 11, 58, 47, -53, -60, -15, 6,
	8, -56, -1, 48, 10, -62, -48, 9, 49, 10,
	-15, -4, 58, -15, -4, 14, -45, -58, 6, -45,
	60, 10, -65, 11, -63, 40, 6, 6, -65, 14,
	-26, 14, 14, -50, 17, 16, -45, -45, 14, -10,
	25, -15, -54, -15, -65, -65, 11, 4, 49, 10,
	-56, -15, 40, 14, 52, 11, 60, -57, -23, 10,
	-1, -57, -45, 14, 17, 16, -1, -45, 14, -50,
	14, -45, 8, 14, 47, 48, -15, -1, -32, -31,
	15, 18, 27, 28, 60, 11, 60, -65, -15, -1,
	-15, 10, -1, 6, -65, 47, 48, -15, -1, -27,
	-47, -20, -10, 31, -65, 14, -1, -45, -45, 14,
	-45, 4, -1, -1, 15, 18, 15, 18, -1, 6,
	-65, 14, 14, -10, 15, 18, -59, 14, -45, 6,
	47, 47, 48, -1, -1, -1, -1, 4, 49, -1,
	-1, -45, -54, -54, -54, -1, -15, 14, -33, 17,
	16, 14, -33, 14, -65, 11, -65, 11, 14, 17,
	16, -1, -54, 14, -65, -65, -1, -54, -54, -54,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 18, 19, -2, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 33, 34, 35, 36, 37, 38, 0, 0, 0,
	0, 141, 142, 77, 11, 0, 56, 176, 0, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, -2, 59,
	65, 77, 0, 0, 74, 81, 82, 19, -2, 21,
	22, 23, 30, 13, -2, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 107, 107, 13, -2,
	13, 0, 0, 131, 132, 133, 134, 13, 0, 184,
	188, 75, 0, 125, 126, 0, 123, 124, 0, 0,
	75, 79, 147, 0, 77, 0, 217, 13, 164, 60,
	66, 68, 70, 135, 136, 137, 138, 139, 140, 166,
	0, 215, 216, 168, 0, 78, 0, 75, 117, 129,
	0, 0, 13, 159, 13, 0, 0, 118, 119, 120,
	121, 61, 67, 69, 165, 167, 64, 11, 101, 107,
	108, 103, 0, 0, 11, 0, 0, 0, 122, 130,
	0, 13, 13, 14, 15, 16, 17, 0, 0, 192,
	127, 128, 143, 0, 144, 12, 11, 11, 0, 19,
	-2, 0, 177, 178, 179, 62, 63, -2, 0, 55,
	83, 84, 71, 86, 87, 154, 0, 0, 160, 0,
	157, 58, 94, 0, 0, 0, 104, 106, 94, 110,
	13, 112, 162, 0, 0, 13, 0, 0, 180, 185,
	13, 76, 0, 80, 0, 0, 11, 0, 0, 57,
	0, 190, 0, 155, 158, 0, 156, 11, 109, 102,
	105, 11, 0, 163, 0, 13, 13, 175, 169, 0,
	171, 181, 13, 191, 193, 194, 39, 196, 197, 198,
	0, 0, 200, 203, 145, 0, 146, 0, 39, 11,
	151, 73, 72, 161, 0, 95, 96, 39, 98, 99,
	0, 91, 186, 0, 0, 111, 13, 173, 174, 170,
	182, 0, 13, 0, 0, 0, 0, 0, 0, 0,
	148, 88, 100, 187, 0, 0, 189, 89, 172, 13,
	192, -2, -2, 201, 202, 204, 205, 0, 0, 92,
	93, 183, 0, 0, 0, 11, 11, 206, 0, 0,
	192, 208, 0, 210, 149, 11, 152, 11, 207, 0,
	192, 192, 199, 209, 150, 153, 192, 199, 199, 199,
}
var RubyTok1 = []int{

	1,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64,
}
var RubyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var RubyDebug = 0

type RubyLexer interface {
	Lex(lval *RubySymType) int
	Error(s string)
}

const RubyFlag = -1000

func RubyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(RubyToknames) {
		if RubyToknames[c-4] != "" {
			return RubyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func RubyStatname(s int) string {
	if s >= 0 && s < len(RubyStatenames) {
		if RubyStatenames[s] != "" {
			return RubyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Rubylex1(lex RubyLexer, lval *RubySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = RubyTok1[0]
		goto out
	}
	if char < len(RubyTok1) {
		c = RubyTok1[char]
		goto out
	}
	if char >= RubyPrivate {
		if char < RubyPrivate+len(RubyTok2) {
			c = RubyTok2[char-RubyPrivate]
			goto out
		}
	}
	for i := 0; i < len(RubyTok3); i += 2 {
		c = RubyTok3[i+0]
		if c == char {
			c = RubyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = RubyTok2[1] /* unknown char */
	}
	if RubyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", RubyTokname(c), uint(char))
	}
	return c
}

func RubyParse(Rubylex RubyLexer) int {
	var Rubyn int
	var Rubylval RubySymType
	var RubyVAL RubySymType
	RubyS := make([]RubySymType, RubyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Rubystate := 0
	Rubychar := -1
	Rubyp := -1
	goto Rubystack

ret0:
	return 0

ret1:
	return 1

Rubystack:
	/* put a state and value onto the stack */
	if RubyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", RubyTokname(Rubychar), RubyStatname(Rubystate))
	}

	Rubyp++
	if Rubyp >= len(RubyS) {
		nyys := make([]RubySymType, len(RubyS)*2)
		copy(nyys, RubyS)
		RubyS = nyys
	}
	RubyS[Rubyp] = RubyVAL
	RubyS[Rubyp].yys = Rubystate

Rubynewstate:
	Rubyn = RubyPact[Rubystate]
	if Rubyn <= RubyFlag {
		goto Rubydefault /* simple state */
	}
	if Rubychar < 0 {
		Rubychar = Rubylex1(Rubylex, &Rubylval)
	}
	Rubyn += Rubychar
	if Rubyn < 0 || Rubyn >= RubyLast {
		goto Rubydefault
	}
	Rubyn = RubyAct[Rubyn]
	if RubyChk[Rubyn] == Rubychar { /* valid shift */
		Rubychar = -1
		RubyVAL = Rubylval
		Rubystate = Rubyn
		if Errflag > 0 {
			Errflag--
		}
		goto Rubystack
	}

Rubydefault:
	/* default state action */
	Rubyn = RubyDef[Rubystate]
	if Rubyn == -2 {
		if Rubychar < 0 {
			Rubychar = Rubylex1(Rubylex, &Rubylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if RubyExca[xi+0] == -1 && RubyExca[xi+1] == Rubystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Rubyn = RubyExca[xi+0]
			if Rubyn < 0 || Rubyn == Rubychar {
				break
			}
		}
		Rubyn = RubyExca[xi+1]
		if Rubyn < 0 {
			goto ret0
		}
	}
	if Rubyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Rubylex.Error("syntax error")
			Nerrs++
			if RubyDebug >= 1 {
				__yyfmt__.Printf("%s", RubyStatname(Rubystate))
				__yyfmt__.Printf(" saw %s\n", RubyTokname(Rubychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Rubyp >= 0 {
				Rubyn = RubyPact[RubyS[Rubyp].yys] + RubyErrCode
				if Rubyn >= 0 && Rubyn < RubyLast {
					Rubystate = RubyAct[Rubyn] /* simulate a shift of "error" */
					if RubyChk[Rubystate] == RubyErrCode {
						goto Rubystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if RubyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", RubyS[Rubyp].yys)
				}
				Rubyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if RubyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", RubyTokname(Rubychar))
			}
			if Rubychar == RubyEofCode {
				goto ret1
			}
			Rubychar = -1
			goto Rubynewstate /* try again in the same state */
		}
	}

	/* reduction by production Rubyn */
	if RubyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Rubyn, RubyStatname(Rubystate))
	}

	Rubynt := Rubyn
	Rubypt := Rubyp
	_ = Rubypt // guard against "declared and not used"

	Rubyp -= RubyR2[Rubyn]
	RubyVAL = RubyS[Rubyp+1]

	/* consult goto table to find next state */
	Rubyn = RubyR1[Rubyn]
	Rubyg := RubyPgo[Rubyn]
	Rubyj := Rubyg + RubyS[Rubyp].yys + 1

	if Rubyj >= RubyLast {
		Rubystate = RubyAct[Rubyg]
	} else {
		Rubystate = RubyAct[Rubyj]
		if RubyChk[Rubystate] != -Rubyn {
			Rubystate = RubyAct[Rubyg]
		}
	}
	// dummy call; replaced with literal code
	switch Rubynt {

	case 1:
		//line parser.y:184
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:186
		{
		}
	case 3:
		//line parser.y:188
		{
		}
	case 4:
		//line parser.y:190
		{
		}
	case 5:
		//line parser.y:192
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:194
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:196
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:202
		{
		}
	case 11:
		//line parser.y:204
		{
		}
	case 12:
		//line parser.y:205
		{
		}
	case 13:
		//line parser.y:208
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:210
		{
		}
	case 15:
		//line parser.y:212
		{
		}
	case 16:
		//line parser.y:214
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:216
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 18:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 19:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 20:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 21:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 22:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 23:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 24:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 25:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 26:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 27:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 28:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 29:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 30:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 31:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 32:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 33:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 34:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 35:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 36:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 37:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 38:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 39:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 40:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 41:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 42:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 43:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 44:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 45:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 46:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 47:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 48:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 49:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 50:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 51:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 52:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 53:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 54:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 55:
		//line parser.y:226
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 56:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 57:
		//line parser.y:237
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 58:
		//line parser.y:244
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 59:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 60:
		//line parser.y:258
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 61:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 62:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 63:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 64:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 65:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 66:
		//line parser.y:305
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:313
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:321
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:329
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:337
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:347
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 72:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 73:
		//line parser.y:366
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 74:
		//line parser.y:368
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 75:
		//line parser.y:371
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:373
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:375
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 78:
		//line parser.y:377
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:379
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:381
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:385
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:387
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:389
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:391
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:394
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:396
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:398
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 89:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 90:
		//line parser.y:420
		{
		}
	case 91:
		//line parser.y:422
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 92:
		//line parser.y:424
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 93:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 94:
		//line parser.y:434
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 95:
		//line parser.y:436
		{
		}
	case 96:
		//line parser.y:438
		{
		}
	case 97:
		//line parser.y:440
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:442
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:444
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:446
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericSlice...)
		}
	case 101:
		//line parser.y:449
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 102:
		//line parser.y:451
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 103:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 104:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 105:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 106:
		//line parser.y:460
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 107:
		//line parser.y:462
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 108:
		//line parser.y:464
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:468
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:473
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 111:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 112:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 113:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 114:
		//line parser.y:505
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 115:
		//line parser.y:513
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 116:
		//line parser.y:517
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 117:
		//line parser.y:522
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 118:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 119:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 120:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 121:
		//line parser.y:550
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 122:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 123:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 124:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 125:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 126:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 127:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 128:
		//line parser.y:577
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 129:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 130:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 131:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 132:
		//line parser.y:585
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 133:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 134:
		//line parser.y:587
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 135:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 136:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 137:
		//line parser.y:608
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 138:
		//line parser.y:617
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 139:
		//line parser.y:626
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 140:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 141:
		//line parser.y:643
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 142:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 143:
		//line parser.y:646
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 144:
		//line parser.y:648
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 145:
		//line parser.y:651
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 146:
		//line parser.y:659
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 147:
		//line parser.y:667
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 148:
		//line parser.y:669
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 149:
		//line parser.y:676
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 150:
		//line parser.y:683
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 151:
		//line parser.y:691
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 152:
		//line parser.y:698
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 153:
		//line parser.y:705
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 154:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 155:
		//line parser.y:715
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 156:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 157:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 158:
		//line parser.y:732
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 159:
		//line parser.y:734
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 160:
		//line parser.y:736
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 161:
		//line parser.y:738
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 162:
		//line parser.y:741
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 163:
		//line parser.y:748
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 164:
		//line parser.y:756
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 165:
		//line parser.y:763
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 166:
		//line parser.y:770
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 167:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 168:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 169:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 170:
		//line parser.y:798
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 171:
		//line parser.y:806
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 172:
		//line parser.y:815
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 173:
		//line parser.y:822
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 174:
		//line parser.y:829
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 175:
		//line parser.y:836
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 176:
		//line parser.y:843
		{
		}
	case 177:
		//line parser.y:844
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 178:
		//line parser.y:845
		{
		}
	case 179:
		//line parser.y:848
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 180:
		//line parser.y:851
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 181:
		//line parser.y:859
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 182:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 183:
		//line parser.y:870
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}

			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Var:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
					Class: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				},
			}
		}
	case 184:
		//line parser.y:885
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 185:
		//line parser.y:887
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 186:
		//line parser.y:890
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 187:
		//line parser.y:892
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 188:
		//line parser.y:895
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 189:
		//line parser.y:904
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 190:
		//line parser.y:913
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 191:
		//line parser.y:922
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 192:
		//line parser.y:925
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 193:
		//line parser.y:927
		{
		}
	case 194:
		//line parser.y:929
		{
		}
	case 195:
		//line parser.y:931
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 196:
		//line parser.y:933
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 197:
		//line parser.y:935
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 198:
		//line parser.y:937
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 199:
		//line parser.y:939
		{
		}
	case 200:
		//line parser.y:941
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 201:
		//line parser.y:943
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 202:
		//line parser.y:945
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 203:
		//line parser.y:947
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 204:
		//line parser.y:949
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 205:
		//line parser.y:951
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 206:
		//line parser.y:954
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 207:
		//line parser.y:961
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 208:
		//line parser.y:969
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 209:
		//line parser.y:976
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 210:
		//line parser.y:984
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 211:
		//line parser.y:992
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 212:
		//line parser.y:999
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 213:
		//line parser.y:1006
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 214:
		//line parser.y:1013
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 215:
		//line parser.y:1021
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 216:
		//line parser.y:1024
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 217:
		//line parser.y:1026
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	}
	goto Rubystack /* stack new state and value */
}
