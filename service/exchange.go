package service

import "Vico1993/Wallet/domain"

type Exchange interface {
	Load() (domain.Wallet, error)
}