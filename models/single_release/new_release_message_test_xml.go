package models

type NewReleaseMessageTest struct {
	MessageId                           string `xml:"MessageHeader>MessageId"`
	SenderPartyId                       string `xml:"MessageHeader>MessageSender>PartyId"`
	PartnerName                         string `xml:"MessageHeader>MessageSender>PartnerName"`
	FullName                            string `xml:"MessageHeader>MessageSender>PartyName>FullName"`
	RecipeintPartyId                    string `xml:"MessageHeader>MessageRecipient>PartyId"`
	MessageCreatedDateTime              string `xml:"MessageHeader>MessageCreatedDateTime"`
	MessageControlType                  string `xml:"MessageHeader>MessageControlType"`

	ReleaseDate                         string `xml:"MessageHeader>MessageDetails>ReleaseDate"`
	SongTitle                           string `xml:"MessageHeader>MessageDetails>SongTitle"`
	ArtistName                          string `xml:"MessageHeader>MessageDetails>ArtistName"`

	UpdateIndicator                     string `xml:"UpdateIndicator"`
	SoundRecordingType                  string `xml:"ResourceList>SoundRecording>SoundRecordingType"`

	SoundRecordingId   struct {
					Text          string `xml:",chardata"`
					Isrc          string `xml:"ISRC"`
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
	Isrc                                string `xml:"ReleaseList>Release>ReleaseId>ISRC"`
  	DealList                            string `xml:"DealList"`
}
