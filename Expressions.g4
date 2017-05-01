grammar Expressions;

start : (codeline NEWLINE+)+ EOF ;

codeline : label=ID ':' code=expr ;

expr
    : '(' e=expr ')'                 #ParenExpr
    | unary                        #UnaryExpr
    | a=expr op=(ADD|SUB) b=expr   #AddSubExpr
    | literal                      #LiteralExpr
    ;

unary
    : op=SUB val=expr
    | op=INV val=expr
    ;

literal
    : lit=INT                      #IntLiteral
    ;

INV : '~' ;
ADD : '+' ;
SUB : '-' ;

fragment AlphaNum : ('a'..'z' | 'A'..'Z' | '0'..'9' | '$' | '_') ;
ID : ('a'..'z' | '$' | '_') AlphaNum*;
INT :   [0-9]+ ;         // match integers
NEWLINE:'\r'? '\n' ;     // return newlines to parser (is end-statement signal)
COMMENT : '#' ~( '\r' | '\n' )* -> skip ;
WS  :   [ \t]+ -> skip ; // toss out whitespace
