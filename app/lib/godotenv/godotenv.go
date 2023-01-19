package godotenv

import (
	"fmt"

	"Sayaka/lib"
	"github.com/joho/godotenv"
)

var _ lib.Preparer = &GoDotEnv{}

type GoDotEnv struct {
	localEnvFileName string
}

func NewGoDotEnv() *GoDotEnv {
	return &GoDotEnv{localEnvFileName()}
}

func localEnvFileName() string {
	return ".env.local"
}

func (e *GoDotEnv) Prepare() error {
	lfn := e.localEnvFileName
	if err := godotenv.Load(lfn); err != nil {
		//considered as production environment
		fmt.Println("ℹ️ Considered as production environment. Ensure that environment variables are set.")
		return nil
	}
	return nil
}
