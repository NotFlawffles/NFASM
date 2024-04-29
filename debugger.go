package main

import "fmt"

type Debugger struct {
    enabled bool
}

func NewDebugger(enabled bool) Debugger {
    return Debugger{enabled}
}

func (this *Debugger) Log(informations... any) {
    if !this.enabled {
	return
    }

    fmt.Println(informations...)
}

func (this *Debugger) LogRegisters(registers *[]*uint16) error {
    if !this.enabled {
	return nil
    }

    for index, value := range *registers {
	registerAsString, err := RegisterAsString(uint16(index))

	if err != nil {
	    return err
	}

	this.Log(registerAsString + ":", *value)
    }

    return nil
}
