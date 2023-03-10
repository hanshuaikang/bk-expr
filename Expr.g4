grammar Expr;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
HYPHEN: '-';
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
    ;
AND: 'and' ;
OR: 'or' ;
NOT: 'not';
IN: 'in';

LPAREN : '(' ;
RPAREN : ')' ;
ARGUMENTS : TEXT (',' TEXT)* ;
TEXT : ('a'..'z' | '0'..'9' | 'A'..'Z')+;


FUNCTION
    : '@'[a-zA-Z_] [.a-zA-Z_0-9]* LPAREN (ARGUMENTS (',' ARGUMENTS)*)? RPAREN
    ;

DATE
    :'\'' DIGIT HYPHEN DIGIT HYPHEN DIGIT '\''
    ;

STRING
    : '\'' (ESC|.)*? '\'';

fragment ESC : '\\"' | '\\\\' ;

container
    : VARIABLE
    | STRING
    ;

start : expr EOF;


expr
   : expr op=('*'|'/') expr # MulDiv
   | expr op=('+'|'-') expr # AddSub
   | '(' expr ')'           # Parenthesis // 括号规则
   | expr op=(GT | GE | LT | LE | EQ | NE) expr     # Compare
   | expr IN container                              # InExpr
   | expr NOT IN container                          # NotIn
   | NOT expr                                       # NotExpr
   | expr op=(AND | OR) expr                        # Binary
   | STRING                                         # String
   | DIGIT                                          # Number
   | FLOAT                                          # Float
   | VARIABLE                                       # Variable
   | FUNCTION                                       # Function
   | DATE                                           # DATE
   ;
