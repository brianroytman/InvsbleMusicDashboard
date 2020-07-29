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

	"github.com/gorilla/mux"
)

type SingleJSON struct {
	PartnerName string `json:"PartnerName,omitempty"`
	ReleaseDate string `json:"ReleaseDate,omitempty"`
	SongTitle   string `json:"SongTitle,omitempty"`
	ArtistName  string `json:"ArtistName, omitempty"`
	Genre       string `json:"Genre, omitempty"`
}

type NewReleaseMessage struct {
	PartnerName string `xml:"PartnerName"`
	ReleaseDate string `xml:"ReleaseDate"`
	SongTitle   string `xml:"SongTitle"`
	ArtistName  string `xml:"ArtistName"`
	Genre       string `xml:"Genre`
}

type SinglesJSON []SingleJSON

type SinglesXML []NewReleaseMessage

func LiveDownloadReleaseContent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Endpoint Hit: Download Single Release Content")
	var fileName = "300_Entertainment_Output.xml"

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	//w.Header().Set("Content-Disposition", "attachment; filename=download_single_release_output.xml")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	/*
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
	*/

	//modtime := time.Now()

	//var sX NewReleaseMessage
	//json.Unmarshal(body, &sX)
	//xmlOut, _ := xml.MarshalIndent(sX, "", "\t")

	http.ServeFile(w, r, fileName)

}

func LiveUploadReleaseContent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST Endpoint Hit: Read/Write to XML File - Single Release Content")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	var sJ SingleJSON
	err = json.Unmarshal(body, &sJ)
	if err != nil {
		panic(err)
	}
	log.Println(string(sJ.SongTitle))

	var sX NewReleaseMessage
	json.Unmarshal(body, &sX)
	xmlOut, _ := xml.MarshalIndent(sX, "", "\t")
	modtime := time.Now()

	fmt.Println(string(xmlOut))
	var fileName = "300_Entertainment_Output.xml"

	_ = ioutil.WriteFile(fileName, xmlOut, 0644)

	//var xmlReader, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	/*
		//w.Header().Add("Content-Disposition", "Attachment")
		//w.Header().Set("Content-Disposition", "attachment; filename=300_Entertainment_Output.xml")
		//w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Header().Set("Content-Disposition", "attachment; filename=300_Entertainment_Output.xml")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	*/

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	//w.Header().Set("Content-Disposition", "attachment; filename=download_single_release_output.xml")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	http.ServeContent(w, r, fileName, modtime, bytes.NewReader(xmlOut))

}

func handleRequests() {

	fmt.Println("About to create Mux router")
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Created Mux router!!")

	router.HandleFunc("/uploadLive", LiveUploadReleaseContent).Methods("POST")
	router.HandleFunc("/downloadLive", LiveDownloadReleaseContent).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend-light-bootstrap-dashboard-react/build/static/"))))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend-light-bootstrap-dashboard-react/build/index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Youtube Ingestion Dashboard Application")
	handleRequests()
}
