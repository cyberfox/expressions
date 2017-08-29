grammar Expressions;

options {
  language = Go;
  runtimeImport = 'github.com/wxio/antlr4/runtime/Go/antlr';
}

start : (codeline NEWLINE+)+ EOF ;

codeline : label=ID ':' code=expr ;

expr
    : '(' e=expr ')'               #ParenExpr
    | op=SUB val=expr              #NegateExpr
    | op=INV val=expr              #InvertExpr
    | a=expr op=(ADD|SUB) b=expr   #AddSubExpr
    | literal                      #LiteralExpr
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
