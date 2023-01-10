package lib

import (
	"github.com/joho/godotenv"
)

type Preparer interface {
	Prepare() error
}

type GoDotEnv struct {
	filenames []string
}

var _ Preparer = &GoDotEnv{}

func NewGoDotEnv() *GoDotEnv {
	return &GoDotEnv{[]string{".env.local"}}
}

func (e *GoDotEnv) Prepare() error {
	filenames := e.filenames
	if err := godotenv.Load(filenames...); err != nil {
		return err
	}
	return nil
}
