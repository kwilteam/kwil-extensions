package extension

import "fmt"

type Config struct {
	RpcUrl string `env:"RPC_URL" envDefault:"http://localhost:8545"`
	Port   int    `env:"PORT" envDefault:"50051"`
}

func (c *Config) ListenAddr() string {
	return fmt.Sprintf(":%d", c.Port)
}
