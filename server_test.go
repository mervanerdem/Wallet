package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TStorage struct {
	dataholder map[int]Wallet
}

func (t *TStorage) Update(wallet Wallet) error { return nil }
func (t *TStorage) Get(id int) (*Wallet, error) {
	if res, ok := t.dataholder[id]; ok {
		return &res, nil
	}

	return nil, fmt.Errorf("not found")
}

func TestServer(t *testing.T) {
	server := NewServer(&TStorage{
		dataholder: map[int]Wallet{
			1: {id: 1, balance: 10},
			2: {id: 2, balance: 20},
			3: {id: 3, balance: 30},
			4: {id: 4, balance: 40},
			5: {id: 5, balance: 50},
		},
	})
	// Get Balance
	t.Run("get balance", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/wallets/1/balance", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := struct{ Balance int }{}
		json.NewDecoder(response.Body).Decode(&got)

		want := 10

		if got.Balance != want {
			t.Errorf("got %d , want %d", got, want)
		}
	})
	//Balance id not avaible
	t.Run("get balance not found", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/wallets/22/balance", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Result().StatusCode != http.StatusNotFound {
			t.Errorf("expected code 404, got %d", response.Result().StatusCode)
		}
	})
	//Debit
	t.Run("use debit", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/debit", strings.NewReader(`{"Amount":10}`))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("expected code 200, got %d", response.Result().StatusCode)
		}
	})
	//Negative debit
	t.Run("can not take negative debit", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/debit", strings.NewReader(`{"Amount":-5}`))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("expected code 200, got %d", response.Result().StatusCode)
		}
	})

	//credit
	t.Run("use credit", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/wallets/2/credit", strings.NewReader(`{"Amount":10}`))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("expected code 200, got %d", response.Result().StatusCode)
		}
	})

	//Negative credit
	t.Run("can not take negative credit", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/wallets/2/credit", strings.NewReader(`{"Amount":10}`))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("expected code 200, got %d", response.Result().StatusCode)
		}
	})

}
