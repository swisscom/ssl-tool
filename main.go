package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logger = logrus.New()

var logLevel string

var rootCmd = &cobra.Command{
	Use:     "ssl-tool",
	Short:   "ssl-tool inspects and retrieves SSL/TLS certificates.",
	Version: "1.0.0",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch strings.ToLower(logLevel) {
		case "debug":
			logger.SetLevel(logrus.DebugLevel)
		case "info":
			logger.SetLevel(logrus.InfoLevel)
		case "warning":
			logger.SetLevel(logrus.WarnLevel)
		case "error":
			logger.SetLevel(logrus.ErrorLevel)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "warning", "log level (debug, info, warning, error)")
	rootCmd.AddCommand(certCmd)
	rootCmd.AddCommand(clientCmd)
}
