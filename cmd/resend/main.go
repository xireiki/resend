//go:build !generate

package main

import "github.com/xireiki/resend/log"

func main() {
	if err := mainCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
