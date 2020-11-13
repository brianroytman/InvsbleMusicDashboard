/*!

=========================================================
* Light Bootstrap Dashboard React - v1.3.0
=========================================================

* Product Page: https://www.creative-tim.com/product/light-bootstrap-dashboard-react
* Copyright 2019 Creative Tim (https://www.creative-tim.com)
* Licensed under MIT (https://github.com/creativetimofficial/light-bootstrap-dashboard-react/blob/master/LICENSE.md)

* Coded by Creative Tim

=========================================================

* The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

*/
import React, { Component } from "react";
import {
  Grid,
  Row,
  Col,
  FormGroup,
  ControlLabel,
  FormControl
} from "react-bootstrap";

import { TextField } from '@material-ui/core';

import { makeStyles } from '@material-ui/core/styles';
import FilledInput from '@material-ui/core/FilledInput';
import FormHelperText from '@material-ui/core/FormHelperText';
//import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import { Card } from "components/Card/Card.jsx";
import { FormInputs } from "components/FormInputs/FormInputs.jsx";
import { UserCard } from "components/UserCard/UserCard.jsx";
import Radio from  "components/CustomRadio/CustomRadio.jsx";
//import Button from "components/CustomButton/CustomButton.jsx";
import Select from "components/Select/Select.jsx";
import Button from "components/Button/Button.jsx";
import Input from "components/Input/Input.jsx";

import avatar from "assets/img/faces/face-3.jpg";
import { saveAs } from 'file-saver';
import FileSaver from 'file-saver';

class IngestionTool extends Component {
  constructor(props) {
    super(props);
    this.state = {
        MessageDetails: {
          LanguageAndScriptCode: "en",
          MessageSchemaVersionId: "ern/382",
          Ern: "http://ddex.net/xml/ern/382",
          SchemaLocation: "http://ddex.net/xml/ern/382 http://ddex.net/xml/ern/382/release-notification.xsd",
          MessageHeader: {
            PartnerName: "",
            ParentalWarningType: "",
            ReleaseType: "",
            ReleaseTitleText: "",
            TitleTextArtistName: "",
            Genre: "",
            Isrc: "",
            MessageCreatedDateTime: Date().toLocaleString(),
            MessageSender: {
              PartyName: {
                FullName: "THREEHUNDRED",
              },
              PartyId: "PADPIDA2014082004R",
            },
            MessageControlType: "LiveMessage",
            MessageRecipient: {
              PartyName: {
                FullName: "Youtube",
              },
              PartyId: "PADPIDA2014082004R",
            },
            MessageId: "123456QWERTY",
          },
          UpdateIndicator: "OriginalMessage",
          ResourceList: {
              SoundRecording: {
                SoundRecordingType: "MusicalWorkSoundRecording",
                SoundRecordingId: {
                   Isrc: "QMCE32000050",
                   ProprietaryId: [
                      {
                         Namespace: "YOUTUBE:AT_CUSTOM_ID",
                         text: "THREEHUNDRED-181"
                      },
                      {
                         Namespace: "YOUTUBE:SR_CUSTOM_ID",
                         text: "THREEHUNDRED-181"
                      },
                      {
                         Namespace: "DPID:PADPIDA2014040303H",
                         text: "THREEHUNDRED"
                      },
                   ],
                 },
                 ResourceReference: "A1",
                 ReferenceTitle: {
                   TitleText: "Ramona Flowers",
                   LanguageAndScriptCode: "en"
                 },
                 Duration: "PT00H01M42S",
                 SoundRecordingDetailsByTerritory: {
                   TerritoryCode: "Worldwide",
                   Title: {
                     TitleText: "Ramona Flowers",
                     TitleType: "DisplayTitle",
                   },
                   DisplayArtist: {
                     PartyName: {
                       FullName: "Haroinfather",
                     },
                     ArtistRole: "MainArtist",
                   },
                   LabelName: "Crowned LLC 300 Entertainment",
                   RightsController: {
                     PartyName: {
                       "FullName": "INDMUSIC"
                     },
                     PartyId: "PADPIDA2014040303H",
                     RightsControllerRole: "RightsController",
                     RightSharePercentage: "100.00"
                   },
                   PLine: {
                     Year: "2020",
                     PLineText: "Crowned LLC 300 Entertainment"
                   },
                   Genre: {
                     GenreText: "Hip Hop"
                   },
                   ParentalWarningType: "Explicit",
                   TechnicalSoundRecordingDetails: {
                     TechnicalResourceDetailsReference: "T1",
                     File:  {
                       FileName: "Haroinfather - RAMONA FLOWERS (MAIN)_JL MASTER 2.wav",
                       FilePath: "/",
                       HashSum: {
                         HashSum: "a01cbacff993beab6769f817cc46786b",
                         HashSumAlgorithmType: "MD5"
                       }
                     }
                   }
                 }
              },
              Image: {
                ImageType: "FrontCoverImage",
                ImageId: {
                  ProprietaryId: [
                    {
                      Namespace: "PADPIDA2014040303H",
                      text: "THREEHUNDRED"
                    }
                  ]
                },
                ResourceReference: "A2",
                ImageDetailsByTerritory: {
                  TerritoryCode: "Worldwide",
                  TechnicalImageDetails: {
                    TechnicalResourceDetailsReference: "T2",
                    ImageCodecType: "JPEG",
                    ImageHeight: "2000",
                    ImageWidth: "2000",
                    File: {
                      FileName: "haroinfather - ramona flowers ARTWORK v2.jpg",
                      FilePath: "/",
                      HashSum: {
                        HashSum: "101448ba335e56f53be92d662d6f4b08",
                        HashSumAlgorithmType: "MD5"
                     }
                   }
                  }
                }
              }
            },
            ReleaseList: {
              Release: [
                {
                  ReleaseId: {
                    Icpn: "810043680172",
                    ProprietaryId: [
                      {
                        Namespace: "PADPIDA2014040303H",
                        Text: "THREEHUNDRED"
                      }
                    ],
                  },
                  ReleaseReference: "R0",
                  ReferenceTitle: {
                    TitleText: "Ramona Flowers",
                    LanguageAndScriptCode: "en"
                  },
                  ReleaseResourceReferenceList: {
                    ReleaseResourceReference: [
                      {
                        ReleaseResourceType: "PrimaryResource",
                        Text: "A1"
                      },
                      {
                        ReleaseResourceType: "SecondaryResource",
                        Text: "A2"
                      }
                    ]
                  },
                  ReleaseType: "Album",
                  ReleaseDetailsByTerritory: {
                    TerritoryCode: "Worldwide",
                    LabelName: "Crowned LLC 300 Entertainment",
                    Title: [
                      {
                        TitleText: "Ramona Flowers",
                        TitleType: "DisplayTitle"
                      },
                      {
                        TitleText: "Ramona Flowers",
                        LanguageAndScriptCode: "en",
                        TitleType: "FormalTitle"
                      }
                   ],
                  DisplayArtist: {
                    PartyName: {
                      FullName: "Haroinfather"
                    },
                    ArtistRole: "MainArtist",
                  },
                  ParentalWarningType: "Explicit",
                  ResourceGroup: {
                    ResourceGroup: {
                      Title: [
                        {
                          TitleText: "Ramona Flowers",
                          TitleType: "GroupingTitle"
                        }
                      ],
                      SequenceNumber: "1",
                      ResourceGroupContentItem: {
                        SequenceNumber: "1",
                        ResourceType: "SoundRecording",
                        ReleaseResourceReference: [
                          {
                            ReleaseResourceType: "SecondaryResource",
                            text: "A2"
                          }
                        ]
                      }
                    }
                  },
                  Genre: {
                    GenreText: "Hip Hop"
                  },
                  OriginalReleaseDate: "2020-02-20"
                },
                PLine: {
                  Year: "2020",
                  PLineText: "Crownded LLC 300 Entertainment"
                }
              },
              {
                ReleaseId: {
                  Isrc: "QMCE32000050",
                  ProprietaryId: [
                    {
                      Namespace: "PADPIDA2014040303H",
                      Text: "THREEHUNDRED"
                    }
                  ]
                },
                ReleaseReference: "R1",
                ReferenceTitle: {
                  TitleText: "Ramona Flowers",
                  LanguageAndScriptCode: "en"
                },
                ReleaseResourceReferenceList: {
                  ReleaseResourceReference: [
                    {
                      ReleaseResourceType: "PrimaryResource",
                      Text: "A1"
                    }
                  ]
                },
                ReleaseType: "TrackRelease",
                ReleaseDetailsByTerritory: {
                  TerritoryCode: "Worldwide",
                  LabelName: "Crowned LLC 300 Entertainment",
                  Title: [
                    {
                      TitleText: "Ramona Flowers"
                    }
                  ],
                  DisplayArtist: {
                    PartyName: {
                      FullName: "Haroinfather"
                    },
                    ArtistRole: "MainArtist"
                  },
                  ParentalWarningType: "Explicit",
                  Genre: {
                    GenreText: "Hip Hop"
                  },
                  OriginalReleaseDate: "2020-02-20"
                },
                PLine: {
                  Year: "2020",
                  PLineText: "Crowned LLC 300 Entertainment"
                },
                CLine: {
                  Year: "2020",
                  CLineText: "Crowned LLC 300 Entertainment"
                }
              }
            ]
          },
          DealList: {
            ReleaseDeal: {
              DealReleaseReference: "R1",
              Deal: {
                DealReference: "YT_MATCH_POLICY:Monetize in all countries"
              }
            }
          }
        },
        messageRecipientNameOptions: ["Youtube", "YouTube_ContentID"],
        partnerOptions: ["300 Entertainment", "The Royt Sound"],
        parentalWarningTypeOptions: ["Explicit", "Not Explicit"],
        releaseTypeOptions: ["Album", "EP", "Single"],

        /*
        MessageDetails: {
          messageId: "1-a",
          messageCreatedDateTime: Date().toLocaleString(),
          messageControlType: "",
          senderPartyId: "",
          isrc: "",
          titleTextArtistName: "",
          //duration: "",
          fullName: "",
          labelName: "",
          pLineText: "",
          pLineYear: "2020",
          genre: "",
          parentalWarningType: "",
          release: {
            releaseType: releaseType,
          },
          fileNameData: "",
          fileNameImage: "",
          icpn: "",
          propertyIdTitleText: "",
          resourceReleaseType: "",
          resourceReference: "A1",
          territoryCode: "",
          releaseLabelName: "",
          releaseTitleText: "",
          releaseIsrc: "",
          languageTitleText: "",
          languageAndScriptCode: "en",
          displayArtistName: "",
          displayParentalWarningType: "",
          originalReleaseDate: "",
        },
        partnerOptions: ["300 Entertainment", "The Royt Sound"],
        parentalWarningTypeOptions: ["Explicit", "Not Explicit"],
        releaseTypeOptions: ["Album", "EP", "Single"],
        territoryCodeOptions: ["Worldwide", "Other"],
        */

        /*
        MessageDetails: {
          messageId: "",
          messageCreatedDateTime: "",
          titleTextArtistName: "",
          //duration: "",
          fullName: "",
          labelName: "",
          year: "",
          pLineText: "",
          genre: "",
          parentalWarningType: "",
          releaseType: "",
          fileNameData: "",
          fileNameImage: "",
          icpn: "",
          propertyIdTitleText: "",
          resourceReleaseType: "",
          territoryCode: "",
          releaseLabelName: "",
          releaseTitleText: "",
          languageTitleText: "",
          displayArtistName: "",
          displayParentalWarningType: "",
          originalReleaseDate: "",
        },
        */

        //partnerOptions: ["300 Entertainment", "The Royt Sound"],
        //parentalWarningTypeOptions: ["Explicit", "Not Explicit"],
        //releaseTypeOptions: ["Album", "EP", "Single"],
        //territoryCodeOptions: ["Worldwide", "Other"],

    };
    this.handleReleaseDate = this.handleReleaseDate.bind(this);
    this.handleFormSubmit = this.handleFormSubmit.bind(this);
    this.handleClearForm = this.handleClearForm.bind(this);
    this.handleInput = this.handleInput.bind(this);
  }

  handleReleaseDate(e) {
    let value = e.target.value;
    this.setState(
      prevState => ({
        MessageDetails: {
          ...prevState.MessageDetails,
          originalReleaseDate: value
        }
      }),
      () => console.log(this.state.MessageDetails)
    );
  }

  handleInput(e) {
    let value = e.target.value;
    let name = e.target.name;
    this.setState(
      prevState => ({
        MessageDetails: {
          ...prevState.MessageDetails,
          [name]: value
        }
      }),
      () => console.log(this.state.MessageDetails)
    );
  }

  handleFormSubmit(e) {
    e.preventDefault();
    let releaseData = this.state.MessageDetails;

    fetch("/uploadLive", {
      method: "POST",
      //mode: "no-cors",
      body: JSON.stringify(releaseData),
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      }
    }).then(function(response) {
          return response.blob();
    }).then(function(blob) {
          FileSaver.saveAs(blob, 'xml_output_300.xml');
    });
     /*.then(response => {
      response.json().then(data => {
        console.log("Successful" + data);
      });
    }).catch(error => {
        alert(error.message);
    });
  }
  */
  }

  handleClearForm(e) {
    e.preventDefault();
    this.setState({
      MessageDetails: {
        PartnerName: "",
        ParentalWarningType: "",
        ReleaseType: "",
        ReleaseTitleText: "",
        TitleTextArtistName: "",
        Genre: "",
        Isrc: "",
        MessageHeader: {
          MessageCreatedDateTime: Date().toLocaleString(),
        },
        MessageSender: {
          PartyName: {
            FullName: "THREEHUNDRED",
          },
          PartyId: "PADPIDA2014082004R",
        },
        MessageControlType: "LiveMessage",
        MessageRecipient: {
          PartyName: {
            FullName: "Youtube",
          },
          PartyId: "PADPIDA2014082004R",
        },
        MessageId: "123456QWERTY",
      }

      /*
        MessageDetails: {
          messageId: "1-a",
          messageCreatedDateTime: Date().toLocaleString(),
          messageControlType: "LiveMessage",
          UpdateIndicator: "OriginalMessage",
          senderPartyId: "PADPIDA2014082004R",
          isrc: "",
          titleTextArtistName: "",
          //duration: "",
          fullName: "",
          labelName: "",
          pLineText: "",
          pLineYear: "",
          genre: "",
          parentalWarningType: "",
          releaseType: "",
          fileNameData: "",
          fileNameImage: "",
          icpn: "",
          propertyIdTitleText: "",
          resourceReleaseType: "",
          resourceReference: "A1",
          territoryCode: "",
          releaseLabelName: "",
          releaseIsrc: "",
          releaseTitleText: "",
          languageTitleText: "",
          languageAndScriptCode: "",
          displayArtistName: "",
          displayParentalWarningType: "",
          originalReleaseDate: "",
        }
        */
        /*
        MessageDetails: {
          partnerName: "",
          releaseDate: "",
          songTitle: "",
          artistName: "",
          genre:"",
          isrc: "",
        }
        */
    });
  }

  render() {
    //const []

    return (
      <form className="container-fluid" onSubmit={this.handleFormSubmit}>
        <Select
          title={"Partner Name"}
          name={"partnerName"}
          options={this.state.partnerOptions}
          value={this.state.MessageDetails.partnerName}
          placeholder={"Select Partner"}
          handleChange={this.handleInput}
        />{" "}
        {/* Name of the partner organization */}
        <Select
          title={"Parental Warning Type"}
          name={"parentalWarningType"}
          options={this.state.parentalWarningTypeOptions}
          value={this.state.MessageDetails.parentalWarningType}
          placeholder={"Select Parental Warning Type"}
          handleChange={this.handleInput}
        />{" "}
        {/* Name of the partner organization */}
        <Select
          title={"Release Type"}
          name={"releaseType"}
          options={this.state.releaseTypeOptions}
          value={this.state.MessageDetails.releaseType}
          placeholder={"Select Release Type"}
          handleChange={this.handleInput}
          //handleChange={setReleaseType(value... )}
        />{" "}
        {/* Name of the partner organization */}
        <Row>
          <Col xs="6">
            <Input
              inputType={"text"}
              title={"Song Title"}
              name={"releaseTitleText"}
              value={this.state.MessageDetails.releaseTitleText}
              placeholder={"Enter the song title"}
              handleChange={this.handleInput}
            />{" "}
            {/* Name of the song */}
          </Col>
        <Col xs="6">
          <Input
            inputType={"text"}
            title={"Artist Name"}
            name={"titleTextArtistName"}
            value={this.state.MessageDetails.titleTextArtistName}
            placeholder={"Enter the artist name"}
            handleChange={this.handleInput}
          />{" "}
          {/* Name of the artist */}
          </Col>
        </Row>
        <Row>
          <Col xs="6">
            <Input
              inputType={"text"}
              title={"Genre"}
              name={"genre"}
              value={this.state.MessageDetails.genre}
              placeholder={"Enter the genre"}
              handleChange={this.handleInput}
            />{" "}
            {/* Name of the genre */}
          </Col>
          <Col xs="6">
            <Input
              inputType={"text"}
              title={"ISRC"}
              name={"isrc"}
              value={this.state.MessageDetails.isrc}
              placeholder={"Enter the ISRC code"}
              handleChange={this.handleInput}
            />{" "}
            {/* ISRC code for the song */}
          </Col>
        </Row>
        <Button
          action={this.handleFormSubmit}
          type={"primary"}
          title={"Submit"}
          style={buttonStyle}
        />{" "}
        {/*Submit */}
        <Button
          action={this.handleClearForm}
          type={"secondary"}
          title={"Clear"}
          style={buttonStyle}
        />{" "}
        {/* Clear the form */}
      </form>
    );
  }
}

const buttonStyle = {
  margin: "10px 10px 10px 10px"
};

export default IngestionTool;
