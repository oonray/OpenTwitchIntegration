package main

import (
	"github.com/gorilla/mux"
	twitch "github.com/gempir/go-twitch-irc/v2"
)

var web WebServer

type Emote struct {
	
}

type Message struct {
	Name string `json:"name"`
	Message string `json:"message"`
	Emotes []Emote `json:"emotes"`
	Bits	int `json:"bits"`
}

type WebServer struct {
	router *mux.Router
	messages []twitch.PrivateMessage
}

func (w *WebServer) GetData(){
	for{
		data:= <-Irc.Messages_out
		w.messages = append(w.messages,data)
	}
}

func (w *WebServer) Init(){
	w.router = mux.NewRouter()
	go w.GetData()
}
