// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Connection interface {
	IsConnection()
	GetEdge() Edge
}

type Edge interface {
	IsEdge()
	GetCursor() string
	GetNode() Node
}

type Node interface {
	IsNode()
	GetID() string
}

type Post interface {
	IsNode()
	IsPost()
	GetID() string
	GetTitle() string
	GetBody() string
}

type FacebookPost struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	FacebookID string `json:"facebookId"`
}

func (FacebookPost) IsPost()               {}
func (this FacebookPost) GetID() string    { return this.ID }
func (this FacebookPost) GetTitle() string { return this.Title }
func (this FacebookPost) GetBody() string  { return this.Body }

func (FacebookPost) IsNode() {}

type InstagramPost struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	InstagramID string `json:"instagramId"`
}

func (InstagramPost) IsPost()               {}
func (this InstagramPost) GetID() string    { return this.ID }
func (this InstagramPost) GetTitle() string { return this.Title }
func (this InstagramPost) GetBody() string  { return this.Body }

func (InstagramPost) IsNode() {}

type PostConnection struct {
	Edge *PostEdge `json:"edge"`
}

func (PostConnection) IsConnection()      {}
func (this PostConnection) GetEdge() Edge { return *this.Edge }

type PostEdge struct {
	Cursor string `json:"cursor"`
	Node   Post   `json:"node"`
}

func (PostEdge) IsEdge()                {}
func (this PostEdge) GetCursor() string { return this.Cursor }
func (this PostEdge) GetNode() Node     { return this.Node }

type TweetPost struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	TweetID string `json:"tweetId"`
}

func (TweetPost) IsPost()               {}
func (this TweetPost) GetID() string    { return this.ID }
func (this TweetPost) GetTitle() string { return this.Title }
func (this TweetPost) GetBody() string  { return this.Body }

func (TweetPost) IsNode() {}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Posts []Post `json:"posts"`
}

type Device string

const (
	DeviceIphone Device = "IPHONE"
)

var AllDevice = []Device{
	DeviceIphone,
}

func (e Device) IsValid() bool {
	switch e {
	case DeviceIphone:
		return true
	}
	return false
}

func (e Device) String() string {
	return string(e)
}

func (e *Device) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Device(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Device", str)
	}
	return nil
}

func (e Device) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}