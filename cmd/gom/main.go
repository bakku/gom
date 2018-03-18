package main

import (
	"fmt"
	"os"

	"github.com/bakku/gom/commands"
)

func main() {
	command, err := commands.Select(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = command.Run(os.Args[2:]...); err != nil {
		fmt.Println(err)
	}
}
