package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestWallet(t *testing.T) {

	//Wallet Balance Test
	t.Run("Wallet balance test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(15))

		got := dummy_wallet.Balance()
		wanted := 15.0

		if got != wanted {
			t.Errorf("expected %v got %v", wanted, got)
		}
	})
	//Wallet Credit Test
	t.Run("Wallet credit test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(15))

		dummy_wallet.Credit(decimal.NewFromFloat(5.0))

		got := dummy_wallet.Balance()
		wanted := 10.0

		if got != wanted {
			t.Errorf("expected %v, got %v", wanted, got)
		}

	})
	//Negative Credit Test
	t.Run("Wallet negative credit test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(15))

		got := dummy_wallet.Credit(decimal.NewFromFloat(-5))

		if got == nil {
			t.Errorf("expected error got nil")
		}

	})
	//Wallet Debit Test
	t.Run("Wallet debit test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(15))

		dummy_wallet.Debit(decimal.NewFromFloat(5))

		got := dummy_wallet.Balance()
		wanted := 20.0

		if got != wanted {
			t.Errorf("expected %v, got %v", wanted, got)
		}
	})
	//Wallet credit can not higher than balance Test
	t.Run("Credit can not higher than balance", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(5))

		got := dummy_wallet.Credit(decimal.NewFromFloat(10))

		if got == nil {
			t.Errorf("expected error got nil")
		}
	})
	//Wallet Credit can not be negative Test
	t.Run("Credit can not higher than balance", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(5))

		got := dummy_wallet.Credit(decimal.NewFromFloat(-5))

		if got == nil {
			t.Errorf("expected error got nil")
		}
	})
}

func getEmptyWallet(id int, amonut decimal.Decimal) Wallet {
	return Wallet{
		ID:             id,
		Wallet_balance: amonut,
	}
}
