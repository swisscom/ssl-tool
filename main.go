package main

import (
	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"
	"strings"
)

var logger = logrus.New()

var args struct {
	Certificate *CertificateCmd `arg:"subcommand:cert"`
	Client      *ClientCmd      `arg:"subcommand:client"`
	LogLevel    string          `arg:"-l,--log-level" default:"warning"`
}

const specifySubcommand = "please specify a subcommand"

func main() {
	arg.MustParse(&args)

	switch strings.ToLower(args.LogLevel) {
	case "warning":
		logger.SetLevel(logrus.WarnLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	}

	if args.Certificate != nil {
		doCertificateCmd(args.Certificate)
	} else if args.Client != nil {
		doClientCmd()
	} else {
		logger.Errorf(specifySubcommand)
	}
}
