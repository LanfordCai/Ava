package dogecoin

import (
	"github.com/btcsuite/btcutil/base58"
)

// ░░░░░░░░░▄░░░░░░░░░░░░░░▄
// ░░░░░░░░▌▒█░░░░░░░░░░░▄▀▒▌
// ░░░░░░░░▌▒▒█░░░░░░░░▄▀▒▒▒▐
// ░░░░░░░▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
// ░░░░░▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
// ░░░▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
// ░░▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌
// ░░▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐
// ░▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌
// ░▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌
// ▌▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒▐
// ▐▒▒▐▀▐▀▒░▄▄▒▄▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
// ▐▒▒▒▀▀▄▄▒▒▒▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▐
// ░▌▒▒▒▒▒▒▀▀▀▒▒▒▒▒▒░▒░▒░▒░▒▒▒▌
// ░▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▄▒▒▐
// ░░▀▄▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▄▒▒▒▒▌
// ░░░░▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀
// ░░░░░░▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀
// ░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▀▀
// 👆 A cute comment from https://dogechain.info/

// IsValidAddress ...
// SEE: https://github.com/dogecoin/dogecoin/blob/master/src/chainparams.cpp#L132
func IsValidAddress(address string, isTestnet bool) bool {
	return IsP2PKH(address, isTestnet) || IsP2SH(address, isTestnet)
}

// IsP2PKH ...
func IsP2PKH(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == 113
	}
	return version == 30
}

// IsP2SH ...
func IsP2SH(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == 196
	}
	return version == 22
}
