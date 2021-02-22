/*
Copyright © 2021 Michael Bruskov <mixanemca@yandex.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
