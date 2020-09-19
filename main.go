package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	configTypes "main/configStructs"
	"main/handler"

	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		ldata, err := ioutil.ReadFile("/app/config/config.json")
		if err != nil {
			fmt.Println("error:", err)
		}

		data = ldata
	}

	var config configTypes.AppConfig

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error:", err)
	}

	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	listener := make(chan *docker.APIEvents)
	err = client.AddEventListener(listener)
	if err != nil {
		log.Fatal(err)
	}

	handler.OrchestrateBuildConfig(client, config)

	for {
		select {
		case msg := <-listener:
			handler.HandleDockerEvent(msg, client, config)
		}
	}
}
