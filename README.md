# BK-EXPR go语言表达式解析工具

#### 支持变量渲染

变量方式: $env

### 数值计算

- 加
- 减
- 乘
- 除
- 取余

#### 布尔计算

- (==)
- (>)
- (<)
- (>=)
- (<=)

_左值和右值必须为同类型_

#### 逻辑运算符

&&
||
not

### 包含表达式

- in
- not in 

示例:

```go
// 运算符
expression := "1+1-1*1/2"

// 比较运算符
expression := "1>1"

//逻辑运算符
expression := "1>1 && 1==1"

// 包含表达式(现阶段只支持变量)
expression := "'本科' in $list "
ctx := map[string]interface{}{"list": []interface{}{"本科", "硕士"}, "age": 7, "money": float64(10)}



```