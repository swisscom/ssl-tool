package main

import (
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse [certfile]",
	Short: "Parse and display a PEM certificate file",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			logger.Errorf("please specify a certificate file")
			return nil
		}
		doParseCertificate(args[0])
		return nil
	},
}

func doParseCertificate(certFile string) {
	fileBytes, err := os.ReadFile(certFile)
	if err != nil {
		logger.Errorf("unable to read certificate file: %v", err)
		return
	}

	showCert(fileBytes)
}

func showCert(fileBytes []byte) {
	input := fileBytes

	var certs []*x509.Certificate
	for {
		pemBlock, rest := pem.Decode(input)
		if pemBlock == nil {
			break
		}
		if pemBlock.Type != "CERTIFICATE" {
			logger.Errorf("invalid pem block type: %s", pemBlock.Type)
			return
		}
		input = rest

		cert, err := x509.ParseCertificate(pemBlock.Bytes)
		if err != nil {
			logger.Errorf("unable to parse certificate: %v", err)
			return
		}

		certs = append(certs, cert)
	}

	printCertChain(certs)
}

func printCertificate(cert *x509.Certificate) []*simpletable.Cell {
	if cert == nil {
		logger.Fatalf("cert cannot be nil!")
	}

	return []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: cert.Subject.CommonName},
		{Align: simpletable.AlignLeft, Text: cert.Issuer.CommonName},
		{Align: simpletable.AlignLeft, Text: expiryStyle(cert.NotAfter)(formatDateWithExpiry(cert.NotAfter))},
	}
}
