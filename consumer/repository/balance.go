package repository

import (
	"encoding/json"
	"github.com/astomiwanda/go-coinbit/consumer/model"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

type BalanceRepository interface {
	Update(walletID string, balance *model.Balance) error
}

type balanceRepository struct {
	db *leveldb.DB
}

func (c *balanceRepository) Update(walletID string, balance *model.Balance) error {
	balanceByte, err := json.Marshal(balance)
	if err != nil {
		logrus.Errorf("Error Marshal: %v", err)
		return err
	}

	err = c.db.Put([]byte(walletID), balanceByte, nil)
	if err != nil {
		logrus.Errorf("Error Put: %v", err)
		return err
	}

	return nil
}

func NewBalanceRepository(db *leveldb.DB) BalanceRepository {
	return &balanceRepository{db: db}
}
