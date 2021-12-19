package main

import (
	log "github.com/sirupsen/logrus"
)

func main(){
	err := Irc.Init()	
	if err != nil {
		log.Fatal(err)
	}

	go func (){
		for{
			data:= <-Irc.Messages_out
			log.Infof("%s: %s",data.User.Name,data.Message)
		}
	}()

	err = Irc.Connect()	
	if err != nil {
		log.Fatal(err)
	}

}
