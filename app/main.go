package main

import (
	"fmt"

	"Sayaka/lib"
	"Sayaka/lib/godotenv"
)

func main() {
	if err := preloadModules(libraries()); err != nil {
		panic("External modules setup unsuccessful")
	}
	fmt.Println("🎉Successful external modules setup")

	server := NewServer()
	server.Init()
}

func libraries() []lib.Preparer {
	return []lib.Preparer{
		godotenv.NewGoDotEnv(),
	}
}

func preloadModules(libs []lib.Preparer) error {
	for k := range libs {
		if err := libs[k].Prepare(); err != nil {
			return err
		}
	}
	return nil
}
