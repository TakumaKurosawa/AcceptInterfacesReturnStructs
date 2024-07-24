package uid

import (
	"github.com/google/uuid"
)

type Generator interface {
	NewUUIDV7() (string, error)
}

type generator struct{}

func NewGenerator() *generator {
	return &generator{}
}

func (g *generator) NewUUIDV7() (string, error) {
	result, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
