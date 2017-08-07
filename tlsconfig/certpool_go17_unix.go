// +build go1.7,!windows

package tlsconfig

import "crypto/x509"

// SystemCertPool returns a copy of the system cert pool,
// returns an error if failed to load.
func SystemCertPool() (*x509.CertPool, error) {
	return x509.SystemCertPool()
}
