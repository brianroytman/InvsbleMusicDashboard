package middleware

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"InvsbleMusicDashboard/models" // models package where User schema is defined
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable
	"strconv"  // package used to covert string into int type
	"github.com/gorilla/mux" // used to get the params from the route
  )

	func DownloadYouTubeIngestionXML(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET Endpoint Hit: Download Single Release Content")
		var fileName = "300_Entertainment_Output.xml"

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Expires", "0")

		http.ServeFile(w, r, fileName)
	}

	func UploadReleaseContent(w http.ResponseWriter, r *http.Request) {
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

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Expires", "0")

		http.ServeContent(w, r, fileName, modtime, bytes.NewReader(xmlOut))
	}
