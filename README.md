# Simple Exchange

A really simple (and probably full of jargon misuage) exchange, to help me learn go and be base example of a future distributed systems guild to help learn different distributed systems pitfalls, etc.

## Run Locally

```sh
go run order-service/*.go
go run matching-service/*.go
```

Example order JSON:
```json
{
    "Id": "d60d7f82-e2d9-4905-bc51-b0d7ebfb5389",
    "Symbol": {
        "BidAsset": "ETH",
        "AskAsset": "BTC"
    },
    "BidPrice": 1,
    "AskPrice": 28258
}
```

Calling endpoints
```sh
curl http://localhost:8080/orders
curl -d "@order.json" -X POST http://localhost:8080/order 
curl -d "@order.json" -X POST http://localhost:8081/match
```
