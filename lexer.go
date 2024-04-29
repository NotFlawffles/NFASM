package main

type Lexer struct {
    content string
    span Span
    current byte
}

func NewLexer(stream, content string) Lexer {
    return Lexer{
	content,
	NewSpan(stream, 0, 1, 1, uint64(len(content))),
	content[0],
    }
}

func (this *Lexer) LexNext() (Token, error) {
    switch this.SkipWhitespace() {
    case '_':
	return this.LexIdentifier()

    case '"':
	return this.LexString()

    case ',':
	return this.LexComma()

    case ':':
	return this.LexColon()

    default:
	if this.IsAlpha() {
	    return this.LexIdentifier()
	} else if this.IsDigit() {
	    return this.LexInteger()
	}
    }

    return this.LexEndOfFile()
}

func (this *Lexer) IsAlpha() bool {
    return (this.current >= 'A' && this.current <= 'Z') || (this.current >= 'a' && this.current <= 'z')
}

func (this *Lexer) IsAlnum() bool {
    return this.IsAlpha() || this.IsDigit()
}

func (this *Lexer) IsDigit() bool {
    return this.current >= '0' && this.current <= '9'
}

func (this *Lexer) LexIdentifier() (Token, error) {
    span := this.span
    var value string

    for this.IsAlnum() || this.current == '_' {
	value += string(this.current)
	this.Advance()
    }

    return NewToken(TokenIdentifier, value, *span.WithLength(uint64(len(value)))), nil
}

func (this *Lexer) LexInteger() (Token, error) {
    span := this.span
    var value string

    for this.IsDigit() {
	value += string(this.current)
	this.Advance()
    }

    return NewToken(TokenInteger, value, *span.WithLength(uint64(len(value)))), nil
}

func (this *Lexer) LexString() (Token, error) {
    span := this.span
    var value string

    this.Advance()

    for this.current != '"' {
	value += string(this.current)
	this.Advance()
    }

    return this.AdvanceWithToken(NewToken(TokenString, value, *span.WithLength(uint64(len(value))))), nil
}

func (this *Lexer) LexComma() (Token, error) {
    return this.AdvanceWithToken(NewToken(TokenComma, string(","), this.span)), nil
}

func (this *Lexer) LexColon() (Token, error) {
    return this.AdvanceWithToken(NewToken(TokenColon, string(":"), this.span)), nil
}

func (this *Lexer) LexUnhandled() (Token, error) {
    return this.AdvanceWithToken(NewToken(TokenUnhandled, string(this.current), this.span)), nil
}

func (this *Lexer) LexEndOfFile() (Token, error) {
    return NewToken(TokenEndOfFile, "<EndOfFile>", this.span), nil
}

func (this *Lexer) SkipWhitespace() byte {
    for this.current == ' ' || this.current == '\t' || this.current == '\n' {
	if this.current == '\n' {
	    this.span.row++
	    this.span.column = 0
	}

	this.Advance()
    }

    return this.current
}

func (this *Lexer) AdvanceWithToken(token Token) Token {
    this.Advance()
    return token
}

func (this *Lexer) Advance() byte {
    this.span.column++
    this.span.index++
    
    if this.span.index == uint64(len(this.content)) {
	this.current = 0
    } else {
	this.current = this.content[this.span.index]
    }

    return this.current
}
