package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/mail"
)

type BodyType string

const (
	BodyTypePlain BodyType = "plain"
	BodyTypeHtml  BodyType = "html"
)

type Email struct {
	From    *mail.Address
	To      string
	Subject string
	Body    string
	Type    BodyType //default is plain text
}

//Email API for email.cologic.ai

type Client struct {
	//Base API URL
	Url string `json:"url" yaml:"url"`
	//Token
	Token string `json:"token" yaml:"token"`
}

func New(token, url string) *Client {
	return &Client{
		Url:   url,
		Token: token,
	}
}

func (e *Client) SendEmail(email Email) error {
	return e.Send(email)
}

func (e *Client) Send(email Email) error {
	//Create the POST request to URL
	req, err := http.NewRequest("POST", e.Url+"/v1/email", nil)
	if err != nil {
		return err
	}
	//Set the API Key
	req.Header.Set("X-Api-Key", e.Token)
	//Set the Content-Type
	req.Header.Set("Content-Type", "application/json")
	//unmarshal the email to json
	body, err := json.Marshal(email)
	if err != nil {
		return err
	}
	//Set the body of the request
	req.Body = io.NopCloser(bytes.NewReader(body))
	//Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	//Check the status code
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}
