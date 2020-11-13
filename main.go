package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"bytes"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)


type SingleJSON struct {
	PartnerName string `json:"PartnerName,omitempty"`
	ReleaseDate string `json:"ReleaseDate,omitempty"`
	SongTitle   string `json:"SongTitle,omitempty"`
	ArtistName  string `json:"ArtistName,omitempty"`
	Genre       string `json:"Genre,omitempty"`
}

type NewReleaseMessage struct {
	XMLName                xml.Name `xml:"NewReleaseMessage"`
	Text                   string   `xml:",chardata"`
	LanguageAndScriptCode  string   `xml:"LanguageAndScriptCode,attr"`
	MessageSchemaVersionId string   `xml:"MessageSchemaVersionId,attr"`
	Ern                    string   `xml:"ern,attr"`
	Xs                     string   `xml:"xs,attr"`
	SchemaLocation         string   `xml:"schemaLocation,attr"`
	MessageHeader          struct {
		Text          string `xml:",chardata"`
		MessageId     string `xml:"MessageId"`
		MessageSender struct {
			Text      string `xml:",chardata"`
			PartyId   string `xml:"PartyId"`
			PartyName struct {
				Text     string `xml:",chardata"`
				FullName string `xml:"FullName"`
			} `xml:"PartyName"`
		} `xml:"MessageSender"`
		MessageRecipient struct {
			Text      string `xml:",chardata"`
			PartyId   string `xml:"PartyId"`
			PartyName struct {
				Text     string `xml:",chardata"`
				FullName string `xml:"FullName"`
			} `xml:"PartyName"`
		} `xml:"MessageRecipient"`
		MessageCreatedDateTime string `xml:"MessageCreatedDateTime"`
		MessageControlType     string `xml:"MessageControlType"`
	} `xml:"MessageHeader"`
	UpdateIndicator string `xml:"UpdateIndicator"`
	ResourceList    struct {
		Text           string `xml:",chardata"`
		SoundRecording struct {
			Text               string `xml:",chardata"`
			SoundRecordingType string `xml:"SoundRecordingType"`
			SoundRecordingId   struct {
				Text          string `xml:",chardata"`
				ISRC          string `xml:"ISRC"`
				ProprietaryId []struct {
					Text      string `xml:",chardata"`
					Namespace string `xml:"Namespace,attr"`
				} `xml:"ProprietaryId"`
			} `xml:"SoundRecordingId"`
			ResourceReference string `xml:"ResourceReference"`
			ReferenceTitle    struct {
				Text                  string `xml:",chardata"`
				LanguageAndScriptCode string `xml:"LanguageAndScriptCode,attr"`
				TitleText             string `xml:"TitleText"`
			} `xml:"ReferenceTitle"`
			Duration                         string `xml:"Duration"`
			SoundRecordingDetailsByTerritory struct {
				Text          string `xml:",chardata"`
				TerritoryCode string `xml:"TerritoryCode"`
				Title         struct {
					Text      string `xml:",chardata"`
					TitleType string `xml:"TitleType,attr"`
					TitleText string `xml:"TitleText"`
				} `xml:"Title"`
				DisplayArtist struct {
					Text      string `xml:",chardata"`
					PartyName struct {
						Text     string `xml:",chardata"`
						FullName string `xml:"FullName"`
					} `xml:"PartyName"`
					ArtistRole string `xml:"ArtistRole"`
				} `xml:"DisplayArtist"`
				LabelName        string `xml:"LabelName"`
				RightsController struct {
					Text      string `xml:",chardata"`
					PartyName struct {
						Text     string `xml:",chardata"`
						FullName string `xml:"FullName"`
					} `xml:"PartyName"`
					PartyId              string `xml:"PartyId"`
					RightsControllerRole string `xml:"RightsControllerRole"`
					RightSharePercentage string `xml:"RightSharePercentage"`
				} `xml:"RightsController"`
				PLine struct {
					Text      string `xml:",chardata"`
					Year      string `xml:"Year"`
					PLineText string `xml:"PLineText"`
				} `xml:"PLine"`
				Genre struct {
					Text      string `xml:",chardata"`
					GenreText string `xml:"GenreText"`
				} `xml:"Genre"`
				ParentalWarningType            string `xml:"ParentalWarningType"`
				TechnicalSoundRecordingDetails struct {
					Text                              string `xml:",chardata"`
					TechnicalResourceDetailsReference string `xml:"TechnicalResourceDetailsReference"`
					File                              struct {
						Text     string `xml:",chardata"`
						FileName string `xml:"FileName"`
						FilePath string `xml:"FilePath"`
					} `xml:"File"`
				} `xml:"TechnicalSoundRecordingDetails"`
			} `xml:"SoundRecordingDetailsByTerritory"`
		} `xml:"SoundRecording"`
		Image struct {
			Text      string `xml:",chardata"`
			ImageType string `xml:"ImageType"`
			ImageId   struct {
				Text          string `xml:",chardata"`
				ProprietaryId struct {
					Text      string `xml:",chardata"`
					Namespace string `xml:"Namespace,attr"`
				} `xml:"ProprietaryId"`
			} `xml:"ImageId"`
			ResourceReference       string `xml:"ResourceReference"`
			ImageDetailsByTerritory struct {
				Text                  string `xml:",chardata"`
				TerritoryCode         string `xml:"TerritoryCode"`
				TechnicalImageDetails struct {
					Text                              string `xml:",chardata"`
					TechnicalResourceDetailsReference string `xml:"TechnicalResourceDetailsReference"`
					ImageCodecType                    string `xml:"ImageCodecType"`
					ImageHeight                       string `xml:"ImageHeight"`
					ImageWidth                        string `xml:"ImageWidth"`
					File                              struct {
						Text     string `xml:",chardata"`
						FileName string `xml:"FileName"`
						FilePath string `xml:"FilePath"`
					} `xml:"File"`
				} `xml:"TechnicalImageDetails"`
			} `xml:"ImageDetailsByTerritory"`
		} `xml:"Image"`
	} `xml:"ResourceList"`
	ReleaseList            struct {
		Text    string `xml:",chardata"`
		Release []struct {
			Text           string `xml:",chardata"`
			ReferenceTitle struct {
				Text                  string `xml:",chardata"`
				LanguageAndScriptCode string `xml:"LanguageAndScriptCode"`
				TitleText             string `xml:"TitleText"`
			} `xml:"ReferenceTitle"`
			ReleaseResourceReferenceList struct {
				Text                     string `xml:",chardata"`
				ReleaseResourceReference []struct {
					Chardata            string `xml:",chardata"`
					Text                string `xml:"text"`
					ReleaseResourceType string `xml:"ReleaseResourceType"`
				} `xml:"ReleaseResourceReference"`
			} `xml:"ReleaseResourceReferenceList"`
			ReleaseDetailsByTerritory struct {
				Text          string `xml:",chardata"`
				ResourceGroup struct {
					Text          string `xml:",chardata"`
					ResourceGroup struct {
						Text                     string `xml:",chardata"`
						ResourceGroupContentItem struct {
							Text                     string `xml:",chardata"`
							ResourceType             string `xml:"ResourceType"`
							ReleaseResourceReference struct {
								Chardata            string `xml:",chardata"`
								Text                string `xml:"text"`
								ReleaseResourceType string `xml:"ReleaseResourceType"`
							} `xml:"ReleaseResourceReference"`
							SequenceNumber string `xml:"SequenceNumber"`
						} `xml:"ResourceGroupContentItem"`
						SequenceNumber string `xml:"SequenceNumber"`
						Title          struct {
							Text      string `xml:",chardata"`
							TitleType string `xml:"TitleType"`
							TitleText string `xml:"TitleText"`
						} `xml:"Title"`
					} `xml:"ResourceGroup"`
					ResourceGroupContentItem struct {
						Text                     string `xml:",chardata"`
						ResourceType             string `xml:"ResourceType"`
						ReleaseResourceReference struct {
							Chardata            string `xml:",chardata"`
							Text                string `xml:"text"`
							ReleaseResourceType string `xml:"ReleaseResourceType"`
						} `xml:"ReleaseResourceReference"`
						SequenceNumber string `xml:"SequenceNumber"`
					} `xml:"ResourceGroupContentItem"`
				} `xml:"ResourceGroup"`
				OriginalReleaseDate string `xml:"OriginalReleaseDate"`
				TerritoryCode       string `xml:"TerritoryCode"`
				DisplayArtist       struct {
					Text      string `xml:",chardata"`
					PartyName struct {
						Text     string `xml:",chardata"`
						FullName string `xml:"FullName"`
					} `xml:"PartyName"`
					ArtistRole string `xml:"ArtistRole"`
				} `xml:"DisplayArtist"`
				LabelName string `xml:"LabelName"`
				Title     []struct {
					Text                  string `xml:",chardata"`
					TitleType             string `xml:"TitleType"`
					TitleText             string `xml:"TitleText"`
					LanguageAndScriptCode string `xml:"LanguageAndScriptCode"`
				} `xml:"Title"`
				Genre struct {
					Text      string `xml:",chardata"`
					GenreText string `xml:"GenreText"`
				} `xml:"Genre"`
				ParentalWarningType string `xml:"ParentalWarningType"`
			} `xml:"ReleaseDetailsByTerritory"`
			PLine struct {
				Text      string `xml:",chardata"`
				PLineText string `xml:"PLineText"`
				Year      string `xml:"Year"`
			} `xml:"PLine"`
			ReleaseType string `xml:"ReleaseType"`
			ReleaseId   struct {
				Text          string `xml:",chardata"`
				ICPN          string `xml:"ICPN"`
				ProprietaryId struct {
					Chardata  string `xml:",chardata"`
					Text      string `xml:"text"`
					Namespace string `xml:"Namespace"`
				} `xml:"ProprietaryId"`
				ISRC string `xml:"ISRC"`
			} `xml:"ReleaseId"`
			ReleaseReference string `xml:"ReleaseReference"`
			CLine            struct {
				Text      string `xml:",chardata"`
				CLineText string `xml:"CLineText"`
				Year      string `xml:"Year"`
			} `xml:"CLine"`
		} `xml:"Release"`
	} `xml:"ReleaseList"`
	DealList struct {
		Text        string `xml:",chardata"`
		ReleaseDeal struct {
			Text                 string `xml:",chardata"`
			DealReleaseReference string `xml:"DealReleaseReference"`
			Deal                 struct {
				Text          string `xml:",chardata"`
				DealReference string `xml:"DealReference"`
			} `xml:"Deal"`
		} `xml:"ReleaseDeal"`
	} `xml:"DealList"`
}

/*
type NewReleaseMessage struct {
  MessageId                           string `xml:"MessageHeader>MessageId"`
  SenderPartyId                       string `xml:"MessageHeader>MessageSender>PartyId"`
  PartnerName                         string `xml:"MessageHeader>MessageSender>PartyName>FullName"`
  RecipeintPartyId                    string `xml:"MessageHeader>MessageRecipient>PartyId"`
  MessageCreatedDateTime              string `xml:"MessageHeader>MessageCreatedDateTime"`
  MessageControlType                  string `xml:"MessageHeader>MessageControlType"`

	//OriginalReleaseDate                 string `xml:"MessageHeader>MessageDetails>ReleaseDate"`
	//ReleaseTitleText                    string `xml:"MessageHeader>MessageDetails>SongTitle"`
	//TitleTextArtistName                 string `xml:"MessageHeader>MessageDetails>ArtistName"`

	UpdateIndicator                     string `xml:"UpdateIndicator"`
	SoundRecordingType                  string `xml:"ResourceList>SoundRecording>SoundRecordingType"`
	Isrc						 										string `xml:"ResourceList>SoundRecording>SoundRecordingId>ISRC"`

  ResourceReference                   string `xml:"ResourceList>SoundRecording>ResourceReference"`
	ReleaseTitleText										string `xml:"ResourceList>SoundRecording>ReferenceTitle"`

  Duration                            string `xml:"ResourceList>SoundRecording>Duration"`
  TerritoryCode                       string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>TerritoryCode"`
  //TitleTextArtistName                 string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>Title" TitleType='DisplayTitle'">TitleText"`
  TitleTextArtistName               	string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>DisplayArtist>PartyName>FullName"`
  ArtistRole                          string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>DisplayArtist>ArtistRole"`
  LabelName                           string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>LabelName"`
  RightsControllerFullName            string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>RightsController>PartyName>FullName"`
  RightsControllerPartyId             string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>RightsController>PartyId"`
  RightsControllerRole                string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>RightsController>RightsControllerRole"`
  RightSharePercentage                string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>RightsController>RightSharePercentage"`
  PLineYear                           string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>PLine>Year"`
  PLineText                           string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>PLine>PLineText"`
	Genre                               string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>Genre>GenreText"`
	ParentalWarningType			            string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>ParentalWarningType"`
	TechnicalResourceDetailsReference   string `xml:"ResourceList>SoundRecording>SoundRecordingDetailsByTerritory>TechnicalSoundRecordingDetails>TechnicalSoundRecordingDetailsReference"`
	ReleaseIsrc                         string `xml:"ReleaseList>Release>ReleaseId>ISRC"`
  DealList                            string `xml:"DealList"`
}
*/


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
	router.HandleFunc("/createClaimsEvent", CreateEventClaimsTool).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend-light-bootstrap-dashboard-react/build/static/"))))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend-light-bootstrap-dashboard-react/build/index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
        // The file token.json stores the user's access and refresh tokens, and is
        // created automatically when the authorization flow completes for the first
        // time.
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        defer f.Close()
        json.NewEncoder(f).Encode(token)
}

/*
func CreateEvent() {
}
*/

func CreateEventClaimsTool(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := getClient(config)

        srv, err := calendar.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Calendar client: %v", err)
        }

				event := &calendar.Event{
				  Summary: "Google I/O 2015",
				  Location: "800 Howard St., San Francisco, CA 94103",
				  Description: "A chance to hear more about Google's developer products.",
				  Start: &calendar.EventDateTime{
				    DateTime: "2020-10-28T09:00:00-07:00",
				    TimeZone: "America/Los_Angeles",
				  },
				  End: &calendar.EventDateTime{
				    DateTime: "2020-10-28T17:00:00-07:00",
				    TimeZone: "America/Los_Angeles",
				  },
				  Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
				  Attendees: []*calendar.EventAttendee{
				    &calendar.EventAttendee{Email:"lpage@example.com"},
				    &calendar.EventAttendee{Email:"sbrin@example.com"},
				  },
				}

				calendarId := "primary"
				event, err = srv.Events.Insert(calendarId, event).Do()
				if err != nil {
				  log.Fatalf("Unable to create event. %v\n", err)
				}
				fmt.Printf("Event created: %s\n", event.HtmlLink)



				/*
        t := time.Now().Format(time.RFC3339)
        events, err := srv.Events.List("primary").ShowDeleted(false).
                SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
        if err != nil {
                log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
        }
        fmt.Println("Upcoming events:")
        if len(events.Items) == 0 {
                fmt.Println("No upcoming events found.")
        } else {
                for _, item := range events.Items {
                        date := item.Start.DateTime
                        if date == "" {
                                date = item.Start.Date
                        }
                        fmt.Printf("%v (%v)\n", item.Summary, date)
                }
        }
				*/
}

func main() {
	fmt.Println("Youtube Ingestion Dashboard Application")
	handleRequests()
}
