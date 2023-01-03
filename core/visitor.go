package core

import (
	"fmt"
	"github.com/TencentBlueKing/bk-expr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
	"strings"
	"time"
)

type ExprVisitor struct {
	parser.BaseExprVisitor
	context map[string]interface{}
	funcHub map[string]Function
}

func (v *ExprVisitor) setContext(ctx map[string]interface{}) {
	v.context = ctx
}

func (v *ExprVisitor) installFunction(functionName string, function Function) {
	if v.funcHub == nil {
		v.funcHub = map[string]Function{}
	}
	v.funcHub[functionName] = function
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

func (v *ExprVisitor) getFunction(funcName string, args []string) Expression {
	function, ok := v.funcHub[funcName]
	if !ok {
		panic("the function is not in context")
	}
	value, err := function.Execute(args)
	if err != nil {
		panic(fmt.Sprintf("the function is execute failed, err=%s", err))
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
	case parser.ExprLexerHYPHEN:
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

func (v *ExprVisitor) VisitNotIn(ctx *parser.NotInContext) interface{} {
	// not in
	left := ctx.Expr().Accept(v).(Expression)
	right := ctx.Container().Accept(v).(Expression)
	return &ArrayExpr{left, right, NewOP("not in", NotIn)}
}

func (v *ExprVisitor) VisitDATE(ctx *parser.DATEContext) interface{} {
	length := len(ctx.GetText())
	val, _ := time.Parse("2006-01-02", ctx.GetText()[1:length-1])
	return NewValueExpression(val)
}

func (v *ExprVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	funcName := strings.Split(ctx.GetText()[1:], "(")[0]
	argsText := strings.TrimRight(strings.Split(ctx.GetText()[1:], "(")[1], ")")
	argsList := strings.Split(argsText, ",")
	return v.getFunction(funcName, argsList)
}
