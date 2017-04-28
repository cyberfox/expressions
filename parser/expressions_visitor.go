// Generated from Expressions.g4 by ANTLR 4.7.

package parser // Expressions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExpressionsParser.
type ExpressionsVisitor interface {
	antlr.ParseTreeVisitor

	// Provide the default result that will be returned from GetChildren (default: nil)
	DefaultResult() interface{}

	// Decide if we should continue visiting the rule nodes, based on the result so far.
	ShouldVisitNextChild(node antlr.RuleNode, resultSoFar interface{}) bool

	// Given the results found in the children so far, and the new result, determine the new result value.
	AggregateResult(resultSoFar interface{}, childResult interface{}) interface{}

	// Visit a parse tree produced by ExpressionsParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#codeline.
	VisitCodeline(ctx *CodelineContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#LiteralExpr.
	VisitLiteralExpr(ctx *LiteralExprContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}
}
