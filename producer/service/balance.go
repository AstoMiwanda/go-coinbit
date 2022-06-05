package service

import (
	"github.com/astomiwanda/go-coinbit/consumer/model"
	"github.com/astomiwanda/go-coinbit/producer/repository"
)

type BalanceService interface {
	GetBalance(walletID string) (*model.Balance, error)
}

type balanceService struct {
	balance           map[int]*model.Balance
	balanceRepository repository.BalanceRepository
}

func (b *balanceService) GetBalance(walletID string) (*model.Balance, error) {
	return b.balanceRepository.GetBalance(walletID)
}

func NewBalanceService(balanceRepository repository.BalanceRepository) BalanceService {
	return &balanceService{
		balanceRepository: balanceRepository,
	}
}
