package sdk

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type cache struct {
	Data      map[string]*Meta `json:"d"`
	Namespace string           `json:"n"`
	lock      sync.RWMutex
}

func newCache(namespace string) *cache {
	return &cache{Data: make(map[string]*Meta), Namespace: namespace}
}

func (c *cache) get(key string) *Meta {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.Data[key]
}

func (c *cache) add(key, value string) {
	c.lock.Lock()
	c.Data[key] = &Meta{Key: key, Value: value}
	c.lock.Unlock()
}

func (c *cache) delete(key string) {
	c.lock.Lock()
	delete(c.Data, key)
	c.lock.Unlock()
}

func (c *cache) update(key, value string) {
	c.lock.Lock()
	if meta, ok := c.Data[key]; ok {
		meta.Value = value
	} else {
		c.Data[key] = &Meta{Key: key, Value: value}
	}
	c.lock.Unlock()
}

func (c *cache) save(path string) error {
	c.lock.RLock()
	content, err := json.Marshal(c)
	c.lock.RUnlock()
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s.json", path, c.Namespace)
	return os.WriteFile(filePath, content, 0644)
}

type caches struct {
	data map[string]*cache
	path string
}

func newCaches() *caches {
	return &caches{data: make(map[string]*cache)}
}

func (c *caches) load(namespace string) {
	if content, err := os.ReadFile(fmt.Sprintf("%s/%s.json", c.path, namespace)); err == nil {
		var ca cache
		if err := json.Unmarshal(content, &ca); err == nil {
			c.data[namespace] = &ca
		}
	}
}

func (c *caches) get(namespace string) *cache {
	ch := newCache(namespace)
	if ch, ok := c.data[namespace]; ok {
		return ch
	}

	c.data[namespace] = ch
	return ch
}

func (c *caches) save() error {
	stat, err := os.Stat(c.path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(c.path, 0755)
		}

		if !os.IsExist(err) {
			return err
		}
	}

	if !stat.IsDir() {
		return fmt.Errorf("path[%s] not directory", c.path)
	}

	for _, ca := range c.data {
		if err := ca.save(c.path); err != nil {
			return err
		}
	}

	return nil
}
