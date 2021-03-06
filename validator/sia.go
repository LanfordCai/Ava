package validator

import (
	"encoding/hex"

	"github.com/LanfordCai/ava/crypto"
)

// Sia ...
type Sia struct{}

var _ Validator = (*Sia)(nil)

// ValidateAddress - Check a Sia address is valid or not
func (s *Sia) ValidateAddress(addr string, network NetworkType) *Result {
	unlockhashWithChecksum, err := hex.DecodeString(addr)
	if err != nil || len(unlockhashWithChecksum) != 38 {
		return &Result{Success, false, Unknown, ""}
	}

	unlockhash := unlockhashWithChecksum[:32]
	checksum256 := crypto.Blake2b256(unlockhash)

	var validChecksum [6]byte
	copy(validChecksum[:], checksum256[:6])

	var checksum [6]byte
	copy(checksum[:], unlockhashWithChecksum[32:])
	if checksum == validChecksum {
		return &Result{Success, true, Normal, ""}
	}
	return &Result{Success, false, Unknown, ""}
}
