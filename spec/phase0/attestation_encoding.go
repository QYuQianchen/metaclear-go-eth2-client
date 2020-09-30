// Code generated by fastssz. DO NOT EDIT.
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Attestation object
func (a *Attestation) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the Attestation object to a target array
func (a *Attestation) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(228)

	// Offset (0) 'AggregationBits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(a.AggregationBits)

	// Field (1) 'Data'
	if a.Data == nil {
		a.Data = new(AttestationData)
	}
	if dst, err = a.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Signature'
	if len(a.Signature) != 96 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, a.Signature...)

	// Field (0) 'AggregationBits'
	if len(a.AggregationBits) > 2048 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, a.AggregationBits...)

	return
}

// UnmarshalSSZ ssz unmarshals the Attestation object
func (a *Attestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 228 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'AggregationBits'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (1) 'Data'
	if a.Data == nil {
		a.Data = new(AttestationData)
	}
	if err = a.Data.UnmarshalSSZ(buf[4:132]); err != nil {
		return err
	}

	// Field (2) 'Signature'
	if cap(a.Signature) == 0 {
		a.Signature = make([]byte, 0, len(buf[132:228]))
	}
	a.Signature = append(a.Signature, buf[132:228]...)

	// Field (0) 'AggregationBits'
	{
		buf = tail[o0:]
		if err = ssz.ValidateBitlist(buf, 2048); err != nil {
			return err
		}
		if cap(a.AggregationBits) == 0 {
			a.AggregationBits = make([]byte, 0, len(buf))
		}
		a.AggregationBits = append(a.AggregationBits, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Attestation object
func (a *Attestation) SizeSSZ() (size int) {
	size = 228

	// Field (0) 'AggregationBits'
	size += len(a.AggregationBits)

	return
}

// HashTreeRoot ssz hashes the Attestation object
func (a *Attestation) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the Attestation object with a hasher
func (a *Attestation) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'AggregationBits'
	hh.PutBitlist(a.AggregationBits, 2048)

	// Field (1) 'Data'
	if err = a.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Signature'
	if len(a.Signature) != 96 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(a.Signature)

	hh.Merkleize(indx)
	return
}