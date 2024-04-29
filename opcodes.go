package main

import "errors"

const (
    OpcodeNop = 0
    OpcodeMov = 1
    OpcodeAdd = 2
    OpcodeSub = 3
    OpcodeMul = 4
    OpcodeDiv = 5
    OpcodeRem = 6
    OpcodeOr = 7
    OpcodeXor = 8
    OpcodeAnd = 9
    OpcodeNot = 10
    OpcodeLa = 11
    OpcodeLas = 12
    OpcodeStr = 13
    OpcodeSyscall = 14
    OpcodeJmp = 15			// we need more opcodes
    OpcodeJmpl = 16
    OpcodePush = 17
    OpcodePop = 18
    OpcodeRet = 19
    OpcodeInc = 20
    OpcodeDec = 21
    OpcodeCmp = 22
    OpcodeCount = 23
)

func OpcodeAsString(opcode uint16) (string, error) {
    if opcode >= OpcodeCount {
	return "", errors.New("opcode out of bounds")
    } else {
	return []string{"nop", "mov", "add", "sub", "mul", "div", "rem", "or", "xor", "and", "not", "la", "las", "str", "syscall", "jmp", "jmpl", "push", "pop", "ret", "inc", "dec", "cmp"}[opcode], nil
    }
}

func OpcodeAsInt(opcode string) (uint16, error) {
    for index, value := range []string{"nop", "mov", "add", "sub", "mul", "div", "rem", "or", "xor", "and", "not", "la", "las", "str", "syscall", "jmp", "jmpl", "push", "pop", "ret", "inc", "dec", "cmp"} {
	if opcode == value  {
	    return uint16(index), nil
	}
    }

    return 0, errors.New("opcode out of bounds")
}
