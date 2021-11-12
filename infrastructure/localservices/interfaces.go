package localservices

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=localservices
type ICutter interface {
	Generate() string
}

type myRandInterface interface {
	Read([]byte) (int, error)
}
