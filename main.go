package main

import (
	"flag"
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
        fmt.Println("\nNUtils Menu\t\tVersion: 1.0.0")
        fmt.Println("[1] TypeScript Configuration")
        fmt.Println("[2] Pomodoro Timer")
        fmt.Println("[3] Exit")
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
            var duration string
            fmt.Println("\nEnter duration of timer: ")
            fmt.Scanln(&duration)
            utils.StartTimmer(duration)
        } else if choice == "3" {
            os.Exit(3)
        } else {
            fmt.Println()
        }
    }
}

func main() {
    // Creates the flags
    templateFlag := flag.String("template", "", "The template you would like to use")
    versionFlag := flag.Bool("version", false, "See the current version of NUtils")
    pomodoroFlag := flag.Int64("pomodoro", 0, "The time you would like to set the timer for")

    // Parses the flags into the variables
    flag.Parse()

    if *templateFlag != "" {
        // Gets the exe location
        exePath, err := os.Executable()
        if err != nil {
            panic(err.Error)
        }
        utils.CreateTSConfig(exePath, *templateFlag)
    }

    if *versionFlag {
      fmt.Println("NUtils " + version)
    }

    if *pomodoroFlag != 0 {
        utils.StartTimmer(fmt.Sprint(*pomodoroFlag))
    }

    // Checks argument length and starts the menu if the arguments are not more then 2
    if len(os.Args) < 2 {
        initMainMenu()
    }
}
