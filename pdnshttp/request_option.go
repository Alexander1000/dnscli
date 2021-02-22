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

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// RequestOption is a special type of function that can be passed to most HTTP
// request functions in this package; it is used to modify an HTTP request and
// to implement special request logic.
type RequestOption func(*http.Request) error

// WithJSONRequestBody adds a JSON body to a request. The input type can be
// anything, as long as it can be marshaled by "json.Marshal". This method will
// also automatically set the correct content type and content-length.
func WithJSONRequestBody(in interface{}) RequestOption {
	return func(req *http.Request) error {
		if in == nil {
			return nil
		}

		buf := bytes.Buffer{}
		enc := json.NewEncoder(&buf)
		err := enc.Encode(in)

		if err != nil {
			return err
		}

		rc := ioutil.NopCloser(&buf)

		copyBuf := buf.Bytes()

		req.Body = rc
		req.Header.Set("Content-Type", "application/json")
		req.ContentLength = int64(buf.Len())
		req.GetBody = func() (io.ReadCloser, error) {
			r := bytes.NewReader(copyBuf)
			return ioutil.NopCloser(r), nil
		}

		return nil
	}
}

// WithQueryValue adds a query parameter to a request's URL.
func WithQueryValue(key, value string) RequestOption {
	return func(req *http.Request) error {
		q := req.URL.Query()
		q.Set(key, value)

		req.URL.RawQuery = q.Encode()
		return nil
	}
}
