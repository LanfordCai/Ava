package validator

// Iris ...
type Iris struct{}

var _ BitcoinLike = (*Iris)(nil)

// ValidateAddress returns validate result of bytom address
func (v *Iris) ValidateAddress(addr string, network NetworkType) *Result {
	hrp, program, err := bech32AddrDecode(addr)
	if err != nil || hrp != v.AddressHrp(network) {
		return &Result{false, Unknown, ""}
	}

	if len(program) == v.SegwitProgramLength(P2WPKH) {
		return &Result{true, P2WPKH, ""}
	}

	return &Result{false, Unknown, ""}
}

// AddressVersion returns bytom address version according to the address type and
// network type
func (v *Iris) AddressVersion(addrType AddressType, network NetworkType) byte {
	panic(ErrUnsupported.Error())
}

// AddressHrp returns hrps of bytom according to the network
func (v *Iris) AddressHrp(network NetworkType) string {
	if network == Mainnet {
		return "iaa"
	}
	panic(ErrUnsupported.Error())
}

// SegwitProgramLength returns segwit program length of bytom
func (v *Iris) SegwitProgramLength(addrType AddressType) int {
	if addrType == P2WPKH {
		return 20
	} else if addrType == P2WSH {
		panic(ErrUnsupported.Error())
	}

	panic(ErrInvalidAddrType.Error())
}