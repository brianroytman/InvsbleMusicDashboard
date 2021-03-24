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
          partnerName: "",
          releaseDate: "",
          songTitle: "",
          artistName: "",
          genre:"",
          isrc: ""
        },
        partnerOptions: ["300 Entertainment", "The Royt Sound"],
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
          releaseDate: value
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
          partnerName: "",
          releaseDate: "",
          songTitle: "",
          artistName: "",
          genre:"",
          isrc: ""
        }
    });
  }

  render() {
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
        <Input
          inputType={"date"}
          name={"releaseDate"}
          title={"Release Date"}
          value={this.state.MessageDetails.releaseDate}
          placeholder={"Enter the release date"}
          handleChange={this.handleReleaseDate}
        />{" "}
        {/* Release Date */}
        <Row>
          <Col xs="6">
            <Input
              inputType={"text"}
              title={"Song Title"}
              name={"songTitle"}
              value={this.state.MessageDetails.songTitle}
              placeholder={"Enter the song title"}
              handleChange={this.handleInput}
            />{" "}
            {/* Name of the song */}
          </Col>
        <Col xs="6">
          <Input
            inputType={"text"}
            title={"Artist Name"}
            name={"artistName"}
            value={this.state.MessageDetails.artistName}
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
