package main

import "testing"

func TestWallet(t *testing.T) {

	//Wallet Balance Test
	t.Run("Wallet balance test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, 15)

		got := dummy_wallet.Balance()
		wanted := 15

		if got != wanted {
			t.Errorf("expected %d, got %d", wanted, got)
		}
	})
	//Wallet Credit Test
	t.Run("Wallet credit test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, 15)

		dummy_wallet.Credit(5)

		got := dummy_wallet.Balance()
		wanted := 20

		if got != wanted {
			t.Errorf("expected %d, got %d", wanted, got)
		}

	})
	//Negative Credit Test
	t.Run("Wallet negative credit test", func(t *testing.T) {
		dummy_wallet := getEmptyWallet(1, 15)

		got := dummy_wallet.Credit(-5)

		if got == nil {
			t.Errorf("expected error got nil")
		}

	})

}

func getEmptyWallet(id, amonut int) Wallet {
	return Wallet{
		id:      id,
		balance: amonut,
	}
}
