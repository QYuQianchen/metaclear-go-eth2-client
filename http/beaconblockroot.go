// Copyright © 2021 - 2024 Attestant Limited.
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

type beaconBlockRootJSON struct {
	Root phase0.Root `json:"root"`
}

// BeaconBlockRoot fetches a block's root given a set of options.
func (s *Service) BeaconBlockRoot(ctx context.Context,
	opts *api.BeaconBlockRootOpts,
) (
	*api.Response[*phase0.Root],
	error,
) {
	if err := s.assertIsActive(ctx); err != nil {
		return nil, err
	}
	if opts == nil {
		return nil, client.ErrNoOptions
	}
	if opts.Block == "" {
		return nil, errors.Join(errors.New("no block specified"), client.ErrInvalidOptions)
	}

	endpoint := fmt.Sprintf("/eth/v1/beacon/blocks/%s/root", opts.Block)
	httpResponse, err := s.get(ctx, endpoint, "", &opts.Common, false)
	if err != nil {
		return nil, err
	}

	data, metadata, err := decodeJSONResponse(bytes.NewReader(httpResponse.body), beaconBlockRootJSON{})
	if err != nil {
		return nil, err
	}

	return &api.Response[*phase0.Root]{
		Data:     &data.Root,
		Metadata: metadata,
	}, nil
}
