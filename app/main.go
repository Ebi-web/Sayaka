package main

import (
	"fmt"
	"net/http"

	"Sayaka/controllers"
	"Sayaka/lib"
	"Sayaka/lib/godotenv"
)

func main() {
	if err := preloadModules(libraries()); err != nil {
		panic("External modules setup unsuccessful")
	}
	fmt.Println("🎉Successful external modules setup")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		s, reply, err := controllers.Respond(w, r)
		w.WriteHeader(s)
		if err != nil {
			fmt.Fprintf(w, "I'm sorry, but would you retry the request?")
		}
		fmt.Fprintf(w, reply)
	})

	http.ListenAndServe(":8080", nil)
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
