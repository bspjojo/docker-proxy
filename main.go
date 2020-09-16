package main

import (
	"log"
	"main/handler"

	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	listener := make(chan *docker.APIEvents)
	err = client.AddEventListener(listener)
	if err != nil {
		log.Fatal(err)
	}

	handler.OrchestrateBuildConfig(client)

	for {
		select {
		case msg := <-listener:
			handler.HandleDockerEvent(msg, client)
		}
	}
}
