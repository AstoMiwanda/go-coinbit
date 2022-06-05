package service

import (
	"encoding/json"
	"github.com/astomiwanda/go-coinbit/consumer/model"
	"github.com/astomiwanda/go-coinbit/consumer/repository"
	"github.com/astomiwanda/go-coinbit/lib"
	"github.com/sirupsen/logrus"
	"time"
)

type BalanceService interface {
	UpdateBalance(data []byte) error
}

type balanceService struct {
	balance           map[int]*model.Balance
	balanceRepository repository.BalanceRepository
}

func (c *balanceService) UpdateBalance(data []byte) error {
	var api model.BalanceAPI
	err := json.Unmarshal(data, &api)
	if err != nil {
		logrus.Errorf("Error Unmashar Service: %v", err)
		return err
	}

	var balance = &model.Balance{}
	err = lib.Merge(api, balance)
	if err != nil {
		logrus.Errorf("Error Merger Service: %v", err)
		return err
	}

	balance.UpdatedAt = lib.TimePtr(time.Now())
	return c.balanceRepository.Update(balance.WalletID.String(), balance)
}

func NewBalanceService(balanceRepository repository.BalanceRepository) BalanceService {
	return &balanceService{
		balanceRepository: balanceRepository,
	}
}
