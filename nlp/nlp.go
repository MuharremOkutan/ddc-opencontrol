package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/ddc-opencontrol/nlp/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
