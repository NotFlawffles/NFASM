package main

const (
    AstInstruction = iota
    AstLabel
    AstDeclaration
)

type Ast struct {
    kind int
    name, destination, source string
    userStates uint16
}

func NewAst(kind int, name, destination, source string) Ast {
    return Ast{kind, name, destination, source, 0}
}
