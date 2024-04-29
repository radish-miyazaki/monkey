package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL はトークンが未知の文字を含んでいることを示します。
	ILLEGAL = "ILLEGAL"
	// EOF は入力の終端を示します。
	EOF = "EOF"
	// IDENT は識別子を示します。
	IDENT = "IDENT"
	// INT は整数リテラルを示します。
	INT = "INT"
	// ASSIGN は代入演算子を示します。
	ASSIGN = "="
	// PLUS は加算演算子を示します。
	PLUS = "+"
	// MINUS は減算演算子を示します。
	MINUS = "-"
	// BANG はビット反転演算子を示します。
	BANG = "!"
	// ASTERISK は乗算演算子を示します。
	ASTERISK = "*"
	// SLASH は除算演算子を示します。
	SLASH = "/"
	// LT は小なり演算子を示します。
	LT = "<"
	// GT は大なり演算子を示します。
	GT = ">"
	// EQ は等価演算子を示します。
	EQ = "=="
	// NOT_EQ は不等価演算子を示します。
	NOT_EQ = "!="
	//COMMA はカンマを示します。
	COMMA = ","
	// SEMICOLON はセミコロンを示します。
	SEMICOLON = ";"
	// LPAREN は左括弧を示します。
	LPAREN = "("
	// RPAREN は右括弧を示します。
	RPAREN = ")"
	// LBRACE は左波括弧を示します。
	LBRACE = "{"
	// RBRACE は右波括弧を示します。
	RBRACE = "}"
	// FUNCTION は関数を示します。
	FUNCTION = "function"
	// STRING は文字列リテラルを示します。
	STRING = "string"
	// LET は変数束縛を示します。
	LET = "let"
	// IF は条件分岐を示します。
	IF = "if"
	// ELSE は条件分岐のそれ以外を示します。
	ELSE = "else"
	// TRUE は真を示します。
	TRUE = "true"
	// FALSE は偽を示します。
	FALSE = "false"
	// RETURN は戻り値を示します。
	RETURN = "return"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
