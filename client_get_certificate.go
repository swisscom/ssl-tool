package main

import (
	"errors"
	"fmt"
	"net/url"

	ssl_tool "github.com/swisscom/ssl-tool/pkg"
	"github.com/spf13/cobra"
)

var getCertificateCmd = &cobra.Command{
	Use:   "get-certificate <url>",
	Short: "Retrieve SSL/TLS certificates from a remote host",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		doGetCertificateCmd(args[0])
		return nil
	},
}

func doGetCertificateCmd(rawURL string) {
	u, err := url.Parse(rawURL)
	if err != nil {
		logger.Fatalf("unable to parse URL: %v", err)
	}

	if u.Scheme != "https" {
		logger.Fatalf("invalid scheme %s, only https is supported", u.Scheme)
	}

	certs, err := ssl_tool.GetCerts(u.String())
	if errors.Is(err, ssl_tool.ErrInvalidCert) {
		fmt.Println(invalidCertStyle.Render("! Invalid Certificate"))
	} else if err != nil {
		logger.Fatalf("unable to get certs: %v", err)
	}
	printCertChain(certs)
}
