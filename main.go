package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"bytes"
	"time"

	"InvsbleMusicDashboard/routes"
	//"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Youtube XML Ingestion Dashboard Application")
	//handleRequests()

	r := routes.Router()
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
