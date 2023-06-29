# Kwil fractal-eth-extension example

This is a Kwil extension that expose fractal ethereum smart contract functions(external view/pure) to Kwil, developer can call those functions in Kuneiform. 

## configuration

This extension will read configuration from system environment variables.

```bash
export RPC_URL=xxx
```

## run

```bash
# run with go
go run main.go

# or docker
make docker
# put your configuration in .env file
docker compose up -d
```