package repository

import (
	"encoding/json"
	"fmt"
	"github.com/astomiwanda/go-coinbit/consumer/model"
	"github.com/astomiwanda/go-coinbit/lib"
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
		logrus.Errorf("Error Get Balance: %v", err)
		return nil, err
	}

	balance = &model.Balance{}
	err = json.Unmarshal(data, balance)
	if err != nil {
		logrus.Errorf("Error unmarshal balance repository: %v", err)
		return nil, err
	}

	thresholdKey := fmt.Sprintf("%s%s", lib.BalanceThreshold, walletID)
	dataThreshold, err := c.db.Get([]byte(thresholdKey), nil)
	if err != nil {
		logrus.Errorf("Error Get Threshold: %v", err)
		return nil, err
	}

	threshold := &model.Threshold{}
	err = json.Unmarshal(dataThreshold, threshold)
	if err != nil {
		logrus.Errorf("Error unmarshal threshold repository: %v", err)
		return nil, err
	}

	balance.IsThreshold = threshold.IsThreshold
	logrus.Infof("Balance: %v", balance)
	return
}

func NewBalanceRepository(db *leveldb.DB) BalanceRepository {
	return &balanceRepository{db: db}
}
