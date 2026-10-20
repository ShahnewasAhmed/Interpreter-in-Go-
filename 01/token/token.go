package token

// We use a separate TokenType instead of using string directly, so the type system can distinguish a token type from an ordinary string.
type TokenType string

// A token has a token type and the exact text we read from the source code.
// The literal is kept as a string for now; later the parser will turn things like "5" into the appropriate value.
type Token struct {
	//Token Type → What kind/category is this?
	//Literal    → What exact text did we read?
	Type    TokenType
	Literal string
}

// Store keywords here so we can distinguish them from normal identifiers.
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// If the identifier is a keyword, return its keyword token.
// Otherwise, treat it as a normal identifier.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Define the token types used by the lexer.
// Tokens can represent identifiers, literals, keywords, operators, delimiters, etc.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	FLOAT  = "FLOAT"
	STRING = "STRING"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)
