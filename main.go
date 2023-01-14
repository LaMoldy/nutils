package main

import (
	"fmt"
	"os"
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
        fmt.Println("There wasn't enough arguments given!")
        fmt.Println("Please try again.")
    } else {
        fmt.Println("Has proper arguments")
    }
}
