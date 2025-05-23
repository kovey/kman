package cache

import (
	"encoding/json"
	"fmt"
	"os"
)

type Meta struct {
	Key   string `json:"k"`
	Value string `json:"v"`
	Valid bool   `json:"vl"`
}

type metas struct {
	Data      map[string]*Meta `json:"data"`
	Namespace string           `json:"namespace"`
}

func newMetas() *metas {
	return &metas{Data: make(map[string]*Meta)}
}

func (m *metas) cached(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0755)
		}

		if !os.IsExist(err) {
			return err
		}
	}

	if !stat.IsDir() {
		return fmt.Errorf("path[%s] not directory", path)
	}

	content, err := json.Marshal(m)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s.json", path, m.Namespace)
	return os.WriteFile(filePath, content, 0644)
}

func (m *metas) has(key string) bool {
	_, ok := m.Data[key]
	return ok
}

func (m *metas) get(key string) *Meta {
	return m.Data[key]
}

func (m *metas) add(key, value string) bool {
	if meta, ok := m.Data[key]; ok {
		if meta.Valid && meta.Value == value {
			return false
		}

		meta.Value = value
		if !meta.Valid {
			meta.Valid = true
		}
		return true
	}

	m.Data[key] = &Meta{Key: key, Value: value, Valid: true}
	return true
}

func (m *metas) addInvalid(key, value string) bool {
	if _, ok := m.Data[key]; ok {
		return false
	}

	m.Data[key] = &Meta{Key: key, Value: value, Valid: false}
	return true
}
