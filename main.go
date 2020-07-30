package main

import (
	"fmt"
	"github.com/filecoin-project/specs-actors/actors/crypto"
	"github.com/qiluge/filecoin-golang-sdk/sdk"
)

func main() {
	fil, err := sdk.NewFileCoinSDK("./")
	if err != nil {
		fmt.Print(err)
		return
	}
	addr, err := fil.NewAccount(uint8(crypto.SigTypeBLS))
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(addr)
}
