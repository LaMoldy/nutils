package main

import (
	"fmt"
	"os"

	utils "github.com/lamoldy/nutils/utils"
)

const version string = "1.0.0"

func displayAsciiTItle() {
    fmt.Println("************************************************")
    fmt.Println("*  _   _ _    _ _______ _____ _       _____    *")
    fmt.Println("* | \\ | | |  | |__   __|_   _| |     / ____|   *")
    fmt.Println("* |  \\| | |  | |  | |    | | | |    | (___     *")
    fmt.Println("* |     | |  | |  | |    | | | |     \\___ \\    *")
    fmt.Println("* | |\\  | |__| |  | |   _| |_| |____ ____) |   *")
    fmt.Println("* |_| \\_|\\____/   |_|  |_____|______|_____/    *")
    fmt.Println("*                                              *")
    fmt.Println("************************************************")
}

func initMainMenu() {
    var choice string

    for {
        fmt.Println("\nNUTils Menu\t\tVersion: 1.0.0")
        fmt.Println("[1] TypeScript Configuration")
        fmt.Println("[2] Display Help")
        fmt.Print("\nEnter: ")
        fmt.Scanln(&choice)

        if choice == "1" {
            var templateName string
            fmt.Print("\nEnter template: ")
            fmt.Scanln(&templateName)

            exePath, err := os.Executable()
            if err != nil {
                panic(err.Error())
            }
            utils.CreateTSConfig(exePath, templateName)
            break
        } else if choice == "2" {
            break
        } else {
            fmt.Println()
        }
    }
}

func main() {
    // Checks argument length
    if len(os.Args) > 2 {
        fmt.Println("Too many arguments have been given!")
        fmt.Println("Please try again.")
    } else if len(os.Args) < 2 {
        initMainMenu()
    } else {
        if os.Args[1] == "-version" {
            fmt.Println("NUtils " + version)
        } else if os.Args[1] == "-help" {
            fmt.Println("-version")
            fmt.Println("\tDisplays the version")
        } else {
            fmt.Println("Error, invalid argument: " + os.Args[1])
            fmt.Println("Use '-help' for a list of arguments")
        }
    }
}
