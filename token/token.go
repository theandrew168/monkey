// NOTE: will likely collide with desired "token" var name
package token

// NOTE: seems redundant when already in the "token" namespace
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
