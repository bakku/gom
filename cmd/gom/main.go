package main

import (
	"fmt"
	"os"

	"github.com/bakku/gom/commands"
)

func main() {
	if len(os.Args) <= 1 {
		printUsage()
		return
	}

	command, err := commands.Select(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = command.Run(os.Args[2:]...); err != nil {
		fmt.Println(err)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tgom init\t\t\t-\tInitializes the db folder and the schema_migrations table")
	fmt.Println("\tgom generate <migration>\t-\tCreate up.sql and down.sql for the specified migration name")
	fmt.Println("\tgom migrate\t\t\t-\tExecutes all up.sql files of all migrations that were not migrated yet")
	fmt.Println("\tgom rollback\t\t\t-\tExecutes the down.sql file of the last migrated migration")
}
