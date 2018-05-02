# Status

Public Apis

- [x] Depth
- [x] Candlestick
- [x] Ticker
- [x] Transactions

Private Apis

- [] Assets
- [] Order
- [] Trade
- [] Withdraw

Realtime Apis

- [] Ticker
- [] Depth
- [] Candlestick
- [] Transactions

# How to use

```
timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

client := bitbank.New().Auth(API_KEY, API_SECRET)
depth, _ := client.Depth.Get(timeout, "btc_jpy")
```
