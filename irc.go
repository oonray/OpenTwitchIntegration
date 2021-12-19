package main

import (
	twitch "github.com/gempir/go-twitch-irc/v2"
	toml "github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

var (
	Conffile string = path.Join(os.Getenv("HOME"),".twitch.toml")
    Conf config = config{}
	Irc  twitchirc = twitchirc{Config:&Conf,Connected:false}
)

type config struct {
	User 	 string `toml:"usr"`
	Password string `toml:"pass"`
	is_loaded bool  `toml:-`
}

func (c *config)Load() error {
	reader , err := os.Open(Conffile)
	if err != nil {return err}
	data, err := ioutil.ReadAll(reader)
	if err != nil {return err}
	err = toml.Unmarshal(data,c)
	if err != nil {return err}
	c.is_loaded = true
	return nil
}

type twitchirc struct {
	Client *twitch.Client
	Config *config
	Connected bool
	Messages_out chan twitch.PrivateMessage
}

func (t *twitchirc)OnMesage(m twitch.PrivateMessage){
	t.Messages_out <- m
}

func (t *twitchirc)OnConnect(){
	log.Info("Connected")
	t.Connected = true
}

func (t *twitchirc)Init() error {
	t.Messages_out = make(chan twitch.PrivateMessage)
	if t.Config == nil {
		t.Config = &Conf
	}

	if !t.Config.is_loaded {
		t.Config.Load()
	}

	if t.Client == nil {
		t.Client = twitch.NewClient(t.Config.User,t.Config.Password)
	}

	t.Client.OnPrivateMessage(t.OnMesage)
	t.Client.OnConnect(t.OnConnect)

	return nil
}

func (t *twitchirc)Connect() error {
	t.Client.Join("oonray")
	err := t.Client.Connect()
	return err
}
