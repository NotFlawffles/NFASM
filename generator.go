package main

import (
	"fmt"
	"strconv"
)

func ReferenceLabel(labels *[]Label, source string) *Label {
    for _, label := range *labels {
	if label.name == source {
	    return &label
	}
    }

    return nil
}

func CalculateSyntaxSize(ast *Ast) uint16 {
    var size uint16

    switch ast.kind {
    case AstInstruction:
	size = 2

	if ast.destination != "<no value>" {
	    size += 1
	}

	if ast.source != "<no value>" {
	    size += 1
	}

	break

    case AstDeclaration:
	if ast.source == "String" {
	    size = uint16(len(ast.destination))
	} else if ast.source == "Integer" {
	    size = 1
	}

	break

    default:
	size = 0
	break
    }

    return size
}

func CollectLabels(tree *[]Ast, labels *[]Label) {
    var generationSize uint16

    for _, ast := range *tree {
	if ast.kind == AstLabel {
	    *labels = append(*labels, NewLabel(ast.name, generationSize))
	} else {
	    generationSize += CalculateSyntaxSize(&ast)
	}
    }
}

func Generate(tree []Ast) ([]uint16, error) {
    var generation []uint16
    var labels []Label

    CollectLabels(&tree, &labels)

    for _, ast := range tree {
	switch ast.kind {
	case AstInstruction:
	    opcode, err := OpcodeAsInt(ast.name)

	    if err != nil {
		return generation, err
	    }

	    generation = append(generation, opcode)
	    generation = append(generation, ast.userStates)

	    if ast.destination != "<no value>" {
		register, err := RegisterAsInt(ast.destination)

		if err != nil {
		    return generation, err
		}

		generation = append(generation, register)
	    }

	    if ast.source != "<no value>" {
		if ast.userStates & UserStateImmediate != 0x0000 {
		    convert, err := strconv.ParseUint(ast.source, 10, 16)

		    if err != nil {
			return generation, err
		    }

		    generation = append(generation, uint16(convert))
		} else {
		    register, err := RegisterAsInt(ast.source)

		    if err != nil {
			label := ReferenceLabel(&labels, ast.source)

			if label != nil {
			    if ast.destination != "<no value>" {
				generation[len(generation) - 2] |= UserStateImmediate
			    } else {
				generation[len(generation) - 1] |= UserStateImmediate
			    }
			    generation = append(generation, label.address)
			} else {
			    fmt.Println("NOT IMPLEMENTED: label not found: " + ast.source)
			}
		    } else {
			generation = append(generation, register)
		    }
		}
	    }

	    break

	case AstLabel:
	    break

	case AstDeclaration:
	    if ast.userStates & UserStateImmediate != 0x0000 {
		convert, err := strconv.ParseUint(ast.destination, 10, 16)

		if err != nil {
		    return generation, err
		}

		generation = append(generation, uint16(convert))
	    } else {
		bytes := []byte(ast.destination)

		for _, b := range bytes {
		    generation = append(generation, uint16(b))
		}

		break
	    }

	default:
	    break
	}
    }

    return generation, nil
}
