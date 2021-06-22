// Copyright © 2021 Attestant Limited.
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

package multi

import (
	"context"

	eth2client "github.com/attestantio/go-eth2-client"
	api "github.com/attestantio/go-eth2-client/api/v1"
)

// NodeSyncing provides the syncing information for the node.
func (s *Service) NodeSyncing(ctx context.Context) (*api.SyncState, error) {
	res, err := s.doCall(ctx, func(ctx context.Context, client eth2client.Service) (interface{}, error) {
		nodeSyncing, err := client.(eth2client.NodeSyncingProvider).NodeSyncing(ctx)
		if err != nil {
			return nil, err
		}
		return nodeSyncing, nil
	})
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*api.SyncState), nil
}
