package Model

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestWallet(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want decimal.Decimal) {
		t.Helper()
		if got.Equal(want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	}

	//Model Balance Test
	t.Run("Model balance test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(15))

		got := dummy_wallet.Balance()
		want := decimal.NewFromFloat(15)

		assertCorrectMessage(t, got, want)
	})
	//Model Credit Test
	t.Run("Model credit test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(15.0))

		dummy_wallet.Credit(decimal.NewFromFloat(5.0))

		got := dummy_wallet.Balance()
		want := decimal.NewFromFloat(10.0)

		assertCorrectMessage(t, got, want)
	})
	//Negative Credit Test
	t.Run("Model negative credit test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(15))

		got := dummy_wallet.Credit(decimal.NewFromFloat(-5))

		if got == nil {
			t.Errorf("expected error got nil")
		}

	})
	//Model Debit Test
	t.Run("Model debit test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(15))

		dummy_wallet.Debit(decimal.NewFromFloat(5))

		got := dummy_wallet.Balance()
		wanted := decimal.NewFromFloat(20.0)

		if got.Equal(wanted) == false {
			t.Errorf("expected %v, got %v", wanted, got)
		}
	})
	//Model credit can not higher than balance Test
	t.Run("Credit can not higher than balance", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, decimal.NewFromFloat(5))

		got := dummy_wallet.Credit(decimal.NewFromFloat(10))

		if got == nil {
			t.Errorf("expected error got nil")
		}
	})
	//Model Credit can not be negative Test
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
