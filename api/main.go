package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Upload()

	fmt.Println(config.ConnectionString)
	fmt.Println("Running the API!")

	r := router.Generate()

	fmt.Println(config.SecretKey)

	fmt.Printf("Running on the port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
