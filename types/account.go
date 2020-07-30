package types

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

type Account struct {
	*types.KeyInfo
	Address address.Address
}
