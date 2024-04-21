package main

import (
	"time"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const (
	PROMODORI_TIME int = 25*60
	BREAK_TIME     int = 5*60
	MAX_ITERATIONS int = 5
)

type Timer int64
const (
	Promodori Timer = iota
	Break Timer = iota
)

func main() {

	//TODO: os.args

	//TODO: maybe insert countdown

	for iter := 0; iter < MAX_ITERATIONS; iter++ {

		// TODO: alert on promodori
		// promodori timer
		timerType := Promodori
		for timer := PROMODORI_TIME; timer > -1; timer-- {
			printTimer(timer, timerType)
			time.Sleep(time.Second)
		}

		// TODO: alert on break
		// break timer
		timerType := Break
		for {
			printTimer(timer, timerType)
			time.Sleep(time.Second)
		}
	}
}

func printTimer(time int, timerType Timer) {
	var typeStr string
	if timerType == Promodori {
		typeStr = "Promodori"
	} else {
		typeStr = "Break"
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Printf("%s: %s\n", typeStr, formatTimer(time))
}

func formatTimer(seconds int) string {
	return strconv.Itoa(seconds / 60) + ":" + strconv.Itoa(seconds % 60)
}
