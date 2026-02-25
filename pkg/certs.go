package ssl_tool

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrInvalidCert = errors.New("invalid certificate")

// GetCerts returns the full chain of certificates for the provided url
func GetCerts(u string) ([]*x509.Certificate, error) {
	req, err := http.NewRequest(http.MethodHead, u, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Do(req)
	if err != nil {
		var certVerifyErr *tls.CertificateVerificationError
		if errors.As(err, &certVerifyErr) {
			return certVerifyErr.UnverifiedCertificates, ErrInvalidCert
		}
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	if res.TLS == nil {
		return nil, fmt.Errorf("no TLS connection state available for %s", u)
	}

	return res.TLS.PeerCertificates, nil
}
