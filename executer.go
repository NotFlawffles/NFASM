package main

import "os"

func Execute(path string, arguments []string) {
    cpu := NewCPU(false)
    cpu.LoadProgramFromFile(path)
    err := cpu.Run(arguments)

    if err != nil {
	panic(err)
    }

    cpu.debugger.Log("exit code: ", cpu.b)
    os.Exit(int(cpu.b))
}
