// Generated from Expressions.g4 by ANTLR 4.7.

package parser // Expressions
import (
	"fmt"
	"reflect"
	"strconv"
)
import "github.com/antlr/antlr4/runtime/Go/antlr"



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
	ExpressionsParserEOF = antlr.TokenEOF
	ExpressionsParserT__0 = 1
	ExpressionsParserT__1 = 2
	ExpressionsParserT__2 = 3
	ExpressionsParserINV = 4
	ExpressionsParserADD = 5
	ExpressionsParserSUB = 6
	ExpressionsParserID = 7
	ExpressionsParserINT = 8
	ExpressionsParserNEWLINE = 9
	ExpressionsParserCOMMENT = 10
	ExpressionsParserWS = 11
)

// ExpressionsParser rules.
const (
	ExpressionsParserRULE_start = 0
	ExpressionsParserRULE_codeline = 1
	ExpressionsParserRULE_expr = 2
	ExpressionsParserRULE_unary = 3
	ExpressionsParserRULE_literal = 4
)

type IStartContextInternal interface {
    // embed exported interface
    IStartContext

    //  ruleGetterDecl
    //  ruleListGetterDecl
    AllCodeline() []ICodelineContext 
    //  ruleListIndexedGetterDecl
    Codeline(i int) ICodelineContext 
    //  ruleContextDecls
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IStartContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

    //Gets for labeled elements
    //tokenDecls	
    //tokenTypeDecls
    //tokenListDecls
    //attributeDecls
    //tokenGetterDecl
   EOF() antlr.TerminalNode
    //tokenListGetterDecl
   AllNEWLINE() []antlr.TerminalNode
    //tokenListIndexedGetterDecl
   NEWLINE(i int) antlr.TerminalNode

    // IsStartContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
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

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *StartContext) EOF() antlr.TerminalNode {
    return s.GetToken(ExpressionsParserEOF, 0)
}


func (s *StartContext) AllCodeline() []ICodelineContext {
//ContextRuleListGetterDecl
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

//provideCopyFrom
func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(StartEntryListener); ok {
        listenerT.EnterStart(s)
    }
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(StartExitListener); ok {
        listenerT.ExitStart(s)
    }
}

func (s *StartContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case StartContextVisitor:
		return t.VisitStart(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//extensionMembers



func (p *ExpressionsParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpressionsParserRULE_start)
	var //TokenTypeDecl
	_la int


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


type ICodelineContextInternal interface {
    // embed exported interface
    ICodelineContext

    //  ruleGetterDecl
    Expr() IExprContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
     GetCode() IExprContext 
     
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type ICodelineContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

    //Gets for labeled elements
    //tokenDecls	
     GetLabel() antlr.Token
     
    //tokenTypeDecls
    //tokenListDecls
    //attributeDecls
    //tokenGetterDecl
   ID() antlr.TerminalNode
    //tokenListGetterDecl
    //tokenListIndexedGetterDecl

    // IsCodelineContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
}

type CodelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	label antlr.Token
	//RuleContextDecl
	code IExprContext
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

//StructDecl tokenDecls
func (s *CodelineContext) GetLabel() antlr.Token { return s.label }
func (s *CodelineContext) SetLabel(v antlr.Token) { s.label = v }


//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls
func (s *CodelineContext) GetCode() IExprContext { return s.code }
func (s *CodelineContext) SetCode(v IExprContext) { s.code = v }


//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
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

//provideCopyFrom
func (s *CodelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *CodelineContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(CodelineEntryListener); ok {
        listenerT.EnterCodeline(s)
    }
}

func (s *CodelineContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(CodelineExitListener); ok {
        listenerT.ExitCodeline(s)
    }
}

func (s *CodelineContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case CodelineContextVisitor:
		return t.VisitCodeline(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//extensionMembers



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


type IExprContextInternal interface {
    // embed exported interface
    IExprContext

    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IExprContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

    //Gets for labeled elements
    //tokenDecls	
    //tokenTypeDecls
    //tokenListDecls
    //attributeDecls
    //tokenGetterDecl
    //tokenListGetterDecl
    //tokenListIndexedGetterDecl

    // IsExprContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
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

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers




//Begin AltLabelStructDecl


type AddSubExprContext struct {
	*ExprContext
	//RuleContextDecl
	a IExprContext
	//TokenDecl
	op antlr.Token
	//RuleContextDecl
	b IExprContext
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IAddSubExprContextInternal interface {
    IAddSubExprContext
    //Gets for raw elements
    //  ruleGetterDecl
    //  ruleListGetterDecl
    AllExpr() []IExprContext 
    //  ruleListIndexedGetterDecl
    Expr(i int) IExprContext 

    //  tokenGetterDecl
    ADD() antlr.TerminalNode
    SUB() antlr.TerminalNode
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IAddSubExprContext interface {
    //Current rule
    IExprContext

    //Gets for labeled elements
    //  tokenDecls
    GetOp() antlr.Token
     
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
     GetA() IExprContext 
      GetB() IExprContext 
     
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*AddSubExprContext) IsAddSubExprContext() {}

//AltLabelStructDecl tokenDecls
func (s *AddSubExprContext) GetOp() antlr.Token { return s.op }
func (s *AddSubExprContext) SetOp(v antlr.Token) { s.op = v }


//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls
func (s *AddSubExprContext) GetA() IExprContext { return s.a }
func (s *AddSubExprContext) SetA(v IExprContext) { s.a = v }


func (s *AddSubExprContext) GetB() IExprContext { return s.b }
func (s *AddSubExprContext) SetB(v IExprContext) { s.b = v }


//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *AddSubExprContext) AllExpr() []IExprContext {
//ContextRuleListGetterDecl
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


func (s *AddSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(AddSubExprEntryListener); ok {
        listenerT.EnterAddSubExpr(s)
    }
}


func (s *AddSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(AddSubExprExitListener); ok {
        listenerT.ExitAddSubExpr(s)
    }
}


func (s *AddSubExprContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case AddSubExprContextVisitor:
		return t.VisitAddSubExpr(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




//Begin AltLabelStructDecl


type ParenExprContext struct {
	*ExprContext
	//RuleContextDecl
	e IExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IParenExprContextInternal interface {
    IParenExprContext
    //Gets for raw elements
    //  ruleGetterDecl
    Expr() IExprContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IParenExprContext interface {
    //Current rule
    IExprContext

    //Gets for labeled elements
    //  tokenDecls
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
     GetE() IExprContext 
     
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*ParenExprContext) IsParenExprContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls
func (s *ParenExprContext) GetE() IExprContext { return s.e }
func (s *ParenExprContext) SetE(v IExprContext) { s.e = v }


//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ParenExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(ParenExprEntryListener); ok {
        listenerT.EnterParenExpr(s)
    }
}


func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(ParenExprExitListener); ok {
        listenerT.ExitParenExpr(s)
    }
}


func (s *ParenExprContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ParenExprContextVisitor:
		return t.VisitParenExpr(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




//Begin AltLabelStructDecl


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

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type ILiteralExprContextInternal interface {
    ILiteralExprContext
    //Gets for raw elements
    //  ruleGetterDecl
    Literal() ILiteralContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type ILiteralExprContext interface {
    //Current rule
    IExprContext

    //Gets for labeled elements
    //  tokenDecls
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*LiteralExprContext) IsLiteralExprContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *LiteralExprContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}


func (s *LiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(LiteralExprEntryListener); ok {
        listenerT.EnterLiteralExpr(s)
    }
}


func (s *LiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(LiteralExprExitListener); ok {
        listenerT.ExitLiteralExpr(s)
    }
}


func (s *LiteralExprContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case LiteralExprContextVisitor:
		return t.VisitLiteralExpr(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




//Begin AltLabelStructDecl


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

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IUnaryExprContextInternal interface {
    IUnaryExprContext
    //Gets for raw elements
    //  ruleGetterDecl
    Unary() IUnaryContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IUnaryExprContext interface {
    //Current rule
    IExprContext

    //Gets for labeled elements
    //  tokenDecls
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*UnaryExprContext) IsUnaryExprContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *UnaryExprContext) Unary() IUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}


func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(UnaryExprEntryListener); ok {
        listenerT.EnterUnaryExpr(s)
    }
}


func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(UnaryExprExitListener); ok {
        listenerT.ExitUnaryExpr(s)
    }
}


func (s *UnaryExprContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case UnaryExprContextVisitor:
		return t.VisitUnaryExpr(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




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
	var //TokenTypeDecl
	_la int


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

			var _x = p.expr(0)
		        localctx.(*ParenExprContext).e = _x
		     


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


type IUnaryContextInternal interface {
    // embed exported interface
    IUnaryContext

    //  ruleGetterDecl
    Expr() IExprContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
     GetVal() IExprContext 
     
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IUnaryContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

    //Gets for labeled elements
    //tokenDecls	
     GetOp() antlr.Token
     
    //tokenTypeDecls
    //tokenListDecls
    //attributeDecls
    //tokenGetterDecl
   SUB() antlr.TerminalNode
   INV() antlr.TerminalNode
    //tokenListGetterDecl
    //tokenListIndexedGetterDecl

    // IsUnaryContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
}

type UnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	op antlr.Token
	//RuleContextDecl
	val IExprContext
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

//StructDecl tokenDecls
func (s *UnaryContext) GetOp() antlr.Token { return s.op }
func (s *UnaryContext) SetOp(v antlr.Token) { s.op = v }


//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls
func (s *UnaryContext) GetVal() IExprContext { return s.val }
func (s *UnaryContext) SetVal(v IExprContext) { s.val = v }


//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
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

//provideCopyFrom
func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(UnaryEntryListener); ok {
        listenerT.EnterUnary(s)
    }
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(UnaryExitListener); ok {
        listenerT.ExitUnary(s)
    }
}

func (s *UnaryContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case UnaryContextVisitor:
		return t.VisitUnary(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//extensionMembers



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


type ILiteralContextInternal interface {
    // embed exported interface
    ILiteralContext

    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type ILiteralContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

    //Gets for labeled elements
    //tokenDecls	
    //tokenTypeDecls
    //tokenListDecls
    //attributeDecls
    //tokenGetterDecl
    //tokenListGetterDecl
    //tokenListIndexedGetterDecl

    // IsLiteralContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
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

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers



//Begin AltLabelStructDecl


type IntLiteralContext struct {
	*LiteralContext
	//TokenDecl
	lit antlr.Token
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IIntLiteralContextInternal interface {
    IIntLiteralContext
    //Gets for raw elements
    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    INT() antlr.TerminalNode
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IIntLiteralContext interface {
    //Current rule
    ILiteralContext

    //Gets for labeled elements
    //  tokenDecls
    GetLit() antlr.Token
     
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*IntLiteralContext) IsIntLiteralContext() {}

//AltLabelStructDecl tokenDecls
func (s *IntLiteralContext) GetLit() antlr.Token { return s.lit }
func (s *IntLiteralContext) SetLit(v antlr.Token) { s.lit = v }


//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *IntLiteralContext) INT() antlr.TerminalNode {
    return s.GetToken(ExpressionsParserINT, 0)
}


func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(IntLiteralEntryListener); ok {
        listenerT.EnterIntLiteral(s)
    }
}


func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(IntLiteralExitListener); ok {
        listenerT.ExitIntLiteral(s)
    }
}


func (s *IntLiteralContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case IntLiteralContextVisitor:
		return t.VisitIntLiteral(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




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
			if localctx != nil { t = localctx.(*ExprContext) }
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

