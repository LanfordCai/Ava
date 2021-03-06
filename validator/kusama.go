package validator

import (
	"bytes"

	"github.com/LanfordCai/ava/base58check"
	"github.com/LanfordCai/ava/crypto"
)

// Kusama ...
type Kusama struct{}

var _ SS58 = (*Kusama)(nil)

// ValidateAddress returns validate result of aeternity address
// see: https://github.com/paritytech/substrate/wiki/External-Address-Format-(SS58)
func (v *Kusama) ValidateAddress(addr string, network NetworkType) *Result {
	decoded := base58check.BitcoinEncoder.Decode(addr)
	dataLen := len(decoded)
	// 1 byte type + 3 bytes address format
	if dataLen < 4 {
		return &Result{Success, false, Unknown, ""}
	}
	ss58AddrType := decoded[0]
	if ss58AddrType != v.AddressType() {
		return &Result{Success, false, Unknown, ""}
	}

	checksum := decoded[dataLen-v.ChecksumLen():]
	payload := decoded[1 : dataLen-v.ChecksumLen()]

	expectedChecksum := v.CalcChecksum(payload)
	if bytes.Compare(checksum, expectedChecksum) == 0 {
		return &Result{Success, true, Normal, ""}
	}
	return &Result{Success, false, Unknown, ""}
}

// AddressType ...
func (v *Kusama) AddressType() byte {
	return 2
}

// AccountIdxLen ...
func (v *Kusama) AccountIdxLen() int {
	return 32
}

// ChecksumLen ...
func (v *Kusama) ChecksumLen() int {
	return 2
}

// CalcChecksum ...
func (v *Kusama) CalcChecksum(payload []byte) []byte {
	input := []byte("SS58PRE")
	input = append(input, v.AddressType())
	input = append(input, payload...)
	return crypto.Blake2b512(input)[:v.ChecksumLen()]
}
