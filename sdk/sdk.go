package sdk

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/filecoin-project/lotus/node/repo"
	"github.com/filecoin-project/specs-actors/actors/crypto"
)

type FileCoinSDK struct {
	Wallet *wallet.Wallet
}

func NewFileCoinSDK(walletPath string) (*FileCoinSDK, error) {
	fsrepo, err := repo.NewFS(walletPath)
	if err != nil {
		return nil, err
	}

	lkrepo, err := fsrepo.Lock(repo.FullNode)
	if err != nil {
		return nil, err
	}

	defer lkrepo.Close()
	keystore, err := lkrepo.KeyStore()
	if err != nil {
		return nil, err
	}
	w, err := wallet.NewWallet(keystore)
	if err != nil {
		return nil, err
	}
	return &FileCoinSDK{Wallet: w}, nil
}

func (this *FileCoinSDK) NewAccount(sigType uint8) (address.Address, error) {
	filSigType := crypto.SigType(sigType)
	return this.Wallet.GenerateKey(filSigType)
}
