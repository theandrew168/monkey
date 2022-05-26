package token

type Token struct {
	Kind    string
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]string{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) string {
	if kind, ok := keywords[ident]; ok {
		return kind
	}
	return IDENT
}
