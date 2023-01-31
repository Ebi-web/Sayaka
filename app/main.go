package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"Sayaka/api"
	"Sayaka/lib"
	"Sayaka/lib/godotenv"
)

func main() {
	var port int
	var databaseDatasource string

	flag.IntVar(&port, "port", 8080, "API server port")
	flag.StringVar(&databaseDatasource, "databaseDatasource", "mysql://root:password@tcp(127.0.0.1:3306)/sayaka?sslmode=disable", "")
	flag.Parse()

	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	if err := preloadModules(libraries()); err != nil {
		panic("🚫External modules setup unsuccessful")
	}
	fmt.Println("🎉Successful external modules setup")

	server := api.NewServer()
	if err := server.Init(databaseDatasource); err != nil {
		log.Fatal(err)
	}
	server.Run(port)
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
