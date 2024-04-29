package main

import (
	"os"
	"strings"
	"encoding/binary"
)

func Compile(path string) error {
    buffer, err := ReadFile(path)

    if err != nil {
	return err
    }

    lexer := NewLexer(path, buffer)
    parser, err := NewParser(&lexer)

    if err != nil {
	return err
    }

    tree, err := parser.Parse()

    if err != nil {
	return err
    }

    generation, err := Generate(tree)

    if err != nil {
	return err
    }

    /* get the output file name from the source file name without the extension, assumed (.s) */
    pathSplitted := strings.Split(path, ".")
    outName := pathSplitted[:len(pathSplitted) - 1][0]

    file, err := os.Create(outName)

    if err != nil {
	return err
    }

    defer file.Close()

    err = binary.Write(file, binary.LittleEndian, generation)

    if err != nil {
	return err
    }

    return nil
}
