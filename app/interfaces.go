package app

import "url-changer/domain"

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=app
type KeyGenerator interface {
	MakeKey(string) (string, error)
	GetURL(string) (string, error)
}

//// repo
type UrlSaver interface {
	Save(domain.LongURL, string) error
	GetFullString(string) (string, error)
}
