// Code generated from JsonQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // JsonQuery

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JsonQueryParser struct {
	*antlr.BaseParser
}

var JsonQueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonqueryParserInit() {
	staticData := &JsonQueryParserStaticData
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
		"query", "attrPathValue", "subAttrValue", "attrPath", "subAttr", "value",
		"listStrings", "listDoubles", "listAttrPaths", "listInts",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 148, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 3, 0, 23, 8, 0, 1, 0, 3, 0, 26, 8, 0, 1, 0, 1, 0, 3, 0, 30, 8, 0, 1,
		0, 1, 0, 3, 0, 34, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 64, 8, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 5, 0, 71, 8, 0, 10, 0, 12, 0, 74, 9, 0, 1, 1, 1, 1,
		3, 1, 78, 8, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 85, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 96, 8, 5, 1, 5, 1, 5,
		3, 5, 100, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 105, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 5, 6, 111, 8, 6, 10, 6, 12, 6, 114, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 122, 8, 7, 10, 7, 12, 7, 125, 9, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 5, 8, 132, 8, 8, 10, 8, 12, 8, 135, 9, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 5, 9, 141, 8, 9, 10, 9, 12, 9, 144, 9, 9, 1, 9, 1, 9, 1, 9, 0, 1, 0,
		10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 3, 1, 0, 12, 21, 1, 0, 22, 25,
		1, 0, 13, 18, 162, 0, 63, 1, 0, 0, 0, 2, 75, 1, 0, 0, 0, 4, 79, 1, 0, 0,
		0, 6, 82, 1, 0, 0, 0, 8, 86, 1, 0, 0, 0, 10, 104, 1, 0, 0, 0, 12, 106,
		1, 0, 0, 0, 14, 117, 1, 0, 0, 0, 16, 128, 1, 0, 0, 0, 18, 136, 1, 0, 0,
		0, 20, 22, 6, 0, -1, 0, 21, 23, 5, 8, 0, 0, 22, 21, 1, 0, 0, 0, 22, 23,
		1, 0, 0, 0, 23, 25, 1, 0, 0, 0, 24, 26, 5, 34, 0, 0, 25, 24, 1, 0, 0, 0,
		25, 26, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 29, 5, 1, 0, 0, 28, 30, 5,
		34, 0, 0, 29, 28, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31,
		33, 3, 0, 0, 0, 32, 34, 5, 34, 0, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1, 0,
		0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 5, 2, 0, 0, 36, 64, 1, 0, 0, 0, 37, 38,
		3, 6, 3, 0, 38, 39, 5, 34, 0, 0, 39, 40, 5, 3, 0, 0, 40, 64, 1, 0, 0, 0,
		41, 42, 3, 6, 3, 0, 42, 43, 5, 34, 0, 0, 43, 44, 7, 0, 0, 0, 44, 45, 5,
		34, 0, 0, 45, 46, 3, 10, 5, 0, 46, 64, 1, 0, 0, 0, 47, 48, 3, 6, 3, 0,
		48, 49, 5, 34, 0, 0, 49, 50, 7, 0, 0, 0, 50, 51, 5, 34, 0, 0, 51, 52, 3,
		2, 1, 0, 52, 64, 1, 0, 0, 0, 53, 54, 7, 1, 0, 0, 54, 55, 5, 34, 0, 0, 55,
		56, 5, 1, 0, 0, 56, 57, 3, 16, 8, 0, 57, 58, 5, 2, 0, 0, 58, 59, 5, 34,
		0, 0, 59, 60, 7, 2, 0, 0, 60, 61, 5, 34, 0, 0, 61, 62, 3, 10, 5, 0, 62,
		64, 1, 0, 0, 0, 63, 20, 1, 0, 0, 0, 63, 37, 1, 0, 0, 0, 63, 41, 1, 0, 0,
		0, 63, 47, 1, 0, 0, 0, 63, 53, 1, 0, 0, 0, 64, 72, 1, 0, 0, 0, 65, 66,
		10, 5, 0, 0, 66, 67, 5, 34, 0, 0, 67, 68, 5, 9, 0, 0, 68, 69, 5, 34, 0,
		0, 69, 71, 3, 0, 0, 6, 70, 65, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70,
		1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 1, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0,
		75, 77, 5, 26, 0, 0, 76, 78, 3, 4, 2, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1,
		0, 0, 0, 78, 3, 1, 0, 0, 0, 79, 80, 5, 4, 0, 0, 80, 81, 3, 2, 1, 0, 81,
		5, 1, 0, 0, 0, 82, 84, 5, 26, 0, 0, 83, 85, 3, 8, 4, 0, 84, 83, 1, 0, 0,
		0, 84, 85, 1, 0, 0, 0, 85, 7, 1, 0, 0, 0, 86, 87, 5, 4, 0, 0, 87, 88, 3,
		6, 3, 0, 88, 9, 1, 0, 0, 0, 89, 105, 5, 10, 0, 0, 90, 105, 5, 11, 0, 0,
		91, 105, 5, 27, 0, 0, 92, 105, 5, 28, 0, 0, 93, 105, 5, 29, 0, 0, 94, 96,
		5, 5, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 99, 5, 30, 0, 0, 98, 100, 5, 31, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100,
		1, 0, 0, 0, 100, 105, 1, 0, 0, 0, 101, 105, 3, 18, 9, 0, 102, 105, 3, 14,
		7, 0, 103, 105, 3, 12, 6, 0, 104, 89, 1, 0, 0, 0, 104, 90, 1, 0, 0, 0,
		104, 91, 1, 0, 0, 0, 104, 92, 1, 0, 0, 0, 104, 93, 1, 0, 0, 0, 104, 95,
		1, 0, 0, 0, 104, 101, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 103, 1, 0,
		0, 0, 105, 11, 1, 0, 0, 0, 106, 107, 5, 6, 0, 0, 107, 112, 5, 28, 0, 0,
		108, 109, 5, 33, 0, 0, 109, 111, 5, 28, 0, 0, 110, 108, 1, 0, 0, 0, 111,
		114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 115,
		1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 7, 0, 0, 116, 13, 1, 0,
		0, 0, 117, 118, 5, 6, 0, 0, 118, 123, 5, 29, 0, 0, 119, 120, 5, 33, 0,
		0, 120, 122, 5, 29, 0, 0, 121, 119, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123,
		121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 123,
		1, 0, 0, 0, 126, 127, 5, 7, 0, 0, 127, 15, 1, 0, 0, 0, 128, 133, 3, 6,
		3, 0, 129, 130, 5, 33, 0, 0, 130, 132, 3, 6, 3, 0, 131, 129, 1, 0, 0, 0,
		132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134,
		17, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 137, 5, 6, 0, 0, 137, 142, 5,
		30, 0, 0, 138, 139, 5, 33, 0, 0, 139, 141, 5, 30, 0, 0, 140, 138, 1, 0,
		0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0,
		143, 145, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 145, 146, 5, 7, 0, 0, 146,
		19, 1, 0, 0, 0, 15, 22, 25, 29, 33, 63, 72, 77, 84, 95, 99, 104, 112, 123,
		133, 142,
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

// JsonQueryParserInit initializes any static state used to implement JsonQueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonQueryParserInit() {
	staticData := &JsonQueryParserStaticData
	staticData.once.Do(jsonqueryParserInit)
}

// NewJsonQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonQueryParser(input antlr.TokenStream) *JsonQueryParser {
	JsonQueryParserInit()
	this := new(JsonQueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JsonQueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JsonQuery.g4"

	return this
}

// JsonQueryParser tokens.
const (
	JsonQueryParserEOF              = antlr.TokenEOF
	JsonQueryParserT__0             = 1
	JsonQueryParserT__1             = 2
	JsonQueryParserT__2             = 3
	JsonQueryParserT__3             = 4
	JsonQueryParserT__4             = 5
	JsonQueryParserT__5             = 6
	JsonQueryParserT__6             = 7
	JsonQueryParserNOT              = 8
	JsonQueryParserLOGICAL_OPERATOR = 9
	JsonQueryParserBOOLEAN          = 10
	JsonQueryParserNULL             = 11
	JsonQueryParserIN               = 12
	JsonQueryParserEQ               = 13
	JsonQueryParserNE               = 14
	JsonQueryParserGT               = 15
	JsonQueryParserLT               = 16
	JsonQueryParserGE               = 17
	JsonQueryParserLE               = 18
	JsonQueryParserCO               = 19
	JsonQueryParserSW               = 20
	JsonQueryParserEW               = 21
	JsonQueryParserASTERISK         = 22
	JsonQueryParserPLUS             = 23
	JsonQueryParserMINUS            = 24
	JsonQueryParserDIVISON          = 25
	JsonQueryParserATTRNAME         = 26
	JsonQueryParserVERSION          = 27
	JsonQueryParserSTRING           = 28
	JsonQueryParserDOUBLE           = 29
	JsonQueryParserINT              = 30
	JsonQueryParserEXP              = 31
	JsonQueryParserNEWLINE          = 32
	JsonQueryParserCOMMA            = 33
	JsonQueryParserSP               = 34
)

// JsonQueryParser rules.
const (
	JsonQueryParserRULE_query         = 0
	JsonQueryParserRULE_attrPathValue = 1
	JsonQueryParserRULE_subAttrValue  = 2
	JsonQueryParserRULE_attrPath      = 3
	JsonQueryParserRULE_subAttr       = 4
	JsonQueryParserRULE_value         = 5
	JsonQueryParserRULE_listStrings   = 6
	JsonQueryParserRULE_listDoubles   = 7
	JsonQueryParserRULE_listAttrPaths = 8
	JsonQueryParserRULE_listInts      = 9
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyAll(ctx *QueryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	QueryContext
	op antlr.Token
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *CompareExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *CompareExpContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CompareExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *CompareExpContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *CompareExpContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGT, 0)
}

func (s *CompareExpContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLT, 0)
}

func (s *CompareExpContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGE, 0)
}

func (s *CompareExpContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLE, 0)
}

func (s *CompareExpContext) CO() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCO, 0)
}

func (s *CompareExpContext) SW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSW, 0)
}

func (s *CompareExpContext) EW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEW, 0)
}

func (s *CompareExpContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIN, 0)
}

func (s *CompareExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitCompareExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulSumExpContext struct {
	QueryContext
	op antlr.Token
}

func NewMulSumExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulSumExpContext {
	var p = new(MulSumExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *MulSumExpContext) GetOp() antlr.Token { return s.op }

func (s *MulSumExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulSumExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulSumExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *MulSumExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *MulSumExpContext) ListAttrPaths() IListAttrPathsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAttrPathsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListAttrPathsContext)
}

func (s *MulSumExpContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MulSumExpContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserASTERISK, 0)
}

func (s *MulSumExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserPLUS, 0)
}

func (s *MulSumExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserMINUS, 0)
}

func (s *MulSumExpContext) DIVISON() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDIVISON, 0)
}

func (s *MulSumExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *MulSumExpContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *MulSumExpContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGT, 0)
}

func (s *MulSumExpContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLT, 0)
}

func (s *MulSumExpContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGE, 0)
}

func (s *MulSumExpContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLE, 0)
}

func (s *MulSumExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitMulSumExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	QueryContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *ParenExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNOT, 0)
}

func (s *ParenExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *ParenExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PresentExpContext struct {
	QueryContext
}

func NewPresentExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PresentExpContext {
	var p = new(PresentExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *PresentExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresentExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *PresentExpContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, 0)
}

func (s *PresentExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitPresentExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalExpContext struct {
	QueryContext
}

func NewLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpContext {
	var p = new(LogicalExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *LogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *LogicalExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *LogicalExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *LogicalExpContext) LOGICAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLOGICAL_OPERATOR, 0)
}

func (s *LogicalExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLogicalExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareExpAttrPathContext struct {
	QueryContext
	op antlr.Token
}

func NewCompareExpAttrPathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpAttrPathContext {
	var p = new(CompareExpAttrPathContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *CompareExpAttrPathContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpAttrPathContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpAttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpAttrPathContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpAttrPathContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *CompareExpAttrPathContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *CompareExpAttrPathContext) AttrPathValue() IAttrPathValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathValueContext)
}

func (s *CompareExpAttrPathContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *CompareExpAttrPathContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *CompareExpAttrPathContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGT, 0)
}

func (s *CompareExpAttrPathContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLT, 0)
}

func (s *CompareExpAttrPathContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGE, 0)
}

func (s *CompareExpAttrPathContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLE, 0)
}

func (s *CompareExpAttrPathContext) CO() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCO, 0)
}

func (s *CompareExpAttrPathContext) SW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSW, 0)
}

func (s *CompareExpAttrPathContext) EW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEW, 0)
}

func (s *CompareExpAttrPathContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIN, 0)
}

func (s *CompareExpAttrPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitCompareExpAttrPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *JsonQueryParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, JsonQueryParserRULE_query, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserNOT {
			{
				p.SetState(21)
				p.Match(JsonQueryParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(24)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(27)
			p.Match(JsonQueryParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(28)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(31)
			p.query(0)
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(32)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(35)
			p.Match(JsonQueryParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewPresentExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.AttrPath()
		}
		{
			p.SetState(38)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(JsonQueryParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			p.AttrPath()
		}
		{
			p.SetState(42)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4190208) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(44)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Value()
		}

	case 4:
		localctx = NewCompareExpAttrPathContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.AttrPath()
		}
		{
			p.SetState(48)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpAttrPathContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4190208) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpAttrPathContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(50)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.AttrPathValue()
		}

	case 5:
		localctx = NewMulSumExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62914560) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(54)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(JsonQueryParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.ListAttrPaths()
		}
		{
			p.SetState(57)
			p.Match(JsonQueryParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MulSumExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&516096) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MulSumExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(60)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, JsonQueryParserRULE_query)
			p.SetState(65)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				goto errorExit
			}
			{
				p.SetState(66)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(67)
				p.Match(JsonQueryParserLOGICAL_OPERATOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(68)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(69)
				p.query(6)
			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttrPathValueContext is an interface to support dynamic dispatch.
type IAttrPathValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRNAME() antlr.TerminalNode
	SubAttrValue() ISubAttrValueContext

	// IsAttrPathValueContext differentiates from other interfaces.
	IsAttrPathValueContext()
}

type AttrPathValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathValueContext() *AttrPathValueContext {
	var p = new(AttrPathValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPathValue
	return p
}

func InitEmptyAttrPathValueContext(p *AttrPathValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPathValue
}

func (*AttrPathValueContext) IsAttrPathValueContext() {}

func NewAttrPathValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathValueContext {
	var p = new(AttrPathValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_attrPathValue

	return p
}

func (s *AttrPathValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathValueContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserATTRNAME, 0)
}

func (s *AttrPathValueContext) SubAttrValue() ISubAttrValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubAttrValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubAttrValueContext)
}

func (s *AttrPathValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitAttrPathValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) AttrPathValue() (localctx IAttrPathValueContext) {
	localctx = NewAttrPathValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonQueryParserRULE_attrPathValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(JsonQueryParserATTRNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(76)
			p.SubAttrValue()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubAttrValueContext is an interface to support dynamic dispatch.
type ISubAttrValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrPathValue() IAttrPathValueContext

	// IsSubAttrValueContext differentiates from other interfaces.
	IsSubAttrValueContext()
}

type SubAttrValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrValueContext() *SubAttrValueContext {
	var p = new(SubAttrValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttrValue
	return p
}

func InitEmptySubAttrValueContext(p *SubAttrValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttrValue
}

func (*SubAttrValueContext) IsSubAttrValueContext() {}

func NewSubAttrValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrValueContext {
	var p = new(SubAttrValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subAttrValue

	return p
}

func (s *SubAttrValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrValueContext) AttrPathValue() IAttrPathValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathValueContext)
}

func (s *SubAttrValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubAttrValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubAttrValue() (localctx ISubAttrValueContext) {
	localctx = NewSubAttrValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonQueryParserRULE_subAttrValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(JsonQueryParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.AttrPathValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRNAME() antlr.TerminalNode
	SubAttr() ISubAttrContext

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
	return p
}

func InitEmptyAttrPathContext(p *AttrPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserATTRNAME, 0)
}

func (s *AttrPathContext) SubAttr() ISubAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitAttrPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonQueryParserRULE_attrPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(JsonQueryParserATTRNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonQueryParserT__3 {
		{
			p.SetState(83)
			p.SubAttr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrPath() IAttrPathContext

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
	return p
}

func InitEmptySubAttrContext(p *SubAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonQueryParserRULE_subAttr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(JsonQueryParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.AttrPath()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListOfDoublesContext struct {
	ValueContext
}

func NewListOfDoublesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfDoublesContext {
	var p = new(ListOfDoublesContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfDoublesContext) ListDoubles() IListDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListDoublesContext)
}

func (s *ListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfStringsContext struct {
	ValueContext
}

func NewListOfStringsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfStringsContext {
	var p = new(ListOfStringsContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfStringsContext) ListStrings() IListStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringsContext)
}

func (s *ListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserBOOLEAN, 0)
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	ValueContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNULL, 0)
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleContext struct {
	ValueContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *DoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type VersionContext struct {
	ValueContext
}

func NewVersionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VersionContext {
	var p = new(VersionContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) VERSION() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserVERSION, 0)
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

type LongContext struct {
	ValueContext
}

func NewLongContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LongContext {
	var p = new(LongContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *LongContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *LongContext) EXP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEXP, 0)
}

func (s *LongContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLong(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfIntsContext struct {
	ValueContext
}

func NewListOfIntsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfIntsContext {
	var p = new(ListOfIntsContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfIntsContext) ListInts() IListIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListIntsContext)
}

func (s *ListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonQueryParserRULE_value)
	var _la int

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Match(JsonQueryParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(JsonQueryParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVersionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Match(JsonQueryParserVERSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(93)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewLongContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserT__4 {
			{
				p.SetState(94)
				p.Match(JsonQueryParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(97)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(98)
				p.Match(JsonQueryParserEXP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewListOfIntsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(101)
			p.ListInts()
		}

	case 8:
		localctx = NewListOfDoublesContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(102)
			p.ListDoubles()
		}

	case 9:
		localctx = NewListOfStringsContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(103)
			p.ListStrings()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStringsContext is an interface to support dynamic dispatch.
type IListStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListStringsContext differentiates from other interfaces.
	IsListStringsContext()
}

type ListStringsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringsContext() *ListStringsContext {
	var p = new(ListStringsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
	return p
}

func InitEmptyListStringsContext(p *ListStringsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
}

func (*ListStringsContext) IsListStringsContext() {}

func NewListStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringsContext {
	var p = new(ListStringsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listStrings

	return p
}

func (s *ListStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringsContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSTRING)
}

func (s *ListStringsContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, i)
}

func (s *ListStringsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserCOMMA)
}

func (s *ListStringsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, i)
}

func (s *ListStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListStrings() (localctx IListStringsContext) {
	localctx = NewListStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonQueryParserRULE_listStrings)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(JsonQueryParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JsonQueryParserCOMMA {
		{
			p.SetState(108)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(JsonQueryParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListDoublesContext is an interface to support dynamic dispatch.
type IListDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDOUBLE() []antlr.TerminalNode
	DOUBLE(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListDoublesContext differentiates from other interfaces.
	IsListDoublesContext()
}

type ListDoublesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListDoublesContext() *ListDoublesContext {
	var p = new(ListDoublesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
	return p
}

func InitEmptyListDoublesContext(p *ListDoublesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
}

func (*ListDoublesContext) IsListDoublesContext() {}

func NewListDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListDoublesContext {
	var p = new(ListDoublesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listDoubles

	return p
}

func (s *ListDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListDoublesContext) AllDOUBLE() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserDOUBLE)
}

func (s *ListDoublesContext) DOUBLE(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, i)
}

func (s *ListDoublesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserCOMMA)
}

func (s *ListDoublesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, i)
}

func (s *ListDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListDoubles() (localctx IListDoublesContext) {
	localctx = NewListDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonQueryParserRULE_listDoubles)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(JsonQueryParserDOUBLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JsonQueryParserCOMMA {
		{
			p.SetState(119)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(JsonQueryParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListAttrPathsContext is an interface to support dynamic dispatch.
type IListAttrPathsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttrPath() []IAttrPathContext
	AttrPath(i int) IAttrPathContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListAttrPathsContext differentiates from other interfaces.
	IsListAttrPathsContext()
}

type ListAttrPathsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListAttrPathsContext() *ListAttrPathsContext {
	var p = new(ListAttrPathsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listAttrPaths
	return p
}

func InitEmptyListAttrPathsContext(p *ListAttrPathsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listAttrPaths
}

func (*ListAttrPathsContext) IsListAttrPathsContext() {}

func NewListAttrPathsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAttrPathsContext {
	var p = new(ListAttrPathsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listAttrPaths

	return p
}

func (s *ListAttrPathsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListAttrPathsContext) AllAttrPath() []IAttrPathContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttrPathContext); ok {
			len++
		}
	}

	tst := make([]IAttrPathContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttrPathContext); ok {
			tst[i] = t.(IAttrPathContext)
			i++
		}
	}

	return tst
}

func (s *ListAttrPathsContext) AttrPath(i int) IAttrPathContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *ListAttrPathsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserCOMMA)
}

func (s *ListAttrPathsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, i)
}

func (s *ListAttrPathsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAttrPathsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListAttrPathsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListAttrPaths(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListAttrPaths() (localctx IListAttrPathsContext) {
	localctx = NewListAttrPathsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonQueryParserRULE_listAttrPaths)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.AttrPath()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JsonQueryParserCOMMA {
		{
			p.SetState(129)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.AttrPath()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListIntsContext is an interface to support dynamic dispatch.
type IListIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListIntsContext differentiates from other interfaces.
	IsListIntsContext()
}

type ListIntsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListIntsContext() *ListIntsContext {
	var p = new(ListIntsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
	return p
}

func InitEmptyListIntsContext(p *ListIntsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
}

func (*ListIntsContext) IsListIntsContext() {}

func NewListIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIntsContext {
	var p = new(ListIntsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listInts

	return p
}

func (s *ListIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIntsContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserINT)
}

func (s *ListIntsContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, i)
}

func (s *ListIntsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserCOMMA)
}

func (s *ListIntsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, i)
}

func (s *ListIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListInts() (localctx IListIntsContext) {
	localctx = NewListIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonQueryParserRULE_listInts)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(JsonQueryParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JsonQueryParserCOMMA {
		{
			p.SetState(138)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
		p.Match(JsonQueryParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *JsonQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonQueryParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
