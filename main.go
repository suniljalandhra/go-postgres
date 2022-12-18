package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suniljalandhra/go-postgres/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
