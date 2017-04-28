package expressions

import (
	"github.com/cyberfox/expressions/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ExpressionsVisitor struct {
	parser.BaseExpressionsVisitor
	Expressions map[string]parser.IExprContext
}

func NewExpressionsVisitor() *ExpressionsVisitor {
	visitor := new(ExpressionsVisitor)
	visitor.SetSuper(visitor)
	visitor.Expressions = make(map[string]parser.IExprContext)

	return visitor
}

func (ev *ExpressionsVisitor) VisitCodeline(ctx *parser.CodelineContext) interface{} {
	ev.Expressions[ctx.GetLabel().GetText()] = ctx.GetCode()

	return nil
}

func GetExpressions(filename string) map[string]parser.IExprContext {
	input := antlr.NewFileStream(filename)
	lexer := parser.NewExpressionsLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExpressionsParser(stream)
	tree := p.Start()

	tv := NewExpressionsVisitor()
	tree.Accept(tv)

	return tv.Expressions
}
