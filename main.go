//go:build !test
// +build !test

package main

import (
	"github.com/webmalc/contacts-extractor/cmd"
	"github.com/webmalc/contacts-extractor/common/config"
	"github.com/webmalc/contacts-extractor/common/logger"
	"github.com/webmalc/contacts-extractor/extractor"
	"github.com/webmalc/contacts-extractor/formatters"
	"github.com/webmalc/contacts-extractor/validators"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	formatter := formatters.NewCSV()
	ext := extractor.NewExtractor(
		validators.NewPhoneValidator(),
		validators.NewEmailValidator(),
		validators.NewVKValidator(),
	)
	cmdRouter := cmd.NewCommandRouter(log, ext, formatter)
	cmdRouter.Run()
}
