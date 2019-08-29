package main

import (
	"encoding/json"
	"flag"
	"github.com/jpillora/chisel/client"
	chshare "github.com/jpillora/chisel/share"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
)

var configPath = flag.String("c", "", "config file path")
var generate = flag.String("g", "", "generate config")

type ConfigFile struct {
	Server  string   `json:"server"`
	Remotes []string `json:"remotes"`
}

func main() {
	flag.Parse()
	if *configPath != "" {
		b, err := ioutil.ReadFile(*configPath)
		if err != nil {
			logrus.Panic(err)
		}
		config := ConfigFile{}
		err = json.Unmarshal(b, &config)
		if err != nil {
			logrus.Panic(err)
		}

		c, err := chclient.NewClient(&chclient.Config{
			KeepAlive:        0,
			MaxRetryCount:    -1,
			MaxRetryInterval: 0,
			Server:           config.Server,
			Remotes:          config.Remotes,
		})
		go chshare.GoStats()
		if err = c.Run(); err != nil {
			log.Fatal(err)
		}

		sigs := make(chan os.Signal, 1)
		<-sigs
	}
	if *generate != "" {
		logrus.Println("Writing config file", *generate)
		c := ConfigFile{
			Server: "http://host:port",
			Remotes: []string{
				"localhost:8080:remotehost:8080",
				"localhost:8081:remotehost:8081",
			},
		}
		b, _ := json.Marshal(c)
		ioutil.WriteFile(*generate, b, 0644)
	}
	logrus.Println("Nothing left to do")
}
