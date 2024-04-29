package main

import "fmt"

type InstructionWrapper struct {
    Instruction func(*CPU) error
}

var Instructions = []InstructionWrapper{{(*CPU).Nop}, {(*CPU).Mov}, {(*CPU).Add}, {(*CPU).Sub}, {(*CPU).Mul}, {(*CPU).Div}, {(*CPU).Rem}, {(*CPU).Or}, {(*CPU).Xor}, {(*CPU).And}, {(*CPU).Not}, {(*CPU).La}, {(*CPU).Las}, {(*CPU).Str}, {(*CPU).Syscall}, {(*CPU).Jmp}, {(*CPU).Jmpl}, {(*CPU).Push}, {(*CPU).Pop}, {(*CPU).Ret}, {(*CPU).Inc}, {(*CPU).Dec}, {(*CPU).Cmp}}

/*
/
/ Instruction set:
/	before each function call of the register, syntaxes will be explicit to cover all the cases
/	all instuctions can be marked with user states, like "mov a, 69, eq", where this instruction would only execute if the user state equal/zero is on
/	available user states marks that can be used are: [z, nz, c, o]
/						   or: [eq, ne, gt, lt]
/
/	notes:
/		<destination> can only be a register
/		<source> can be either a register, a name or an immediate value
/		cannot have multiple user states marks (for now)
/		instructions do not update user states (for now)
/
/ Bytecode format:
/	*<opcode> *<user states> <destination> <source>
/	the starred ones are required, the others depend on the instruction type
/
/ Instruction sizes (since the machine is 16 bit, the calculation unit here is word):
/	no arged: 2 words (4 bytes),  format: <opcode> <user states>
/	one arged: 3 words (6 bytes), format: <opcode> <user states> <destination/source>
/	two arged: 4 words (8 bytes), format: <opcode> <user states> <destination> <source>
/
*/

/*
/
/ Nop:
/	syntaxes:
/		nop
/
/	behavior:
/		does nothing
/
/	examples:
/		nop
/
*/
func (this *CPU) Nop() error {
    return nil
}

/*
/
/ Mov:
/	syntaxes:
/		mov destination, source
/
/	behavior:
/		moves source value into destination
/		loads source address into destination
/
/	examples:
/		mov a, b
/		mov a, message // declared somewhere as data
/		mov a, 69
/
*/
func (this *CPU) Mov() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination = source
	}
    }

    return nil
}

/*
/
/ Add:
/	syntaxes:
/		add destination, source
/
/	behavior:
/		adds source into destination
/
/	examples:
/		add a, b
/		add a, message // declared somewhere as data
/		add a, 69
/
*/
func (this *CPU) Add() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination += source
	}
    }

    return nil
}

/*
/
/ Sub:
/	syntaxes:
/		sub destination, source
/
/	behavior:
/		subtracts destination from source
/
/	examples:
/		sub a, b
/		sub a, message // declared somewhere as data
/		sub a, 69
/
*/
func (this *CPU) Sub() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination -= source
	}
    }

    return nil
}

/*
/
/ Mul:
/	syntaxes:
/		mul destination, source
/
/	behavior:
/		multiples destination by source
/
/	examples:
/		mul a, b
/		mul a, message // declared somewhere as data
/		mul a, 69
/
*/
func (this *CPU) Mul() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination *= source
	}
    }

    return nil
}

/*
/
/ Div:
/	syntaxes:
/		div destination, source
/
/	behavior:
/		divides destination by source
/
/	examples:
/		div a, b
/		div a, message // declared somewhere as data
/		div a, 69
/
*/
func (this *CPU) Div() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination /= source
	}
    }

    return nil
}

/*
/
/ Rem:
/	syntaxes:
/		rem destination, source
/
/	behavior:
/		gets the reminder of a division of destination by source
/
/	examples:
/		rem a, b
/		rem a, message // declared somewhere as data
/		rem a, 69
/
*/
func (this *CPU) Rem() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination %= source
	}
    }

    return nil
}

/*
/
/ Or:
/	syntaxes:
/		or destination, source
/
/	behavior:
/		ors destination by source (bitwise operation)
/
/	examples:
/		or a, b
/		or a, message // declared somewhere as data
/		or a, 69
/
*/
func (this *CPU) Or() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination |= source
	}
    }

    return nil
}

/*
/
/ Xor:
/	syntaxes:
/		xor destination, source
/
/	behavior:
/		xors destination by source (bitwise operation)
/
/	examples:
/		xor a, b
/		xor a, message // declared somewhere as data
/		xor a, 69
/
*/
func (this *CPU) Xor() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination ^= source
	}
    }

    return nil
}

/*
/
/ And:
/	syntaxes:
/		and destination, source
/
/	behavior:
/		ands destination by source (bitwise operation)
/
/	examples:
/		and a, b
/		and a, message // declared somewhere as data
/		and a, 69
/
*/
func (this *CPU) And() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination &= source
	}
    }

    return nil
}

/*
/
/ Not:
/	syntaxes:
/		not source
/
/	behavior:
/		nots a register, a data or an immediate value
/
/	examples:
/		not a
/		not message // declared somewhere as data
/		not 69
/
*/
func (this *CPU) Not() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	*destination = ^*destination
    }

    return nil
}

/*
/
/ La:
/	syntaxes:
/		la destination, source
/
/	behavior:
/		loads effective address in source into destination
/
/	examples:
/		la a, b
/		la a, message // declared somewhere as data
/		la 69
/
*/
func (this *CPU) La() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    *destination = this.mainMemory[source]
	}
    }

    return nil
}

/*
/
/ Las:
/	syntaxes:
/		las destination
/
/	behavior:
/		loads effective address in destination into destination
/
/	examples:
/		las a
/
*/
func (this *CPU) Las() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	*destination = this.mainMemory[*destination]
    }

    return nil
}

/*
/
/ Str:
/	syntaxes:
/		str destination, source
/
/	behavior:
/		stores source into address of destination
/
/	examples:
/		str a, b
/		str a, message // declared somewhere as data
/		str a, 69
/
*/
func (this *CPU) Str() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    this.mainMemory[*destination] = source
	}
    }

    return nil
}

/*
/
/ Syscall:
/	syntaxes:
/		syscall
/
/	behavior:
/		triggers the software interrupt, system call
/
/	examples:
/		syscall
/
*/
func (this *CPU) Syscall() error {
    switch this.a {
    case SyscallReset:
	break

    case SyscallExit:
	this.rsr ^= ReservedStateRunning
	break

    case SyscallWrite:
	for i := 0; i < int(this.d); i++ {
	    if this.b == 1 {
		fmt.Printf("%c", this.mainMemory[int(this.c) + i])
	    }
	}

	break
    }

    return nil
}

/*
/
/ Jmp:
/	syntaxes:
/		jmp source
/
/	behavior:
/		sets the reserved ip (instruction pointer) value to source
/
/	examples:
/		jmp a
/		jmp puts // a label
/		jmp 69
/
*/
func (this *CPU) Jmp() error {
    source, err := this.GetSource()

    if err != nil {
	return err
    } else {
	this.ip = source
    }

    return nil
}

/*
/
/ Jmpl:
/	syntaxes:
/		jmpl source
/
/	behavior:
/		sets lr (link register) to the reserved ip (instruction pointer) and sets the reserved ip (instruction pointer) value to source
/
/	examples:
/		jmpl a
/		jmpl puts // a label
/		jmpl 69
/
*/
func (this *CPU) Jmpl() error {
    source, err := this.GetSource()

    if err != nil {
	return err
    } else {
	this.lr = this.ip
	this.ip = source
    }

    return nil
}

/*
/
/ Push:
/	syntaxes:
/		push source
/
/	behavior:
/		pushes source onto the stack, decrementing the reserved sp (stack pointer) value by one
/
/	examples:
/		push a
/		push message // declared somewhere as data
/		push 69
/
*/
func (this *CPU) Push() error {
    source, err := this.GetSource()
    
    if err != nil {
	return err
    }

    this.sp--

    this.mainMemory[this.sp] = source

    return nil
}

/*
/
/ Pop:
/	syntaxes:
/		pop destination
/
/	behavior:
/		pops the last value from the stack into destination, incrementing the reserved sp (stack pointer) value by one
/
/	examples:
/		pop a
/
*/
func (this *CPU) Pop() error {
    destination, err := this.GetRegisterReferenceFromEncoding()
    
    if err != nil {
	return err
    }

    *destination = this.mainMemory[this.sp]
    this.sp++

    return nil
}

/*
/
/ Ret:
/	syntaxes:
/		ret
/
/	behavior:
/		sets the reserved ip (instruction pointer) to the reserved lr (link register)
/
/	examples:
/		ret
/
*/
func (this *CPU) Ret() error {
    this.ip = this.lr
    return nil   
}

/*
/
/ Inc:
/	syntaxes:
/		inc destination
/
/	behavior:
/		increments destination by one
/
/	examples:
/		inc a
/
*/
func (this *CPU) Inc() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	*destination++
    }

    return nil   
}

/*
/
/ Dec:
/	syntaxes:
/		dec destination
/
/	behavior:
/		decrements destination by one
/
/	examples:
/		dec a
/
*/
func (this *CPU) Dec() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	*destination--
    }

    return nil   
}

/*
/
/ Cmp:
/	syntaxes:
/		cmp destination, source
/
/	behavior:
/		compares destination with source, changing user states
/
/	examples:
/		cmp a, b
/		cmp a, index // declared somewhere as data
/		cmp a, 69
/
*/
func (this *CPU) Cmp() error {
    destination, err := this.GetRegisterReferenceFromEncoding()

    if err != nil {
	return err
    } else {
	source, err := this.GetSource()
	
	if err != nil {
	    return err
	} else {
	    result := int16(*destination - source)

	    if result == 0 {
		this.usr |= UserStateZero
	    } else if result > 0 {
		this.usr |= UserStateCarry
	    } else if result < 0 {
		this.usr |= UserStateOverflow
	    }
	}
    }

    return nil
}
