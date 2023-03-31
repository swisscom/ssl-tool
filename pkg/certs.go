package ssl_tool

import (
	"crypto/x509"
	"net/http"
)

// GetCerts returns the full chain of certificates for the provided url
func GetCerts(url string) ([]*x509.Certificate, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// TODO: Support invalid certificates?
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res.TLS.PeerCertificates, nil
}
