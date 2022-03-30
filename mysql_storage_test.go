package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestMYSQLStorage(t *testing.T) {

	storage, db, err := NewMStorage("'root':Mervan.1907@tcp(127.0.0.1:3306)/Wallet_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Error(err)
		return
	}
	db.AutoMigrate(&Wallet{})

	storage.New(Wallet{ID: 42, Wallet_balance: decimal.NewFromInt(1000)})
	if wallet, _ := storage.Get(42); wallet == nil {
		t.Error("Expected wallet get nil")
	}

	storage.New(Wallet{ID: 30, Wallet_balance: decimal.NewFromInt(500)})
	if wallet, _ := storage.Get(30); wallet == nil {
		t.Error("expected wallet, got nil")
	}

	if _, err := storage.Get(2); err == nil {
		t.Error("expected error")
	}

	storage.Update(Wallet{ID: 42, Wallet_balance: decimal.NewFromInt(800)})
	if wallet, err := storage.Get(42); err == nil {
		if wallet.Balance() != 800 {
			t.Errorf("expected balance 800, got %f", wallet.Balance())
		}
	} else {
		t.Error("expected wallet, got nil, error: " + err.Error())
	}

}
