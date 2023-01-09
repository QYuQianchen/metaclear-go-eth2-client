// Code generated by fastssz. DO NOT EDIT.
// Hash: 1d3f1d08772acfc74e6f425ba169c073ce1426b49fe20892e24b06981e78812e
// Version: 0.1.3-dev
package capella

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the HistoricalSummary object
func (h *HistoricalSummary) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the HistoricalSummary object to a target array
func (h *HistoricalSummary) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'BlockSummaryRoot'
	dst = append(dst, h.BlockSummaryRoot[:]...)

	// Field (1) 'StateSummaryRoot'
	dst = append(dst, h.StateSummaryRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the HistoricalSummary object
func (h *HistoricalSummary) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'BlockSummaryRoot'
	copy(h.BlockSummaryRoot[:], buf[0:32])

	// Field (1) 'StateSummaryRoot'
	copy(h.StateSummaryRoot[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the HistoricalSummary object
func (h *HistoricalSummary) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the HistoricalSummary object
func (h *HistoricalSummary) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the HistoricalSummary object with a hasher
func (h *HistoricalSummary) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'BlockSummaryRoot'
	hh.PutBytes(h.BlockSummaryRoot[:])

	// Field (1) 'StateSummaryRoot'
	hh.PutBytes(h.StateSummaryRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the HistoricalSummary object
func (h *HistoricalSummary) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(h)
}
