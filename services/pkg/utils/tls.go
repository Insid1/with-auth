package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type CredentialParams struct {
	CAClientCertPath string
	CAServerCertPath string

	CertPath    string
	CertKeyPath string
}

func LoadTLSCredentials(params *CredentialParams) (credentials.TransportCredentials, error) {
	// if no params return insecure connection
	if params == nil {
		return insecure.NewCredentials(), nil
	}

	// Create the credentials to return it
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	// Load certificate of the CA who signed server's certificate
	if params.CAClientCertPath != "" {
		pemServerCA, err := os.ReadFile(params.CAClientCertPath)
		if err != nil {
			return nil, err
		}

		// Create cert Pool to store it in credentials
		certPool := x509.NewCertPool()
		// Adds custom CA certificate
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			return nil, fmt.Errorf("failed to add server CA's certificate")
		}
		cfg.ClientCAs = certPool
	}

	// Load certificate of the CA who signed client's certificate
	if params.CAServerCertPath != "" {
		pemServerCA, err := os.ReadFile(params.CAServerCertPath)
		if err != nil {
			return nil, err
		}

		// Create cert Pool to store it in credentials
		certPool := x509.NewCertPool()
		// Adds custom CA certificate
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			return nil, fmt.Errorf("failed to add server CA's certificate")
		}
		cfg.RootCAs = certPool
	}

	// Load  certificate and private key
	if params.CertPath != "" && params.CertKeyPath != "" {
		// Load  certificate and private key
		serverCert, err := tls.LoadX509KeyPair(params.CertPath, params.CertKeyPath)
		if err != nil {
			return nil, err
		}

		// Store server's certificate
		cfg.Certificates = []tls.Certificate{serverCert}
	}

	return credentials.NewTLS(cfg), nil
}
