// Generated from Expressions.g4 by ANTLR 4.7.

package parser // Expressions

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 54, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 6,
	2, 15, 10, 2, 13, 2, 14, 2, 16, 6, 2, 19, 10, 2, 13, 2, 14, 2, 20, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 36, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 41, 10, 4, 12, 4, 14, 4, 44, 11,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 50, 10, 5, 3, 6, 3, 6, 3, 6, 2, 3, 6,
	7, 2, 4, 6, 8, 10, 2, 3, 3, 2, 7, 8, 2, 54, 2, 18, 3, 2, 2, 2, 4, 24, 3,
	2, 2, 2, 6, 35, 3, 2, 2, 2, 8, 49, 3, 2, 2, 2, 10, 51, 3, 2, 2, 2, 12,
	14, 5, 4, 3, 2, 13, 15, 7, 11, 2, 2, 14, 13, 3, 2, 2, 2, 15, 16, 3, 2,
	2, 2, 16, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 19, 3, 2, 2, 2, 18, 12,
	3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2,
	21, 22, 3, 2, 2, 2, 22, 23, 7, 2, 2, 3, 23, 3, 3, 2, 2, 2, 24, 25, 7, 9,
	2, 2, 25, 26, 7, 3, 2, 2, 26, 27, 5, 6, 4, 2, 27, 5, 3, 2, 2, 2, 28, 29,
	8, 4, 1, 2, 29, 30, 7, 4, 2, 2, 30, 31, 5, 6, 4, 2, 31, 32, 7, 5, 2, 2,
	32, 36, 3, 2, 2, 2, 33, 36, 5, 8, 5, 2, 34, 36, 5, 10, 6, 2, 35, 28, 3,
	2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 42, 3, 2, 2, 2, 37,
	38, 12, 4, 2, 2, 38, 39, 9, 2, 2, 2, 39, 41, 5, 6, 4, 5, 40, 37, 3, 2,
	2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 7,
	3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 46, 7, 8, 2, 2, 46, 50, 5, 6, 4, 2,
	47, 48, 7, 6, 2, 2, 48, 50, 5, 6, 4, 2, 49, 45, 3, 2, 2, 2, 49, 47, 3,
	2, 2, 2, 50, 9, 3, 2, 2, 2, 51, 52, 7, 10, 2, 2, 52, 11, 3, 2, 2, 2, 7,
	16, 20, 35, 42, 49,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'('", "')'", "'~'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "INV", "ADD", "SUB", "ID", "INT", "NEWLINE", "COMMENT",
	"WS",
}

var ruleNames = []string{
	"start", "codeline", "expr", "unary", "literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ExpressionsParser struct {
	*antlr.BaseParser
}

func NewExpressionsParser(input antlr.TokenStream) *ExpressionsParser {
	this := new(ExpressionsParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Expressions.g4"

	return this
}

// ExpressionsParser tokens.
const (
	ExpressionsParserEOF     = antlr.TokenEOF
	ExpressionsParserT__0    = 1
	ExpressionsParserT__1    = 2
	ExpressionsParserT__2    = 3
	ExpressionsParserINV     = 4
	ExpressionsParserADD     = 5
	ExpressionsParserSUB     = 6
	ExpressionsParserID      = 7
	ExpressionsParserINT     = 8
	ExpressionsParserNEWLINE = 9
	ExpressionsParserCOMMENT = 10
	ExpressionsParserWS      = 11
)

// ExpressionsParser rules.
const (
	ExpressionsParserRULE_start    = 0
	ExpressionsParserRULE_codeline = 1
	ExpressionsParserRULE_expr     = 2
	ExpressionsParserRULE_unary    = 3
	ExpressionsParserRULE_literal  = 4
)

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
	p.RuleIndex = ExpressionsParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserEOF, 0)
}

func (s *StartContext) AllCodeline() []ICodelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICodelineContext)(nil)).Elem())
	var tst = make([]ICodelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICodelineContext)
		}
	}

	return tst
}

func (s *StartContext) Codeline(i int) ICodelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICodelineContext)
}

func (s *StartContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ExpressionsParserNEWLINE)
}

func (s *StartContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionsParserNEWLINE, i)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExpressionsParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpressionsParserRULE_start)
	var _la int

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
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ExpressionsParserID {
		{
			p.SetState(10)
			p.Codeline()
		}
		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ExpressionsParserNEWLINE {
			{
				p.SetState(11)
				p.Match(ExpressionsParserNEWLINE)
			}

			p.SetState(14)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
		p.Match(ExpressionsParserEOF)
	}

	return localctx
}

// ICodelineContext is an interface to support dynamic dispatch.
type ICodelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLabel returns the label token.
	GetLabel() antlr.Token

	// SetLabel sets the label token.
	SetLabel(antlr.Token)

	// GetCode returns the code rule contexts.
	GetCode() IExprContext

	// SetCode sets the code rule contexts.
	SetCode(IExprContext)

	// IsCodelineContext differentiates from other interfaces.
	IsCodelineContext()
}

type CodelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	label  antlr.Token
	code   IExprContext
}

func NewEmptyCodelineContext() *CodelineContext {
	var p = new(CodelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionsParserRULE_codeline
	return p
}

func (*CodelineContext) IsCodelineContext() {}

func NewCodelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodelineContext {
	var p = new(CodelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_codeline

	return p
}

func (s *CodelineContext) GetParser() antlr.Parser { return s.parser }

func (s *CodelineContext) GetLabel() antlr.Token { return s.label }

func (s *CodelineContext) SetLabel(v antlr.Token) { s.label = v }

func (s *CodelineContext) GetCode() IExprContext { return s.code }

func (s *CodelineContext) SetCode(v IExprContext) { s.code = v }

func (s *CodelineContext) ID() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserID, 0)
}

func (s *CodelineContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CodelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitCodeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExpressionsParser) Codeline() (localctx ICodelineContext) {
	localctx = NewCodelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExpressionsParserRULE_codeline)

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
		p.SetState(22)

		var _m = p.Match(ExpressionsParserID)

		localctx.(*CodelineContext).label = _m
	}
	{
		p.SetState(23)
		p.Match(ExpressionsParserT__0)
	}
	{
		p.SetState(24)

		var _x = p.expr(0)

		localctx.(*CodelineContext).code = _x
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
	p.RuleIndex = ExpressionsParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_expr

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

type LiteralExprContext struct {
	*ExprContext
}

func NewLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExprContext {
	var p = new(LiteralExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExprContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitLiteralExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	*ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	*ExprContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Unary() IUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubExprContext struct {
	*ExprContext
	a  IExprContext
	op antlr.Token
	b  IExprContext
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubExprContext) GetOp() antlr.Token { return s.op }

func (s *AddSubExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExprContext) GetA() IExprContext { return s.a }

func (s *AddSubExprContext) GetB() IExprContext { return s.b }

func (s *AddSubExprContext) SetA(v IExprContext) { s.a = v }

func (s *AddSubExprContext) SetB(v IExprContext) { s.b = v }

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserADD, 0)
}

func (s *AddSubExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserSUB, 0)
}

func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitAddSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExpressionsParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ExpressionsParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ExpressionsParserRULE_expr, _p)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionsParserT__1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(27)
			p.Match(ExpressionsParserT__1)
		}
		{
			p.SetState(28)
			p.expr(0)
		}
		{
			p.SetState(29)
			p.Match(ExpressionsParserT__2)
		}

	case ExpressionsParserINV, ExpressionsParserSUB:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.Unary()
		}

	case ExpressionsParserINT:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
			localctx.(*AddSubExprContext).a = _prevctx

			p.PushNewRecursionContext(localctx, _startState, ExpressionsParserRULE_expr)
			p.SetState(35)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(36)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AddSubExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ExpressionsParserADD || _la == ExpressionsParserSUB) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AddSubExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(37)

				var _x = p.expr(3)

				localctx.(*AddSubExprContext).b = _x
			}

		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetVal returns the val rule contexts.
	GetVal() IExprContext

	// SetVal sets the val rule contexts.
	SetVal(IExprContext)

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	val    IExprContext
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionsParserRULE_unary
	return p
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) GetOp() antlr.Token { return s.op }

func (s *UnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryContext) GetVal() IExprContext { return s.val }

func (s *UnaryContext) SetVal(v IExprContext) { s.val = v }

func (s *UnaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserSUB, 0)
}

func (s *UnaryContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryContext) INV() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserINV, 0)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExpressionsParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExpressionsParserRULE_unary)

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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionsParserSUB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)

			var _m = p.Match(ExpressionsParserSUB)

			localctx.(*UnaryContext).op = _m
		}
		{
			p.SetState(44)

			var _x = p.expr(0)

			localctx.(*UnaryContext).val = _x
		}

	case ExpressionsParserINV:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)

			var _m = p.Match(ExpressionsParserINV)

			localctx.(*UnaryContext).op = _m
		}
		{
			p.SetState(46)

			var _x = p.expr(0)

			localctx.(*UnaryContext).val = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionsParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntLiteralContext struct {
	*LiteralContext
	lit antlr.Token
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetLit() antlr.Token { return s.lit }

func (s *IntLiteralContext) SetLit(v antlr.Token) { s.lit = v }

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserINT, 0)
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExpressionsParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExpressionsParserRULE_literal)

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

	localctx = NewIntLiteralContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)

		var _m = p.Match(ExpressionsParserINT)

		localctx.(*IntLiteralContext).lit = _m
	}

	return localctx
}

func (p *ExpressionsParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *ExpressionsParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
