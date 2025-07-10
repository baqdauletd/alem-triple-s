package run

import (
	"fmt"
	"log"
	"net/http"
	"triple-s/base"
	"triple-s/handlers"
)

func Launch() {
	err := base.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	err = handlers.DirInit()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", base.Port),
		Handler: handlers.RouterWays(),
	}

	log.Printf("Starting the server on %d...\n", base.Port)
	log.Printf("Data dir: %s", base.Dir)
	err = server.ListenAndServe()
	log.Fatal(err)
}
