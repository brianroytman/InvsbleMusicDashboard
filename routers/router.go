package routers

import (
    "InvsbleMusicDashboard/middleware"
    "github.com/gorilla/mux"
    "fmt"
  	"log"
  	"net/http"
)

// Router is exported and used in main.go
func Router() *mux.Router {
  	router := mux.NewRouter().StrictSlash(true)

  	fmt.Println("Created Mux router!!")

  	router.HandleFunc("/uploadReleaseContent", middleware.UploadReleaseContent).Methods("POST")
  	router.HandleFunc("/downloadYouTubeIngestionXML", middleware.DownloadYouTubeIngestionXML).Methods("GET")

  	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend-react-dashboard/build/static/"))))
  	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  		http.ServeFile(w, r, "./frontend-react-dashboard/build/index.html")
  	})

  	log.Fatal(http.ListenAndServe(":8080", router))

    return router
}
