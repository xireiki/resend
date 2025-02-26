package main

import (
	"context"
	E "errors"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	globalCtx   context.Context
	key         string
)

var mainCommand = &cobra.Command{
	Use:              "resend",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("Resend: ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		if key == "" {
			key = os.Getenv("RESEND_API_KEY")
			if key == "" {
				log.Fatal(E.New("API Key is missing, please use '--key' or environment variable RESEND_API_KEY to add API Key"))
			}
		}
		log.Printf("API Key is \"%s\"\n", key)
	},
}

func init() {
	globalCtx = context.TODO()
	mainCommand.PersistentFlags().StringVarP(&key, "key", "k", "", "set resend api key")
}
