package main

import (
	log "github.com/sirupsen/logrus"
	"help-ukraine/pkg/api"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	api.Rest()
}
