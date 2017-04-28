// Generated from Expressions.g4 by ANTLR 4.7.

package parser // Expressions

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExpressionsVisitor struct {
	*antlr.BaseParseTreeVisitor
	super ExpressionsVisitor
}

func (v *BaseExpressionsVisitor) SetSuper(newSuper ExpressionsVisitor) {
	v.super = newSuper
}

func (v *BaseExpressionsVisitor) DefaultResult() interface{} {
	return nil
}

func (v *BaseExpressionsVisitor) ShouldVisitNextChild(node antlr.RuleNode, resultSoFar interface{}) bool {
	return true
}

func (v *BaseExpressionsVisitor) AggregateResult(resultSoFar, childResult interface{}) interface{} {
	return childResult
}

func (v *BaseExpressionsVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result interface{} = nil

	if v.super != nil {
		result = v.super.DefaultResult()
	}

	for i := 0; i < node.GetChildCount(); i++ {
		if v.super != nil {
			doNextChild := v.super.ShouldVisitNextChild(node, result)
			if !doNextChild {
				break
			}
		}

		child := node.GetChild(i)
		ruleChild, ok := child.(antlr.RuleNode)
		if !ok {
			terminalChild, ok := child.(antlr.TerminalNode)
			if ok {
				childResult := v.VisitTerminal(terminalChild)
				if v.super != nil {
					result = v.super.AggregateResult(result, childResult)
				} else {
					result = childResult
				}
			} else {
				errChild, ok := child.(antlr.ErrorNode)
				if ok {
					v.VisitErrorNode(errChild)
				}
			}
		} else {
			if v.super != nil {
				childResult := ruleChild.Accept(v.super)
				result = v.super.AggregateResult(result, childResult)
			} else {
				result = ruleChild.Accept(v)
			}
		}
	}
	return result
}

func (v *BaseExpressionsVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitCodeline(ctx *CodelineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitLiteralExpr(ctx *LiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitUnary(ctx *UnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
