package models

import (
	"time"
)

type NewReleaseMessageOld struct {
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
