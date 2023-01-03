// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 53, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 27, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 48, 10,
	4, 12, 4, 14, 4, 51, 11, 4, 3, 4, 2, 3, 6, 5, 2, 4, 6, 2, 6, 3, 2, 3, 4,
	3, 2, 5, 6, 3, 2, 10, 15, 3, 2, 17, 18, 2, 62, 2, 8, 3, 2, 2, 2, 4, 10,
	3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 9, 7, 16, 2, 2, 9, 3, 3, 2, 2, 2, 10,
	11, 5, 6, 4, 2, 11, 12, 7, 2, 2, 3, 12, 5, 3, 2, 2, 2, 13, 14, 8, 4, 1,
	2, 14, 15, 7, 21, 2, 2, 15, 16, 5, 6, 4, 2, 16, 17, 7, 22, 2, 2, 17, 27,
	3, 2, 2, 2, 18, 19, 7, 19, 2, 2, 19, 27, 5, 6, 4, 10, 20, 27, 7, 27, 2,
	2, 21, 27, 7, 7, 2, 2, 22, 27, 7, 8, 2, 2, 23, 27, 7, 16, 2, 2, 24, 27,
	7, 25, 2, 2, 25, 27, 7, 26, 2, 2, 26, 13, 3, 2, 2, 2, 26, 18, 3, 2, 2,
	2, 26, 20, 3, 2, 2, 2, 26, 21, 3, 2, 2, 2, 26, 22, 3, 2, 2, 2, 26, 23,
	3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 25, 3, 2, 2, 2, 27, 49, 3, 2, 2, 2,
	28, 29, 12, 16, 2, 2, 29, 30, 9, 2, 2, 2, 30, 48, 5, 6, 4, 17, 31, 32,
	12, 15, 2, 2, 32, 33, 9, 3, 2, 2, 33, 48, 5, 6, 4, 16, 34, 35, 12, 13,
	2, 2, 35, 36, 9, 4, 2, 2, 36, 48, 5, 6, 4, 14, 37, 38, 12, 9, 2, 2, 38,
	39, 9, 5, 2, 2, 39, 48, 5, 6, 4, 10, 40, 41, 12, 12, 2, 2, 41, 42, 7, 20,
	2, 2, 42, 48, 5, 2, 2, 2, 43, 44, 12, 11, 2, 2, 44, 45, 7, 19, 2, 2, 45,
	46, 7, 20, 2, 2, 46, 48, 5, 2, 2, 2, 47, 28, 3, 2, 2, 2, 47, 31, 3, 2,
	2, 2, 47, 34, 3, 2, 2, 2, 47, 37, 3, 2, 2, 2, 47, 40, 3, 2, 2, 2, 47, 43,
	3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2,
	50, 7, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 5, 26, 47, 49,
}
var literalNames = []string{
	"", "'*'", "'/'", "'+'", "'-'", "", "", "", "'>'", "'>='", "'<'", "'<='",
	"'=='", "'!='", "", "'and'", "'or'", "'not'", "'in'", "'('", "')'",
}
var symbolicNames = []string{
	"", "MUL", "DIV", "ADD", "HYPHEN", "DIGIT", "FLOAT", "WHITESPACE", "GT",
	"GE", "LT", "LE", "EQ", "NE", "VARIABLE", "AND", "OR", "NOT", "IN", "LPAREN",
	"RPAREN", "ARGUMENTS", "TEXT", "FUNCTION", "DATE", "STRING",
}

var ruleNames = []string{
	"container", "start", "expr",
}

type ExprParser struct {
	*antlr.BaseParser
}

// NewExprParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ExprParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewExprParser(input antlr.TokenStream) *ExprParser {
	this := new(ExprParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Expr.g4"

	return this
}

// ExprParser tokens.
const (
	ExprParserEOF        = antlr.TokenEOF
	ExprParserMUL        = 1
	ExprParserDIV        = 2
	ExprParserADD        = 3
	ExprParserHYPHEN     = 4
	ExprParserDIGIT      = 5
	ExprParserFLOAT      = 6
	ExprParserWHITESPACE = 7
	ExprParserGT         = 8
	ExprParserGE         = 9
	ExprParserLT         = 10
	ExprParserLE         = 11
	ExprParserEQ         = 12
	ExprParserNE         = 13
	ExprParserVARIABLE   = 14
	ExprParserAND        = 15
	ExprParserOR         = 16
	ExprParserNOT        = 17
	ExprParserIN         = 18
	ExprParserLPAREN     = 19
	ExprParserRPAREN     = 20
	ExprParserARGUMENTS  = 21
	ExprParserTEXT       = 22
	ExprParserFUNCTION   = 23
	ExprParserDATE       = 24
	ExprParserSTRING     = 25
)

// ExprParser rules.
const (
	ExprParserRULE_container = 0
	ExprParserRULE_start     = 1
	ExprParserRULE_expr      = 2
)

// IContainerContext is an interface to support dynamic dispatch.
type IContainerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContainerContext differentiates from other interfaces.
	IsContainerContext()
}

type ContainerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainerContext() *ContainerContext {
	var p = new(ContainerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_container
	return p
}

func (*ContainerContext) IsContainerContext() {}

func NewContainerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContainerContext {
	var p = new(ContainerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_container

	return p
}

func (s *ContainerContext) GetParser() antlr.Parser { return s.parser }

func (s *ContainerContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ExprParserVARIABLE, 0)
}

func (s *ContainerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContainerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitContainer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Container() (localctx IContainerContext) {
	localctx = NewContainerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExprParserRULE_container)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.Match(ExprParserVARIABLE)
	}

	return localctx
}

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExprParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExprParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.expr(0)
	}
	{
		p.SetState(9)
		p.Match(ExprParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableContext struct {
	*ExprContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ExprParserVARIABLE, 0)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(ExprParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(ExprParserDIV, 0)
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(ExprParserADD, 0)
}

func (s *AddSubContext) HYPHEN() antlr.TerminalNode {
	return s.GetToken(ExprParserHYPHEN, 0)
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotInContext struct {
	*ExprContext
}

func NewNotInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInContext {
	var p = new(NotInContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotInContext) NOT() antlr.TerminalNode {
	return s.GetToken(ExprParserNOT, 0)
}

func (s *NotInContext) IN() antlr.TerminalNode {
	return s.GetToken(ExprParserIN, 0)
}

func (s *NotInContext) Container() IContainerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContainerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContainerContext)
}

func (s *NotInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitNotIn(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*ExprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(ExprParserSTRING, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	*ExprContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ExprParserFLOAT, 0)
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionContext struct {
	*ExprContext
}

func NewFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionContext {
	var p = new(FunctionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(ExprParserFUNCTION, 0)
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type DATEContext struct {
	*ExprContext
}

func NewDATEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DATEContext {
	var p = new(DATEContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DATEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DATEContext) DATE() antlr.TerminalNode {
	return s.GetToken(ExprParserDATE, 0)
}

func (s *DATEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitDATE(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesisContext struct {
	*ExprContext
}

func NewParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserLPAREN, 0)
}

func (s *ParenthesisContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesisContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExprParserRPAREN, 0)
}

func (s *ParenthesisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitParenthesis(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	*ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(ExprParserDIGIT, 0)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareContext struct {
	*ExprContext
	op antlr.Token
}

func NewCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareContext {
	var p = new(CompareContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CompareContext) GetOp() antlr.Token { return s.op }

func (s *CompareContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CompareContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompareContext) GT() antlr.TerminalNode {
	return s.GetToken(ExprParserGT, 0)
}

func (s *CompareContext) GE() antlr.TerminalNode {
	return s.GetToken(ExprParserGE, 0)
}

func (s *CompareContext) LT() antlr.TerminalNode {
	return s.GetToken(ExprParserLT, 0)
}

func (s *CompareContext) LE() antlr.TerminalNode {
	return s.GetToken(ExprParserLE, 0)
}

func (s *CompareContext) EQ() antlr.TerminalNode {
	return s.GetToken(ExprParserEQ, 0)
}

func (s *CompareContext) NE() antlr.TerminalNode {
	return s.GetToken(ExprParserNE, 0)
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(ExprParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryContext struct {
	*ExprContext
	op antlr.Token
}

func NewBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContext {
	var p = new(BinaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BinaryContext) GetOp() antlr.Token { return s.op }

func (s *BinaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BinaryContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinaryContext) AND() antlr.TerminalNode {
	return s.GetToken(ExprParserAND, 0)
}

func (s *BinaryContext) OR() antlr.TerminalNode {
	return s.GetToken(ExprParserOR, 0)
}

func (s *BinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

type InExprContext struct {
	*ExprContext
}

func NewInExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExprContext {
	var p = new(InExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InExprContext) IN() antlr.TerminalNode {
	return s.GetToken(ExprParserIN, 0)
}

func (s *InExprContext) Container() IContainerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContainerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContainerContext)
}

func (s *InExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitInExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ExprParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ExprParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserLPAREN:
		localctx = NewParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(12)
			p.Match(ExprParserLPAREN)
		}
		{
			p.SetState(13)
			p.expr(0)
		}
		{
			p.SetState(14)
			p.Match(ExprParserRPAREN)
		}

	case ExprParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(16)
			p.Match(ExprParserNOT)
		}
		{
			p.SetState(17)
			p.expr(8)
		}

	case ExprParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(ExprParserSTRING)
		}

	case ExprParserDIGIT:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.Match(ExprParserDIGIT)
		}

	case ExprParserFLOAT:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Match(ExprParserFLOAT)
		}

	case ExprParserVARIABLE:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.Match(ExprParserVARIABLE)
		}

	case ExprParserFUNCTION:
		localctx = NewFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(ExprParserFUNCTION)
		}

	case ExprParserDATE:
		localctx = NewDATEContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(ExprParserDATE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expr)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(27)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExprParserMUL || _la == ExprParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(28)
					p.expr(15)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expr)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(30)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExprParserADD || _la == ExprParserHYPHEN) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(31)
					p.expr(14)
				}

			case 3:
				localctx = NewCompareContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expr)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(33)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExprParserGT)|(1<<ExprParserGE)|(1<<ExprParserLT)|(1<<ExprParserLE)|(1<<ExprParserEQ)|(1<<ExprParserNE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(34)
					p.expr(12)
				}

			case 4:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expr)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(36)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExprParserAND || _la == ExprParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(37)
					p.expr(8)
				}

			case 5:
				localctx = NewInExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expr)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(39)
					p.Match(ExprParserIN)
				}
				{
					p.SetState(40)
					p.Container()
				}

			case 6:
				localctx = NewNotInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expr)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(42)
					p.Match(ExprParserNOT)
				}
				{
					p.SetState(43)
					p.Match(ExprParserIN)
				}
				{
					p.SetState(44)
					p.Container()
				}

			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExprParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
