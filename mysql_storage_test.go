package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestMYSQLStorage(t *testing.T) {

	storage, db, err := NewMStorage("root:secret@tcp(127.0.0.1:3307)/Wallet")
	if err != nil {
		t.Error(err)
		return
	}
	db.AutoMigrate(&Wallet{})

	storage.New(Wallet{ID: 3, Wallet_balance: decimal.NewFromInt(1000)})
	if wallet, _ := storage.Get(3); wallet == nil {
		t.Error("Expected wallet get nil")
	}

	storage.New(Wallet{ID: 2, Wallet_balance: decimal.NewFromInt(500)})
	if wallet, _ := storage.Get(2); wallet == nil {
		t.Error("expected wallet, got nil")
	}

	if _, err := storage.Get(200); err == nil {
		t.Error("expected error")
	}

	storage.Update(Wallet{ID: 1, Wallet_balance: decimal.NewFromInt(800)})
	if wallet, err := storage.Get(1); err == nil {
		if wallet.Balance().Equal(decimal.NewFromFloat(800)) == false {
			t.Errorf("expected balance 800, got %v", wallet.Balance())
		}
	} else {
		t.Error("expected wallet, got nil, error: " + err.Error())
	}

}
