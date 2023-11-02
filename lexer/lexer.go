package lexer

type Lexer struct {
	input    string
	position int  // current position in input (points to current char)
	forward  int  // current reading position in input (after current char)
	ch       byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readNextChar() // initialize l.ch
	return l
}

func (l *Lexer) readNextChar() {
	if l.forward >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.forward]
	}
	l.position = l.forward
	l.forward++
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' { // check if the next character is '=', namely '=='
			l.readNextChar()
			literal := "=="
			tok = Token{Type: EQ, Literal: literal}
		} else {
			tok = newToken(ASSIGN, l.ch)
		}
	case '+':
		if l.peekChar() == '=' { // check if the next character is '=', namely '+='
			l.readNextChar()
			literal := "+="
			tok = Token{Type: PLUSEQ, Literal: literal}
		} else {
			tok = newToken(PLUS, l.ch)
		}
	case '-':
		tok = newToken(MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' { // check if the next character is '=', namely '!='
			l.readNextChar()
			literal := "!="
			tok = Token{Type: NOT_EQ, Literal: literal}
		} else {
			tok = newToken(BANG, l.ch)
		}
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case '/':
		tok = newToken(SLASH, l.ch)
	case '<':
		if l.peekChar() == '=' { // check if the next character is '=', namely '<='
			l.readNextChar()
			literal := "<="
			tok = Token{Type: LTEQ, Literal: literal}
		} else {
			tok = newToken(LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' { // check if the next character is '=', namely '>='
			l.readNextChar()
			literal := ">="
			tok = Token{Type: GTEQ, Literal: literal}
		} else {
			tok = newToken(GT, l.ch)
		}
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal) // check if the identifier is a keyword
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		}
	}

	l.readNextChar()
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readNextChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readNextChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readNextChar()
	}
}

// peekChar returns the next character in the input without advancing the lexer's position.
func (l *Lexer) peekChar() byte {
	if l.forward >= len(l.input) {
		return 0
	} else {
		return l.input[l.forward]
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
