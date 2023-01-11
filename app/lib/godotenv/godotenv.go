package godotenv

import (
	"Sayaka/lib"
	"github.com/joho/godotenv"
)

var _ lib.Preparer = &GoDotEnv{}

type GoDotEnv struct {
	filenames []string
}

func NewGoDotEnv() *GoDotEnv {
	return &GoDotEnv{defaultFileNames()}
}

func defaultFileNames() []string {
	return []string{".env.local"}
}

func (e *GoDotEnv) Prepare() error {
	filenames := e.filenames
	if err := godotenv.Load(filenames...); err != nil {
		return err
	}
	return nil
}
