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

package pdnshttp

import "fmt"

// ErrNotFound error not found with URL
type ErrNotFound struct {
	URL string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("not found: %s", e.URL)
}

// ErrUnexpectedStatus error with URL and HTTP status code
type ErrUnexpectedStatus struct {
	URL        string
	StatusCode int
}

func (e ErrUnexpectedStatus) Error() string {
	return fmt.Sprintf("unexpected status code %d: %s", e.StatusCode, e.URL)
}

// IsNotFound ...
func IsNotFound(err error) bool {
	switch err.(type) {
	case ErrNotFound:
		return true
	case *ErrNotFound:
		return true
	}

	return false
}
