package main

import "errors"

const (
	RegisterEncodingA     = 0
	RegisterEncodingB     = 1
	RegisterEncodingC     = 2
	RegisterEncodingD     = 3
	RegisterEncodingE     = 4
	RegisterEncodingIP    = 5
	RegisterEncodingLr    = 6
	RegisterEncodingDP    = 7
	RegisterEncodingHP    = 8
	RegisterEncodingSP    = 9
	RegisterEncodingUSR   = 10
	RegisterEncodingRSR   = 11
	RegisterEncodingOPAR  = 12
	RegisterEncodingUSAR  = 13
	RegisterEncodingCount = 14
)

/*
/
/ Registers:
/	all registers are 16 bit
/
/ They consist of:
/	a, b, c, d, e -> general purpose registers (user, accessible)
/	ip (instruction pointer) (reserved, inaccessible) (for now)
/	lr (link register) (reserved, accessible)
/	dp (data pointer) (reserved, inaccessible)
/	hp (heap pointer) (reserved, inaccessible) // too lazy to implement a heap system
/	sp (stack pointer) (reserved, inaccessible) (for now)
/	usr (user states register) (reserved, inaccessible) (for now)
/	rsr (reserverd states register) (reserved, inaccessible)
/	opar (opcode/operand addressing register) (reserved, inaccessible)
/	usar (user states addressing register) (reserved, inaccessible)
/
/ Breaking down (listing bitwisely in the corresponding order): (for now)
/	the usr (user states register) holds these flags: [overflow, carry, zero, immediate]
/	the rsr (reserved states register) hold these flags: [running]
/	the opar (opcode/operand addressing register) is a temporary register used to store opcodes/operands in the Fetch/Decode processes
/	the usar (user states addressing register) is a reserved register used to store the states of the current running instruction (in Decode process)
/
*/

func RegisterAsInt(register string) (uint16, error) {
	for index, value := range []string{"a", "b", "c", "d", "e", "ip", "lr", "dp", "hp", "sp", "usr", "rsr", "opar", "usar"} {
		if register == value {
			return uint16(index), nil
		}
	}

	return 0, errors.New("register encoding out of bounds")
}

func RegisterAsString(register uint16) (string, error) {
    for index1 := range RegisterEncodingCount {
	if register == uint16(index1) {
	    for index2, value := range []string{"a", "b", "c", "d", "e", "ip", "lr", "dp", "hp", "sp", "usr", "rsr", "opar", "usar"} {
		if index1 == index2 {
		    return value, nil
		}
	    }
	}
    }

    return "", errors.New("register encoding out of bounds")
}
