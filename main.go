package main

import (
	log "github.com/sirupsen/logrus"
)

func main(){
	err := Irc.Init()	
	if err != nil {
		log.Fatal(err)
	}

	err = Irc.Connect()	
	if err != nil {
		log.Fatal(err)
	}


}
