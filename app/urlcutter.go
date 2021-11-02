package app

import (
	"url-changer/domain"
	"url-changer/infrastructure/localservices"
)

type urlCutterService struct {
	keyGenerator localservices.KeyGenerator
	repo         UrlSaver
}

func NewUrlCutterService(keyGenerator localservices.KeyGenerator, repo UrlSaver) *urlCutterService {
	return &urlCutterService{keyGenerator: keyGenerator, repo: repo}
}

func (u *urlCutterService) MakeKey(s string) (string, error) {
	key := u.keyGenerator.Generate()
	url := domain.LongURL{LongURLData: s}
	err := u.repo.Save(url, key)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (u *urlCutterService) GetURL(s string) (string, error) {
	url, err := u.repo.Get(s)
	if err != nil {
		return "", err
	}
	return url, nil
}
