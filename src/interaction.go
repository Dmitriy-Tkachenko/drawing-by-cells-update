package main

import (
	"encoding/json"
	"mime"
	"net/http"
)

const Version = "1.0"

type RequestType string

const (
	SimpleUtterance RequestType = "SimpleUtterance"
	ButtonPressed   RequestType = "ButtonPressed"
)

type NLU struct {
	Tokens   []string `json:"tokens"`
	Entities []string `json:"entities"`
}

const (
	OnStart = ""
	OnInterrupt = "on_interrupt"
)

type RequestIn struct {
	Command string `json:"command"`
	OriginalUtterance string `json:"original_utterance"`
	Type RequestType `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
	NLU NLU `json:"nlu"`
}

type Screen struct{}

type Interfaces struct {
	Screen *Screen `json:"screen,omitempty"`
}

func (i *Interfaces) IsScreen() bool {
	return i.Screen != nil
}

type Meta struct {
	ClientID string `json:"client_id"`
	Locale string `json:"locale"`
	Timezone string `json:"timezone"`
	Interfaces Interfaces `json:"interfaces"`
	CityRu string `json:"_city_ru,omitempty"`
}

type Session struct {
	SessionID string `json:"session_id"`
	UserID string `json:"user_id"`
	SkillID string `json:"skill_id"`
	New bool `json:"new"`
	MessageID int `json:"message_id"`
}

type Request struct {
	Meta Meta `json:"meta"`
	Request RequestIn `json:"request"`
	Session Session `json:"session"`
	Version string `json:"version"`
}

type BindingType string

const (
	BindingTypeSuggest BindingType = "suggest"
)

type DefaultPayload struct {
	BindingType    BindingType `json:"binding_type"`
	Index          int         `json:"index"`
	TargetPhraseID string      `json:"target_phrase_id"`
}

type Button struct {
	Title   string      `json:"title"`
	Payload interface{} `json:"payload,omitempty"`
	URL     string      `json:"url,omitempty"`
}

type CardType string

const (
	BigImage CardType = "BigImage"
	ItemsList CardType = "ItemsList"
	Link CardType = "Link"
)

type CardItem struct {
	ImageID int `json:"image_id"`
}
type Card struct {
	Type CardType `json:"type"`
	Title string `json:"title,omitempty"`
	Text string `json:"text,omitempty"`
	Description string `json:"description,omitempty"`
	ImageID int `json:"image_id,omitempty"`
	Items []CardItem `json:"items,omitempty"`
	URL string `json:"url,omitempty"`
}

func NewBigImage(imageID int) *Card {
	return &Card{
		Type:        BigImage,
		ImageID:     imageID,
	}
}

func NewItemsList(items ...CardItem) *Card {
	return &Card{
		Type:  ItemsList,
		Items: items,
	}
}

func NewImageList(imageIDs ...int) *Card {
	items := make([]CardItem, len(imageIDs))

	for i := 0; i < len(imageIDs); i++ {
		items[i].ImageID = imageIDs[i]
	}

	return NewItemsList(items...)
}

func NewLink(url, title, text string, imageID int) *Card {
	return &Card{
		Type:    Link,
		URL:     url,
		Title:   title,
		Text:    text,
		ImageID: imageID,
	}
}

type Response struct {
	Text string `json:"text"`
	TTS string `json:"tts,omitempty"`
	TTSType string `json:"tts_type"`
	SSML string `json:"ssml"`
	Buttons []Button `json:"buttons,omitempty"`
	EndSession bool `json:"end_session"`
	Card *Card `json:"card,omitempty"`
}

func (r *Response) AddURL(title string, url string) {
	if r.Buttons == nil {
		r.Buttons = make([]Button, 0)
	}

	r.Buttons = append(r.Buttons, Button{
		Title: title,
		URL:   url,
	})
}

func (r *Response) AddButton(title string, payload interface{}) {
	if r.Buttons == nil {
		r.Buttons = make([]Button, 0)
	}

	r.Buttons = append(r.Buttons, Button{
		Title:   title,
		Payload: payload,
	})
}

type responseSession struct {
	SessionID string `json:"session_id"`
	MessageID int    `json:"message_id"`
	UserID    string `json:"user_id"`
}

type response struct {
	Response Response        `json:"response"`
	Session  responseSession `json:"session"`
	Version  string          `json:"version"`
}

type Webhook struct {
	event func(r Request) Response
}

func NewWebhook() *Webhook {
	return &Webhook{}
}

func (wh *Webhook) OnEvent(f func(r Request) Response) {
	wh.event = f
}

func (wh *Webhook) HandleFunc(w http.ResponseWriter, r *http.Request) {
	mediatype, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if mediatype != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	resp := wh.event(req)

	fullResponse := response{
		Response: resp,
		Session: responseSession{
			SessionID: req.Session.SessionID,
			MessageID: req.Session.MessageID,
			UserID:    req.Session.UserID,
		},
		Version: Version,
	}
	w.Header().Add("Content-Type", "application/json; encoding=utf-8")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(fullResponse)
}
