package repository

import (
	"errors"
	"github.com/rmrachmanfauzan/bank_assessment/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	TopupAccount(string, float64) (*model.Account, error)
	WithdrawAccount(string, float64) (*model.Account, error)
	GetSaldo(no_rekening string) (*model.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) TopupAccount(noRekening string, nominal float64) (*model.Account,error) {	
	var account model.Account

	
	err := r.db.Where("no_rekening = ?", noRekening).First(&account).Error
	if err != nil {
		return nil, errors.New("rekening not found")
	}

	// Top up the balance
	account.Saldo += nominal

	// Save
	if err := r.db.Save(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}


func (r *accountRepository) WithdrawAccount(noRekening string, nominal float64) (*model.Account,error) {	
	var account model.Account

	
	err := r.db.Where("no_rekening = ?", noRekening).First(&account).Error
	if err != nil {
		return nil, errors.New("rekening not found")
	}
	// Top up the balance
	if nominal > account.Saldo {	
		return nil, errors.New("insufficient balance")		
	}
	
	account.Saldo -= nominal

	// Save
	if err := r.db.Save(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
func (r *accountRepository) GetSaldo(no_rekening string) (*model.Account, error) {
	var account model.Account
	err := r.db.Where("no_rekening = ?", no_rekening).First(&account).Error
	if err != nil {
		return nil, errors.New("rekening not found")
	}
	return &account, nil
}


