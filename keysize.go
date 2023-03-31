package main

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"errors"
)

func getKeySize(cert *x509.Certificate) (int, error) {
	switch pub := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		return pub.Size() * 8, nil
	case *ed25519.PublicKey:
		return len(*pub) * 8, nil
	case *ecdsa.PublicKey:
		return pub.Curve.Params().BitSize, nil
	default:
		return 0, errors.New("unsupported public key algorithm")
	}
}
