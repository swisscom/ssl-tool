package main

import (
	"crypto/x509"
	"fmt"
	"github.com/alexeyco/simpletable"
	"strconv"
	"strings"
	"time"
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
	addKVStyled(t, "Not After", formatDateWithExpiry(leaf.NotAfter), expiryStyle(leaf.NotAfter))
	addKV(t, "Public Key Algorithm", leaf.PublicKeyAlgorithm.String())

	var keySizeStr string
	keySize, err := getKeySize(leaf)
	if err != nil {
		keySizeStr = "err"
	} else {
		keySizeStr = strconv.Itoa(keySize)
	}

	addKV(t, "Public Key Size (bits)", keySizeStr)

	t.Println()
	fmt.Println()

	// Print Cert Chain
	printTitle("Certificate Chain")
	t = certTable()
	for i, cert := range certs {
		c := printCertificate(cert)
		if c != nil {
			c[0].Text = chainIndexStyle.Render(fmt.Sprintf("[%d]", i)) + " " + c[0].Text
		}
		t.Body.Cells = append(t.Body.Cells, c)
	}
	t.Println()
}

// expiryStyle returns the appropriate lipgloss render function based on how far
// away the expiry date is: red if already expired, yellow if within 30 days,
// green otherwise.
func expiryStyle(notAfter time.Time) func(string) string {
	now := time.Now()
	if now.After(notAfter) {
		return func(s string) string { return expiredStyle.Render(s) }
	}
	if notAfter.Before(now.Add(30 * 24 * time.Hour)) {
		return func(s string) string { return warnStyle.Render(s) }
	}
	return func(s string) string { return validStyle.Render(s) }
}

func addKV(t *simpletable.Table, key string, value string) {
	t.Body.Cells = append(t.Body.Cells, []*simpletable.Cell{
		{Text: keyStyle.Render(key)},
		{Text: valueStyle.Render(value)},
	})
}

func addKVStyled(t *simpletable.Table, key string, value string, renderFn func(string) string) {
	t.Body.Cells = append(t.Body.Cells, []*simpletable.Cell{
		{Text: keyStyle.Render(key)},
		{Text: renderFn(value)},
	})
}
