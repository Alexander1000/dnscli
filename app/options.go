package app

import (
	"os"
	"time"
)

// WithBaseURL sets an app's base URL
func WithBaseURL(baseURL string) Option {
	return func(a *app) error {
		a.baseURL = baseURL
		return nil
	}
}

// WithTLS sets an app's TLS settings
func WithTLS(tlsEnable bool, tlsCAPath, tlsCertPath, tlsKeyPath string) Option {
	return func(a *app) error {
		a.tlsEnable = tlsEnable
		a.tlsCAPath = tlsCAPath
		a.tlsCertPath = tlsCertPath
		a.tlsKeyPath = tlsKeyPath
		return nil
	}
}

// WithTimeout sets an app's timeout
func WithTimeout(t int64) Option {
	return func(a *app) error {
		a.httpClient.Timeout = time.Duration(t) * time.Second
		return nil
	}
}

// WithDebuggingOutput can be used to supply an io.Writer to the client into which all
// outgoing HTTP requests and their responses will be logged. Useful for debugging.
func WithDebuggingOutput(yes bool) Option {
	return func(a *app) error {
		if yes {
			a.debugOutput = os.Stderr
		}
		return nil
	}
}
