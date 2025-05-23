package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type namespaces struct {
	data      map[string]*metas
	fs        map[string]*files
	cachePath string
}

func newNamespaces() *namespaces {
	return &namespaces{data: make(map[string]*metas), fs: make(map[string]*files)}
}

func (n *namespaces) loadFromCache(namespace string) error {
	if _, ok := n.data[namespace]; ok {
		return nil
	}

	n.cachePath = os.Getenv("CACHE_PATH")
	file := fmt.Sprintf("%s/%s.json", n.cachePath, namespace)
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	n.data[namespace] = newMetas()
	return json.Unmarshal(content, n.data[namespace])
}

func (n *namespaces) init() error {
	n.cachePath = os.Getenv("CACHE_PATH")
	namespace := strings.Split(os.Getenv("CONFIG_NAMESPACE"), ",")
	path := strings.Split(os.Getenv("CONFIG_PATH"), ",")
	for index, ns := range namespace {
		if _, ok := n.data[ns]; ok {
			continue
		}

		ns = strings.Trim(ns, " ")
		n.data[ns] = newMetas()
		n.data[ns].Namespace = ns
		n.fs[ns] = newFiles(strings.Trim(path[index], " "))
		if err := n.fs[ns].load(); err != nil {
			return err
		}
	}

	return nil
}

func (n *namespaces) get(namespace, key string) *Meta {
	if m, ok := n.data[namespace]; ok {
		return m.get(key)
	}

	return nil
}

func (n *namespaces) add(namespace, key, value string) error {
	if _, ok := n.data[namespace]; !ok {
		return nil
	}

	if !n.data[namespace].add(key, value) {
		return nil
	}

	return n.fs[namespace].flush(key, value)
}

func (n *namespaces) cached() error {
	if os.Getenv("CACHE_OPEN") != "On" {
		return nil
	}

	for _, meta := range n.data {
		if err := meta.cached(n.cachePath); err != nil {
			return err
		}
	}

	return nil
}

func (n *namespaces) addToFile(namespace, key, value string) {
	n.fs[namespace].add(key, value)
}

func (n *namespaces) parse() error {
	for _, fis := range n.fs {
		if err := fis.parse(); err != nil {
			return err
		}
	}

	return nil
}

var ns = newNamespaces()

func Init() error {
	return ns.init()
}

func Add(namespace, key, value string) error {
	return ns.add(namespace, key, value)
}

func Get(namespace, key string) *Meta {
	return ns.get(namespace, key)
}

func Parse() error {
	return ns.parse()
}

func LoadFromCache(namespace string) error {
	return ns.loadFromCache(namespace)
}

func AddToFile(namespace, key, value string) {
	ns.add(namespace, key, value)
}

func Cached() error {
	return ns.cached()
}
