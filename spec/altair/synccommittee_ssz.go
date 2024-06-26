// Code generated by fastssz. DO NOT EDIT.
// Hash: 3788614e00eec4f1965525d5d47a3f0341227525e7e3c2bd84f6500e73c4c20d
// Version: 0.1.3
package altair

import (
	"github.com/QYuQianchen/metaclear-go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SyncCommittee object
func (s *SyncCommittee) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SyncCommittee object to a target array
func (s *SyncCommittee) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Pubkeys'
	if size := len(s.Pubkeys); size != 512 {
		err = ssz.ErrVectorLengthFn("SyncCommittee.Pubkeys", size, 512)
		return
	}
	for ii := 0; ii < 512; ii++ {
		if size := len(s.Pubkeys[ii]); size != 48 {
			err = ssz.ErrBytesLengthFn("SyncCommittee.Pubkeys[ii]", size, 48)
			return
		}
		dst = append(dst, s.Pubkeys[ii][:]...)
	}

	// Field (1) 'AggregatePubkey'
	dst = append(dst, s.AggregatePubkey[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SyncCommittee object
func (s *SyncCommittee) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 24624 {
		return ssz.ErrSize
	}

	// Field (0) 'Pubkeys'
	s.Pubkeys = make([]phase0.BLSPubKey, 512)
	for ii := 0; ii < 512; ii++ {
		copy(s.Pubkeys[ii][:], buf[0:24576][ii*48:(ii+1)*48])
	}

	// Field (1) 'AggregatePubkey'
	copy(s.AggregatePubkey[:], buf[24576:24624])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncCommittee object
func (s *SyncCommittee) SizeSSZ() (size int) {
	size = 24624
	return
}

// HashTreeRoot ssz hashes the SyncCommittee object
func (s *SyncCommittee) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncCommittee object with a hasher
func (s *SyncCommittee) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkeys'
	{
		if size := len(s.Pubkeys); size != 512 {
			err = ssz.ErrVectorLengthFn("SyncCommittee.Pubkeys", size, 512)
			return
		}
		subIndx := hh.Index()
		for _, i := range s.Pubkeys {
			if len(i) != 48 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (1) 'AggregatePubkey'
	hh.PutBytes(s.AggregatePubkey[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SyncCommittee object
func (s *SyncCommittee) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
