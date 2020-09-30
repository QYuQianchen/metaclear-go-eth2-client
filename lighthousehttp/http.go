// Copyright © 2020 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lighthousehttp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// get sends an HTTP get request and returns the body.
func (s *Service) get(ctx context.Context, endpoint string) (io.Reader, error) {
	log.Trace().Str("endpoint", endpoint).Msg("GET request")

	reference, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint")
	}
	url := s.base.ResolveReference(reference).String()
	log.Trace().Str("url", url).Msg("GET request")
	opCtx, cancel := context.WithTimeout(ctx, s.timeout)
	req, err := http.NewRequestWithContext(opCtx, http.MethodGet, url, nil)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to create GET request")
	}
	resp, err := s.client.Do(req)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to connect to GET endpoint")
	}
	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		cancel()
		return nil, fmt.Errorf("HTTP call to lighthouse endpoint %s failed with status %d", url, resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to read GET response")
	}
	cancel()

	log.Trace().Str("response", string(data)).Msg("GET response")

	return bytes.NewReader(data), nil
}

// post sends an HTTP post request and returns the body.
func (s *Service) post(ctx context.Context, endpoint string, body io.Reader) (io.Reader, error) {
	reference, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint")
	}
	url := s.base.ResolveReference(reference).String()

	if e := log.Trace(); e.Enabled() {
		if body == nil {
			e.Str("url", url).Msg("POST request")
		} else {
			bodyBytes, err := ioutil.ReadAll(body)
			if err != nil {
				return nil, errors.New("failed to read request body")
			}
			body = bytes.NewReader(bodyBytes)
			e.Str("url", url).Str("body", string(bodyBytes)).Msg("POST request")
		}
	}

	opCtx, cancel := context.WithTimeout(ctx, s.timeout)
	req, err := http.NewRequestWithContext(opCtx, http.MethodPost, url, body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to create POST request")
	}
	resp, err := s.client.Do(req)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to connect to lighthouse HTTP endpoint")
	}
	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		cancel()
		return nil, fmt.Errorf("HTTP call to lighthouse endpoint failed with status %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to read POST response")
	}
	cancel()

	return bytes.NewReader(data), nil
}