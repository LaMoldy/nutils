package utils

import (
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func StartTimmer(duration string) {
    minutes, err := strconv.Atoi(duration)
    if err != nil {
        panic(err.Error())
    }

    timer := time.NewTimer(time.Duration(minutes) * time.Minute)
    <- timer.C

    displayWindowsNotification()
}


func displayWindowsNotification() {
    exePath, err := os.Executable()
    if err != nil {
        panic(err.Error())
    }

    var iconPath string
    if runtime.GOOS == "windows" {
        iconPath = exePath[:len(exePath) - 10] + "assets\\icon.png"
    } else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
        iconPath = exePath[:len(exePath) - 6] + "assets/icon.png"
    }

    err = beeep.Notify("NUtils", "Take a break, the timer has finished", iconPath)
    if err != nil {
        panic(err.Error())
    }
}
