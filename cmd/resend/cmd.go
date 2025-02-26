package main

import (
	"context"
	E "errors"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/xireiki/resend/log"
)

var (
	globalCtx    context.Context
	disableColor bool
	key          string
)

var mainCommand = &cobra.Command{
	Use:              "resend",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if key == "" {
			key = os.Getenv("RESEND_API_KEY")
			if key == "" {
				log.Fatal(E.New("API Key is missing, please use \"--key\" or environment variable RESEND_API_KEY to add API Key"))
			}
		}
		if disableColor {
			log.SetStdLogger(log.NewDefaultFactory(context.Background(), log.Formatter{BaseTime: time.Now(), DisableColors: true}, os.Stderr, "", nil, false).Logger())
		}
		log.Debug("API Key is ", key)
	},
}

func init() {
	globalCtx = context.TODO()
	mainCommand.PersistentFlags().StringVarP(&key, "key", "k", "", "set resend api key")
}
