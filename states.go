package main

const (
    // User
    UserStateDefault = 0x0000
    UserStateImmediate = 0x0001
    UserStateZero = 0x0002
    UserStateCarry = 0x0004
    UserStateOverflow  = 0x0008
    UserStateCount = 0x0005

    // Reserved
    ReservedStateDefault = 0x0000
    ReservedStateRunning = 0x0001
    ReservedStateCount = 0x0002
)
