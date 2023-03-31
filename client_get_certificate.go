package main

import (
	ssl_tool "github.com/swisscom/ssl-tool/pkg"
	"net/url"
)

type GetCertificateCmd struct {
	Url string `arg:"positional,required"`
}

func doGetCertificateCmd() {
	cmd := args.Client.GetCertificate
	u, err := url.Parse(cmd.Url)
	if err != nil {
		logger.Fatalf("unable to parse URL: %v", err)
	}

	if u.Scheme != "https" {
		logger.Fatalf("invalid scheme %s, only https is supported", u.Scheme)
	}

	certs, err := ssl_tool.GetCerts(u.String())
	if err != nil {
		logger.Fatalf("unable to get certs: %v", err)
	}
	printCertChain(certs)
}
