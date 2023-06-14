package client

import "time"

type ClientOpt func(*ExtensionClient)

func WithTimeout(timeout time.Duration) ClientOpt {
	return func(c *ExtensionClient) {
		c.timeout = timeout
	}
}
