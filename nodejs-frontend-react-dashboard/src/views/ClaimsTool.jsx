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
  FormControl,
  Form
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



class ClaimsTool extends Component {
  constructor(props) {
    super(props);

    this.state = {
        ClaimDetails: {
          claimLength: "",
          artistName: "",
          songTitle: "",
          claimUrl: "",
          emailThread: ""
        },
        claimLengthOptions: ["30 Days"],
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
        ClaimDetails: {
          ...prevState.ClaimDetails,
          releaseDate: value
        }
      }),
      () => console.log(this.state.ClaimDetailsDetails)
    );
  }

  handleInput(e) {
    let value = e.target.value;
    let name = e.target.name;
    this.setState(
      prevState => ({
        ClaimDetails: {
          ...prevState.ClaimDetails,
          [name]: value
        }
      }),
      () => console.log(this.state.ClaimDetails)
    );
  }

  handleFormSubmit(e) {
    e.preventDefault();
    let claimData = this.state.ClaimDetails;

    //create API endpoint "createClaimsEvent" on the Go side to POST to

    fetch("/createClaimsEvent", {
      method: "POST",
      //mode: "no-cors",
      body: JSON.stringify(claimData),
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      }
    }).then(function(response) {
          return response.blob();
    }).then(function(blob) {
          FileSaver.saveAs(blob, 'claim_details_300.xml');
    });

  }

  handleClearForm(e) {
    e.preventDefault();
    this.setState({
        ClaimDetails: {
          claimLength: "",
          artistName: "",
          songTitle: "",
          claimUrl: "",
          emailThread: ""
        }
    });
  }

  render() {
    return (
      <form className="container-fluid" onSubmit={this.handleFormSubmit}>
        <Select
          title={"Claim Length (Days)"}
          name={"claimLength"}
          options={this.state.claimLengthOptions}
          value={this.state.ClaimDetails.claimLength}
          placeholder={"Select Claim Length"}
          handleChange={this.handleInput}
        />{" "}
        {/* # of days for the claim */}
        <Row>
          <Col xs="6">
            <Input
              inputType={"text"}
              title={"Artist Name"}
              name={"artistName"}
              value={this.state.ClaimDetails.artistName}
              placeholder={"Enter the artist name"}
              handleChange={this.handleInput}
            />{" "}
            {/* Name of the artist */}
          </Col>
          <Col xs="6">
            <Input
              inputType={"text"}
              title={"Song Title"}
              name={"songTitle"}
              value={this.state.ClaimDetails.songTitle}
              placeholder={"Enter the song title"}
              handleChange={this.handleInput}
            />{" "}
            {/* Name of the song */}
          </Col>
        </Row>
        <Row>
          <Col xs="6">
            <Input
              inputType={"text"}
              title={"Claim URL"}
              name={"claimUrl"}
              value={this.state.ClaimDetails.claimUrl}
              placeholder={"Enter the claim URL"}
              handleChange={this.handleInput}
            />{" "}
            {/* Url to send the claim */}
          </Col>
          <Col xs="6">
            <Input
              inputType={"text"}
              title={"Email Thread"}
              name={"emailThread"}
              value={this.state.ClaimDetails.emailThread}
              placeholder={"Enter the email thread"}
              handleChange={this.handleInput}
            />{" "}
            {/* Email thread to send the notification of claim */}
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

export default ClaimsTool;
