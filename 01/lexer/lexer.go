package lexer

import "interpreter-go/01/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// Here we are converting our input into a Lexer (new data type).
// We only initialize its input field.
// The other fields position, readPosition and ch are automatically initialized to 0.
// Then readChar() is called to load the first character into l.ch.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Generating a Token from a token type and a character/literal.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Increasing the position and moving to the next char.
// A common question is: why do we use readPosition for l.ch?
// When we initialize our Lexer, position and readPosition are both 0.
// So we use readPosition to read the current char, then set position to that readPosition to keep track of it, and finally increase readPosition by 1 to point to the next char.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// Checks for operators that are made of 2 characters, like == and !=.
// It looks at the next character without moving the current position.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Checking whether the character is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isDot(ch byte) bool {
	return '.' == ch
}

// Reading until the end of a number.
func (l *Lexer) readNumber() (string, token.TokenType) {
	position := l.position

	dotCount := 0
	// readNumber() keeps moving forward while the current character is a digit.
	for isDigit(l.ch) || isDot(l.ch) {
		if isDot(l.ch) {
			dotCount++
		}
		l.readChar()
	}
	literal := l.input[position:l.position]

	if dotCount == 0 {
		return literal, token.INT
	}

	if dotCount > 1 {
		return literal, token.ILLEGAL
	}

	if literal[len(literal)-1] == '.' {
		return literal, token.ILLEGAL
	}

	return literal, token.FLOAT
}

// Skipping whitespace, new lines, tabs, etc.
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Checking whether the character is a letter or underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// If the current character is a letter, keep reading until the identifier ends.
func (l *Lexer) readIdentifier() string {
	position := l.position

	// Basically, we read the whole word so we can find out whether it is an identifier or a keyword.
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()

		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

// Returns the Token for the current character/token.
// At the end, l.readChar() moves the Lexer towards the next character.
// For identifiers and numbers, readIdentifier/readNumber already move it forward, so we return early in those cases.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) || isDot(l.ch) {
			tok.Literal, tok.Type = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
