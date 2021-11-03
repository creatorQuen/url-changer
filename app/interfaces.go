package app

import "url-changer/domain"

type KeyGenerator interface {
	MakeKey(string) (string, error)
	GetURL(string) (string, error)
}

type UrlSaver interface {
	Save(domain.LongURL, string) error
	GetFullString(string) (string, error)
}
