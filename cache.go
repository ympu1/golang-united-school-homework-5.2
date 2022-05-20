package cache

import "time"

type Cache struct {
	m map[string]string
}

func NewCache() Cache {
	c := Cache{}
	c.m = make(map[string]string)

	return c
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.m[key]
	return value, ok
}

func (c *Cache) Put(key, value string) {
	c.m[key] = value
}

func (c *Cache) Keys() []string {
	var keys []string

	for key := range c.m {
		keys = append(keys, key)
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.m[key] = value

	time.Sleep(time.Until(deadline))
	delete(c.m, key)
}
