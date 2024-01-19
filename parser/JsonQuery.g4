grammar JsonQuery;

query
   : NOT? SP? '(' SP? query SP? ')'                                                                             #parenExp
   | query SP LOGICAL_OPERATOR SP query                                                                         #logicalExp
   | attrPath SP 'pr'                                                                                           #presentExp
   | attrPath SP op=( EQ | NE | GT | LT | GE | LE | CO | SW | EW | IN ) SP value                                #compareExp
   | attrPath SP op=( EQ | NE | GT | LT | GE | LE ) SP TIME_NOW_ADD SP '(' value ')'                           #timeNowAddExp
   | (ASTERISK|PLUS|MINUS|DIVISON) SP '(' listAttrPaths ')' SP op=( EQ | NE | GT | LT | GE | LE ) SP value      #mulSumExp
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
TIME_NOW_ADD        : 'TIME_NOW_ADD' ;

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
   | DATETIME          #datetime
   | '-'? INT EXP?     #long
   | listInts          #listOfInts
   | listDoubles       #listOfDoubles
   | listStrings       #listOfStrings
   ;

DATETIME
   : DIGIT DIGIT DIGIT DIGIT  '-' DIGIT DIGIT  '-' DIGIT DIGIT 'T' DIGIT DIGIT ':' DIGIT DIGIT
   ;

VERSION
   : INT '.' INT '.' INT
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;

listStrings
   : '[' subListOfStrings
   ;

subListOfStrings
   : STRING COMMA subListOfStrings
   | STRING ']';

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
   : '[' subListOfDoubles
   ;

subListOfDoubles
   : DOUBLE COMMA subListOfDoubles
   | DOUBLE ']';

listAttrPaths
   :  subListOfAttrPaths
   ;

subListOfAttrPaths
   : attrPath COMMA subListOfAttrPaths
   | attrPath ;


listInts
   : '[' subListOfInts
   ;

subListOfInts
   : INT COMMA subListOfInts
   | INT ']';

// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;

// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

NEWLINE
   : '\n' ;

COMMA
   : ',' ' '*;
SP
   : ' ' NEWLINE*
   ;
