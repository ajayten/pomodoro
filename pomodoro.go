package main

import (
	"fmt"
	"time"
)

const (
	redColor   = "\033[31m"
	greenColor = "\033[32m"
	resetColor = "\033[0m"
)

//func showWorkMessage(duration time.Duration) {
//	fmt.Printf("%s %v Minutes of Work!-Work!-Work!%s\n", redColor, duration, resetColor)
//}

func pomodoroTimer() {
	workDuration := 20 * time.Minute
	breakDuration := 5 * time.Minute

	for {

		fmt.Printf("%sWork! Work! Work!%s\n", redColor, resetColor)
		time.Sleep(workDuration)

		fmt.Printf("%sBreak! Yeah!%s\n", greenColor, resetColor)
		time.Sleep(breakDuration)
	}
}

func main() {
	pomodoroTimer()
}
