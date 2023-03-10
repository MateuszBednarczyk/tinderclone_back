package services

type ICountrier interface {
}

type countrier struct {
}

func NewCountrier() *countrier {
	return &countrier{}
}
