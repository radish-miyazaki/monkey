package token

type TokenType string

type Token struct {
	Type    TokenType
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
	// COMMA はカンマを示します。
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
	// LET は変数束縛を示します。
	LET = "let"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
