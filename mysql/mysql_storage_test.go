package mysql

import (
	"testing"
	"wallet/Model"

	"github.com/shopspring/decimal"
)

func TestMYSQLStorage(t *testing.T) {

	storage, db, err := NewMStorage("tester:secret@tcp(db:3306)/test")
	if err != nil {
		t.Error(err)
		return
	}
	db.AutoMigrate(&Model.Wallet{})

	storage.New(Model.Wallet{ID: 42, Wallet_balance: decimal.NewFromInt(1000)})
	if wallet, _ := storage.Get(42); wallet == nil {
		t.Error("Expected wallet get nil")
	}

	storage.New(Model.Wallet{ID: 30, Wallet_balance: decimal.NewFromInt(500)})
	if wallet, _ := storage.Get(30); wallet == nil {
		t.Error("expected wallet, got nil")
	}

	if _, err := storage.Get(200); err == nil {
		t.Error("expected error")
	}

	storage.Update(Model.Wallet{ID: 42, Wallet_balance: decimal.NewFromInt(800)})
	if wallet, err := storage.Get(42); err == nil {
		if wallet.Balance().Equal(decimal.NewFromFloat(800)) == false {
			t.Errorf("expected balance 800, got %v", wallet.Balance())
		}
	} else {
		t.Error("expected wallet, got nil, error: " + err.Error())
	}

}
