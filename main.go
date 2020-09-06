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
    MessageId                           string `xml:"MessageHeader>MessageId"`
    SenderPartyId                       string `xml:"MessageHeader>MessageSender>PartyId"`
	PartnerName                         string `xml:"MessageHeader>MessageSender>PartnerName"`
    FullName                            string `xml:"MessageHeader>MessageSender>PartyName>FullName"`
    RecipeintPartyId                    string `xml:"MessageHeader>MessageRecipient>PartyId"`
    MessageCreatedDateTime              string `xml:"MessageHeader>MessageCreatedDateTime"`
    MessageControlType                  string `xml:"MessageHeader>MessageControlType"`
    /*
	ReleaseDate                         string `xml:"MessageHeader>MessageDetails>ReleaseDate"`
	SongTitle                           string `xml:"MessageHeader>MessageDetails>SongTitle"`
	ArtistName                          string `xml:"MessageHeader>MessageDetails>ArtistName"`
	*/
	UpdateIndicator                     string `xml:"UpdateIndicator"`
	SoundRecordingType                  string `xml:"ResourceList>SoundRecording>SoundRecordingType"`

    SoundRecordingId   struct {
    				Text          string `xml:",chardata"`
    				ISRC          string `xml:"ISRC"`
    				ProprietaryId []struct {
    					Text      string `xml:",chardata"`
    					Namespace string `xml:"Namespace,attr"`
    				} `xml:"ProprietaryId"`
    			} `xml:"ResourceList>SoundRecording>SoundRecordingId"`


    ResourceReference                   string `xml:"ResourceList>SoundRecording>ResourceReference"`
    ReferenceTitle    struct {
    				Text                  string `xml:",chardata"`
    				LanguageAndScriptCode string `xml:"LanguageAndScriptCode,attr"`
    				TitleText             string `xml:"TitleText"`
    			} `xml:"ResourceList>SoundRecording>ReferenceTitle"`

    Duration                            string `xml:"ResourceList>SoundRecording>Duration"`
    TerritoryCode                       string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>TerritoryCode"`
    Title         struct {
    					Text      string `xml:",chardata"`
    					TitleType string `xml:"TitleType,attr"`
    					TitleText string `xml:"TitleText"`
    				} `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>Title"`
    //TitleText                           string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>Title" TitleType='DisplayTitle'">TitleText"`
    DisplayArtistFullName               string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>DisplayArtist>PartyName>FullName"`
    ArtistRole                          string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>DisplayArtist>ArtistRole"`
    LabelName                           string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>LabelName"`
    RightsControllerFullName            string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>RightsController>PartyName>FullName"`
    RightsControllerPartyId             string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>RightsController>PartyId"`
    RightsControllerRole                string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>RightsController>RightsControllerRole"`
    RightSharePercentage                string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>RightsController>RightSharePercentage"`
    PLineYear                           string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>PLine>Year"`
    PLineText                           string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>PLine>PLineText"`
	Genre                               string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>Genre>GenreText"`
	ParentalExplicitWarning             string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>ParentalWarningType"`
	TechnicalResourceDetailsReference   string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>TechnicalSoundRecordingDetails>TechnicalSoundRecordingDetailsReference"`
	ReleaseList                         string `xml:"ReleaseList"`
    DealList                            string `xml:"DealList"`
}


/*
type NewReleaseMessage struct {
    MessageId               string          `xml:"MessageHeader>MessageId"`
    SenderPartyId           string          `xml:"MessageHeader>MessageSender>PartyId"`
    FullName                string          `xml:"MessageHeader>MessageSender>PartyName>FullName"`
    RecipeintPartyId        string          `xml:"MessageHeader>MessageRecipient>PartyId"`
    MessageCreatedDateTime  string          `xml:"MessageHeader>MessageCreatedDateTime"`

    PartnerName             string          `xml:"MessageHeader>MessageDetails>PartnerName"`
    ReleaseDate             string          `xml:"MessageHeader>MessageDetails>ReleaseDate"`
    SongTitle               string          `xml:"MessageHeader>MessageDetails>SongTitle"`
    ArtistName              string          `xml:"MessageHeader>MessageDetails>ArtistName"`

    UpdateIndicator         string          `xml:"UpdateIndicator"`
    SoundRecordingType      string          `xml:"ResourceList>SoundRecording>SoundRecordingType"`
    ISRC                    string          `xml:"ResourceList>SoundRecording>SoundRecordingId>ISRC"`
    ResourceReference       string          `xml:"ResourceList>SoundRecording>ResourceReference"`
    ReferenceTitleText      string          `xml:"ResourceList>SoundRecording>ReferenceTitle LanguageAndScriptCode='en'>TitleText"`
    Duration                string          `xml:"ResourceList>SoundRecording>Duration"`
    TerritoryCode           string          `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>TerritoryCode"`
    TitleText               string          `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>Title TitleType='DisplayTitle'>TitleText"`
    DisplayArtistFullName   string          `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>DisplayArtist>PartyName>FullName"`
    LabelName               string          `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>LabelName"`
    Genre                   string          `xml:"ResourceList>SoundRecording>Genre"`
    ReleaseList             string          `xml:"ReleaseList"`
    DealList                string          `xml:"DealList"`
}
*/

type SinglesJSON []SingleJSON

//type SinglesXML []NewReleaseMessage

func LiveDownloadReleaseContent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Endpoint Hit: Download Single Release Content")
	var fileName = "300_Entertainment_Output.xml"

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
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
	//result := &Results{}
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
