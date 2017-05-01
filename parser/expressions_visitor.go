// Generated from Expressions.g4 by ANTLR 4.7.

package parser // Expressions
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by ExpressionsParser.
type ExpressionsVisitor interface {
    antlr.ParseTreeVisitor
    StartContextVisitor
    CodelineContextVisitor
    AddSubExprContextVisitor
    ParenExprContextVisitor
    LiteralExprContextVisitor
    UnaryExprContextVisitor
    UnaryContextVisitor
    IntLiteralContextVisitor
}

type StartContextVisitor interface {
    VisitStart(ctx IStartContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type CodelineContextVisitor interface {
    VisitCodeline(ctx ICodelineContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type AddSubExprContextVisitor interface {
    VisitAddSubExpr(ctx IAddSubExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ParenExprContextVisitor interface {
    VisitParenExpr(ctx IParenExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type LiteralExprContextVisitor interface {
    VisitLiteralExpr(ctx ILiteralExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type UnaryExprContextVisitor interface {
    VisitUnaryExpr(ctx IUnaryExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type UnaryContextVisitor interface {
    VisitUnary(ctx IUnaryContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type IntLiteralContextVisitor interface {
    VisitIntLiteral(ctx IIntLiteralContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}