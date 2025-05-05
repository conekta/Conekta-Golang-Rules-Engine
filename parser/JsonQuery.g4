grammar JsonQuery;

query
   : NOT? '(' query ')'                                                                                 #parenExp
   | query LOGICAL_OPERATOR query                                                                             #logicalExp
   | attrPath 'pr'                                                                                               #presentExp
   | attrPath op=( EQ | NE | GT | LT | GE | LE | CO | SW | EW | IN ) value                                    #compareExp
   | attrPath op=( EQ | NE | GT | LT | GE | LE | CO | SW | EW | IN ) attrPathValue                            #compareExpAttrPath
   | (ASTERISK|PLUS|MINUS|DIVISON) '(' listAttrPaths ')' op=( EQ | NE | GT | LT | GE | LE ) value          #mulSumExp
   ;

NOT
   : 'not' | 'NOT'
   ;

LOGICAL_OPERATOR
   : 'and' | 'or'
   ;

BOOLEAN
   : 'true' | 'false'
   ;

NULL
   : 'null'
   ;

IN:  'IN' | 'in';
EQ : 'eq' | 'EQ' | '==';
NE : 'ne' | 'NE' | '!=';
GT : 'gt' | 'GT' | '>';
LT : 'lt' | 'LT' | '<';
GE : 'ge' | 'GE' | '>=';
LE : 'le' | 'LE' | '<=';
CO : 'co' | 'CO';
SW : 'sw' | 'SW';
EW : 'ew' | 'EW';

attrPathValue
   : ATTRNAME subAttrValue?
   ;

subAttrValue
   : '.' attrPathValue
   ;

attrPath
   : ATTRNAME subAttr?
   ;

subAttr
   : '.' attrPath
   ;

ASTERISK            : 'MLP' ;
PLUS                : 'SUM' ;
MINUS               : 'SUBTRACT' ;
DIVISON             : 'DIV' ;

ATTRNAME
   : ALPHA ATTR_NAME_CHAR* ;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA | COMMA
   ;

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

value
   : BOOLEAN           #boolean
   | NULL              #null
   | VERSION           #version
   | STRING            #string
   | DOUBLE            #double
   | '-'? INT EXP?     #long
   | listInts          #listOfInts
   | listDoubles       #listOfDoubles
   | listStrings       #listOfStrings
   ;

VERSION
   : INT '.' INT '.' INT
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;

listStrings
   : '[' STRING (COMMA STRING)* ']'
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

DOUBLE
   : '-'? INT '.' [0-9] + EXP?
   ;

listDoubles
   : '[' DOUBLE (COMMA DOUBLE)* ']'
   ;

listAttrPaths
   : attrPath (COMMA attrPath)*
   ;

listInts
   : '[' INT (COMMA INT)* ']'
   ;

// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;

// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

COMMA
   : ',' ;

WS
   : [ \t\r\n]+ -> skip
   ;
