package service

import "Vico1993/Wallet/domain/wallet"

type Exchange interface {
	Load() (wallet.Wallet, error)
}