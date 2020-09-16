package handler

import (
	"log"

	"main/handler/stringBuilder"
	"main/handler/writer"

	docker "github.com/fsouza/go-dockerclient"
)

func HandleDockerEvent(msg *docker.APIEvents, client *docker.Client) {
	if msg.Type == "container" {
		log.Println("=============================")
		log.Println(msg)
		image := msg.Actor.Attributes["image"]
		log.Println("image ", image)

		switch msg.Action {
		case "create",
			"destroy",
			"die",
			"exec_create",
			"exec_detach",
			"exec_die",
			"exec_start",
			"kill",
			"pause",
			"restart",
			"start",
			"stop",
			"unpause",
			"update":
			log.Println(msg.Action)

			log.Println("Action ", msg.Action)
			log.Println("Type ", msg.Type)
			log.Println("Actor ", msg.Actor)
			log.Println("Status ", msg.Status)
			log.Println("From ", msg.From)
			log.Println("Time", msg.Time)
			log.Println("TimeNano", msg.TimeNano)
			log.Println("=============================")

			OrchestrateBuildConfig(client)
		}
	}
}

func OrchestrateBuildConfig( client *docker.Client) {
	content := stringBuilder.BuildString(client)
	writer.WriteContent(content)
}
