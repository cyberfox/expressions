package expressions

import (
	"github.com/cyberfox/expressions/parser"
	"strconv"
)

type ExprVisitor struct {
	parser.BaseExpressionsVisitor
}

func (v *ExprVisitor) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	result, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err == nil {
		return int64(result)
	}
	return nil
}

func (v *ExprVisitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *ExprVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	op := ctx.GetOp().GetText()
	left := ctx.GetA().Accept(v)
	right := ctx.GetB().Accept(v)

	switch left.(type) {
	default:
		return nil
	case int64:
		if op == "+" {
			return left.(int64) + right.(int64)
		} else {
			return left.(int64) - right.(int64)
		}
	}

	return nil
}

func (v *ExprVisitor) VisitUnary(ctx *parser.UnaryContext)interface{} {
	val := ctx.GetVal().Accept(v)

	switch ctx.GetOp().GetText() {
	case "-":
		switch val.(type) {
		case int64:
			return -val.(int64);
		}
	case "~":
		switch val.(type) {
		case int64:
			return ^val.(int64);
		}
	}
	return nil;
}

func NewEvaluator() *ExprVisitor {
	visitor := new(ExprVisitor)
	visitor.SetSuper(visitor)
	return visitor
}
