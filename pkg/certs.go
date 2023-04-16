package ssl_tool

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
)

var ErrInvalidCert = fmt.Errorf("invalid certificate")

// GetCerts returns the full chain of certificates for the provided url
func GetCerts(u string) ([]*x509.Certificate, error) {
	req, err := http.NewRequest(http.MethodHead, u, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Do(req)
	if err != nil {
		switch err.(type) {
		case *url.Error:
			urlErrInner := err.(*url.Error).Err
			switch urlErrInner.(type) {
			case *tls.CertificateVerificationError:
				return urlErrInner.(*tls.CertificateVerificationError).UnverifiedCertificates, ErrInvalidCert
			}
		}
		return nil, err
	} else {
		return res.TLS.PeerCertificates, nil
	}
}
