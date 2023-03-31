package main

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/alexeyco/simpletable"
	"io"
	"os"
)

type ParseCmd struct {
	CertFile string `arg:"positional" description:"A certificate in .pem format"`
}

func doParseCertificate(cmd *ParseCmd) {
	if cmd == nil {
		logger.Fatalf("cmd cannot be nil")
	}

	if cmd.CertFile == "" {
		logger.Errorf("please specify a certificate file")
		return
	}

	f, err := os.Open(cmd.CertFile)
	if err != nil {
		logger.Errorf("unable to open certificate file: %v", err)
		return
	}

	var fileBytes []byte
	fileBytes, err = io.ReadAll(f)
	if err != nil {
		logger.Errorf("unable to read file: %v", err)
		return
	}
	err = f.Close()
	if err != nil {
		logger.Errorf("unable to close file: %v", err)
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
		if rest == nil {
			break
		}
		input = rest
		if pemBlock.Type != "CERTIFICATE" {
			logger.Errorf("invalid pem block type: %s", pemBlock.Type)
			return
		}

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
		logger.Errorf("cert cannot be nil!")
		return nil
	}

	return []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: cert.Subject.CommonName},
		{Align: simpletable.AlignLeft, Text: cert.Issuer.CommonName},
	}
}
