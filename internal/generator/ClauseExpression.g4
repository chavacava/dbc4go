// Calc.g4
grammar ClauseExpression;

// Tokens
FORALL                  : '@forall';
EXISTS                  : '@exists';
IMPLIES                 : '==>';
IN                      : '@in';
INDEXOF                 : '@indexof';
DOT                     : '.';
WHITESPACE              : [ \r\n\t]+ -> skip;
ID                      :  LETTER (LETTER|[0-9])*; 
fragment LETTER         : [a-zA-Z_];
DECIMAL_LIT             : ('0' | [1-9] ('_'? [0-9])*);
RAW_STRING_LIT          : '`' ~'`'* '`';
INTERPRETED_STRING_LIT  : '"' ~'"'* '"';

// Rules
root
   : clauseExpression EOF
   ;

clauseExpression
   : clauseExpression IMPLIES clauseExpression # Implies
   | FORALL iterator IN collection ':' clauseExpression # ForallElement
   | FORALL iterator INDEXOF collection ':' clauseExpression # ForallIndex
   | EXISTS iterator IN collection ':' clauseExpression # ExistsElement
   | EXISTS iterator INDEXOF collection ':' clauseExpression # ExistsIndex
   | completeGoExpression # PlainGoExpression
   | '!'?'(' clauseExpression ')' # ExprInParens
   ;

iterator
    : ID
    ;

collection
    : qualifiedIdentifier (functionCallArguments | sliceIndex )?
    ;

completeGoExpression
    : goExpression
    ;

goExpression
    : primaryExpression (('==' | '!=' | '||' | '&&' |'==' | '!='|'>' | '<' | '>=' | '<='|'+' | '-'|'*' | '/' | '%') primaryExpression)*
    ;

primaryExpression
    : '!' primaryExpression
    | '(' goExpression ')'
    | number
    | string
    | qualifiedIdentifier (functionCallArguments | sliceIndex )?
    ;

qualifiedIdentifier
   : ID ( '.' ID)* 
   ;

functionCallArguments
    : '(' (goExpression (',' goExpression)*)? ')'
    ;

sliceIndex
    : '[' goExpression ']'
    ;

number
    : DECIMAL_LIT
    ;

string
    : RAW_STRING_LIT
    | INTERPRETED_STRING_LIT
    ;
   
