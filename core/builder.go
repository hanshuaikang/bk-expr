package core

import (
	"github.com/TencentBlueKing/bk-expr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Build(expression string, ctx map[string]interface{}) Expression {
	is := antlr.NewInputStream(expression)
	lexer := parser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewExprParser(stream)
	visitor := ExprVisitor{}
	visitor.setContext(ctx)
	expr := p.Start().Accept(&visitor).(Expression)
	return expr

}
