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
	logger       log.ContextLogger
)

var (
	disableColor bool
	logLevel     string
	key          string
)

var mainCommand = &cobra.Command{
	Use:              "resend",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		loggerOptions := log.Options{
			Context:       context.Background(),
			Options:       log.LogOptions{},
			Observable:    false,
			DefaultWriter: os.Stderr,
			BaseTime:      time.Now(),
		}
		if disableColor {
			loggerOptions.Options.DisableColor = disableColor
		}
		if logLevel != "" {
			_, err := log.ParseLevel(logLevel)
			if err != nil {
				log.Fatal("Unknown log level: ", logLevel)
			}
			loggerOptions.Options.Level = logLevel
		}
		logger, err := log.New(loggerOptions)
		if err != nil {
			log.Fatal("Failed to create logger: ", err)
		}
		log.SetStdLogger(logger.Logger())
		if key == "" {
			key = os.Getenv("RESEND_API_KEY")
			if key == "" {
				log.Fatal(E.New("API Key is missing, please use \"--key\" or environment variable RESEND_API_KEY to add API Key"))
			}
		}
		log.Trace("API Key is ", key)
	},
}

func init() {
	globalCtx = context.TODO()
	mainCommand.PersistentFlags().StringVarP(&key, "key", "k", "", "set resend api key")
	mainCommand.PersistentFlags().StringVarP(&logLevel, "log-level", "L", "warn", "set log level")
	mainCommand.PersistentFlags().BoolVarP(&disableColor, "disable-color", "", false, "Disable log color")
}
