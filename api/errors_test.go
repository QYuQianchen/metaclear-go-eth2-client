// Copyright © 2023 Attestant Limited.
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

package api_test

import (
	"testing"

	"github.com/QYuQianchen/metaclear-go-eth2-client/api"
	"github.com/QYuQianchen/metaclear-go-eth2-client/spec"
	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	v1 := &api.VersionedProposal{
		Version: spec.DataVersionAltair,
	}

	_, err := v1.KZGProofs()
	require.ErrorIs(t, err, api.ErrUnsupportedVersion)

	v2 := &api.VersionedProposal{
		Version: spec.DataVersionDeneb,
	}

	_, err = v2.KZGProofs()
	require.ErrorIs(t, err, api.ErrDataMissing)
}
