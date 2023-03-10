package services

type ICitier interface {
	SaveNewCity() *Result
}

type citier struct {
}

func NewCitier() *citier {
	return &citier{}
}
