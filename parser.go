package main

import (
	"errors"
)

type Parser struct {
	lexer   *Lexer
	current Token
}

func NewParser(lexer *Lexer) (Parser, error) {
	token, err := lexer.LexNext()
	return Parser{lexer, token}, err
}

func (this *Parser) Parse() ([]Ast, error) {
	var tree []Ast

	for this.current.kind != TokenEndOfFile {
		if ast, err := this.ParseNext(); err != nil {
			return tree, err
		} else {
			tree = append(tree, ast)
		}
	}

	return tree, nil
}

func (this *Parser) ParseNext() (Ast, error) {
	switch this.current.kind {
	case TokenIdentifier:
		return this.ParseIdentifier()

	default:
		return NewAst(0, "", "", ""), nil
	}
}

func (this *Parser) ParseIdentifier() (Ast, error) {
	if this.IsInstruction() {
		return this.ParseInstruction()
	} else if this.IsDeclarator() {
		return this.ParseDeclaration()
	} else {
		return this.ParseName()
	}
}

func (this *Parser) IsInstruction() bool {
	for index := range OpcodeCount {
		if opcode, err := OpcodeAsString(uint16(index)); err != nil {
			return false
		} else {
			if this.current.value == opcode {
				return true
			}
		}
	}

	return false
}

func (this *Parser) ParseInstruction() (Ast, error) {
	var ast Ast
	var err error
	var parsed bool

	for _, value := range []string{"nop", "syscall", "ret"} {
		if this.current.value == value {
			ast, err = this.ParseNoArged()
			parsed = true
			break
		}
	}

	for _, value := range []string{"not", "jmp", "jmpl", "push", "las", "inc", "dec"} {
		if this.current.value == value && !parsed {
			ast, err = this.ParseOneArgedSource()
			parsed = true
			break
		}
	}

	for _, value := range []string{"pop"} {
		if this.current.value == value && !parsed {
			ast, err = this.ParseOneArgedDestination()
			parsed = true
			break
		}
	}

	for _, value := range []string{"mov", "add", "sub", "mul", "div", "rem", "or", "xor", "and", "la", "str", "cmp"} {
		if this.current.value == value && !parsed {
			ast, err = this.ParseTwoArged()
			break
		}
	}

	if this.current.kind == TokenComma {
		if _, err := this.Eat([]int{TokenComma}); err != nil {
			return ast, err
		}

		if this.IsUserState() {
			ast.userStates |= this.GetUserState()
		}

		this.Eat([]int{TokenIdentifier})
	}

	if ast.destination == "" {
		ast.destination = "<no value>"
	}

	if ast.source == "" {
		ast.source = "<no value>"
	}

	return ast, err
}

func (this *Parser) ParseNoArged() (Ast, error) {
	if name, err := this.Eat([]int{TokenIdentifier}); err != nil {
		return NewAst(0, "", "", ""), err
	} else {
		return NewAst(AstInstruction, name.value, "", ""), nil
	}
}

func (this *Parser) ParseOneArgedSource() (Ast, error) {
	var ast Ast
	name, err := this.Eat([]int{TokenIdentifier})

	if err != nil {
		return ast, err
	}

	source, err := this.Eat([]int{TokenIdentifier, TokenInteger})

	if source.kind == TokenInteger {
		ast.userStates |= 0x0001
	}

	if err != nil {
		return ast, err
	}

	ast.kind = AstInstruction
	ast.name = name.value
	ast.source = source.value

	return ast, err
}

func (this *Parser) ParseOneArgedDestination() (Ast, error) {
	var ast Ast
	name, err := this.Eat([]int{TokenIdentifier})

	if err != nil {
		return ast, err
	}

	destination, err := this.Eat([]int{TokenIdentifier})

	if err != nil {
		return ast, err
	}

	ast.kind = AstInstruction
	ast.name = name.value
	ast.source = destination.value

	return ast, err
}

func (this *Parser) ParseTwoArged() (Ast, error) {
	var ast Ast
	name, err := this.Eat([]int{TokenIdentifier})

	if err != nil {
		return ast, err
	}

	destination, err := this.Eat([]int{TokenIdentifier})

	if err != nil {
		return ast, err
	}

	if _, err := this.Eat([]int{TokenComma}); err != nil {
		return ast, err
	}

	source, err := this.Eat([]int{TokenIdentifier, TokenInteger})

	if source.kind == TokenInteger {
		ast.userStates |= 0x0001
	}

	if err != nil {
		return ast, err
	}

	ast.kind = AstInstruction
	ast.name = name.value
	ast.destination = destination.value
	ast.source = source.value

	return ast, err
}

func (this *Parser) IsUserState() bool {
	if this.current.kind == TokenIdentifier {
		for _, value := range []string{"eq", "ne", "gt", "lt", "z", "nz", "c", "o"} {
			if this.current.value == value {
				return true
			}
		}
	}

	return false
}

func (this *Parser) GetUserState() uint16 {
	switch this.current.value {
	case "eq":
	    return 0x0002

	case "ne":
	    return 0x0002

	case "gt":
	    return 0x0004

	case "lt":
	    return 0x0008

	case "z":
	    return 0x0002

	case "nz":
	    return 0x0002

	case "c":
	    return 0x0004

	case "o":
	    return 0x0008

	default:
	    return UserStateDefault
	}
}

func (this *Parser) IsDeclarator() bool {
	for _, value := range []string{"db", "dw"} {
		if this.current.value == value {
			return true
		}
	}

	return false
}

func (this *Parser) ParseDeclaration() (Ast, error) {
	var ast Ast
	var err error

	name, err := this.Eat([]int{TokenIdentifier})

	if err != nil {
		return ast, err
	}

	ast.source = TokenKindAsString(this.current.kind)

	value, err := this.Eat([]int{TokenString, TokenInteger})

	if err != nil {
		return ast, err
	}

	ast.kind = AstDeclaration
	ast.name = name.value
	ast.destination = value.value

	if value.kind == TokenInteger {
		ast.userStates |= UserStateImmediate
	}

	return ast, err
}

func (this *Parser) ParseName() (Ast, error) {
	var ast Ast
	var err error

	name, err := this.Eat([]int{TokenIdentifier})

	if err != nil {
		return ast, err
	}

	if _, err := this.Eat([]int{TokenColon}); err != nil {
		return ast, err
	}

	ast = NewAst(AstLabel, name.value, "", "")
	return ast, err
}

func (this *Parser) Eat(tokenKinds []int) (Token, error) {
	for _, kind := range tokenKinds {
		if this.current.kind == kind {
			token := this.current
			this.current, _ = this.lexer.LexNext()
			return token, nil
		}
	}

	return NewToken(0, "", NewSpan("", 0, 0, 0, 0)), errors.New("unexpected token: " + this.current.value)
}
