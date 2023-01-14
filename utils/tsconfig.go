package utils

import (
	"bufio"
	"fmt"
	"os"
)

func getConfigFromTemplate(filePath string) []string {
    file, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }

    // Closes file when everything is finished
    defer file.Close()

    // Creates a new scanner to read the file
    scanner := bufio.NewScanner(file)

    // Reads the file and saves into array
    var fileContent []string
    for scanner.Scan() {
        fileContent = append(fileContent, scanner.Text())
    }

    if scanner.Err() != nil {
        panic(err)
    }

    return fileContent
}

func CreateTSConfig(exePath, template string) {
    file, err := os.Create("tsconfig.json")

    if err != nil {
        panic(err)
    }

    defer file.Close()

    filePath := exePath + "templates\\" + template + ".txt"
    fmt.Println(getConfigFromTemplate(filePath))
}
