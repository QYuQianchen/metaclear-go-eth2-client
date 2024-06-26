// Code generated by fastssz. DO NOT EDIT.
// Hash: 3788614e00eec4f1965525d5d47a3f0341227525e7e3c2bd84f6500e73c4c20d
// Version: 0.1.3
package altair

import (
	"github.com/QYuQianchen/metaclear-go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SyncCommitteeMessage object
func (s *SyncCommitteeMessage) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SyncCommitteeMessage object to a target array
func (s *SyncCommitteeMessage) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(s.Slot))

	// Field (1) 'BeaconBlockRoot'
	dst = append(dst, s.BeaconBlockRoot[:]...)

	// Field (2) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, uint64(s.ValidatorIndex))

	// Field (3) 'Signature'
	dst = append(dst, s.Signature[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SyncCommitteeMessage object
func (s *SyncCommitteeMessage) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 144 {
		return ssz.ErrSize
	}

	// Field (0) 'Slot'
	s.Slot = phase0.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'BeaconBlockRoot'
	copy(s.BeaconBlockRoot[:], buf[8:40])

	// Field (2) 'ValidatorIndex'
	s.ValidatorIndex = phase0.ValidatorIndex(ssz.UnmarshallUint64(buf[40:48]))

	// Field (3) 'Signature'
	copy(s.Signature[:], buf[48:144])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncCommitteeMessage object
func (s *SyncCommitteeMessage) SizeSSZ() (size int) {
	size = 144
	return
}

// HashTreeRoot ssz hashes the SyncCommitteeMessage object
func (s *SyncCommitteeMessage) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncCommitteeMessage object with a hasher
func (s *SyncCommitteeMessage) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(s.Slot))

	// Field (1) 'BeaconBlockRoot'
	hh.PutBytes(s.BeaconBlockRoot[:])

	// Field (2) 'ValidatorIndex'
	hh.PutUint64(uint64(s.ValidatorIndex))

	// Field (3) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SyncCommitteeMessage object
func (s *SyncCommitteeMessage) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
