package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type CPU struct {
	mainMemory                                                           []uint16
	videoMemory                                                          []uint16
	programSize, a, b, c, d, e, ip, lr, dp, hp, sp, usr, rsr, opar, usar uint16
	registers                                                            []*uint16
	debugger                                                             Debugger
}

func NewCPU(debug bool) *CPU {
	return &CPU{
		make([]uint16, MemorySize),
		make([]uint16, VideoMemorySize),
		0, 0, 0, 0, 0, 0, SegmentTextStart, SegmentTextStart, SegmentDataSize, SegmentHeapStart, SegmentStackStart, UserStateDefault, ReservedStateDefault, 0, 0,
		make([]*uint16, 14),
		NewDebugger(debug),
	}
}

/* initializes basic registering system, as this.registers is used among different calls, *required to be called before running */
func (this *CPU) LoadRegisters() {
	this.registers = []*uint16{&this.a, &this.b, &this.c, &this.d, &this.e, &this.ip, &this.lr, &this.dp, &this.hp, &this.sp, &this.usr, &this.rsr, &this.opar, &this.usar}
	this.debugger.Log("loaded registers: ", this.registers)
}

func (this *CPU) LoadArguments(arguments []string) {
    for _, argument := range arguments {
	address := this.Allocate(uint16(len(argument)))

	for index, character := range argument {
	    this.mainMemory[int(address) + index] = uint16(character)
	}

	this.sp--
	this.mainMemory[this.sp] = address
    }

    this.sp--
    this.mainMemory[this.sp] = uint16(len(arguments))
}

func (this *CPU) Allocate(size uint16) uint16 {
    address := this.hp
    this.dp += size + 1
    return address
}

func (this *CPU) ClearMemory() {
	for i := 0; i < MemorySize; i++ {
		this.mainMemory[i] = 0x0000
	}

	this.debugger.Log("cleared memory: [0:", MemorySize, "]")
}

func (this *CPU) ClearProgram() {
	for i := 0; i < SegmentTextSize; i++ {
		this.mainMemory[i] = 0x0000
	}

	this.debugger.Log("cleared text segment: [", SegmentTextStart, ":", SegmentTextSize, "]")
}

func (this *CPU) LoadProgramFromMemory(program []uint16) error {
	if len(program) >= SegmentTextSize {
		this.debugger.Log("failed to load program: ", program)
		return errors.New("failed to load program: len(program) >= MemorySize")
	} else {
		for index, value := range program {
			this.mainMemory[SegmentTextStart+index] = value
		}

		this.programSize = uint16(len(program))
		this.debugger.Log("loaded program: ", program, "size: ", len(program))
		return nil
	}
}

func (this *CPU) LoadProgramFromFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()
	program := make([]uint16, fileSize/2)
	err = binary.Read(file, binary.LittleEndian, &program)

	if err != nil {
		return err
	}

	this.LoadProgramFromMemory(program)
	return nil
}

func (this *CPU) Fetch() (uint16, error) {
	if this.ip >= SegmentTextStart+this.programSize {
		return 0, errors.New("ip out of bounds")
	}

	this.opar = this.mainMemory[SegmentTextStart+this.ip]
	this.debugger.Log("at:", this.ip, "fetched:", this.mainMemory[this.ip])
	this.ip++
	return this.opar, nil
}

func (this *CPU) Decode() {
	this.usar = this.mainMemory[SegmentTextStart+this.ip]
	this.ip++
}

func (this *CPU) Execute() {
	if this.Conditioned() {
		if this.UserStatesMatches() {
			Instructions[this.opar].Instruction(this)
		}
	} else {
		Instructions[this.opar].Instruction(this)
	}
}

func (this *CPU) Run(args []string) error {
	this.LoadRegisters()
	this.LoadArguments(args)
	this.rsr |= ReservedStateRunning

	for this.rsr&ReservedStateRunning != 0x0000 {
		_, err := this.Fetch()

		if err != nil {
			if this.debugger.enabled {
				this.debugger.Log("program halted")
				break
			} else {
				this.ip -= 1
			}
		}

		this.Decode()
		this.Execute()
	}

	return this.debugger.LogRegisters(&this.registers)
}

func (this *CPU) Conditioned() bool {
	return this.usar&0xfffe != 0x0000
}

func (this *CPU) UserStatesMatches() bool {
	return (this.usar&0xfffe)&(this.usr&0xfffe) != 0x0000
}

func (this *CPU) GetRegisterReferenceFromEncoding() (*uint16, error) {
	if this.registers[0] == nil {
		return nil, errors.New("failed to get register value: registers not loaded")
	} else {
		register, _ := this.Fetch()
		return this.registers[register], nil
	}
}

func (this *CPU) GetSource() (uint16, error) {
	if this.usar&UserStateImmediate != 0x0000 {
		return this.Fetch()
	} else {
		register, err := this.GetRegisterReferenceFromEncoding()

		if err != nil {
			return 0, err
		} else {
			return *register, nil
		}
	}
}

func (this *CPU) Debug() {
	fmt.Println("registers:\n\ta: ", this.a, "\n\tb: ", this.b, "\n\tc: ", this.c, "\n\td: ", this.d, "\n\te: ", this.e, "\n\tip: ", this.ip, "\n\tdp: ", this.dp, "\n\thp: ", this.hp, "\n\tsp: ", this.sp, "\n\tusr (states, user): ", this.usr, "\n\trsr (states, reserved): ", this.rsr, "\n\topar (opcode, addressing): ", this.opar, "\n\tusar (states, addressing): ", this.usar)
}
