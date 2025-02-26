package main

import (
	"log"

	"github.com/spf13/cobra"
)

var emailCommand = &cobra.Command{
	Use:   "email",
	Short: "Set of e-mail related commands",
	Run:   func(cmd *cobra.Command, args []string) {
		log.Println("Use '--help' to see the subcommands")
	},
}

func init() {
	mainCommand.AddCommand(emailCommand)
}
