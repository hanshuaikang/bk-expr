// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExprVisitor) VisitContainer(ctx *ContainerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitNotIn(ctx *NotInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitDATE(ctx *DATEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitParenthesis(ctx *ParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitCompare(ctx *CompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitInExpr(ctx *InExprContext) interface{} {
	return v.VisitChildren(ctx)
}
