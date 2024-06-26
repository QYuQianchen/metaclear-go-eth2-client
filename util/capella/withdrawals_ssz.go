// Code generated by fastssz. DO NOT EDIT.
// Hash: f9cb2c9b5cf06463c18e9b55a69b1be3d8219cd96f8ca4ea87175a23c967b5bf
package capella

import (
	"github.com/QYuQianchen/metaclear-go-eth2-client/spec/capella"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Withdrawals object
func (w *ExecutionPayloadWithdrawals) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(w)
}

// MarshalSSZTo ssz marshals the Withdrawals object to a target array
func (w *ExecutionPayloadWithdrawals) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Withdrawals'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(w.Withdrawals) * 44

	// Field (0) 'Withdrawals'
	if size := len(w.Withdrawals); size > 16 {
		err = ssz.ErrListTooBigFn("Withdrawals.Withdrawals", size, 16)
		return
	}
	for ii := 0; ii < len(w.Withdrawals); ii++ {
		if dst, err = w.Withdrawals[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Withdrawals object
func (w *ExecutionPayloadWithdrawals) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Withdrawals'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Withdrawals'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 44, 16)
		if err != nil {
			return err
		}
		w.Withdrawals = make([]*capella.Withdrawal, num)
		for ii := 0; ii < num; ii++ {
			if w.Withdrawals[ii] == nil {
				w.Withdrawals[ii] = new(capella.Withdrawal)
			}
			if err = w.Withdrawals[ii].UnmarshalSSZ(buf[ii*44 : (ii+1)*44]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Withdrawals object
func (w *ExecutionPayloadWithdrawals) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Withdrawals'
	size += len(w.Withdrawals) * 44

	return
}

// HashTreeRoot ssz hashes the Withdrawals object
func (w *ExecutionPayloadWithdrawals) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(w)
}

// HashTreeRootWith ssz hashes the Withdrawals object with a hasher
func (w *ExecutionPayloadWithdrawals) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Withdrawals'
	{
		subIndx := hh.Index()
		num := uint64(len(w.Withdrawals))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range w.Withdrawals {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Withdrawals object
func (w *ExecutionPayloadWithdrawals) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(w)
}
