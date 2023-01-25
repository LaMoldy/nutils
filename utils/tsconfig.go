package utils 

import (
	"bufio"
	"os"
	"runtime"
)

// Opens and file and saves it line by line
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

// Creates a tsconfig file with content from specified template
func CreateTSConfig(exePath, template string) {
    // Creates the file
    file, err := os.Create("tsconfig.json")
    if err != nil {
        panic(err.Error())
    }

    // Closes the file when it is done
    defer file.Close()

    // Gets the template file and checks os's for proper slashes
    var filePath string
    if runtime.GOOS == "windows" {
        filePath = exePath[:len(exePath) - 10] + "templates\\" + template + ".txt"
    } else if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
      filePath = exePath[:len(exePath) - 6] + "templates/" + template + ".txt"
    }

    // Gets the templates content to write to the file
    fileContents := getConfigFromTemplate(filePath)

    // Writes the strings into the file
    for i := 0; i < len(fileContents); i++ {
        file.WriteString(fileContents[i] + "\n")
    }
}
