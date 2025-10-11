package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}
func main() {

	for {
		now := time.Now()
		current := now.Format("15:04:05")

		ClearTerminal()
		myFigure := figure.NewColorFigure(current, "starwars", "cyan", true)
		myFigure.Print()
		time.Sleep(time.Second)

	}
}
