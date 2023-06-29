package extension

import (
	"fmt"
	"github.com/hashicorp/golang-lru/v2"
	"github.com/kwilteam/extension-fractal-demo/extension/registry"
	"log"
)

type registryGetter interface {
	GetRegistryInstance(chain string, address string) (*registry.Registry, error)
}

type cacheRegistryGetter struct {
	registryGetter
	cache *lru.Cache[string, *registry.Registry]
}

func newDefaultRegistryGetter(registryGetter registryGetter) *cacheRegistryGetter {
	cache, err := lru.New[string, *registry.Registry](5)
	if err != nil {
		log.Fatalf("Error creating cache: %s", err)
	}
	return &cacheRegistryGetter{
		registryGetter: registryGetter,
		cache:          cache}
}

func (c *cacheRegistryGetter) GetRegistryInstance(chain string, address string) *registry.Registry {
	key := fmt.Sprintf("%s-%s", chain, address)
	if v, ok := c.cache.Get(key); ok {
		log.Printf("Registry [%s] found in cache", chain+address)
		return v
	}

	log.Printf("Key [%s] not found in cache", key)
	v, err := c.registryGetter.GetRegistryInstance(chain, address)
	if err != nil {
		log.Fatalf("Error getting registry: %s", err)
	}
	c.cache.Add(key, v)
	return v
}
