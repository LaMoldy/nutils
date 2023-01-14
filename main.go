package main

import (
	"fmt"
	"os"
	"runtime"

	utils "github.com/lamoldy/nutils/utils"
)

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
        fmt.Println("\nNUTils Menu")
        fmt.Println("[1] TypeScript Configuration")
        fmt.Println("[2] Display Help")
        fmt.Print("\nEnter: ")
        fmt.Scanln(&choice)

        if choice == "1" {
            if runtime.GOOS == "darwin" {
                exePath, err := os.Executable()
                if err != nil {
                    panic(err.Error())
                }
                
                var templateName string
                fmt.Print("\nEnter template: ")
                fmt.Scanln(&templateName)
                utils.CreateTSConfig(exePath, templateName)
            }

            if runtime.GOOS == "windows" {
                exePath := os.Args[0]
                
                var templateName string
                fmt.Print("\nEnter template: ")
                fmt.Scanln(&templateName)
                utils.CreateTSConfig(exePath[:len(exePath) - 10], templateName)
            }
            break
        } else if choice == "2" {
            break
        } else {
            fmt.Println()
        }
    }
}

func main() {
    displayAsciiTItle()

    for index, arg := range os.Args {
        fmt.Printf("Index: %v Argument: %v\n", index, arg)
    }

    // Checks argument length
    if len(os.Args) > 2 {
        fmt.Println("Too many arguments have been given!")
        fmt.Println("Please try again.")
    } else if len(os.Args) < 2 {
        initMainMenu()
    } else {
        fmt.Println("Here")
    }
}
