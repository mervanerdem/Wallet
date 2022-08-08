# Wallet
## Context
Wallets of the players of an online casino, getting and updating players account balances API's.

## How to Install

### Firstly Clone the repository

```go
git clone github.com/mervanerdem/Wallet
```
After cloning if you don't have used library, install them.(It's not about using library but you can see in that [link](https://pkg.go.dev/cmd/go/internal/get))

### Build or Run it
For building
```go
go build ./
```
Or Runing
```Go
go run ./
```
## How to Use
In this project use the localhost.
```
baseURL: localhost:8080
```
### Get wallet balance
In  this case with wallet id get balance of the wallet. 
```
GET {{baseURL}}/api/v1/wallets/{wallet_id}/balance
```
The place of {wallet_id} write in to which wallet balance do you need.

The responce be like
```JSON
{
    "Balance": "800",
    "Wallet_ID": 42
}
```
### Update Balance With Credit
In this case Send Credit  from wallet
```
POST {{baseURL}}/api/v1/wallets/{wallet_id}/credit
```
The Body Type is JSON format:
```JSON
{
    "Amount": 10
}
```
The responce is like:
```JSON
{
    "Balance": "450",
    "Credit": "10",
    "Wallet_ID": 30
}
```
### Update Balance With Debit
In this case Take Debit from another wallet
```
POST {{baseURL}}/api/v1/wallets/{wallet_id}/debit
```
The Body Type is JSON format:
```JSON
{
    "Amount": 10
}
```
The responce is like:
```JSON
{
    "Balance": "460",
    "Credit": "10",
    "Wallet_ID": 30
}
```

## Library
Used Library

[HTTP](https://github.com/gin-gonic/gin) : Gin is a web framework written in Go (Golang)

[GORM](https://github.com/go-gorm/gorm) : The fantastic ORM library for Golang, aims to be developer friendly.

[Decimal](https://github.com/shopspring/decimal) :Arbitrary-precision fixed-point decimal numbers in go.


## Licanse

Apache-2.0 License



