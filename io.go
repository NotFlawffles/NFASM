package main

import (
	"os"
)

func ReadFile(path string) (string, error) {
    file, err := os.Open(path)

    if err != nil {
	return "", err
    }

    defer file.Close()

    fileInfo, err := file.Stat()

    if err != nil {
	return "", err
    }

    fileSize := fileInfo.Size()
    buffer := make([]byte, fileSize)
    file.Read(buffer)
    
    return string(buffer), err
}
