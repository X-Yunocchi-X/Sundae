package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT   = "<"
	GT   = ">"
	LTEQ = "<="
	GTEQ = ">="

	EQ     = "=="
	NOT_EQ = "!="

	PLUSEQ = "+="

	COMMA     = ","
	SEMICOLON = ";"

	QUATE = "\""

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	FOR      = "FOR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywordMap = map[string]TokenType{
	"for":    FOR,
	"return": RETURN,
}

func newToken(tokenType TokenType, literal byte) Token {
	return Token{Type: tokenType, Literal: string(literal)}
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywordMap[ident]; ok {
		return tok
	}
	return IDENT
}
