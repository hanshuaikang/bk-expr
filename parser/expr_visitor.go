// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#container.
	VisitContainer(ctx *ContainerContext) interface{}

	// Visit a parse tree produced by ExprParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by ExprParser#Float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by ExprParser#Parenthesis.
	VisitParenthesis(ctx *ParenthesisContext) interface{}

	// Visit a parse tree produced by ExprParser#Variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by ExprParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by ExprParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by ExprParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by ExprParser#NotIn.
	VisitNotIn(ctx *NotInContext) interface{}

	// Visit a parse tree produced by ExprParser#Compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by ExprParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by ExprParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by ExprParser#Binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by ExprParser#InExpr.
	VisitInExpr(ctx *InExprContext) interface{}
}
