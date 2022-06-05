package repository

import (
	"encoding/json"
	"github.com/astomiwanda/go-coinbit/consumer/model"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

type BalanceRepository interface {
	GetBalance(walletID string) (*model.Balance, error)
}

type balanceRepository struct {
	db *leveldb.DB
}

func (c *balanceRepository) GetBalance(walletID string) (balance *model.Balance, err error) {
	data, err := c.db.Get([]byte(walletID), nil)
	if err != nil {
		logrus.Errorf("Error Get: %v", err)
		return nil, err
	}

	balance = &model.Balance{}
	err = json.Unmarshal(data, balance)
	if err != nil {
		logrus.Errorf("Error unmarshal repository: %v", err)
		return nil, err
	}

	return
}

func NewBalanceRepository(db *leveldb.DB) BalanceRepository {
	return &balanceRepository{db: db}
}
