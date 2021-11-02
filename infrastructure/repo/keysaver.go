package repo

import (
	"errors"
	"sync"
	"url-changer/domain"
)

type keySaver struct {
	repo123 TemporaryRepo
	mu      sync.RWMutex
}

func NewKeySaver() *keySaver {
	repo := make(map[string]string, 0)
	return &keySaver{repo123: TemporaryRepo{repo}}
}

func (k *keySaver) Get(s string) (string, error) {
	url, ok := k.repo123.Repo[s]
	if !ok {
		return "", errors.New("no such url")
	}
	return url, nil
}

func (k *keySaver) Save(url domain.LongURL, key string) error {
	k.mu.RLock()
	k.repo123.Repo[key] = url.LongURLData
	k.mu.RUnlock()
	return nil
}

type TemporaryRepo struct {
	Repo map[string]string
}
