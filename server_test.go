package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
			42: {id: 42, balance: 20},
		},
	})

	t.Run("get balance", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/wallets/42/balance", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := struct{ Balance int }{}
		json.NewDecoder(response.Body).Decode(&got)

		want := 20

		if got.Balance != want {
			t.Errorf("got %d , want %d", got, want)
		}
	})
}
