package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("a walrus appears")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("a walrus appears")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("a walrus appears")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("a walrus appears")
}
