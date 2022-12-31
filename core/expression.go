package core

import (
	"fmt"
	"reflect"
	"regexp"
)

type Expression interface {
	Eval() (interface{}, error)
	String() string
}

type ValueExpression struct {
	value interface{}
	kind  reflect.Kind
}

func (e *ValueExpression) Eval() (interface{}, error) {
	return e.value, nil
}

func (e *ValueExpression) String() string {
	return fmt.Sprintf("%v", e.value)
}

type OP struct {
	Symbol string
	Func   opFunc
}

func (op *OP) Eval(left, right interface{}) (interface{}, error) {
	return op.Func(left, right), nil
}

type opFunc func(left, right interface{}) interface{}

type BiExpr struct {
	left  Expression
	right Expression
	op    *OP
}

func (e *BiExpr) String() string {
	return fmt.Sprintf("%s %s %s", e.left.String(), e.op.Symbol, e.right.String())
}

func (e *BiExpr) Eval() (interface{}, error) {
	lv, err := e.left.Eval()
	if err != nil {
		return nil, err
	}
	rv, err := e.right.Eval()
	if err != nil {
		return nil, err
	}

	ret, err := e.op.Eval(lv, rv)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type ArrayExpr struct {
	left  Expression
	right Expression
	op    *OP
}

func (e *ArrayExpr) Eval() (interface{}, error) {

	lv, err := e.left.Eval()
	if err != nil {
		return nil, err
	}
	rv, err := e.right.Eval()
	if err != nil {
		return nil, err
	}

	ret, err := e.op.Eval(lv, rv)
	if err != nil {
		return nil, err
	}
	return ret, nil

}

type NotExpr struct {
	right Expression
}

func (e *NotExpr) Eval() (interface{}, error) {
	rightValue, err := e.right.Eval()
	if err != nil {
		return nil, err
	}

	return !rightValue.(bool), nil
}

type RegexExpr struct {
	left  Expression
	regex Expression
}

func (e *RegexExpr) String() string {
	return fmt.Sprintf("left: %s, regex:%s", e.left.String(), e.regex.String())
}

func (e *RegexExpr) Eval() (interface{}, error) {
	result, err := regexp.MatchString(e.regex.String(), e.left.String())
	return result, err
}

func (e *NotExpr) String() string {
	return fmt.Sprintf("%s", e.right.String())
}

func (e *ArrayExpr) String() string {
	return fmt.Sprintf("%s %s %s", e.left.String(), e.op.Symbol, e.right.String())
}

func NewOP(symbol string, fun opFunc) *OP {
	return &OP{symbol, fun}
}

func NewBiCalcExpr(left, right Expression, op *OP) Expression {
	return &BiExpr{left, right, op}
}

func NewValueExpression(value interface{}) *ValueExpression {
	kind := reflect.TypeOf(value).Kind()
	return &ValueExpression{value, kind}
}
