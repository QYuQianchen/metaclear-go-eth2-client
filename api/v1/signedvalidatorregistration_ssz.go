// Code generated by fastssz. DO NOT EDIT.
// Hash: 4138c3a9facc5ccd31f1e1607b26966b6ecfd78af99f1a5ed49e386fc2cf1ac8
// Version: 0.1.3
package v1

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SignedValidatorRegistration object
func (s *SignedValidatorRegistration) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedValidatorRegistration object to a target array
func (s *SignedValidatorRegistration) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(ValidatorRegistration)
	}
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Signature'
	dst = append(dst, s.Signature[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SignedValidatorRegistration object
func (s *SignedValidatorRegistration) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 180 {
		return ssz.ErrSize
	}

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(ValidatorRegistration)
	}
	if err = s.Message.UnmarshalSSZ(buf[0:84]); err != nil {
		return err
	}

	// Field (1) 'Signature'
	copy(s.Signature[:], buf[84:180])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedValidatorRegistration object
func (s *SignedValidatorRegistration) SizeSSZ() (size int) {
	size = 180
	return
}

// HashTreeRoot ssz hashes the SignedValidatorRegistration object
func (s *SignedValidatorRegistration) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedValidatorRegistration object with a hasher
func (s *SignedValidatorRegistration) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(ValidatorRegistration)
	}
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedValidatorRegistration object
func (s *SignedValidatorRegistration) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}