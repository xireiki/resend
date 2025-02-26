package main

import (
	"github.com/spf13/cobra"
	"github.com/xireiki/resend/log"
)

var emailCommand = &cobra.Command{
	Use:   "email",
	Short: "Set of e-mail related commands",
	Run:   func(cmd *cobra.Command, args []string) {
		log.Error("Use \"--help\" to see the subcommands")
	},
}

func init() {
	mainCommand.AddCommand(emailCommand)
}
