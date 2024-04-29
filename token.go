package main

const (
    TokenIdentifier = iota
    TokenInteger
    TokenString
    TokenComma
    TokenColon
    TokenUnhandled
    TokenEndOfFile
)

type Token struct {
    kind int
    value string
    span Span
}

func NewToken(kind int, value string, span Span) Token {
    return Token{kind, value, span}
}

func TokenKindAsString(kind int) string {
    switch kind {
    case TokenIdentifier:
	return "Identifier"
	
    case TokenInteger:
	return "Integer"

    case TokenString:
	return "String"

    case TokenComma:
	return "Comma"

    case TokenColon:
	return "Colon"

    case TokenUnhandled:
	return "Unhandled"

    case TokenEndOfFile:
	return "EndOfFile"

    default:
	return "Unreachable"
    }
}
