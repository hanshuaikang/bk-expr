package core

import (
	"github.com/TencentBlueKing/bk-expr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type ExprVisitor struct {
	parser.BaseExprVisitor
	context map[string]interface{}
}

func (v *ExprVisitor) setContext(ctx map[string]interface{}) {
	v.context = ctx
}

func (v *ExprVisitor) getVariable(varKey string) Expression {
	// 获取某个变量的值
	valKey := varKey[1:]
	value, ok := v.context[valKey]
	if !ok {
		panic("the variable is not in context")
	}
	return NewValueExpression(value)
}

func (v *ExprVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *ExprVisitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	// 访问数字
	val, _ := strconv.Atoi(ctx.GetText())
	return NewValueExpression(val)
}

func (v *ExprVisitor) VisitFloat(ctx *parser.FloatContext) interface{} {
	// 访问Float
	val, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return NewValueExpression(val)
}

func (v *ExprVisitor) VisitString(ctx *parser.StringContext) interface{} {
	// 访问字符串
	l := len(ctx.GetText())
	s := ctx.GetText()
	return NewValueExpression(s[1 : l-1])
}

func (v *ExprVisitor) VisitParenthesis(ctx *parser.ParenthesisContext) interface{} {
	// 访问括号
	return ctx.Expr().Accept(v)
}

func (v *ExprVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	// 访问加减
	left := ctx.Expr(0).Accept(v).(Expression)
	right := ctx.Expr(1).Accept(v).(Expression)
	var t antlr.Token = ctx.GetOp()
	switch t.GetTokenType() {
	case parser.ExprLexerADD:
		return NewBiCalcExpr(left, right, NewOP("+", Plus))
	case parser.ExprLexerSUB:
		return NewBiCalcExpr(left, right, NewOP("-", Sub))
	default:
		panic("Temporarily unsupported operation types")
	}
}

func (v *ExprVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	// 访问乘除
	left := ctx.Expr(0).Accept(v).(Expression)
	right := ctx.Expr(1).Accept(v).(Expression)
	var t antlr.Token = ctx.GetOp()
	switch t.GetTokenType() {
	case parser.ExprLexerMUL:
		return NewBiCalcExpr(left, right, NewOP("*", Mul))
	case parser.ExprLexerDIV:
		return NewBiCalcExpr(left, right, NewOP("/", Div))
	default:
		panic("Temporarily unsupported operation types")

	}
	return nil
}

func (v *ExprVisitor) VisitCompare(ctx *parser.CompareContext) interface{} {
	// 其他 bool 运算
	left := ctx.Expr(0).Accept(v).(Expression)
	right := ctx.Expr(1).Accept(v).(Expression)
	var t antlr.Token = ctx.GetOp()
	switch t.GetTokenType() {
	case parser.ExprParserGE:
		return NewBiCalcExpr(left, right, NewOP(">=", Ge))
	case parser.ExprParserGT:
		return NewBiCalcExpr(left, right, NewOP(">", Gt))
	case parser.ExprParserLT:
		return NewBiCalcExpr(left, right, NewOP("<", Lt))
	case parser.ExprParserLE:
		return NewBiCalcExpr(left, right, NewOP("<=", Le))
	case parser.ExprParserNE:
		return NewBiCalcExpr(left, right, NewOP("!=", Ne))
	case parser.ExprParserEQ:
		return NewBiCalcExpr(left, right, NewOP("==", Equal))
	default:
		panic("Temporarily unsupported operation types")

	}
}

func (v *ExprVisitor) VisitVariable(ctx *parser.VariableContext) interface{} {
	// 访问变量
	return v.getVariable(ctx.GetText())
}

func (v *ExprVisitor) VisitBinary(ctx *parser.BinaryContext) interface{} {
	// 访问二或运算
	left := ctx.Expr(0).Accept(v).(Expression)
	right := ctx.Expr(1).Accept(v).(Expression)
	var t antlr.Token = ctx.GetOp()
	switch t.GetTokenType() {
	case parser.ExprParserAND:
		return NewBiCalcExpr(left, right, NewOP("and", And))
	case parser.ExprParserOR:
		return NewBiCalcExpr(left, right, NewOP("or", Or))
	default:
		panic("should not happen")

	}
}
func (v *ExprVisitor) VisitContainer(ctx *parser.ContainerContext) interface{} {
	// 访问容器
	if ctx.VARIABLE() != nil {
		return v.getVariable(ctx.VARIABLE().GetText())
	}
	return nil
}

func (v *ExprVisitor) VisitInExpr(ctx *parser.InExprContext) interface{} {
	// in 操作
	left := ctx.Expr().Accept(v).(Expression)
	right := ctx.Container().Accept(v).(Expression)
	return &ArrayExpr{left, right, NewOP("in", In)}
}

func (v *ExprVisitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	// not 不是
	right := ctx.Expr().Accept(v).(Expression)
	return &NotExpr{right: right}
}
