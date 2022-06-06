# GO Coinbit

This is repository is result of my test engineering at coinbit

## How to run

```bash
# clone this repo
git clone git@github.com:AstoMiwanda/go-coinbit.git
# or
git clone https://github.com/AstoMiwanda/go-coinbit.git

cd go-coinbit

# running docker compose
docker-compose up

# run program
go run main.go
```

## How to access endpoint

### Post
- url:
localhost:8080/api/v1/balance
- body request:
```json
{
  "wallet_id": "06ae27cd-ddd6-4c9c-92f8-461af9dbb367",
  "amount": 2000
}
```

### Get
- url:
  localhost:8080/api/v1/balance/06ae27cd-ddd6-4c9c-92f8-461af9dbb367
