// Copyright © 2020 - 2024 Attestant Limited.
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

package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	client "github.com/attestantio/go-eth2-client"
	"github.com/QYuQianchen/metaclear-go-eth2-client/api"
	"github.com/QYuQianchen/metaclear-go-eth2-client/spec/phase0"
)

type beaconStateRootJSON struct {
	Root phase0.Root `json:"root"`
}

// BeaconStateRoot fetches the beacon state root given a set of options.
func (s *Service) BeaconStateRoot(ctx context.Context, opts *api.BeaconStateRootOpts) (*api.Response[*phase0.Root], error) {
	if err := s.assertIsActive(ctx); err != nil {
		return nil, err
	}
	if opts == nil {
		return nil, client.ErrNoOptions
	}
	if opts.State == "" {
		return nil, errors.Join(errors.New("no state specified"), client.ErrInvalidOptions)
	}

	endpoint := fmt.Sprintf("/eth/v1/beacon/states/%s/root", opts.State)
	httpResponse, err := s.get(ctx, endpoint, "", &opts.Common, false)
	if err != nil {
		return nil, err
	}

	data, metadata, err := decodeJSONResponse(bytes.NewReader(httpResponse.body), beaconStateRootJSON{})
	if err != nil {
		return nil, err
	}

	return &api.Response[*phase0.Root]{
		Data:     &data.Root,
		Metadata: metadata,
	}, nil
}
