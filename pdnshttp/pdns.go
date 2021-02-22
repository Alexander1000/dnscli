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
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

// PDNSClient is the client for PowerDNS API
type PDNSClient struct {
	baseURL     string
	httpClient  *http.Client
	debugOutput io.Writer
}

// NewPDNSClient creates a new PDNSClient
func NewPDNSClient(baseURL string, hc *http.Client, debugOutput io.Writer) *PDNSClient {
	return &PDNSClient{
		baseURL:     baseURL,
		httpClient:  hc,
		debugOutput: debugOutput,
	}
}

// NewRequest builds a new request. Usually, this method should not be used;
// prefer using the "Get", "Post", ... methods if possible.
func (pc *PDNSClient) NewRequest(method string, path string, body io.Reader) (*http.Request, error) {
	path = strings.TrimPrefix(path, "/")
	req, err := http.NewRequest(method, pc.baseURL+"/"+path, body)
	if err != nil {
		return nil, err
	}

	return req, err
}

// Get executes a GET request
func (pc *PDNSClient) Get(path string, out interface{}, opts ...RequestOption) error {
	return pc.doRequest(http.MethodGet, path, out, opts...)
}

// Post executes a POST request
func (pc *PDNSClient) Post(path string, out interface{}, opts ...RequestOption) error {
	return pc.doRequest(http.MethodPost, path, out, opts...)
}

// Patch executes a PATCH request
func (pc *PDNSClient) Patch(path string, out interface{}, opts ...RequestOption) error {
	return pc.doRequest(http.MethodPatch, path, out, opts...)
}

// Put executes a PUT request
func (pc *PDNSClient) Put(path string, out interface{}, opts ...RequestOption) error {
	return pc.doRequest(http.MethodPut, path, out, opts...)
}

// Delete executes a DELETE request
func (pc *PDNSClient) Delete(path string, out interface{}, opts ...RequestOption) error {
	return pc.doRequest(http.MethodDelete, path, out, opts...)
}

func (pc *PDNSClient) doRequest(method string, path string, out interface{}, opts ...RequestOption) error {
	req, err := pc.NewRequest(method, path, nil)
	if err != nil {
		return err
	}

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return err
		}
	}

	reqDump, _ := httputil.DumpRequestOut(req, true)
	pc.debugOutput.Write(reqDump)

	res, err := pc.httpClient.Do(req)
	if err != nil {
		return err
	}

	resDump, _ := httputil.DumpResponse(res, true)
	pc.debugOutput.Write(resDump)

	if res.StatusCode == http.StatusNotFound {
		return ErrNotFound{URL: req.URL.String()}
	} else if res.StatusCode >= 400 {
		return ErrUnexpectedStatus{URL: req.URL.String(), StatusCode: res.StatusCode}
	}

	if out != nil {
		if w, ok := out.(io.Writer); ok {
			_, err := io.Copy(w, res.Body)
			return err
		}

		dec := json.NewDecoder(res.Body)
		err = dec.Decode(out)
		if err != nil {
			return err
		}
	}

	return nil
}
