package main

import (
	"crypto/x509"
	"fmt"
	"github.com/alexeyco/simpletable"
	"strings"
)

func printCertChain(certs []*x509.Certificate) {
	if len(certs) == 0 {
		logger.Warnf("no certificates")
		return
	}

	// Print Leaf
	leaf := certs[0]

	printTitle("Certificate")
	t := simpletable.New()
	t.SetStyle(simpletable.StyleCompactClassic)
	t.Header.Cells = []*simpletable.Cell{
		{Text: boldStyle.Render("NAME")},
		{Text: boldStyle.Render("VALUE")},
	}
	addKV(t, "Subject", leaf.Subject.ToRDNSequence().String())
	addKV(t, "Issuer", leaf.Issuer.ToRDNSequence().String())
	addKV(t, "DNS Names", strings.Join(leaf.DNSNames, "\n"))
	addKV(t, "Not Before", formatDate(leaf.NotBefore))
	addKV(t, "Not After", formatDate(leaf.NotAfter))
	addKV(t, "Public Key Algorithm", fmt.Sprintf("%s", leaf.PublicKeyAlgorithm))

	var keySizeStr string
	keySize, err := getKeySize(leaf)
	if err != nil {
		keySizeStr = "err"
	} else {
		keySizeStr = fmt.Sprintf("%d", keySize)
	}

	addKV(t, "Public Key Size (bits)", keySizeStr)

	t.Println()
	fmt.Println()

	// Print Cert Chain
	printTitle("Certificate Chain")
	t = certTable()
	for _, cert := range certs {
		c := printCertificate(cert)
		t.Body.Cells = append(t.Body.Cells, c)
	}
	t.Println()
}

func addKV(t *simpletable.Table, key string, value string) {
	t.Body.Cells = append(t.Body.Cells, []*simpletable.Cell{
		{Text: key},
		{Text: style.Render(value)},
	})
}
