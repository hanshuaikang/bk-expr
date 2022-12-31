grammar Expr;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
DIGIT: [0-9]+;
FLOAT: DIGIT? '.' DIGIT*;
WHITESPACE: [ \r\n\t]+ -> skip;
GT: '>';
GE: '>=';
LT: '<';
LE: '<=';
EQ: '==';
NE: '!=' ;
VARIABLE
    : '$'[a-zA-Z_] [.a-zA-Z_0-9]*
    | '@'[a-zA-Z_] [.a-zA-Z_0-9]*
    ;
AND: 'and' ;
OR: 'or' ;
NOT: 'not';
IN: 'in';
COMMA: ',';
LBRACKET: '[';
RBRACKET: ']';

STRING
    : '\'' (ESC|.)*? '\'';

fragment ESC : '\\"' | '\\\\' ;

container: VARIABLE;

// Rules
start : expr EOF;

expr
   : expr op=('*'|'/') expr # MulDiv
   | expr op=('+'|'-') expr # AddSub
   | '(' expr ')'           # Parenthesis // 括号规则
   | expr op=(GT | GE | LT | LE | EQ | NE) expr     # Compare
   | expr op=(AND | OR) expr                        # Binary
   | expr IN container                              # InExpr
   | NOT expr                                       # NotExpr
   | STRING                                         # String
   | DIGIT                                          # Number
   | FLOAT                                          # Float
   | VARIABLE                                       # Variable
   ;
