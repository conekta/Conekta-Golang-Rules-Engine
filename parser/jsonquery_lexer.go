// Code generated from JsonQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JsonQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var JsonQueryLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonquerylexerLexerInit() {
	staticData := &JsonQueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'pr'", "'.'", "'-'", "'['", "']'", "", "", "", "'null'",
		"", "", "", "", "", "", "", "", "", "", "'MLP'", "'SUM'", "'SUBTRACT'",
		"'DIV'", "", "", "", "", "", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "NOT", "LOGICAL_OPERATOR", "BOOLEAN",
		"NULL", "IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW", "EW",
		"ASTERISK", "PLUS", "MINUS", "DIVISON", "ATTRNAME", "VERSION", "STRING",
		"DOUBLE", "INT", "EXP", "NEWLINE", "COMMA", "SP",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "NOT", "LOGICAL_OPERATOR",
		"BOOLEAN", "NULL", "IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW",
		"EW", "ASTERISK", "PLUS", "MINUS", "DIVISON", "ATTRNAME", "ATTR_NAME_CHAR",
		"DIGIT", "ALPHA", "VERSION", "STRING", "ESC", "UNICODE", "HEX", "DOUBLE",
		"INT", "EXP", "NEWLINE", "COMMA", "SP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 309, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 103, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 110, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 3, 9, 121, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 3, 11, 132, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 140, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		148, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 155, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 162, 8, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 170, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 3, 17, 178, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 184, 8, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 190, 8, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 196, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 5, 25, 221, 8, 25, 10, 25, 12, 25,
		224, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 230, 8, 26, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		30, 5, 30, 245, 8, 30, 10, 30, 12, 30, 248, 9, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 3, 31, 255, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 34, 3, 34, 266, 8, 34, 1, 34, 1, 34, 1, 34, 4, 34,
		271, 8, 34, 11, 34, 12, 34, 272, 1, 34, 3, 34, 276, 8, 34, 1, 35, 1, 35,
		1, 35, 5, 35, 281, 8, 35, 10, 35, 12, 35, 284, 9, 35, 3, 35, 286, 8, 35,
		1, 36, 1, 36, 3, 36, 290, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1,
		38, 5, 38, 298, 8, 38, 10, 38, 12, 38, 301, 9, 38, 1, 39, 1, 39, 5, 39,
		305, 8, 39, 10, 39, 12, 39, 308, 9, 39, 0, 0, 40, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 0, 55, 0, 57, 0, 59, 27, 61, 28, 63, 0,
		65, 0, 67, 0, 69, 29, 71, 30, 73, 31, 75, 32, 77, 33, 79, 34, 1, 0, 9,
		3, 0, 45, 45, 58, 58, 95, 95, 2, 0, 65, 90, 97, 122, 2, 0, 34, 34, 92,
		92, 8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114,
		116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 48, 57, 1, 0, 49, 57, 2,
		0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 336, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 1, 81, 1, 0, 0, 0, 3, 83, 1, 0, 0, 0,
		5, 85, 1, 0, 0, 0, 7, 88, 1, 0, 0, 0, 9, 90, 1, 0, 0, 0, 11, 92, 1, 0,
		0, 0, 13, 94, 1, 0, 0, 0, 15, 102, 1, 0, 0, 0, 17, 109, 1, 0, 0, 0, 19,
		120, 1, 0, 0, 0, 21, 122, 1, 0, 0, 0, 23, 131, 1, 0, 0, 0, 25, 139, 1,
		0, 0, 0, 27, 147, 1, 0, 0, 0, 29, 154, 1, 0, 0, 0, 31, 161, 1, 0, 0, 0,
		33, 169, 1, 0, 0, 0, 35, 177, 1, 0, 0, 0, 37, 183, 1, 0, 0, 0, 39, 189,
		1, 0, 0, 0, 41, 195, 1, 0, 0, 0, 43, 197, 1, 0, 0, 0, 45, 201, 1, 0, 0,
		0, 47, 205, 1, 0, 0, 0, 49, 214, 1, 0, 0, 0, 51, 218, 1, 0, 0, 0, 53, 229,
		1, 0, 0, 0, 55, 231, 1, 0, 0, 0, 57, 233, 1, 0, 0, 0, 59, 235, 1, 0, 0,
		0, 61, 241, 1, 0, 0, 0, 63, 251, 1, 0, 0, 0, 65, 256, 1, 0, 0, 0, 67, 262,
		1, 0, 0, 0, 69, 265, 1, 0, 0, 0, 71, 285, 1, 0, 0, 0, 73, 287, 1, 0, 0,
		0, 75, 293, 1, 0, 0, 0, 77, 295, 1, 0, 0, 0, 79, 302, 1, 0, 0, 0, 81, 82,
		5, 40, 0, 0, 82, 2, 1, 0, 0, 0, 83, 84, 5, 41, 0, 0, 84, 4, 1, 0, 0, 0,
		85, 86, 5, 112, 0, 0, 86, 87, 5, 114, 0, 0, 87, 6, 1, 0, 0, 0, 88, 89,
		5, 46, 0, 0, 89, 8, 1, 0, 0, 0, 90, 91, 5, 45, 0, 0, 91, 10, 1, 0, 0, 0,
		92, 93, 5, 91, 0, 0, 93, 12, 1, 0, 0, 0, 94, 95, 5, 93, 0, 0, 95, 14, 1,
		0, 0, 0, 96, 97, 5, 110, 0, 0, 97, 98, 5, 111, 0, 0, 98, 103, 5, 116, 0,
		0, 99, 100, 5, 78, 0, 0, 100, 101, 5, 79, 0, 0, 101, 103, 5, 84, 0, 0,
		102, 96, 1, 0, 0, 0, 102, 99, 1, 0, 0, 0, 103, 16, 1, 0, 0, 0, 104, 105,
		5, 97, 0, 0, 105, 106, 5, 110, 0, 0, 106, 110, 5, 100, 0, 0, 107, 108,
		5, 111, 0, 0, 108, 110, 5, 114, 0, 0, 109, 104, 1, 0, 0, 0, 109, 107, 1,
		0, 0, 0, 110, 18, 1, 0, 0, 0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 114,
		0, 0, 113, 114, 5, 117, 0, 0, 114, 121, 5, 101, 0, 0, 115, 116, 5, 102,
		0, 0, 116, 117, 5, 97, 0, 0, 117, 118, 5, 108, 0, 0, 118, 119, 5, 115,
		0, 0, 119, 121, 5, 101, 0, 0, 120, 111, 1, 0, 0, 0, 120, 115, 1, 0, 0,
		0, 121, 20, 1, 0, 0, 0, 122, 123, 5, 110, 0, 0, 123, 124, 5, 117, 0, 0,
		124, 125, 5, 108, 0, 0, 125, 126, 5, 108, 0, 0, 126, 22, 1, 0, 0, 0, 127,
		128, 5, 73, 0, 0, 128, 132, 5, 78, 0, 0, 129, 130, 5, 105, 0, 0, 130, 132,
		5, 110, 0, 0, 131, 127, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 24, 1, 0,
		0, 0, 133, 134, 5, 101, 0, 0, 134, 140, 5, 113, 0, 0, 135, 136, 5, 69,
		0, 0, 136, 140, 5, 81, 0, 0, 137, 138, 5, 61, 0, 0, 138, 140, 5, 61, 0,
		0, 139, 133, 1, 0, 0, 0, 139, 135, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140,
		26, 1, 0, 0, 0, 141, 142, 5, 110, 0, 0, 142, 148, 5, 101, 0, 0, 143, 144,
		5, 78, 0, 0, 144, 148, 5, 69, 0, 0, 145, 146, 5, 33, 0, 0, 146, 148, 5,
		61, 0, 0, 147, 141, 1, 0, 0, 0, 147, 143, 1, 0, 0, 0, 147, 145, 1, 0, 0,
		0, 148, 28, 1, 0, 0, 0, 149, 150, 5, 103, 0, 0, 150, 155, 5, 116, 0, 0,
		151, 152, 5, 71, 0, 0, 152, 155, 5, 84, 0, 0, 153, 155, 5, 62, 0, 0, 154,
		149, 1, 0, 0, 0, 154, 151, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 30, 1,
		0, 0, 0, 156, 157, 5, 108, 0, 0, 157, 162, 5, 116, 0, 0, 158, 159, 5, 76,
		0, 0, 159, 162, 5, 84, 0, 0, 160, 162, 5, 60, 0, 0, 161, 156, 1, 0, 0,
		0, 161, 158, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 32, 1, 0, 0, 0, 163,
		164, 5, 103, 0, 0, 164, 170, 5, 101, 0, 0, 165, 166, 5, 71, 0, 0, 166,
		170, 5, 69, 0, 0, 167, 168, 5, 62, 0, 0, 168, 170, 5, 61, 0, 0, 169, 163,
		1, 0, 0, 0, 169, 165, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 34, 1, 0,
		0, 0, 171, 172, 5, 108, 0, 0, 172, 178, 5, 101, 0, 0, 173, 174, 5, 76,
		0, 0, 174, 178, 5, 69, 0, 0, 175, 176, 5, 60, 0, 0, 176, 178, 5, 61, 0,
		0, 177, 171, 1, 0, 0, 0, 177, 173, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178,
		36, 1, 0, 0, 0, 179, 180, 5, 99, 0, 0, 180, 184, 5, 111, 0, 0, 181, 182,
		5, 67, 0, 0, 182, 184, 5, 79, 0, 0, 183, 179, 1, 0, 0, 0, 183, 181, 1,
		0, 0, 0, 184, 38, 1, 0, 0, 0, 185, 186, 5, 115, 0, 0, 186, 190, 5, 119,
		0, 0, 187, 188, 5, 83, 0, 0, 188, 190, 5, 87, 0, 0, 189, 185, 1, 0, 0,
		0, 189, 187, 1, 0, 0, 0, 190, 40, 1, 0, 0, 0, 191, 192, 5, 101, 0, 0, 192,
		196, 5, 119, 0, 0, 193, 194, 5, 69, 0, 0, 194, 196, 5, 87, 0, 0, 195, 191,
		1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 42, 1, 0, 0, 0, 197, 198, 5, 77,
		0, 0, 198, 199, 5, 76, 0, 0, 199, 200, 5, 80, 0, 0, 200, 44, 1, 0, 0, 0,
		201, 202, 5, 83, 0, 0, 202, 203, 5, 85, 0, 0, 203, 204, 5, 77, 0, 0, 204,
		46, 1, 0, 0, 0, 205, 206, 5, 83, 0, 0, 206, 207, 5, 85, 0, 0, 207, 208,
		5, 66, 0, 0, 208, 209, 5, 84, 0, 0, 209, 210, 5, 82, 0, 0, 210, 211, 5,
		65, 0, 0, 211, 212, 5, 67, 0, 0, 212, 213, 5, 84, 0, 0, 213, 48, 1, 0,
		0, 0, 214, 215, 5, 68, 0, 0, 215, 216, 5, 73, 0, 0, 216, 217, 5, 86, 0,
		0, 217, 50, 1, 0, 0, 0, 218, 222, 3, 57, 28, 0, 219, 221, 3, 53, 26, 0,
		220, 219, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222,
		223, 1, 0, 0, 0, 223, 52, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 230, 7,
		0, 0, 0, 226, 230, 3, 55, 27, 0, 227, 230, 3, 57, 28, 0, 228, 230, 3, 77,
		38, 0, 229, 225, 1, 0, 0, 0, 229, 226, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0,
		229, 228, 1, 0, 0, 0, 230, 54, 1, 0, 0, 0, 231, 232, 2, 48, 57, 0, 232,
		56, 1, 0, 0, 0, 233, 234, 7, 1, 0, 0, 234, 58, 1, 0, 0, 0, 235, 236, 3,
		71, 35, 0, 236, 237, 5, 46, 0, 0, 237, 238, 3, 71, 35, 0, 238, 239, 5,
		46, 0, 0, 239, 240, 3, 71, 35, 0, 240, 60, 1, 0, 0, 0, 241, 246, 5, 34,
		0, 0, 242, 245, 3, 63, 31, 0, 243, 245, 8, 2, 0, 0, 244, 242, 1, 0, 0,
		0, 244, 243, 1, 0, 0, 0, 245, 248, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246,
		247, 1, 0, 0, 0, 247, 249, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 249, 250,
		5, 34, 0, 0, 250, 62, 1, 0, 0, 0, 251, 254, 5, 92, 0, 0, 252, 255, 7, 3,
		0, 0, 253, 255, 3, 65, 32, 0, 254, 252, 1, 0, 0, 0, 254, 253, 1, 0, 0,
		0, 255, 64, 1, 0, 0, 0, 256, 257, 5, 117, 0, 0, 257, 258, 3, 67, 33, 0,
		258, 259, 3, 67, 33, 0, 259, 260, 3, 67, 33, 0, 260, 261, 3, 67, 33, 0,
		261, 66, 1, 0, 0, 0, 262, 263, 7, 4, 0, 0, 263, 68, 1, 0, 0, 0, 264, 266,
		5, 45, 0, 0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1, 0,
		0, 0, 267, 268, 3, 71, 35, 0, 268, 270, 5, 46, 0, 0, 269, 271, 7, 5, 0,
		0, 270, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272,
		273, 1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 276, 3, 73, 36, 0, 275, 274,
		1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 70, 1, 0, 0, 0, 277, 286, 5, 48,
		0, 0, 278, 282, 7, 6, 0, 0, 279, 281, 7, 5, 0, 0, 280, 279, 1, 0, 0, 0,
		281, 284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283,
		286, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 285, 277, 1, 0, 0, 0, 285, 278,
		1, 0, 0, 0, 286, 72, 1, 0, 0, 0, 287, 289, 7, 7, 0, 0, 288, 290, 7, 8,
		0, 0, 289, 288, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0,
		291, 292, 3, 71, 35, 0, 292, 74, 1, 0, 0, 0, 293, 294, 5, 10, 0, 0, 294,
		76, 1, 0, 0, 0, 295, 299, 5, 44, 0, 0, 296, 298, 5, 32, 0, 0, 297, 296,
		1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0,
		0, 0, 300, 78, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302, 306, 5, 32, 0, 0,
		303, 305, 3, 75, 37, 0, 304, 303, 1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306,
		304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 80, 1, 0, 0, 0, 308, 306, 1,
		0, 0, 0, 27, 0, 102, 109, 120, 131, 139, 147, 154, 161, 169, 177, 183,
		189, 195, 222, 229, 244, 246, 254, 265, 272, 275, 282, 285, 289, 299, 306,
		0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JsonQueryLexerInit initializes any static state used to implement JsonQueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonQueryLexerInit() {
	staticData := &JsonQueryLexerLexerStaticData
	staticData.once.Do(jsonquerylexerLexerInit)
}

// NewJsonQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonQueryLexer(input antlr.CharStream) *JsonQueryLexer {
	JsonQueryLexerInit()
	l := new(JsonQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &JsonQueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "JsonQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonQueryLexer tokens.
const (
	JsonQueryLexerT__0             = 1
	JsonQueryLexerT__1             = 2
	JsonQueryLexerT__2             = 3
	JsonQueryLexerT__3             = 4
	JsonQueryLexerT__4             = 5
	JsonQueryLexerT__5             = 6
	JsonQueryLexerT__6             = 7
	JsonQueryLexerNOT              = 8
	JsonQueryLexerLOGICAL_OPERATOR = 9
	JsonQueryLexerBOOLEAN          = 10
	JsonQueryLexerNULL             = 11
	JsonQueryLexerIN               = 12
	JsonQueryLexerEQ               = 13
	JsonQueryLexerNE               = 14
	JsonQueryLexerGT               = 15
	JsonQueryLexerLT               = 16
	JsonQueryLexerGE               = 17
	JsonQueryLexerLE               = 18
	JsonQueryLexerCO               = 19
	JsonQueryLexerSW               = 20
	JsonQueryLexerEW               = 21
	JsonQueryLexerASTERISK         = 22
	JsonQueryLexerPLUS             = 23
	JsonQueryLexerMINUS            = 24
	JsonQueryLexerDIVISON          = 25
	JsonQueryLexerATTRNAME         = 26
	JsonQueryLexerVERSION          = 27
	JsonQueryLexerSTRING           = 28
	JsonQueryLexerDOUBLE           = 29
	JsonQueryLexerINT              = 30
	JsonQueryLexerEXP              = 31
	JsonQueryLexerNEWLINE          = 32
	JsonQueryLexerCOMMA            = 33
	JsonQueryLexerSP               = 34
)
