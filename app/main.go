package main

import (
	"fmt"
	"net/http"

	"Sayaka/controllers"
	"Sayaka/lib"
)

var (
	goDotEnv = lib.NewGoDotEnv()
)

func main() {
	if err := preloadModules([]lib.Preparer{
		goDotEnv,
	}); err != nil {
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

func preloadModules(libs []lib.Preparer) error {
	for k := range libs {
		if err := libs[k].Prepare(); err != nil {
			return err
		}
	}
	return nil
}
