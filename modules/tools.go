package modules

import (
	"github.com/interstellar/kelp/support/utils"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// Path2Assets returns the assets inside a PaymentPath
func Path2Assets(path *PaymentPath) (horizon.Asset, horizon.Asset, horizon.Asset) {
	holdAsset := path.HoldAsset
	pathAssetA := path.PathAssetA
	pathAssetB := path.PathAssetB
	return holdAsset, pathAssetA, pathAssetB
}

// ParseAsset returns a horizon asset from strings
func ParseAsset(code string, issuer string) horizon.Asset {
	if code != "XLM" && issuer == "" {
		panic("error: issuer can only be empty if asset is XLM")
	}

	if code == "XLM" && issuer != "" {
		panic("error: issuer needs to be empty if asset is XLM")
	}

	if code == "XLM" {
		asset := utils.Asset2Asset2(build.NativeAsset())
		return asset
	}

	asset := utils.Asset2Asset2(build.CreditAsset(code, issuer))
	return asset
}

// Pair2Assets splits a TradingPair into its assets
func Pair2Assets(pair TradingPair) (horizon.Asset, horizon.Asset) {
	base := pair.Base
	quote := pair.Quote
	return base, quote
}
