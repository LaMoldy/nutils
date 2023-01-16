package utils

import (
	"bufio"
	"os"
  "runtime"
)

/**
 * @param Filepath of the template file
 */
func getConfigFromTemplate(filePath string) []string {
    file, err := os.Open(filePath)
    if err != nil {
        panic(err.Error())
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

    // Checks if the scanner has an error
    if scanner.Err() != nil {
        panic(err.Error())
    }

    return fileContent
}

func CreateTSConfig(exePath, template string) {
    file, err := os.Create("tsconfig.json")

    if err != nil {
        panic(err.Error())
    }

    defer file.Close()


    var filePath string

    if runtime.GOOS == "windows" {
        filePath = exePath[:len(exePath) - 10] + "templates\\" + template + ".txt"
    } else if runtime.GOOS == "darwin" {
      filePath = exePath[:len(exePath) - 6] + "templates/" + template + ".txt"
    }

    fileContents := getConfigFromTemplate(filePath)

    for i := 0; i < len(fileContents); i++ {
        file.WriteString(fileContents[i] + "\n")
    }
}
