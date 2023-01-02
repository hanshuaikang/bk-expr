package main

import (
	"fmt"
	"github.com/TencentBlueKing/bk-expr/core"
	"os"
)

func main() {
	// 参加飞盘活动的前提是，年龄大于18小于45岁本科学历或者很有钱
	//expression := "('本科' in $list and $age>18 and $age<45 ) or ( $age>64 and $money > 100000)"
	expression := "'2022-01-21' - '2022-01-01'"
	ctx := map[string]interface{}{"list": []interface{}{"本科", "硕士"}, "age": 7, "money": float64(10)}
	expr := core.Build(expression, ctx)
	result, err := expr.Eval()
	if err != nil {
		os.Exit(0)
	}
	fmt.Println(result)
}
