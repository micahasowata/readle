package main

import (
	"fmt"
	"log"
	"net/http"

	"micahasowata.com/restapi/internal/handlers"
	"micahasowata.com/restapi/internal/routes"
	"micahasowata.com/restapi/serverconfig"
)

func main() {
	config, err := serverconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}

	handler := handlers.NewHandler()

	mux := http.NewServeMux()
	routes.SetupRoutes(mux, handler)

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	log.Printf("server is up and running at port %s", serverAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed %s\n", err.Error())
	}
}
