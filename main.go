package main

import (
	"fmt"
	"os"
)

const MinimumRequiredArgsCount int = 3

func Usage(executableName string) {
    fmt.Printf("usage: %s [com|exe]\n", executableName)
    os.Exit(1)
}

func main() {
    if len(os.Args) < MinimumRequiredArgsCount {
	Usage(os.Args[0])
    }

    switch os.Args[1] {
    case "com":
	Compile(os.Args[2])
	break

    case "exe":
	Execute(os.Args[2], os.Args[2:])
	break

    default:
	Usage(os.Args[0])
	break
    }
}
