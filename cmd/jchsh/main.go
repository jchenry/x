package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/jchenry/jchenry/pkg/arvelie"
	"github.com/jchenry/jchenry/pkg/neralie"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	PS1 := "[%]: "
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(PS1)
		if input, err := reader.ReadString('\n'); err == nil {
			if err = execute(input); err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
			}
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
		}

	}

}

func execute(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "now":
		t := time.Now()
		fmt.Printf("%s %s\n", arvelie.FromDate(t), neralie.FromTime(t))
		return nil
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
